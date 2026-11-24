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
)

// DeviceDeepLinkService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceDeepLinkService] method instead.
type DeviceDeepLinkService struct {
	options []option.RequestOption
}

// NewDeviceDeepLinkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceDeepLinkService(opts ...option.RequestOption) (r DeviceDeepLinkService) {
	r = DeviceDeepLinkService{}
	r.options = opts
	return
}

// Opens a deep link on the device. On Android the link is dispatched as an intent
// — packageName optionally pins it to a specific app and action overrides the
// default android.intent.action.VIEW. On iOS the URL is opened directly and the
// optional fields must be omitted. Protected packages are rejected.
func (r *DeviceDeepLinkService) ExecuteDeepLink(ctx context.Context, deviceID string, params DeviceDeepLinkExecuteDeepLinkParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/apps/open-deep-link", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

type DeviceDeepLinkExecuteDeepLinkParams struct {
	// Deep link to open (e.g. myapp://path or https://example.com/path)
	DeepLink string `json:"deepLink" api:"required"`
	// Android only: intent action to dispatch. Defaults to android.intent.action.VIEW.
	Action param.Opt[string] `json:"action,omitzero"`
	// Reserved for targeting a specific iOS app; currently rejected as unsupported.
	BundleID param.Opt[string] `json:"bundleId,omitzero"`
	// Android only: package to receive the intent (e.g. com.example.app). Omit to let
	// the system pick the handler.
	PackageName      param.Opt[string] `json:"packageName,omitzero"`
	XDeviceDisplayID param.Opt[int64]  `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceDeepLinkExecuteDeepLinkParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceDeepLinkExecuteDeepLinkParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceDeepLinkExecuteDeepLinkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
