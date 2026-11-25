// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// DeviceNumberService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceNumberService] method instead.
type DeviceNumberService struct {
	options []option.RequestOption
}

// NewDeviceNumberService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceNumberService(opts ...option.RequestOption) (r DeviceNumberService) {
	r = DeviceNumberService{}
	r.options = opts
	return
}

// Returns the cloud number bound to the device and the SIM line the device
// presents (carrier, number).
func (r *DeviceNumberService) Get(ctx context.Context, deviceID string, query DeviceNumberGetParams, opts ...option.RequestOption) (res *DeviceNumberGetResponse, err error) {
	if !param.IsOmitted(query.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", query.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/number", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Binds one of your cloud numbers to the device. The device's SIM shows the number
// and every SMS it receives is delivered into the device's messages app. Available
// on devices whose capabilities report telephony. A number can be bound to one
// device, and a device holds one number; unbind first to move it.
func (r *DeviceNumberService) Bind(ctx context.Context, deviceID string, params DeviceNumberBindParams, opts ...option.RequestOption) (res *DeviceNumberBindResponse, err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/number", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Releases the number bound to the device. The device returns to its own default
// line; SMS to the number stop reaching the device.
func (r *DeviceNumberService) Unbind(ctx context.Context, deviceID string, body DeviceNumberUnbindParams, opts ...option.RequestOption) (res *DeviceNumberUnbindResponse, err error) {
	if !param.IsOmitted(body.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", body.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/number", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type Binding struct {
	BoundAt     time.Time `json:"boundAt" api:"required" format:"date-time"`
	NumberID    string    `json:"numberId" api:"required"`
	PhoneNumber string    `json:"phoneNumber" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoundAt     respjson.Field
		NumberID    respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Binding) RawJSON() string { return r.JSON.raw }
func (r *Binding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimLine struct {
	Carrier       SimLineCarrier `json:"carrier" api:"required"`
	DefaultNumber bool           `json:"defaultNumber" api:"required"`
	// Any of "simulated", "uicc", "esim", "unknown".
	Kind   SimLineKind `json:"kind" api:"required"`
	LineID string      `json:"lineId" api:"required"`
	// Any of "backend_default", "caller", "unknown".
	NumberSource SimLineNumberSource `json:"numberSource" api:"required"`
	// Any of "pending", "active", "inactive", "failed", "unknown".
	State       SimLineState `json:"state" api:"required"`
	Iccid       string       `json:"iccid"`
	PhoneNumber string       `json:"phoneNumber"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier       respjson.Field
		DefaultNumber respjson.Field
		Kind          respjson.Field
		LineID        respjson.Field
		NumberSource  respjson.Field
		State         respjson.Field
		Iccid         respjson.Field
		PhoneNumber   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimLine) RawJSON() string { return r.JSON.raw }
func (r *SimLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimLineCarrier struct {
	CountryISO string `json:"countryIso" api:"required"`
	Mcc        string `json:"mcc" api:"required"`
	Mnc        string `json:"mnc" api:"required"`
	Name       string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryISO  respjson.Field
		Mcc         respjson.Field
		Mnc         respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimLineCarrier) RawJSON() string { return r.JSON.raw }
func (r *SimLineCarrier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimLineKind string

const (
	SimLineKindSimulated SimLineKind = "simulated"
	SimLineKindUicc      SimLineKind = "uicc"
	SimLineKindEsim      SimLineKind = "esim"
	SimLineKindUnknown   SimLineKind = "unknown"
)

type SimLineNumberSource string

const (
	SimLineNumberSourceBackendDefault SimLineNumberSource = "backend_default"
	SimLineNumberSourceCaller         SimLineNumberSource = "caller"
	SimLineNumberSourceUnknown        SimLineNumberSource = "unknown"
)

type SimLineState string

const (
	SimLineStatePending  SimLineState = "pending"
	SimLineStateActive   SimLineState = "active"
	SimLineStateInactive SimLineState = "inactive"
	SimLineStateFailed   SimLineState = "failed"
	SimLineStateUnknown  SimLineState = "unknown"
)

type DeviceNumberGetResponse struct {
	Binding Binding `json:"binding" api:"required"`
	Line    SimLine `json:"line" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Binding     respjson.Field
		Line        respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceNumberGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceNumberGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNumberBindResponse struct {
	Binding Binding `json:"binding" api:"required"`
	Line    SimLine `json:"line" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Binding     respjson.Field
		Line        respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceNumberBindResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceNumberBindResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNumberUnbindResponse struct {
	Binding Binding `json:"binding" api:"required"`
	Line    SimLine `json:"line" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Binding     respjson.Field
		Line        respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceNumberUnbindResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceNumberUnbindResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNumberGetParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

type DeviceNumberBindParams struct {
	// The cloud number's id.
	NumberID         string           `json:"numberId" api:"required" format:"uuid"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceNumberBindParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceNumberBindParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceNumberBindParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNumberUnbindParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}
