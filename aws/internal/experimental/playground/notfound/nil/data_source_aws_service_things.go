package notfound

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/example"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/namevaluesfilters"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/nil/finder"
)

func dataSourceAwsExampleThings() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsExampleThingsRead,

		Schema: map[string]*schema.Schema{
			// All the attributes.
		},
	}
}

func dataSourceAwsExampleThingsRead(d *schema.ResourceData, meta interface{}) error {
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

	// Set all the attributes.
	for _, thing := range things {
		log.Printf("%#v", thing)
	}

	return nil
}
