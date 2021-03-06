// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package efsiface provides an interface to enable mocking the Amazon Elastic File System service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package efsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/efs"
)

// EFSAPI provides an interface to enable mocking the
// efs.EFS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Elastic File System.
//    func myFunc(svc efsiface.EFSAPI) bool {
//        // Make svc.CreateFileSystem request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := efs.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockEFSClient struct {
//        efsiface.EFSAPI
//    }
//    func (m *mockEFSClient) CreateFileSystem(input *efs.CreateFileSystemInput) (*efs.FileSystemDescription, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockEFSClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type EFSAPI interface {
	CreateFileSystem(*efs.CreateFileSystemInput) (*efs.FileSystemDescription, error)
	CreateFileSystemWithContext(aws.Context, *efs.CreateFileSystemInput, ...request.Option) (*efs.FileSystemDescription, error)
	CreateFileSystemRequest(*efs.CreateFileSystemInput) (*request.Request, *efs.FileSystemDescription)

	CreateMountTarget(*efs.CreateMountTargetInput) (*efs.MountTargetDescription, error)
	CreateMountTargetWithContext(aws.Context, *efs.CreateMountTargetInput, ...request.Option) (*efs.MountTargetDescription, error)
	CreateMountTargetRequest(*efs.CreateMountTargetInput) (*request.Request, *efs.MountTargetDescription)

	CreateTags(*efs.CreateTagsInput) (*efs.CreateTagsOutput, error)
	CreateTagsWithContext(aws.Context, *efs.CreateTagsInput, ...request.Option) (*efs.CreateTagsOutput, error)
	CreateTagsRequest(*efs.CreateTagsInput) (*request.Request, *efs.CreateTagsOutput)

	DeleteFileSystem(*efs.DeleteFileSystemInput) (*efs.DeleteFileSystemOutput, error)
	DeleteFileSystemWithContext(aws.Context, *efs.DeleteFileSystemInput, ...request.Option) (*efs.DeleteFileSystemOutput, error)
	DeleteFileSystemRequest(*efs.DeleteFileSystemInput) (*request.Request, *efs.DeleteFileSystemOutput)

	DeleteMountTarget(*efs.DeleteMountTargetInput) (*efs.DeleteMountTargetOutput, error)
	DeleteMountTargetWithContext(aws.Context, *efs.DeleteMountTargetInput, ...request.Option) (*efs.DeleteMountTargetOutput, error)
	DeleteMountTargetRequest(*efs.DeleteMountTargetInput) (*request.Request, *efs.DeleteMountTargetOutput)

	DeleteTags(*efs.DeleteTagsInput) (*efs.DeleteTagsOutput, error)
	DeleteTagsWithContext(aws.Context, *efs.DeleteTagsInput, ...request.Option) (*efs.DeleteTagsOutput, error)
	DeleteTagsRequest(*efs.DeleteTagsInput) (*request.Request, *efs.DeleteTagsOutput)

	DescribeFileSystems(*efs.DescribeFileSystemsInput) (*efs.DescribeFileSystemsOutput, error)
	DescribeFileSystemsWithContext(aws.Context, *efs.DescribeFileSystemsInput, ...request.Option) (*efs.DescribeFileSystemsOutput, error)
	DescribeFileSystemsRequest(*efs.DescribeFileSystemsInput) (*request.Request, *efs.DescribeFileSystemsOutput)

	DescribeMountTargetSecurityGroups(*efs.DescribeMountTargetSecurityGroupsInput) (*efs.DescribeMountTargetSecurityGroupsOutput, error)
	DescribeMountTargetSecurityGroupsWithContext(aws.Context, *efs.DescribeMountTargetSecurityGroupsInput, ...request.Option) (*efs.DescribeMountTargetSecurityGroupsOutput, error)
	DescribeMountTargetSecurityGroupsRequest(*efs.DescribeMountTargetSecurityGroupsInput) (*request.Request, *efs.DescribeMountTargetSecurityGroupsOutput)

	DescribeMountTargets(*efs.DescribeMountTargetsInput) (*efs.DescribeMountTargetsOutput, error)
	DescribeMountTargetsWithContext(aws.Context, *efs.DescribeMountTargetsInput, ...request.Option) (*efs.DescribeMountTargetsOutput, error)
	DescribeMountTargetsRequest(*efs.DescribeMountTargetsInput) (*request.Request, *efs.DescribeMountTargetsOutput)

	DescribeTags(*efs.DescribeTagsInput) (*efs.DescribeTagsOutput, error)
	DescribeTagsWithContext(aws.Context, *efs.DescribeTagsInput, ...request.Option) (*efs.DescribeTagsOutput, error)
	DescribeTagsRequest(*efs.DescribeTagsInput) (*request.Request, *efs.DescribeTagsOutput)

	ModifyMountTargetSecurityGroups(*efs.ModifyMountTargetSecurityGroupsInput) (*efs.ModifyMountTargetSecurityGroupsOutput, error)
	ModifyMountTargetSecurityGroupsWithContext(aws.Context, *efs.ModifyMountTargetSecurityGroupsInput, ...request.Option) (*efs.ModifyMountTargetSecurityGroupsOutput, error)
	ModifyMountTargetSecurityGroupsRequest(*efs.ModifyMountTargetSecurityGroupsInput) (*request.Request, *efs.ModifyMountTargetSecurityGroupsOutput)
}

var _ EFSAPI = (*efs.EFS)(nil)
