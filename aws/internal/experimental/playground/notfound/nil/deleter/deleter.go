package deleter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/service"
)

// ThingByID deletes the thing corresponding to the specified ID.
// Returns nil if no thing is found.
func ThingByID(conn *service.Service, thingID string) error {
	input := &service.DeleteThingInput{
		ThingId: aws.String(thingID),
	}

	_, err := conn.DeleteThing(input)

	if tfawserr.ErrCodeEquals(err, service.ErrCodeResourceNotFoundException) {
		return nil
	}

	return err
}
