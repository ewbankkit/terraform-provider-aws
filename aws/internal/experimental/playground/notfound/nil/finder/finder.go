package finder

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/service"
)

// ThingByID returns the thing corresponding to the specified ID.
// Returns nil if no thing is found.
func ThingByID(conn *service.Service, thingID string) (*service.Thing, error) {
	input := &service.GetThingInput{
		ThingId: aws.String(thingID),
	}

	return Thing(conn, input)
}

// Thing returns the thing corresponding to the specified input.
// Returns nil if no thing is found.
func Thing(conn *service.Service, input *service.GetThingInput) (*service.Thing, error) {
	output, err := conn.GetThing(input)

	// If the AWS API signals that the thing doesn't exist, return nil.
	if tfawserr.ErrCodeEquals(err, service.ErrCodeResourceNotFoundException) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	// Handle any empty result.
	if output == nil || output.Thing == nil {
		return nil, nil
	}

	// If Thing has status(es) indicating that nothing more can be done with the thing
	// and that the thing will eventually be garbage collected by AWS, return nil.
	if status := aws.StringValue(output.Thing.Status); status == service.ThingStatusDeleted {
		return nil, nil
	}

	return output.Thing, nil
}
