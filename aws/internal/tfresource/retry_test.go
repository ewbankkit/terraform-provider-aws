package tfresource

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

func TestRetryOnAWSErrorCodes(t *testing.T) {
	testCases := []struct {
		Name string
		Err  error
		Ok   bool
	}{
		{
			Name: "nil error",
			Err:  nil,
			Ok:   true,
		},
		{
			Name: "other error",
			Err:  errors.New("test"),
			Ok:   false,
		},
		{
			Name: "awserr matching code 1",
			Err:  awserr.New("TestCode1", "TestMessage", nil),
			Ok:   true,
		},
		{
			Name: "awserr matching code 2",
			Err:  awserr.New("TestCode2", "TestMessage", nil),
			Ok:   true,
		},
		{
			Name: "awserr non-matching code",
			Err:  awserr.New("TestCode3", "TestMessage", nil),
			Ok:   false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			n := 0

			f := func(ctx context.Context) (interface{}, error) {
				n = n + 1

				// Alway return success after one retry.
				if n > 1 || testCase.Err == nil {
					return 42, nil
				}

				return nil, testCase.Err
			}

			_, err := RetryOnAWSErrorCodes(context.TODO(), RetryTimeout, f, "TestCode1", "TestCode2")

			if err == nil && !testCase.Ok {
				t.Errorf("unexpected success")
			}

			if err != nil && testCase.Ok {
				t.Errorf("unexpected failure: %#v", err)
			}
		})
	}
}
