// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// DeviceNumberBindingService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceNumberBindingService] method instead.
type DeviceNumberBindingService struct {
	options []option.RequestOption
}

// NewDeviceNumberBindingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeviceNumberBindingService(opts ...option.RequestOption) (r DeviceNumberBindingService) {
	r = DeviceNumberBindingService{}
	r.options = opts
	return
}

func (r *DeviceNumberBindingService) List(ctx context.Context, opts ...option.RequestOption) (res *DeviceNumberBindingListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "devices/number-bindings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *DeviceNumberBindingService) ListCandidates(ctx context.Context, opts ...option.RequestOption) (res *DeviceNumberBindingListCandidatesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "devices/number-bindings/candidates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type DeviceNumberBindingListResponse struct {
	Items []DeviceNumberBindingListResponseItem `json:"items" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceNumberBindingListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceNumberBindingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNumberBindingListResponseItem struct {
	BoundAt    time.Time `json:"boundAt" api:"required" format:"date-time"`
	DeviceID   string    `json:"deviceId" api:"required"`
	DeviceName string    `json:"deviceName" api:"required"`
	NumberID   string    `json:"numberId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoundAt     respjson.Field
		DeviceID    respjson.Field
		DeviceName  respjson.Field
		NumberID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceNumberBindingListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *DeviceNumberBindingListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNumberBindingListCandidatesResponse struct {
	Items []DeviceNumberBindingListCandidatesResponseItem `json:"items" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceNumberBindingListCandidatesResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceNumberBindingListCandidatesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNumberBindingListCandidatesResponseItem struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// Any of "ready".
	State string `json:"state" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceNumberBindingListCandidatesResponseItem) RawJSON() string { return r.JSON.raw }
func (r *DeviceNumberBindingListCandidatesResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
