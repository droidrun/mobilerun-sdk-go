// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiquery"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
	"github.com/droidrun/mobilerun-sdk-go/v6/shared"
)

// ComputeFunctionService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeFunctionService] method instead.
type ComputeFunctionService struct {
	options     []option.RequestOption
	Validations ComputeFunctionValidationService
	Versions    ComputeFunctionVersionService
	Runs        ComputeFunctionRunService
}

// NewComputeFunctionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputeFunctionService(opts ...option.RequestOption) (r ComputeFunctionService) {
	r = ComputeFunctionService{}
	r.options = opts
	r.Validations = NewComputeFunctionValidationService(opts...)
	r.Versions = NewComputeFunctionVersionService(opts...)
	r.Runs = NewComputeFunctionRunService(opts...)
	return
}

// Create a function with its first immutable version
func (r *ComputeFunctionService) New(ctx context.Context, body ComputeFunctionNewParams, opts ...option.RequestOption) (res *ComputeFunctionNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "compute/functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get function metadata
func (r *ComputeFunctionService) Get(ctx context.Context, functionID string, opts ...option.RequestOption) (res *ComputeFunctionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/functions/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Rename a function or update its description
func (r *ComputeFunctionService) Update(ctx context.Context, functionID string, body ComputeFunctionUpdateParams, opts ...option.RequestOption) (res *ComputeFunctionUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/functions/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List functions owned by this organisation
func (r *ComputeFunctionService) List(ctx context.Context, query ComputeFunctionListParams, opts ...option.RequestOption) (res *ComputeFunctionListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "compute/functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a function
func (r *ComputeFunctionService) Delete(ctx context.Context, functionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return err
	}
	path := fmt.Sprintf("compute/functions/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Read verified function source
func (r *ComputeFunctionService) Source(ctx context.Context, functionID string, query ComputeFunctionSourceParams, opts ...option.RequestOption) (res *ComputeFunctionSourceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/functions/%s/source", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type FunctionSource struct {
	FunctionID string `json:"functionId" api:"required" format:"uuid"`
	Sha256     string `json:"sha256" api:"required"`
	SizeBytes  int64  `json:"sizeBytes" api:"required"`
	Source     string `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FunctionID  respjson.Field
		Sha256      respjson.Field
		SizeBytes   respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionSource) RawJSON() string { return r.JSON.raw }
func (r *FunctionSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServerlessFunction struct {
	ID                 string                          `json:"id" api:"required" format:"uuid"`
	CreatedAt          time.Time                       `json:"createdAt" api:"required" format:"date-time"`
	Description        string                          `json:"description" api:"required"`
	LatestReadyVersion int64                           `json:"latestReadyVersion" api:"required"`
	LatestVersion      ServerlessFunctionLatestVersion `json:"latestVersion" api:"required"`
	Name               string                          `json:"name" api:"required"`
	UpdatedAt          time.Time                       `json:"updatedAt" api:"required" format:"date-time"`
	VersionCount       int64                           `json:"versionCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		LatestReadyVersion respjson.Field
		LatestVersion      respjson.Field
		Name               respjson.Field
		UpdatedAt          respjson.Field
		VersionCount       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerlessFunction) RawJSON() string { return r.JSON.raw }
func (r *ServerlessFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServerlessFunctionLatestVersion struct {
	APIKey         ServerlessFunctionLatestVersionAPIKeyUnion  `json:"apiKey" api:"required"`
	CreatedAt      time.Time                                   `json:"createdAt" api:"required" format:"date-time"`
	FunctionID     string                                      `json:"functionId" api:"required" format:"uuid"`
	InputSchema    map[string]any                              `json:"inputSchema" api:"required"`
	InspectError   ServerlessFunctionLatestVersionInspectError `json:"inspectError" api:"required"`
	Manifest       ServerlessFunctionLatestVersionManifest     `json:"manifest" api:"required"`
	OutputKind     string                                      `json:"outputKind" api:"required"`
	OutputSchema   map[string]any                              `json:"outputSchema" api:"required"`
	RuntimeVersion string                                      `json:"runtimeVersion" api:"required"`
	Sha256         string                                      `json:"sha256" api:"required"`
	SizeBytes      int64                                       `json:"sizeBytes" api:"required"`
	// Any of "pending_inspection", "ready", "rejected".
	Status  string `json:"status" api:"required"`
	Version int64  `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey         respjson.Field
		CreatedAt      respjson.Field
		FunctionID     respjson.Field
		InputSchema    respjson.Field
		InspectError   respjson.Field
		Manifest       respjson.Field
		OutputKind     respjson.Field
		OutputSchema   respjson.Field
		RuntimeVersion respjson.Field
		Sha256         respjson.Field
		SizeBytes      respjson.Field
		Status         respjson.Field
		Version        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerlessFunctionLatestVersion) RawJSON() string { return r.JSON.raw }
func (r *ServerlessFunctionLatestVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ServerlessFunctionLatestVersionAPIKeyUnion contains all possible properties and
// values from [ServerlessFunctionLatestVersionAPIKeyPermissions],
// [ServerlessFunctionLatestVersionAPIKeyFull].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ServerlessFunctionLatestVersionAPIKeyUnion struct {
	// This field is from variant [ServerlessFunctionLatestVersionAPIKeyPermissions].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant [ServerlessFunctionLatestVersionAPIKeyFull].
	Full bool `json:"full"`
	JSON struct {
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u ServerlessFunctionLatestVersionAPIKeyUnion) AsServerlessFunctionLatestVersionAPIKeyPermissions() (v ServerlessFunctionLatestVersionAPIKeyPermissions) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ServerlessFunctionLatestVersionAPIKeyUnion) AsServerlessFunctionLatestVersionAPIKeyFull() (v ServerlessFunctionLatestVersionAPIKeyFull) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ServerlessFunctionLatestVersionAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ServerlessFunctionLatestVersionAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServerlessFunctionLatestVersionAPIKeyPermissions struct {
	Permissions map[string][]string `json:"permissions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerlessFunctionLatestVersionAPIKeyPermissions) RawJSON() string { return r.JSON.raw }
func (r *ServerlessFunctionLatestVersionAPIKeyPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServerlessFunctionLatestVersionAPIKeyFull struct {
	// Any of true.
	Full bool `json:"full" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Full        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerlessFunctionLatestVersionAPIKeyFull) RawJSON() string { return r.JSON.raw }
func (r *ServerlessFunctionLatestVersionAPIKeyFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServerlessFunctionLatestVersionInspectError struct {
	Code    string                                             `json:"code" api:"required"`
	Message string                                             `json:"message" api:"required"`
	Errors  []ServerlessFunctionLatestVersionInspectErrorError `json:"errors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Errors      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerlessFunctionLatestVersionInspectError) RawJSON() string { return r.JSON.raw }
func (r *ServerlessFunctionLatestVersionInspectError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServerlessFunctionLatestVersionInspectErrorError struct {
	Code    string                                                 `json:"code" api:"required"`
	Message string                                                 `json:"message" api:"required"`
	Line    int64                                                  `json:"line"`
	Repair  ServerlessFunctionLatestVersionInspectErrorErrorRepair `json:"repair"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Line        respjson.Field
		Repair      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerlessFunctionLatestVersionInspectErrorError) RawJSON() string { return r.JSON.raw }
func (r *ServerlessFunctionLatestVersionInspectErrorError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServerlessFunctionLatestVersionInspectErrorErrorRepair struct {
	Action string `json:"action" api:"required"`
	Hint   string `json:"hint" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Hint        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerlessFunctionLatestVersionInspectErrorErrorRepair) RawJSON() string { return r.JSON.raw }
func (r *ServerlessFunctionLatestVersionInspectErrorErrorRepair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServerlessFunctionLatestVersionManifest struct {
	Credentials []ServerlessFunctionLatestVersionManifestCredential `json:"credentials"`
	Hosts       []string                                            `json:"hosts"`
	Scopes      []string                                            `json:"scopes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Credentials respjson.Field
		Hosts       respjson.Field
		Scopes      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerlessFunctionLatestVersionManifest) RawJSON() string { return r.JSON.raw }
func (r *ServerlessFunctionLatestVersionManifest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServerlessFunctionLatestVersionManifestCredential struct {
	Name        string `json:"name" api:"required"`
	PackageName string `json:"packageName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		PackageName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerlessFunctionLatestVersionManifestCredential) RawJSON() string { return r.JSON.raw }
func (r *ServerlessFunctionLatestVersionManifestCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionNewResponse struct {
	Data ServerlessFunction `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionGetResponse struct {
	Data ServerlessFunction `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionUpdateResponse struct {
	Data ServerlessFunction `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionListResponse struct {
	Data ComputeFunctionListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionListResponseData struct {
	Items      []ServerlessFunction `json:"items" api:"required"`
	Pagination shared.Pagination    `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionSourceResponse struct {
	Data FunctionSource `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionSourceResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionSourceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionNewParams struct {
	Name        string                              `json:"name" api:"required"`
	Source      string                              `json:"source" api:"required"`
	Description param.Opt[string]                   `json:"description,omitzero"`
	APIKey      ComputeFunctionNewParamsAPIKeyUnion `json:"apiKey,omitzero"`
	Manifest    ComputeFunctionNewParamsManifest    `json:"manifest,omitzero"`
	paramObj
}

func (r ComputeFunctionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputeFunctionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeFunctionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ComputeFunctionNewParamsAPIKeyUnion struct {
	OfComputeFunctionNewsAPIKeyPermissions *ComputeFunctionNewParamsAPIKeyPermissions `json:",omitzero,inline"`
	OfComputeFunctionNewsAPIKeyFull        *ComputeFunctionNewParamsAPIKeyFull        `json:",omitzero,inline"`
	paramUnion
}

func (u ComputeFunctionNewParamsAPIKeyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfComputeFunctionNewsAPIKeyPermissions, u.OfComputeFunctionNewsAPIKeyFull)
}
func (u *ComputeFunctionNewParamsAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Permissions is required.
type ComputeFunctionNewParamsAPIKeyPermissions struct {
	Permissions map[string][]string `json:"permissions,omitzero" api:"required"`
	paramObj
}

func (r ComputeFunctionNewParamsAPIKeyPermissions) MarshalJSON() (data []byte, err error) {
	type shadow ComputeFunctionNewParamsAPIKeyPermissions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeFunctionNewParamsAPIKeyPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Full is required.
type ComputeFunctionNewParamsAPIKeyFull struct {
	// Any of true.
	Full bool `json:"full,omitzero" api:"required"`
	paramObj
}

func (r ComputeFunctionNewParamsAPIKeyFull) MarshalJSON() (data []byte, err error) {
	type shadow ComputeFunctionNewParamsAPIKeyFull
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeFunctionNewParamsAPIKeyFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ComputeFunctionNewParamsAPIKeyFull](
		"full", true,
	)
}

type ComputeFunctionNewParamsManifest struct {
	Credentials []ComputeFunctionNewParamsManifestCredential `json:"credentials,omitzero"`
	Hosts       []string                                     `json:"hosts,omitzero"`
	Scopes      []string                                     `json:"scopes,omitzero"`
	paramObj
}

func (r ComputeFunctionNewParamsManifest) MarshalJSON() (data []byte, err error) {
	type shadow ComputeFunctionNewParamsManifest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeFunctionNewParamsManifest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, PackageName are required.
type ComputeFunctionNewParamsManifestCredential struct {
	Name        string `json:"name" api:"required"`
	PackageName string `json:"packageName" api:"required"`
	paramObj
}

func (r ComputeFunctionNewParamsManifestCredential) MarshalJSON() (data []byte, err error) {
	type shadow ComputeFunctionNewParamsManifestCredential
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeFunctionNewParamsManifestCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ComputeFunctionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputeFunctionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeFunctionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Case-insensitive substring match against the function name.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeFunctionListParams]'s query parameters as
// `url.Values`.
func (r ComputeFunctionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ComputeFunctionSourceParams struct {
	Version param.Opt[int64] `query:"version,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeFunctionSourceParams]'s query parameters as
// `url.Values`.
func (r ComputeFunctionSourceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
