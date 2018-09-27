package frontier

import (
	"github.com/digitalbitsorg/go/services/frontier/internal/ledger"
	"github.com/digitalbitsorg/go/services/frontier/internal/resource"
	"github.com/digitalbitsorg/go/support/render/hal"
)

// RootAction provides a summary of the frontier instance and links to various
// useful endpoints
type RootAction struct {
	Action
}

// JSON renders the json response for RootAction
func (action *RootAction) JSON() {
	action.App.UpdateStellarCoreInfo()

	var res resource.Root
	res.Populate(
		action.Ctx,
		ledger.CurrentState(),
		action.App.horizonVersion,
		action.App.coreVersion,
		action.App.networkPassphrase,
		action.App.protocolVersion,
	)

	hal.Render(action.W, res)
}
