package aws

import (
	"context"

	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/tfresource"
)

// Returns true if the error matches all these conditions:
//  * err is of type awserr.Error
//  * Error.Code() matches code
//  * Error.Message() contains message
func isAWSErr(err error, code string, message string) bool {
	var awsErr awserr.Error
	if errors.As(err, &awsErr) {
		return awsErr.Code() == code && strings.Contains(awsErr.Message(), message)
	}
	return false
}

// Returns true if the error matches all these conditions:
//  * err is of type awserr.RequestFailure
//  * RequestFailure.StatusCode() matches status code
// It is always preferable to use isAWSErr() except in older APIs (e.g. S3)
// that sometimes only respond with status codes.
func isAWSErrRequestFailureStatusCode(err error, statusCode int) bool {
	var awsErr awserr.RequestFailure
	if errors.As(err, &awsErr) {
		return awsErr.StatusCode() == statusCode
	}
	return false
}

// retryOnAwsCode retries a function (for up to 2 minutes) when it returns the specified AWS error codes.
// The retried function's return value is this function's return value.
func retryOnAwsCode(code string, f func() (interface{}, error)) (interface{}, error) {
	return tfresource.RetryOnAWSErrorCodes(context.Background(), tfresource.RetryTimeout, func(_ context.Context) (interface{}, error) {
		return f()
	}, code)
}
