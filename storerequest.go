// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"net/http"
	"slices"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// StoreRequestService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoreRequestService] method instead.
type StoreRequestService struct {
	options []option.RequestOption
}

// NewStoreRequestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStoreRequestService(opts ...option.RequestOption) (r StoreRequestService) {
	r = StoreRequestService{}
	r.options = opts
	return
}

// Request a missing store app
func (r *StoreRequestService) New(ctx context.Context, body StoreRequestNewParams, opts ...option.RequestOption) (res *StoreRequestNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "store/requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type StoreRequestNewResponse struct {
	// Any of "ticket", "event".
	Delivered    StoreRequestNewResponseDelivered `json:"delivered" api:"required"`
	TicketNumber int64                            `json:"ticketNumber" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delivered    respjson.Field
		TicketNumber respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StoreRequestNewResponse) RawJSON() string { return r.JSON.raw }
func (r *StoreRequestNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoreRequestNewResponseDelivered string

const (
	StoreRequestNewResponseDeliveredTicket StoreRequestNewResponseDelivered = "ticket"
	StoreRequestNewResponseDeliveredEvent  StoreRequestNewResponseDelivered = "event"
)

type StoreRequestNewParams struct {
	// Any of "android", "ios".
	Platform  StoreRequestNewParamsPlatform `json:"platform,omitzero" api:"required"`
	Requested string                        `json:"requested" api:"required"`
	// Any of "store", "mcp", "api".
	Surface StoreRequestNewParamsSurface `json:"surface,omitzero"`
	paramObj
}

func (r StoreRequestNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StoreRequestNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StoreRequestNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoreRequestNewParamsPlatform string

const (
	StoreRequestNewParamsPlatformAndroid StoreRequestNewParamsPlatform = "android"
	StoreRequestNewParamsPlatformIos     StoreRequestNewParamsPlatform = "ios"
)

type StoreRequestNewParamsSurface string

const (
	StoreRequestNewParamsSurfaceStore StoreRequestNewParamsSurface = "store"
	StoreRequestNewParamsSurfaceMcp   StoreRequestNewParamsSurface = "mcp"
	StoreRequestNewParamsSurfaceAPI   StoreRequestNewParamsSurface = "api"
)
