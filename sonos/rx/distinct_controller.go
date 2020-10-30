package rx

import (
	"context"

	"github.com/LukasKnuth/sonos_ctrl/sonos/models"
)

// Stores state-information for a stateful distinct controller filter
type DistinctController struct {
	controllers map[string]string
}

// Create a new, stateful distinct operation for rxGo, filtering Controllers.
func StatefulDistinctController() *DistinctController {
	return &DistinctController{controllers: make(map[string]string)}
}

// rxGo function to filter an Observable emitting Controllers by distinct ones
// based on IP and USN.
func (state *DistinctController) Distinct(ctx context.Context, item interface{}) (interface{}, error) {
	controller := item.(*models.Controller)
	if currentUsn, ok := state.controllers[controller.IP]; ok {
		if currentUsn == controller.USN {
			// Already know about this one!
			return nil, nil
		} else {
			// Knew the IP, apparently a new deviec!
			state.controllers[controller.IP] = controller.USN
			return nil, nil
		}
	} else {
		// New, unknown controller
		state.controllers[controller.IP] = controller.USN
		return item, nil
	}
}
