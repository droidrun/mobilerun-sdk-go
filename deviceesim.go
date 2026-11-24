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
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiquery"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// DeviceEsimService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceEsimService] method instead.
type DeviceEsimService struct {
	options []option.RequestOption
	Apn     DeviceEsimApnService
}

// NewDeviceEsimService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceEsimService(opts ...option.RequestOption) (r DeviceEsimService) {
	r = DeviceEsimService{}
	r.options = opts
	r.Apn = NewDeviceEsimApnService(opts...)
	return
}

// Returns the eSIM subscriptions currently provisioned on the device.
func (r *DeviceEsimService) List(ctx context.Context, deviceID string, query DeviceEsimListParams, opts ...option.RequestOption) (res *[]DeviceEsimListResponse, err error) {
	if !param.IsOmitted(query.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", query.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/esim", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Download profile and/or enable subscription.
func (r *DeviceEsimService) Activate(ctx context.Context, deviceID string, params DeviceEsimActivateParams, opts ...option.RequestOption) (res *DeviceEsimActivateResponse, err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/esim", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Enables the eSIM subscription identified by the subId in the request body so it
// becomes the active subscription.
func (r *DeviceEsimService) Enable(ctx context.Context, deviceID string, params DeviceEsimEnableParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/esim", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return err
}

// Deletes the eSIM subscription identified by the subId query parameter from the
// device.
func (r *DeviceEsimService) Remove(ctx context.Context, deviceID string, params DeviceEsimRemoveParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/esim", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Enables or disables data roaming for the device's eSIM based on the enabled flag
// in the request body.
func (r *DeviceEsimService) SetRoaming(ctx context.Context, deviceID string, params DeviceEsimSetRoamingParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/esim/roaming", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return err
}

// Returns the connectivity status of the device's eSIM subscriptions.
func (r *DeviceEsimService) Status(ctx context.Context, deviceID string, query DeviceEsimStatusParams, opts ...option.RequestOption) (res *[]DeviceEsimStatusResponse, err error) {
	if !param.IsOmitted(query.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", query.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/esim/status", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type DeviceEsimListResponse struct {
	Carrier     string `json:"carrier" api:"required"`
	DisplayName string `json:"displayName" api:"required"`
	Iccid       string `json:"iccid" api:"required"`
	IsEmbedded  bool   `json:"isEmbedded" api:"required"`
	Slot        int64  `json:"slot" api:"required"`
	SubID       int64  `json:"subId" api:"required"`
	Type        string `json:"type" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier     respjson.Field
		DisplayName respjson.Field
		Iccid       respjson.Field
		IsEmbedded  respjson.Field
		Slot        respjson.Field
		SubID       respjson.Field
		Type        respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceEsimListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceEsimListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceEsimActivateResponse struct {
	Carrier     string `json:"carrier" api:"required"`
	DisplayName string `json:"displayName" api:"required"`
	Iccid       string `json:"iccid" api:"required"`
	IsEmbedded  bool   `json:"isEmbedded" api:"required"`
	Slot        int64  `json:"slot" api:"required"`
	SubID       int64  `json:"subId" api:"required"`
	Type        string `json:"type" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier     respjson.Field
		DisplayName respjson.Field
		Iccid       respjson.Field
		IsEmbedded  respjson.Field
		Slot        respjson.Field
		SubID       respjson.Field
		Type        respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceEsimActivateResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceEsimActivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceEsimStatusResponse struct {
	Carrier            string `json:"carrier" api:"required"`
	DataRoamingEnabled bool   `json:"dataRoamingEnabled" api:"required"`
	DataState          string `json:"dataState" api:"required"`
	IsRoaming          bool   `json:"isRoaming" api:"required"`
	MobileDataEnabled  bool   `json:"mobileDataEnabled" api:"required"`
	NetworkType        string `json:"networkType" api:"required"`
	Operator           string `json:"operator" api:"required"`
	PhoneType          string `json:"phoneType" api:"required"`
	SimState           string `json:"simState" api:"required"`
	SubID              int64  `json:"subId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier            respjson.Field
		DataRoamingEnabled respjson.Field
		DataState          respjson.Field
		IsRoaming          respjson.Field
		MobileDataEnabled  respjson.Field
		NetworkType        respjson.Field
		Operator           respjson.Field
		PhoneType          respjson.Field
		SimState           respjson.Field
		SubID              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceEsimStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceEsimStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceEsimListParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

type DeviceEsimActivateParams struct {
	Enable   bool   `json:"enable" api:"required"`
	SmDpAddr string `json:"smDpAddr" api:"required"`
	// Optional carrier-issued confirmation code (the 4th LPA segment). Required only
	// for plans whose SM-DP+ challenges the device for one. Requires matchingId — the
	// LPA spec only interprets segment 4 when segment 3 is present.
	ConfirmationCode param.Opt[string] `json:"confirmationCode,omitzero"`
	MatchingID       param.Opt[string] `json:"matchingId,omitzero"`
	XDeviceDisplayID param.Opt[int64]  `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceEsimActivateParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceEsimActivateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceEsimActivateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceEsimEnableParams struct {
	SubID            int64            `json:"subId" api:"required"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceEsimEnableParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceEsimEnableParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceEsimEnableParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceEsimRemoveParams struct {
	SubID            int64            `query:"subId" api:"required" json:"-"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceEsimRemoveParams]'s query parameters as `url.Values`.
func (r DeviceEsimRemoveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceEsimSetRoamingParams struct {
	Enabled          bool             `json:"enabled" api:"required"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceEsimSetRoamingParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceEsimSetRoamingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceEsimSetRoamingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceEsimStatusParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}
