// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// DeviceBrowserService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceBrowserService] method instead.
type DeviceBrowserService struct {
	options []option.RequestOption
}

// NewDeviceBrowserService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceBrowserService(opts ...option.RequestOption) (r DeviceBrowserService) {
	r = DeviceBrowserService{}
	r.options = opts
	return
}

// Evaluates a JavaScript expression exactly once in the device's single, confirmed
// foreground Chrome tab via the Chrome DevTools Protocol and returns its
// JSON-serialized result. A responsive browser state without exactly one visible
// page returns DEVICE_NO_BROWSER_TARGET (400). Transient discovery, page-dial, or
// visibility-probe transport failures return CDP_TRANSPORT_UNAVAILABLE (503) with
// Retry-After and an executionOutcome; the service may start Chrome once and retry
// selection before Runtime.evaluate, but dial or probe recovery is skipped for a
// backgrounded Chrome, and a user script is never retried after its write or read.
// Devices without browser support return an unsupported-feature error.
func (r *DeviceBrowserService) ExecuteScript(ctx context.Context, deviceID string, params DeviceBrowserExecuteScriptParams, opts ...option.RequestOption) (res *DeviceBrowserExecuteScriptResponse, err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/browser/execute-script", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type DeviceBrowserExecuteScriptResponse struct {
	// JSON-serialized return value of the script (null if it returned undefined).
	// Non-JSON-serializable numbers (Infinity, NaN, -0) are returned as their string
	// representation.
	Result any `json:"result" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceBrowserExecuteScriptResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceBrowserExecuteScriptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceBrowserExecuteScriptParams struct {
	// JavaScript expression to evaluate in the device's foreground Chrome tab (CDP
	// Runtime.evaluate). It is an expression, not a function body — the expression's
	// value is returned (no top-level 'return'). Must evaluate to a JSON-serializable
	// value; wrap multi-statement logic in an IIFE, e.g. (() => { ... ; return x })().
	Script           string           `json:"script" api:"required"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceBrowserExecuteScriptParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceBrowserExecuteScriptParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceBrowserExecuteScriptParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
