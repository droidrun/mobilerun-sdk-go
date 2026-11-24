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

// DeviceKioskService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceKioskService] method instead.
type DeviceKioskService struct {
	options []option.RequestOption
}

// NewDeviceKioskService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceKioskService(opts ...option.RequestOption) (r DeviceKioskService) {
	r = DeviceKioskService{}
	r.options = opts
	return
}

// Disables Android lock-task (kiosk) mode on the device, releasing it from the
// locked app.
func (r *DeviceKioskService) Disable(ctx context.Context, deviceID string, body DeviceKioskDisableParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", body.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/kiosk", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Locks the device to the package named in the request body using Android
// lock-task (kiosk) mode, preventing the user from leaving the app.
func (r *DeviceKioskService) Enable(ctx context.Context, deviceID string, params DeviceKioskEnableParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/kiosk", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return err
}

type DeviceKioskDisableParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

type DeviceKioskEnableParams struct {
	// Package to lock the device to (Android lock-task mode).
	PackageName      string           `json:"packageName" api:"required"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceKioskEnableParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceKioskEnableParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceKioskEnableParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
