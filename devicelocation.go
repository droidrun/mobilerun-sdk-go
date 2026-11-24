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
	shimjson "github.com/droidrun/mobilerun-sdk-go/v6/internal/encoding/json"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/shared"
)

// DeviceLocationService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceLocationService] method instead.
type DeviceLocationService struct {
	options []option.RequestOption
}

// NewDeviceLocationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceLocationService(opts ...option.RequestOption) (r DeviceLocationService) {
	r = DeviceLocationService{}
	r.options = opts
	return
}

// Returns the device's current simulated GPS location as latitude and longitude.
// Devices without geo support return an unsupported-feature error.
func (r *DeviceLocationService) Get(ctx context.Context, deviceID string, query DeviceLocationGetParams, opts ...option.RequestOption) (res *shared.Location, err error) {
	if !param.IsOmitted(query.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", query.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/location", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Clears any simulated GPS location and restores the device's default location
// behavior. Devices without geo support return an unsupported-feature error.
func (r *DeviceLocationService) Reset(ctx context.Context, deviceID string, body DeviceLocationResetParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", body.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/location", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Sets the device's simulated GPS location to the latitude and longitude in the
// request body. Devices without geo support return an unsupported-feature error.
func (r *DeviceLocationService) Set(ctx context.Context, deviceID string, params DeviceLocationSetParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/location", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

type DeviceLocationGetParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

type DeviceLocationResetParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

type DeviceLocationSetParams struct {
	Location         shared.LocationParam
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceLocationSetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Location)
}
func (r *DeviceLocationSetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
