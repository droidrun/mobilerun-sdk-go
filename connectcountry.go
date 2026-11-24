// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
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

// Mobilerun Connect country coverage information
//
// ConnectCountryService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectCountryService] method instead.
type ConnectCountryService struct {
	options []option.RequestOption
}

// NewConnectCountryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectCountryService(opts ...option.RequestOption) (r ConnectCountryService) {
	r = ConnectCountryService{}
	r.options = opts
	return
}

// Lookup of countries that can be selected when creating a proxy. Each country
// lists the proxy types available there; without a ?type filter, every covered
// country is returned.
func (r *ConnectCountryService) List(ctx context.Context, query ConnectCountryListParams, opts ...option.RequestOption) (res *ConnectCountryListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "connect/countries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A page of countries.
type ConnectCountryListResponse struct {
	Items []ConnectCountryListResponseItem `json:"items" api:"required"`
	// Pagination metadata for a list response.
	Pagination ConnectCountryListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectCountryListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectCountryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectCountryListResponseItem struct {
	// ISO 3166-1 alpha-2 country code (lowercase).
	Code string `json:"code" api:"required"`
	Name string `json:"name" api:"required"`
	// Proxy types available to provision in this country.
	//
	// Any of "dedicated_residential", "residential", "mobile".
	ProxyTypes []string `json:"proxyTypes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		ProxyTypes  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectCountryListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ConnectCountryListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata for a list response.
type ConnectCountryListResponsePagination struct {
	// Whether a next page exists.
	HasNext bool `json:"hasNext" api:"required"`
	// Whether a previous page exists.
	HasPrev bool `json:"hasPrev" api:"required"`
	// Current page number (1-based).
	Page int64 `json:"page" api:"required"`
	// Total number of pages.
	Pages int64 `json:"pages" api:"required"`
	// Number of items per page.
	PageSize int64 `json:"pageSize" api:"required"`
	// Total number of items across all pages.
	Total int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNext     respjson.Field
		HasPrev     respjson.Field
		Page        respjson.Field
		Pages       respjson.Field
		PageSize    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectCountryListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ConnectCountryListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectCountryListParams struct {
	// Page number (1-based).
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of items per page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Filter to countries offering this proxy type.
	//
	// Any of "dedicated_residential", "residential", "mobile".
	Type ConnectCountryListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConnectCountryListParams]'s query parameters as
// `url.Values`.
func (r ConnectCountryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter to countries offering this proxy type.
type ConnectCountryListParamsType string

const (
	ConnectCountryListParamsTypeDedicatedResidential ConnectCountryListParamsType = "dedicated_residential"
	ConnectCountryListParamsTypeResidential          ConnectCountryListParamsType = "residential"
	ConnectCountryListParamsTypeMobile               ConnectCountryListParamsType = "mobile"
)
