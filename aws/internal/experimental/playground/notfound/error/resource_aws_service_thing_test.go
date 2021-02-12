package notfound

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/example"
)

func Test_ResourceRead(t *testing.T) {
	testCases := []struct {
		TestName      string
		Conn          *example.Example
		ThingID       string
		IsNewResource bool
		ExpectError   bool
		ExpectedID    string
	}{
		{
			TestName:      "Valid thing new",
			Conn:          example.New(),
			ThingID:       example.VALID_THING_ID,
			IsNewResource: true,
			ExpectError:   false,
			ExpectedID:    example.VALID_THING_ID,
		},
		{
			TestName:      "Valid thing old",
			Conn:          example.New(),
			ThingID:       example.VALID_THING_ID,
			IsNewResource: false,
			ExpectError:   false,
			ExpectedID:    example.VALID_THING_ID,
		},
		{
			TestName:      "Not found thing new",
			Conn:          example.New(),
			ThingID:       example.NOTFOUND_THING_ID,
			IsNewResource: true,
			ExpectError:   true,
			ExpectedID:    example.NOTFOUND_THING_ID,
		},
		{
			TestName:      "Not found thing old",
			Conn:          example.New(),
			ThingID:       example.NOTFOUND_THING_ID,
			IsNewResource: false,
			ExpectError:   false,
			ExpectedID:    "",
		},
		{
			TestName:      "Empty result thing new",
			Conn:          example.New(),
			ThingID:       example.EMPTY_RESULT_THING_ID,
			IsNewResource: true,
			ExpectError:   true,
			ExpectedID:    example.EMPTY_RESULT_THING_ID,
		},
		{
			TestName:      "Empty result thing old",
			Conn:          example.New(),
			ThingID:       example.EMPTY_RESULT_THING_ID,
			IsNewResource: false,
			ExpectError:   false,
			ExpectedID:    "",
		},
		{
			TestName:      "Erroring thing new",
			Conn:          example.New(),
			ThingID:       example.ERRORING_THING_ID,
			IsNewResource: true,
			ExpectError:   true,
			ExpectedID:    example.ERRORING_THING_ID,
		},
		{
			TestName:      "Erroring thing old",
			Conn:          example.New(),
			ThingID:       example.ERRORING_THING_ID,
			IsNewResource: false,
			ExpectError:   true,
			ExpectedID:    example.ERRORING_THING_ID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			r := resourceAwsExampleThing()
			d := r.Data(nil)
			d.SetId(testCase.ThingID)
			if testCase.IsNewResource {
				d.MarkNewResource()
			}
			err := r.Read(d, testCase.Conn)

			if testCase.ExpectError && err == nil {
				t.Errorf("%s expected an error but got nil", testCase.TestName)
			}

			if !testCase.ExpectError && err != nil {
				t.Errorf("%s did not expect an error but got %q", testCase.TestName, err)
			}

			if testCase.ExpectedID != d.Id() {
				t.Errorf("%s got ID %s, expected %s", testCase.TestName, d.Id(), testCase.ExpectedID)
			}
		})
	}
}

func Test_ResourceDelete(t *testing.T) {
	testCases := []struct {
		TestName    string
		Conn        *example.Example
		ThingID     string
		ExpectError bool
	}{
		{
			TestName:    "Valid thing",
			Conn:        example.New(),
			ThingID:     example.VALID_THING_ID,
			ExpectError: false,
		},
		{
			TestName:    "Not found thing",
			Conn:        example.New(),
			ThingID:     example.NOTFOUND_THING_ID,
			ExpectError: false,
		},
		{
			TestName:    "Erroring thing",
			Conn:        example.New(),
			ThingID:     example.ERRORING_THING_ID,
			ExpectError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			r := resourceAwsExampleThing()
			d := r.Data(nil)
			d.SetId(testCase.ThingID)
			err := r.Delete(d, testCase.Conn)

			if testCase.ExpectError && err == nil {
				t.Errorf("%s expected an error but got nil", testCase.TestName)
			}

			if !testCase.ExpectError && err != nil {
				t.Errorf("%s did not expect an error but got %q", testCase.TestName, err)
			}
		})
	}
}
