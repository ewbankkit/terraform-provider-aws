package waiter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/example"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/nil/finder"
)

func ThingStatus(conn *example.Example, thingID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		thing, err := finder.ThingByID(conn, thingID)

		if err != nil {
			return nil, "", err
		}

		if thing == nil {
			return nil, "", nil
		}

		return thing, aws.StringValue(thing.Status), nil
	}
}
