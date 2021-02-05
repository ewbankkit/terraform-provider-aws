package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

//
// A play AWS service API.
//

type Service struct {
	n int // Number of times that GetThing has been called.
}

type Thing struct {
	ThingId *string
	Status  *string
}

type GetThingInput struct {
	ThingId *string
}

type GetThingOutput struct {
	Thing *Thing
}

type GetThingsInput struct {
	Filters   []*Filter
	ThingIds  []*string
	NextToken *string
}

type GetThingsOutput struct {
	Things    []*Thing
	NextToken *string
}

type DeleteThingInput struct {
	ThingId *string
}

type DeleteThingOutput struct{}

type Filter struct{}

func (c *Service) GetThing(input *GetThingInput) (*GetThingOutput, error) {
	if aws.StringValue(input.ThingId) == EMPTY_RESULT_THING_ID {
		return &GetThingOutput{}, nil
	}

	if aws.StringValue(input.ThingId) == ERRORING_THING_ID {
		return nil, awserr.New(ErrCodeInvalidArgumentException, "erroring", nil)
	}

	if aws.StringValue(input.ThingId) == NOTFOUND_THING_ID {
		return nil, awserr.New(ErrCodeResourceNotFoundException, "not found", nil)
	}

	var status string
	switch c.n {
	case 0:
		status = ThingStatusReady
	case 1:
		status = ThingStatusDeleting
	default:
		status = ThingStatusDeleted
	}

	c.n = c.n + 1

	return &GetThingOutput{Thing: &Thing{ThingId: input.ThingId, Status: aws.String(status)}}, nil
}

func (c *Service) GetThings(input *GetThingsInput) (*GetThingsOutput, error) {
	if len(input.ThingIds) == 0 {
		return &GetThingsOutput{Things: []*Thing{}}, nil
	}

	thingID := aws.StringValue(input.ThingIds[0])

	if thingID == EMPTY_RESULT_THING_ID {
		return &GetThingsOutput{}, nil
	}

	if thingID == ERRORING_THING_ID {
		return nil, awserr.New(ErrCodeInvalidArgumentException, "erroring", nil)
	}

	if thingID == NOTFOUND_THING_ID {
		return nil, awserr.New(ErrCodeResourceNotFoundException, "not found", nil)
	}

	var status string
	switch c.n {
	case 0:
		status = ThingStatusReady
	case 1:
		status = ThingStatusDeleting
	default:
		status = ThingStatusDeleted
	}

	c.n = c.n + 1

	return &GetThingsOutput{Things: []*Thing{{ThingId: aws.String(thingID), Status: aws.String(status)}}}, nil
}

func (c *Service) DeleteThing(input *DeleteThingInput) (*DeleteThingOutput, error) {
	if aws.StringValue(input.ThingId) == ERRORING_THING_ID {
		return nil, awserr.New(ErrCodeInvalidArgumentException, "erroring", nil)
	}

	if aws.StringValue(input.ThingId) == NOTFOUND_THING_ID {
		return nil, awserr.New(ErrCodeResourceNotFoundException, "not found", nil)
	}

	return &DeleteThingOutput{}, nil
}

func (c *Service) GetThingsPages(input *GetThingsInput, fn func(*GetThingsOutput, bool) bool) error {
	for {
		output, err := c.GetThings(input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}

func New() *Service {
	return &Service{}
}

const (
	EMPTY_RESULT_THING_ID = "thing-0"
	ERRORING_THING_ID     = "thing-1"
	NOTFOUND_THING_ID     = "thing-2"
	VALID_THING_ID        = "thing-3"
)

const (
	ErrCodeInvalidArgumentException  = "InvalidArgumentException"
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)

const (
	ThingStatusDeleted  = "DELETED"
	ThingStatusDeleting = "DELETING"
	ThingStatusReady    = "READY"
)
