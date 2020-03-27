package apis

import (
	"strings"
)

// Callback handles all callbacks from inline keyboards.
func (api *API) Callback(cb *Callback) {
	// Get callback action and data.
	cb.Data = strings.TrimSpace(cb.Data)
	parts := strings.Split(cb.Data, "|")
	if len(parts) != 2 {
		api.LogErrorf(cb, "[Callback] invalid callback data: [%s]", cb.Data)
		api.ErrorCallback(cb)
		return
	}
	action, data := parts[0], parts[1]

	switch action {
	case actionDeleteImage:
		api.handleDelete(cb, data)

	case actionDeleteConfirm:
		api.handleDeleteConfirm(cb, data)

	case actionDeleteCancel:
		api.handleDeleteCancel(cb, data)

	default:
		api.UnsupportedCallback(cb)
	}
}

const (
	actionDeleteImage   = "delete"
	actionDeleteConfirm = "delete_confirm"
	actionDeleteCancel  = "delete_cancel"
)
