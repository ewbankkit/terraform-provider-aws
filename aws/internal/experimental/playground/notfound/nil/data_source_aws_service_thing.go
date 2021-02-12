package notfound

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/namevaluesfilters"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/nil/finder"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/example"
)

func dataSourceAwsExampleThing() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsExampleThingRead,

		Schema: map[string]*schema.Schema{
			// All the attributes.
		},
	}
}

func dataSourceAwsExampleThingRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*example.Example)

	filters := namevaluesfilters.New(map[string]string{
		"name-1": "value1",
		"name-2": "value2",
	})

	things, err := finder.ThingsByNameValuesFilters(conn, filters)

	if err != nil {
		return fmt.Errorf("error reading Example Things: %w", err)
	}

	if things == nil {
		return fmt.Errorf("no Example Things matched")
	}

	if n := len(things); n > 0 {
		return fmt.Errorf("%d Example Things matched; use additional constraints to reduce matches to a single Thing", n)
	}

	// Set all the attributes.

	return nil
}
