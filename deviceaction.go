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

// DeviceActionService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceActionService] method instead.
type DeviceActionService struct {
	options []option.RequestOption
}

// NewDeviceActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceActionService(opts ...option.RequestOption) (r DeviceActionService) {
	r = DeviceActionService{}
	r.options = opts
	return
}

// Performs a global system action on the device, such as navigating back or going
// to the home screen, identified by an action code.
func (r *DeviceActionService) Global(ctx context.Context, deviceID string, params DeviceActionGlobalParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/global", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Returns whether the accessibility overlay is currently visible on the device.
func (r *DeviceActionService) OverlayVisible(ctx context.Context, deviceID string, query DeviceActionOverlayVisibleParams, opts ...option.RequestOption) (res *DeviceActionOverlayVisibleResponse, err error) {
	if !param.IsOmitted(query.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", query.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/overlay", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Shows or hides the accessibility overlay on the device based on the visibility
// flag in the request body.
func (r *DeviceActionService) SetOverlayVisible(ctx context.Context, deviceID string, params DeviceActionSetOverlayVisibleParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/overlay", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Swipes from a start coordinate to an end coordinate over the given duration in
// milliseconds. The optional humanize flag (best-effort, also with holdMs) adds a
// human-like path and timing. With holdMs > 0 the request returns 422 when
// duration + holdMs exceeds 9000 ms, or on devices that do not support it.
func (r *DeviceActionService) Swipe(ctx context.Context, deviceID string, params DeviceActionSwipeParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/swipe", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Taps the device screen at the given x/y coordinates. An optional humanize flag
// routes the tap through human-like input on devices that support it.
func (r *DeviceActionService) Tap(ctx context.Context, deviceID string, params DeviceActionTapParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/tap", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

type DeviceActionOverlayVisibleResponse struct {
	Visible bool `json:"visible" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Visible     respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceActionOverlayVisibleResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceActionOverlayVisibleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceActionGlobalParams struct {
	Action           int64            `json:"action" api:"required"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceActionGlobalParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceActionGlobalParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceActionGlobalParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceActionOverlayVisibleParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

type DeviceActionSetOverlayVisibleParams struct {
	Visible          bool             `json:"visible" api:"required"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceActionSetOverlayVisibleParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceActionSetOverlayVisibleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceActionSetOverlayVisibleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceActionSwipeParams struct {
	// Swipe duration in milliseconds
	Duration int64 `json:"duration" api:"required"`
	EndX     int64 `json:"endX" api:"required"`
	EndY     int64 `json:"endY" api:"required"`
	StartX   int64 `json:"startX" api:"required"`
	StartY   int64 `json:"startY" api:"required"`
	// Keep the finger at the end point for holdMs and suppress fling before lifting,
	// so content moves by about the finger travel. Supported on Android cloud phones.
	HoldMs param.Opt[int64] `json:"holdMs,omitzero"`
	// Request human-like input where supported by the device. Defaults to false.
	Humanize param.Opt[bool] `json:"humanize,omitzero"`
	// Deprecated: use humanize. Both names must agree when supplied together.
	Stealth          param.Opt[bool]  `json:"stealth,omitzero"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceActionSwipeParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceActionSwipeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceActionSwipeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceActionTapParams struct {
	X int64 `json:"x" api:"required"`
	Y int64 `json:"y" api:"required"`
	// Request human-like input where supported by the device. Defaults to false.
	Humanize param.Opt[bool] `json:"humanize,omitzero"`
	// Deprecated: use humanize. Both names must agree when supplied together.
	Stealth          param.Opt[bool]  `json:"stealth,omitzero"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceActionTapParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceActionTapParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceActionTapParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
