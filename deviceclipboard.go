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

// DeviceClipboardService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceClipboardService] method instead.
type DeviceClipboardService struct {
	options []option.RequestOption
}

// NewDeviceClipboardService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceClipboardService(opts ...option.RequestOption) (r DeviceClipboardService) {
	r = DeviceClipboardService{}
	r.options = opts
	return
}

// Returns the current text content of the device's clipboard. Devices without
// clipboard support return an unsupported-feature error.
func (r *DeviceClipboardService) Get(ctx context.Context, deviceID string, query DeviceClipboardGetParams, opts ...option.RequestOption) (res *DeviceClipboardGetResponse, err error) {
	if !param.IsOmitted(query.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", query.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/clipboard", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Replaces the device's clipboard content with the text in the request body; an
// empty text clears the clipboard. Devices without clipboard support return an
// unsupported-feature error.
func (r *DeviceClipboardService) Set(ctx context.Context, deviceID string, params DeviceClipboardSetParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/clipboard", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

type DeviceClipboardGetResponse struct {
	Text string `json:"text" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceClipboardGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceClipboardGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceClipboardGetParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

type DeviceClipboardSetParams struct {
	Text             string           `json:"text" api:"required"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceClipboardSetParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceClipboardSetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceClipboardSetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
