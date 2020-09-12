package tfresource

import (
	"context"
	"time"

	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	// Default timeout
	RetryTimeout = 2 * time.Minute
)

// RetryOnAWSErrorCodes retries a function when it returns one of the specified AWS error codes.
// The retried function's return value is this function's return value.
func RetryOnAWSErrorCodes(ctx context.Context, timeout time.Duration, f func(context.Context) (interface{}, error), codes ...string) (interface{}, error) {
	var resp interface{}

	err := resource.RetryContext(ctx, timeout, func() *resource.RetryError {
		var err error

		resp, err = f(ctx)

		if err == nil {
			return nil
		}

		for _, code := range codes {
			if tfawserr.ErrCodeEquals(err, code) {
				return resource.RetryableError(err)
			}
		}

		return resource.NonRetryableError(err)
	})

	return resp, err
}
