// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDBSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the DB instance that you want to create the snapshot of.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBInstance.
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`

	// The identifier for the DB snapshot.
	//
	// Constraints:
	//
	//    * Can't be null, empty, or blank
	//
	//    * Must contain from 1 to 255 letters, numbers, or hyphens
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// Example: my-snapshot-id
	//
	// DBSnapshotIdentifier is a required field
	DBSnapshotIdentifier *string `type:"string" required:"true"`

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateDBSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDBSnapshotInput"}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if s.DBSnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBSnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDBSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB snapshot.
	//
	// This data type is used as a response element in the DescribeDBSnapshots action.
	DBSnapshot *DBSnapshot `type:"structure"`
}

// String returns the string representation
func (s CreateDBSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDBSnapshot = "CreateDBSnapshot"

// CreateDBSnapshotRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Creates a DBSnapshot. The source DBInstance must be in "available" state.
//
//    // Example sending a request using CreateDBSnapshotRequest.
//    req := client.CreateDBSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateDBSnapshot
func (c *Client) CreateDBSnapshotRequest(input *CreateDBSnapshotInput) CreateDBSnapshotRequest {
	op := &aws.Operation{
		Name:       opCreateDBSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBSnapshotInput{}
	}

	req := c.newRequest(op, input, &CreateDBSnapshotOutput{})
	return CreateDBSnapshotRequest{Request: req, Input: input, Copy: c.CreateDBSnapshotRequest}
}

// CreateDBSnapshotRequest is the request type for the
// CreateDBSnapshot API operation.
type CreateDBSnapshotRequest struct {
	*aws.Request
	Input *CreateDBSnapshotInput
	Copy  func(*CreateDBSnapshotInput) CreateDBSnapshotRequest
}

// Send marshals and sends the CreateDBSnapshot API request.
func (r CreateDBSnapshotRequest) Send(ctx context.Context) (*CreateDBSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBSnapshotResponse{
		CreateDBSnapshotOutput: r.Request.Data.(*CreateDBSnapshotOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBSnapshotResponse is the response type for the
// CreateDBSnapshot API operation.
type CreateDBSnapshotResponse struct {
	*CreateDBSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBSnapshot request.
func (r *CreateDBSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
