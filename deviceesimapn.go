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

// DeviceEsimApnService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceEsimApnService] method instead.
type DeviceEsimApnService struct {
	options []option.RequestOption
}

// NewDeviceEsimApnService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceEsimApnService(opts ...option.RequestOption) (r DeviceEsimApnService) {
	r = DeviceEsimApnService{}
	r.options = opts
	return
}

// Returns the access point names (APNs) configured for the device's active eSIM
// subscriptions.
func (r *DeviceEsimApnService) List(ctx context.Context, deviceID string, query DeviceEsimApnListParams, opts ...option.RequestOption) (res *[]DeviceEsimApnListResponse, err error) {
	if !param.IsOmitted(query.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", query.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/esim/apn", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Marks an existing APN, identified by apnId, as the preferred APN for the given
// eSIM subscription in the request body.
func (r *DeviceEsimApnService) Select(ctx context.Context, deviceID string, params DeviceEsimApnSelectParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/esim/apn", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return err
}

// Creates an access point name (APN) from the request body and applies it to the
// given eSIM subscription. Type, protocol, and roaming protocol default to common
// values when omitted.
func (r *DeviceEsimApnService) Set(ctx context.Context, deviceID string, params DeviceEsimApnSetParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/esim/apn", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

type DeviceEsimApnListResponse struct {
	ID              int64  `json:"id" api:"required"`
	Apn             string `json:"apn" api:"required"`
	IsPreferred     bool   `json:"isPreferred" api:"required"`
	Mcc             string `json:"mcc" api:"required"`
	Mnc             string `json:"mnc" api:"required"`
	Name            string `json:"name" api:"required"`
	Protocol        string `json:"protocol" api:"required"`
	RoamingProtocol string `json:"roamingProtocol" api:"required"`
	SubID           int64  `json:"subId" api:"required"`
	Type            string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Apn             respjson.Field
		IsPreferred     respjson.Field
		Mcc             respjson.Field
		Mnc             respjson.Field
		Name            respjson.Field
		Protocol        respjson.Field
		RoamingProtocol respjson.Field
		SubID           respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceEsimApnListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceEsimApnListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceEsimApnListParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

type DeviceEsimApnSelectParams struct {
	ApnID            int64            `json:"apnId" api:"required"`
	SubID            int64            `json:"subId" api:"required"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceEsimApnSelectParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceEsimApnSelectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceEsimApnSelectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceEsimApnSetParams struct {
	Apn              string           `json:"apn" api:"required"`
	Mcc              string           `json:"mcc" api:"required"`
	Mnc              string           `json:"mnc" api:"required"`
	Name             string           `json:"name" api:"required"`
	Protocol         string           `json:"protocol" api:"required"`
	RoamingProtocol  string           `json:"roamingProtocol" api:"required"`
	SubID            int64            `json:"subId" api:"required"`
	Type             string           `json:"type" api:"required"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceEsimApnSetParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceEsimApnSetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceEsimApnSetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
