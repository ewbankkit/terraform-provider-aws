package notfound

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/error/deleter"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/error/finder"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/error/waiter"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/service"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/tfresource"
)

func resourceAwsServiceThing() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsServiceThingCreate,
		Read:   resourceAwsServiceThingRead,
		Update: resourceAwsServiceThingUpdate,
		Delete: resourceAwsServiceThingDelete,

		Schema: map[string]*schema.Schema{
			// All the attributes.
		},
	}
}

func resourceAwsServiceThingCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")

	return resourceAwsServiceThingRead(d, meta)
}

func resourceAwsServiceThingRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*service.Service)

	thing, err := finder.ThingByID(conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] Service Thing (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if err != nil {
		return fmt.Errorf("error reading Service Thing (%s): %w", d.Id(), err)
	}

	log.Printf("%#v", thing)

	// Set all the attributes.

	return nil
}

func resourceAwsServiceThingUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceAwsServiceThingRead(d, meta)
}

func resourceAwsServiceThingDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*service.Service)

	err := deleter.ThingByID(conn, d.Id())

	if tfresource.NotFound(err) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("error deleting Service Thing (%s): %w", d.Id(), err)
	}

	// If the deletion occurs asynchronously, wait.
	_, err = waiter.ThingDeleted(conn, d.Id())

	if err != nil {
		return fmt.Errorf("error waiting for Service Thing (%s) to delete: %w", d.Id(), err)
	}

	return nil
}
