package finder

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/example"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/namevaluesfilters"
)

// ThingByID returns the thing corresponding to the specified ID.
// Returns NotFoundError if no thing is found.
// func ThingByID(conn *example.Example, thingID string) (*example.Thing, error) {
// 	input := &example.GetThingInput{
// 		ThingId: aws.String(thingID),
// 	}

// 	return Thing(conn, input)
// }

// Thing returns the thing corresponding to the specified input.
// Returns NotFoundError if no thing is found.
func Thing(conn *example.Example, input *example.GetThingInput) (*example.Thing, error) {
	output, err := conn.GetThing(input)

	// If the AWS API signals that the thing doesn't exist, return NotFoundError.
	if tfawserr.ErrCodeEquals(err, example.ErrCodeResourceNotFoundException) {
		return nil, &resource.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	// Handle any empty result.
	if output == nil || output.Thing == nil {
		return nil, &resource.NotFoundError{
			Message:     "Empty result",
			LastRequest: input,
		}
	}

	// If Thing has status(es) indicating that nothing more can be done with the thing
	// and that the thing will eventually be garbage collected by AWS, return NotFoundError.
	if status := aws.StringValue(output.Thing.Status); status == example.ThingStatusDeleted {
		return nil, &resource.NotFoundError{
			Message:     status,
			LastRequest: input,
		}
	}

	return output.Thing, nil
}

// ThingByID returns the thing corresponding to the specified ID.
// Returns NotFoundError if no thing is found.
func ThingByID(conn *example.Example, thingID string) (*example.Thing, error) {
	input := &example.GetThingsInput{
		ThingIds: aws.StringSlice([]string{thingID}),
	}

	things, err := Things(conn, input)

	if err != nil {
		return nil, err
	}

	// Handle any empty result.
	if len(things) == 0 {
		return nil, &resource.NotFoundError{
			Message:     "Empty result",
			LastRequest: input,
		}
	}

	return things[0], nil
}

// ThingsByNameValuesFilters returns the things corresponding to the specified filter.
// Returns an empty slice if no thing is found.
func ThingsByNameValuesFilters(conn *example.Example, filters namevaluesfilters.NameValuesFilters) ([]*example.Thing, error) {
	input := &example.GetThingsInput{
		Filters: filters.ExampleFilters(),
	}

	things, err := Things(conn, input)

	if err != nil {
		return nil, err
	}

	return things, nil
}

// Things returns the things corresponding to the specified input.
// Returns an empty slice if no things are found.
func Things(conn *example.Example, input *example.GetThingsInput) ([]*example.Thing, error) {
	var things []*example.Thing

	err := conn.GetThingsPages(input, func(page *example.GetThingsOutput, isLast bool) bool {
		if page == nil {
			return !isLast
		}

		for _, thing := range page.Things {
			if thing == nil {
				continue
			}

			// If Thing has status(es) indicating that nothing more can be done with the thing
			// and that the thing will eventually be garbage collected by AWS, ignore the thing.
			if aws.StringValue(thing.Status) == example.ThingStatusDeleted {
				continue
			}

			things = append(things, thing)
		}

		return !isLast
	})

	// If the AWS API signals that the thing doesn't exist, return an empty slice.
	if tfawserr.ErrCodeEquals(err, example.ErrCodeResourceNotFoundException) {
		return things, nil
	}

	if err != nil {
		return nil, err
	}

	return things, nil
}
