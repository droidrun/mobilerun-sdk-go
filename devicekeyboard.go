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

// DeviceKeyboardService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceKeyboardService] method instead.
type DeviceKeyboardService struct {
	options []option.RequestOption
}

// NewDeviceKeyboardService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceKeyboardService(opts ...option.RequestOption) (r DeviceKeyboardService) {
	r = DeviceKeyboardService{}
	r.options = opts
	return
}

// Clears the contents of the currently focused text input field.
func (r *DeviceKeyboardService) Clear(ctx context.Context, deviceID string, body DeviceKeyboardClearParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", body.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/keyboard", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Sends a single Android key event to the device, identified by its key code.
func (r *DeviceKeyboardService) Key(ctx context.Context, deviceID string, params DeviceKeyboardKeyParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/keyboard", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return err
}

// Types text into the focused input field. The optional completionMode defaults to
// accepted for backwards-compatible low latency; committed additionally waits for
// the complete text or a quiescent UI state.
func (r *DeviceKeyboardService) Write(ctx context.Context, deviceID string, params DeviceKeyboardWriteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/keyboard", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

type DeviceKeyboardClearParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

type DeviceKeyboardKeyParams struct {
	Key              int64            `json:"key" api:"required"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceKeyboardKeyParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceKeyboardKeyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceKeyboardKeyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceKeyboardWriteParams struct {
	Text  string          `json:"text" api:"required"`
	Clear param.Opt[bool] `json:"clear,omitzero"`
	// Per-character mistake rate for humanized typing. -1 uses server default.
	ErrorRate param.Opt[float64] `json:"errorRate,omitzero"`
	// Request human-like input where supported by the device. Defaults to false.
	Humanize param.Opt[bool] `json:"humanize,omitzero"`
	// Deprecated: use humanize. Both names must agree when supplied together.
	Stealth param.Opt[bool] `json:"stealth,omitzero"`
	// Words per minute for humanized typing. 0 uses the device default.
	Wpm              param.Opt[int64] `json:"wpm,omitzero"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	// Completion guarantee. accepted returns after the input provider accepts the
	// operation; committed additionally waits for the focused UI state to contain the
	// complete text or become quiescent.
	//
	// Any of "accepted", "committed".
	CompletionMode DeviceKeyboardWriteParamsCompletionMode `json:"completionMode,omitzero"`
	paramObj
}

func (r DeviceKeyboardWriteParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceKeyboardWriteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceKeyboardWriteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Completion guarantee. accepted returns after the input provider accepts the
// operation; committed additionally waits for the focused UI state to contain the
// complete text or become quiescent.
type DeviceKeyboardWriteParamsCompletionMode string

const (
	DeviceKeyboardWriteParamsCompletionModeAccepted  DeviceKeyboardWriteParamsCompletionMode = "accepted"
	DeviceKeyboardWriteParamsCompletionModeCommitted DeviceKeyboardWriteParamsCompletionMode = "committed"
)
