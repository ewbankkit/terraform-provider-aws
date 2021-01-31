package waiter

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/service"
)

func Test_ThingDeleted(t *testing.T) {
	testCases := []struct {
		TestName    string
		Conn        *service.Service
		ThingID     string
		ExpectError bool
	}{
		{
			TestName:    "Valid thing",
			Conn:        service.New(),
			ThingID:     service.VALID_THING_ID,
			ExpectError: false,
		},
		{
			TestName:    "Not found thing",
			Conn:        service.New(),
			ThingID:     service.NOTFOUND_THING_ID,
			ExpectError: false,
		},
		{
			TestName:    "Erroring thing",
			Conn:        service.New(),
			ThingID:     service.ERRORING_THING_ID,
			ExpectError: true,
		},
		{
			TestName:    "Empty result thing",
			Conn:        service.New(),
			ThingID:     service.EMPTY_RESULT_THING_ID,
			ExpectError: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			_, err := ThingDeleted(testCase.Conn, testCase.ThingID)

			if testCase.ExpectError && err == nil {
				t.Errorf("%s expected an error but got nil", testCase.TestName)
			}

			if !testCase.ExpectError && err != nil {
				t.Errorf("%s did not expect an error but got %q", testCase.TestName, err)
			}
		})
	}
}
