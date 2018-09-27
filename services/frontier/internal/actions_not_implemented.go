package frontier

import (
	hProblem "github.com/digitalbitsorg/go/services/frontier/internal/render/problem"
	"github.com/digitalbitsorg/go/support/render/problem"
)

// NotImplementedAction renders a NotImplemented prblem
type NotImplementedAction struct {
	Action
}

// JSON is a method for actions.JSON
func (action *NotImplementedAction) JSON() {
	problem.Render(action.Ctx, action.W, hProblem.NotImplemented)
}
