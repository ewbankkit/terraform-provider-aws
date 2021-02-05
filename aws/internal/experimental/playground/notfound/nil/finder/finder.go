package finder

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/namevaluesfilters"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/service"
)

// ThingByID returns the thing corresponding to the specified ID.
// Returns nil if no thing is found.
// func ThingByID(conn *service.Service, thingID string) (*service.Thing, error) {
// 	input := &service.GetThingInput{
// 		ThingId: aws.String(thingID),
// 	}

// 	return Thing(conn, input)
// }

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

// ThingByID returns the thing corresponding to the specified ID.
// Returns nil if no thing is found.
func ThingByID(conn *service.Service, thingID string) (*service.Thing, error) {
	input := &service.GetThingsInput{
		ThingIds: aws.StringSlice([]string{thingID}),
	}

	things, err := Things(conn, input)

	if err != nil {
		return nil, err
	}

	if things == nil {
		return nil, nil
	}

	return things[0], nil
}

// ThingsByNameValuesFilters returns the things corresponding to the specified filter.
// Returns nil if no thing is found.
func ThingsByNameValuesFilters(conn *service.Service, filters namevaluesfilters.NameValuesFilters) ([]*service.Thing, error) {
	input := &service.GetThingsInput{
		Filters: filters.ServiceFilters(),
	}

	things, err := Things(conn, input)

	if err != nil {
		return nil, err
	}

	return things, nil
}

// Things returns the things corresponding to the specified input.
// Returns nil if no things are found.
func Things(conn *service.Service, input *service.GetThingsInput) ([]*service.Thing, error) {
	var things []*service.Thing

	err := conn.GetThingsPages(input, func(page *service.GetThingsOutput, isLast bool) bool {
		if page == nil {
			return !isLast
		}

		for _, thing := range page.Things {
			if thing == nil {
				continue
			}

			// If Thing has status(es) indicating that nothing more can be done with the thing
			// and that the thing will eventually be garbage collected by AWS, ignore the thing.
			if aws.StringValue(thing.Status) == service.ThingStatusDeleted {
				continue
			}

			things = append(things, thing)
		}

		return !isLast
	})

	// If the AWS API signals that the thing doesn't exist, return nil.
	if tfawserr.ErrCodeEquals(err, service.ErrCodeResourceNotFoundException) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	// Handle any empty result.
	if len(things) == 0 {
		return nil, nil
	}

	return things, nil
}
