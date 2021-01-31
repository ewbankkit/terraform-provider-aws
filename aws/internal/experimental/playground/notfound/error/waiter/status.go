package waiter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/error/finder"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/service"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/tfresource"
)

func ThingStatus(conn *service.Service, thingID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		thing, err := finder.ThingByID(conn, thingID)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return thing, aws.StringValue(thing.Status), nil
	}
}
