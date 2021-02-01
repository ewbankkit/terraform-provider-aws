package waiter

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/service"
)

const (
	// Maximum amount of time to wait for a Thing to be deleted.
	ThingDeletedTimeout = 5 * time.Minute
)

// ThingDeleted waits for a Thing to be deleted.
func ThingDeleted(conn *service.Service, thingID string) (*service.Thing, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{service.ThingStatusReady, service.ThingStatusDeleting},
		Target:  []string{},
		Refresh: ThingStatus(conn, thingID),
		Timeout: ThingDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if v, ok := outputRaw.(*service.Thing); ok {
		return v, err
	}

	return nil, err
}
