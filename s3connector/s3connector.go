package s3connector

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/golang/glog"
	"io"
	"os"
)

type Connector struct {
	svc        *s3.S3
	cfg        *aws.Config
	session    *session.Session
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func New() (s Connector) {
	s = Connector{}
	s.cfg = aws.NewConfig().WithRegion("eu-west-1")
	// The session the S3 Uploader will use
	s.session = session.Must(session.NewSession(s.cfg))
	s.svc = s3.New(s.session, s.cfg)
	s.uploader = s3manager.NewUploader(s.session)
	s.downloader = s3manager.NewDownloader(s.session)
	return s
}

func (s *Connector) UploadS3(chartName string, fileReader io.Reader, bucket string) (err error) {
	result, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(chartName),
		Body:   fileReader,
	})
	glog.V(4).Infof("Amazon S3 upload result %s", result)
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	return
}

func (s *Connector) DownloadS3(pathChart string, chartName string, bucket string) (err error) {
	f, err := os.Create(pathChart)
	if err != nil {
		return fmt.Errorf("failed to create file %q, %v", pathChart, err)
	}
	// Write the contents of S3 Object to the file
	n, err := s.downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(chartName),
	})
	if err != nil {
		return fmt.Errorf("failed to download file, %v", err)
	}
	glog.V(4).Infof("Amazon S3 download result %d bytes", n)
	return err
}

func (s *Connector) ChecksumS3(chartName string, bucket string) (s3Etag string, err error) {
	obj, err := s.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(chartName),
	})
	if err != nil {
		return s3Etag, fmt.Errorf("failed to get Etag for %s, %s", chartName, err)
	}
	s3Etag = *obj.ETag
	glog.V(4).Infof("Remote S3 object ETag %d", s3Etag)
	return
}
