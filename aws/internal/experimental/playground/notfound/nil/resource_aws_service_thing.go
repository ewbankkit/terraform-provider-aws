package notfound

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/example"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/nil/deleter"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/nil/finder"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/nil/waiter"
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

func testAccCheckAwsExampleThingExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := &example.Example{}
		//conn := testAccProvider.Meta().(*AWSClient).exampleconn

		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		thing, err := finder.ThingByID(conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		if thing == nil {
			return fmt.Errorf("Example Thing %s not found", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckAwsExampleThingDestroy(s *terraform.State) error {
	conn := &example.Example{}
	//conn := testAccProvider.Meta().(*AWSClient).globalacceleratorconn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_example_thing" {
			continue
		}

		thing, err := finder.ThingByID(conn, rs.Primary.ID)

		if thing == nil {
			continue
		}

		if err != nil {
			return err
		}

		return fmt.Errorf("Example Thing %s still exists", rs.Primary.ID)
	}
	return nil
}
