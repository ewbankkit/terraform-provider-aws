package notfound

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/nil/deleter"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/nil/finder"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/nil/waiter"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/example"
)

func resourceAwsExampleThing() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsExampleThingCreate,
		Read:   resourceAwsExampleThingRead,
		Update: resourceAwsExampleThingUpdate,
		Delete: resourceAwsExampleThingDelete,

		Schema: map[string]*schema.Schema{
			// All the attributes.
		},
	}
}

func resourceAwsExampleThingCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")

	return resourceAwsExampleThingRead(d, meta)
}

func resourceAwsExampleThingRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*example.Example)

	thing, err := finder.ThingByID(conn, d.Id())

	if err != nil {
		return fmt.Errorf("error reading Example Thing (%s): %w", d.Id(), err)
	}

	if !d.IsNewResource() && thing == nil {
		log.Printf("[WARN] Example Thing (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if thing == nil {
		return fmt.Errorf("error reading Example Thing (%s): Not found", d.Id())
	}

	log.Printf("%#v", thing)

	// Set all the attributes.

	return nil
}

func resourceAwsExampleThingUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceAwsExampleThingRead(d, meta)
}

func resourceAwsExampleThingDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*example.Example)

	err := deleter.ThingByID(conn, d.Id())

	if err != nil {
		return fmt.Errorf("error deleting Example Thing (%s): %w", d.Id(), err)
	}

	// If the deletion occurs asynchronously, wait.
	_, err = waiter.ThingDeleted(conn, d.Id())

	if err != nil {
		return fmt.Errorf("error waiting for Example Thing (%s) to delete: %w", d.Id(), err)
	}

	return nil
}
