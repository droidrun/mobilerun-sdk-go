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
	"github.com/droidrun/mobilerun-sdk-go/v6/shared"
)

// ComputeFunctionVersionService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeFunctionVersionService] method instead.
type ComputeFunctionVersionService struct {
	options []option.RequestOption
}

// NewComputeFunctionVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewComputeFunctionVersionService(opts ...option.RequestOption) (r ComputeFunctionVersionService) {
	r = ComputeFunctionVersionService{}
	r.options = opts
	return
}

// Upload a new immutable version
func (r *ComputeFunctionVersionService) New(ctx context.Context, functionID string, body ComputeFunctionVersionNewParams, opts ...option.RequestOption) (res *ComputeFunctionVersionNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/functions/%s/versions", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get an exact version
func (r *ComputeFunctionVersionService) Get(ctx context.Context, version int64, query ComputeFunctionVersionGetParams, opts ...option.RequestOption) (res *ComputeFunctionVersionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.FunctionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/functions/%s/versions/%v", query.FunctionID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List active versions, newest first
func (r *ComputeFunctionVersionService) List(ctx context.Context, functionID string, query ComputeFunctionVersionListParams, opts ...option.RequestOption) (res *ComputeFunctionVersionListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/functions/%s/versions", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a version except the last active version
func (r *ComputeFunctionVersionService) Delete(ctx context.Context, version int64, body ComputeFunctionVersionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.FunctionID == "" {
		err = errors.New("missing required functionId parameter")
		return err
	}
	path := fmt.Sprintf("compute/functions/%s/versions/%v", body.FunctionID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ComputeFunctionVersionNewResponse struct {
	Data FunctionVersion `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionVersionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionGetResponse struct {
	Data FunctionVersion `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionVersionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionListResponse struct {
	Data ComputeFunctionVersionListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionVersionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionListResponseData struct {
	Items      []*FunctionVersion `json:"items" api:"required"`
	Pagination shared.Pagination  `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionVersionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionNewParams struct {
	Source   string                                     `json:"source" api:"required"`
	APIKey   ComputeFunctionVersionNewParamsAPIKeyUnion `json:"apiKey,omitzero"`
	Manifest FunctionManifestParam                      `json:"manifest,omitzero"`
	paramObj
}

func (r ComputeFunctionVersionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputeFunctionVersionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeFunctionVersionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ComputeFunctionVersionNewParamsAPIKeyUnion struct {
	OfComputeFunctionVersionNewsAPIKeyPermissions *ComputeFunctionVersionNewParamsAPIKeyPermissions `json:",omitzero,inline"`
	OfComputeFunctionVersionNewsAPIKeyFull        *ComputeFunctionVersionNewParamsAPIKeyFull        `json:",omitzero,inline"`
	paramUnion
}

func (u ComputeFunctionVersionNewParamsAPIKeyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfComputeFunctionVersionNewsAPIKeyPermissions, u.OfComputeFunctionVersionNewsAPIKeyFull)
}
func (u *ComputeFunctionVersionNewParamsAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Permissions is required.
type ComputeFunctionVersionNewParamsAPIKeyPermissions struct {
	Permissions map[string][]string `json:"permissions,omitzero" api:"required"`
	paramObj
}

func (r ComputeFunctionVersionNewParamsAPIKeyPermissions) MarshalJSON() (data []byte, err error) {
	type shadow ComputeFunctionVersionNewParamsAPIKeyPermissions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeFunctionVersionNewParamsAPIKeyPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Full is required.
type ComputeFunctionVersionNewParamsAPIKeyFull struct {
	// Any of true.
	Full bool `json:"full,omitzero" api:"required"`
	paramObj
}

func (r ComputeFunctionVersionNewParamsAPIKeyFull) MarshalJSON() (data []byte, err error) {
	type shadow ComputeFunctionVersionNewParamsAPIKeyFull
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeFunctionVersionNewParamsAPIKeyFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ComputeFunctionVersionNewParamsAPIKeyFull](
		"full", true,
	)
}

type ComputeFunctionVersionGetParams struct {
	FunctionID string `path:"functionId" api:"required" format:"uuid" json:"-"`
	paramObj
}

type ComputeFunctionVersionListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeFunctionVersionListParams]'s query parameters as
// `url.Values`.
func (r ComputeFunctionVersionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ComputeFunctionVersionDeleteParams struct {
	FunctionID string `path:"functionId" api:"required" format:"uuid" json:"-"`
	paramObj
}
