package main

import (
	"path"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"

	"bytes"
	"crypto/md5"

	"github.com/daemonza/helmet/s3connector"
	"github.com/golang/glog"
	"io/ioutil"
)

var (
	url    *string
	host   *string
	port   *string
	charts *string
	bucket *string
	aws    *bool
	s3Conn s3connector.Connector
)

// helm is a wrapper function to execute the helm command on the
// shell.
func helm(arguments []string) (output []byte, err error) {

	command := "helm"
	cmd := exec.Command(command, arguments...)

	// Combine stdout and stderr
	glog.Info("updating helm repository index")
	output, err = cmd.CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
		return output, err
	}

	return output, nil
}

// initRepo initialize a helm repository generating a index.yaml file
func initRepo() error {
	// TODO check if directory is there and create
	// if needed
	err := os.MkdirAll(*charts, 0777)
	if err != nil {
		glog.Error(err.Error())
		return err
	}
	// generate helm index
	_, err = helm([]string{"repo", "index", *charts, "--url", *url})
	if err != nil {
		glog.Error(err.Error())
		return err
	}
	return nil
}

// upload uploads a given file to the the charts directory
func upload(c echo.Context) error {
	chartName := c.Param("chartName")
	// TODO - do some sanitising on chartName

	glog.Info("uploading " + chartName)
	pathCharts := path.Join(*charts, chartName)
	os.Stat(*charts)
	f, err := os.Create(pathCharts)
	defer f.Close()
	if err != nil {
		glog.Error(err.Error())
		return err
	}
	fileContent, err := ioutil.ReadAll(c.Request().Body)
	fileReader := bytes.NewReader(fileContent)
	_, err = io.Copy(f, fileReader)
	defer c.Request().Body.Close()
	if err != nil {
		glog.Error(err.Error())
		return err
	}
	glog.Info("done uploading " + chartName)

	if *aws {
		// Double Upload to S3
		err = s3Conn.UploadS3(chartName, fileReader, *bucket)
		if err != nil {
			glog.Errorf("Failed to upload to S3 %s", err)
		}
	}
	// generate helm index
	initRepo()

	return nil
}

// repo  serves back any files in the charts directory
// with content-type header set to text/yaml
func repo(c echo.Context) (err error) {
	//c.Response().Header().Set("content-type", "text/yaml")
	c.Response().Header().Set("content-type", "text/plain; charset=utf-8")
	chartName := c.Param("*")
	filePath := path.Join(*charts, chartName)

	if *aws {
		localFileSum, err := md5file(filePath)
		if err != nil {
			glog.V(4).Infof("Failed to checksum %s", err)
		}

		s3FileSum, err := s3Conn.ChecksumS3(chartName, *bucket)
		if err != nil {
			glog.V(4).Infof("Failed to check remote checksum %s", err)
		}

		if localFileSum == s3FileSum && localFileSum != "" {
			// local file when checksum match
			return c.File(filePath)
		}
		err = s3Conn.DownloadS3(filePath, chartName, *bucket)
		if err != nil {
			glog.Errorf("Failed to download in S3 %s", err)
		}
	}
	return c.File(filePath)
}

func md5file(filePath string) (md5checksum string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return md5checksum, err
	}
	md5checksum = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func init() {

	// Get command line options
	// repoURL is also the url that get's used to generate the helm repo index file
	url = flag.String("url", "http://localhost:1323/charts/", "The URL where Helmet runs as a repository")
	host = flag.String("host", "0.0.0.0", "The address that Helmet listens on")
	port = flag.String("port", "1323", "The port that Helmet listens on")
	charts = flag.String("charts", "./charts", "Directory where charts get's stored")
	bucket = flag.String("bucket", "charts", "The Bucket where Helmet upload")
	aws = flag.Bool("aws", false, "Use aws s3 as a backend")
	flag.Parse()
	// initialize the helm repository on startup.
	err := initRepo()
	if err != nil {
		glog.Fatal(err.Error())
	}
}

func main() {
	if *aws {
		s3Conn = s3connector.New()
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Endpoints
	e.PUT("/upload/:chartName", upload)

	// Serve the charts directory
	e.GET("/charts/*", repo)

	// Start server
	e.Logger.Fatal(e.Start(*host + ":" + *port))
}
