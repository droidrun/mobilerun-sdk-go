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
	Data ComputeFunctionVersionNewResponseData `json:"data" api:"required"`
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

type ComputeFunctionVersionNewResponseData struct {
	APIKey         ComputeFunctionVersionNewResponseDataAPIKeyUnion  `json:"apiKey" api:"required"`
	CreatedAt      time.Time                                         `json:"createdAt" api:"required" format:"date-time"`
	FunctionID     string                                            `json:"functionId" api:"required" format:"uuid"`
	InputSchema    map[string]any                                    `json:"inputSchema" api:"required"`
	InspectError   ComputeFunctionVersionNewResponseDataInspectError `json:"inspectError" api:"required"`
	Manifest       ComputeFunctionVersionNewResponseDataManifest     `json:"manifest" api:"required"`
	OutputKind     string                                            `json:"outputKind" api:"required"`
	OutputSchema   map[string]any                                    `json:"outputSchema" api:"required"`
	RuntimeVersion string                                            `json:"runtimeVersion" api:"required"`
	Sha256         string                                            `json:"sha256" api:"required"`
	SizeBytes      int64                                             `json:"sizeBytes" api:"required"`
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
func (r ComputeFunctionVersionNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeFunctionVersionNewResponseDataAPIKeyUnion contains all possible
// properties and values from
// [ComputeFunctionVersionNewResponseDataAPIKeyPermissions],
// [ComputeFunctionVersionNewResponseDataAPIKeyFull].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ComputeFunctionVersionNewResponseDataAPIKeyUnion struct {
	// This field is from variant
	// [ComputeFunctionVersionNewResponseDataAPIKeyPermissions].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant [ComputeFunctionVersionNewResponseDataAPIKeyFull].
	Full bool `json:"full"`
	JSON struct {
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u ComputeFunctionVersionNewResponseDataAPIKeyUnion) AsComputeFunctionVersionNewResponseDataAPIKeyPermissions() (v ComputeFunctionVersionNewResponseDataAPIKeyPermissions) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeFunctionVersionNewResponseDataAPIKeyUnion) AsComputeFunctionVersionNewResponseDataAPIKeyFull() (v ComputeFunctionVersionNewResponseDataAPIKeyFull) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeFunctionVersionNewResponseDataAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeFunctionVersionNewResponseDataAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionNewResponseDataAPIKeyPermissions struct {
	Permissions map[string][]string `json:"permissions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionVersionNewResponseDataAPIKeyPermissions) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionNewResponseDataAPIKeyPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionNewResponseDataAPIKeyFull struct {
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
func (r ComputeFunctionVersionNewResponseDataAPIKeyFull) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionNewResponseDataAPIKeyFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionNewResponseDataInspectError struct {
	Code    string                                                   `json:"code" api:"required"`
	Message string                                                   `json:"message" api:"required"`
	Errors  []ComputeFunctionVersionNewResponseDataInspectErrorError `json:"errors"`
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
func (r ComputeFunctionVersionNewResponseDataInspectError) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionNewResponseDataInspectError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionNewResponseDataInspectErrorError struct {
	Code    string                                                       `json:"code" api:"required"`
	Message string                                                       `json:"message" api:"required"`
	Line    int64                                                        `json:"line"`
	Repair  ComputeFunctionVersionNewResponseDataInspectErrorErrorRepair `json:"repair"`
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
func (r ComputeFunctionVersionNewResponseDataInspectErrorError) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionNewResponseDataInspectErrorError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionNewResponseDataInspectErrorErrorRepair struct {
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
func (r ComputeFunctionVersionNewResponseDataInspectErrorErrorRepair) RawJSON() string {
	return r.JSON.raw
}
func (r *ComputeFunctionVersionNewResponseDataInspectErrorErrorRepair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionNewResponseDataManifest struct {
	Credentials []ComputeFunctionVersionNewResponseDataManifestCredential `json:"credentials"`
	Hosts       []string                                                  `json:"hosts"`
	Scopes      []string                                                  `json:"scopes"`
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
func (r ComputeFunctionVersionNewResponseDataManifest) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionNewResponseDataManifest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionNewResponseDataManifestCredential struct {
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
func (r ComputeFunctionVersionNewResponseDataManifestCredential) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionNewResponseDataManifestCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionGetResponse struct {
	Data ComputeFunctionVersionGetResponseData `json:"data" api:"required"`
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

type ComputeFunctionVersionGetResponseData struct {
	APIKey         ComputeFunctionVersionGetResponseDataAPIKeyUnion  `json:"apiKey" api:"required"`
	CreatedAt      time.Time                                         `json:"createdAt" api:"required" format:"date-time"`
	FunctionID     string                                            `json:"functionId" api:"required" format:"uuid"`
	InputSchema    map[string]any                                    `json:"inputSchema" api:"required"`
	InspectError   ComputeFunctionVersionGetResponseDataInspectError `json:"inspectError" api:"required"`
	Manifest       ComputeFunctionVersionGetResponseDataManifest     `json:"manifest" api:"required"`
	OutputKind     string                                            `json:"outputKind" api:"required"`
	OutputSchema   map[string]any                                    `json:"outputSchema" api:"required"`
	RuntimeVersion string                                            `json:"runtimeVersion" api:"required"`
	Sha256         string                                            `json:"sha256" api:"required"`
	SizeBytes      int64                                             `json:"sizeBytes" api:"required"`
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
func (r ComputeFunctionVersionGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeFunctionVersionGetResponseDataAPIKeyUnion contains all possible
// properties and values from
// [ComputeFunctionVersionGetResponseDataAPIKeyPermissions],
// [ComputeFunctionVersionGetResponseDataAPIKeyFull].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ComputeFunctionVersionGetResponseDataAPIKeyUnion struct {
	// This field is from variant
	// [ComputeFunctionVersionGetResponseDataAPIKeyPermissions].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant [ComputeFunctionVersionGetResponseDataAPIKeyFull].
	Full bool `json:"full"`
	JSON struct {
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u ComputeFunctionVersionGetResponseDataAPIKeyUnion) AsComputeFunctionVersionGetResponseDataAPIKeyPermissions() (v ComputeFunctionVersionGetResponseDataAPIKeyPermissions) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeFunctionVersionGetResponseDataAPIKeyUnion) AsComputeFunctionVersionGetResponseDataAPIKeyFull() (v ComputeFunctionVersionGetResponseDataAPIKeyFull) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeFunctionVersionGetResponseDataAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeFunctionVersionGetResponseDataAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionGetResponseDataAPIKeyPermissions struct {
	Permissions map[string][]string `json:"permissions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionVersionGetResponseDataAPIKeyPermissions) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionGetResponseDataAPIKeyPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionGetResponseDataAPIKeyFull struct {
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
func (r ComputeFunctionVersionGetResponseDataAPIKeyFull) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionGetResponseDataAPIKeyFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionGetResponseDataInspectError struct {
	Code    string                                                   `json:"code" api:"required"`
	Message string                                                   `json:"message" api:"required"`
	Errors  []ComputeFunctionVersionGetResponseDataInspectErrorError `json:"errors"`
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
func (r ComputeFunctionVersionGetResponseDataInspectError) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionGetResponseDataInspectError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionGetResponseDataInspectErrorError struct {
	Code    string                                                       `json:"code" api:"required"`
	Message string                                                       `json:"message" api:"required"`
	Line    int64                                                        `json:"line"`
	Repair  ComputeFunctionVersionGetResponseDataInspectErrorErrorRepair `json:"repair"`
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
func (r ComputeFunctionVersionGetResponseDataInspectErrorError) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionGetResponseDataInspectErrorError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionGetResponseDataInspectErrorErrorRepair struct {
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
func (r ComputeFunctionVersionGetResponseDataInspectErrorErrorRepair) RawJSON() string {
	return r.JSON.raw
}
func (r *ComputeFunctionVersionGetResponseDataInspectErrorErrorRepair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionGetResponseDataManifest struct {
	Credentials []ComputeFunctionVersionGetResponseDataManifestCredential `json:"credentials"`
	Hosts       []string                                                  `json:"hosts"`
	Scopes      []string                                                  `json:"scopes"`
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
func (r ComputeFunctionVersionGetResponseDataManifest) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionGetResponseDataManifest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionGetResponseDataManifestCredential struct {
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
func (r ComputeFunctionVersionGetResponseDataManifestCredential) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionGetResponseDataManifestCredential) UnmarshalJSON(data []byte) error {
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
	Items      []*ComputeFunctionVersionListResponseDataItem `json:"items" api:"required"`
	Pagination shared.Pagination                             `json:"pagination" api:"required"`
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

type ComputeFunctionVersionListResponseDataItem struct {
	APIKey         ComputeFunctionVersionListResponseDataItemAPIKeyUnion  `json:"apiKey" api:"required"`
	CreatedAt      time.Time                                              `json:"createdAt" api:"required" format:"date-time"`
	FunctionID     string                                                 `json:"functionId" api:"required" format:"uuid"`
	InputSchema    map[string]any                                         `json:"inputSchema" api:"required"`
	InspectError   ComputeFunctionVersionListResponseDataItemInspectError `json:"inspectError" api:"required"`
	Manifest       ComputeFunctionVersionListResponseDataItemManifest     `json:"manifest" api:"required"`
	OutputKind     string                                                 `json:"outputKind" api:"required"`
	OutputSchema   map[string]any                                         `json:"outputSchema" api:"required"`
	RuntimeVersion string                                                 `json:"runtimeVersion" api:"required"`
	Sha256         string                                                 `json:"sha256" api:"required"`
	SizeBytes      int64                                                  `json:"sizeBytes" api:"required"`
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
func (r ComputeFunctionVersionListResponseDataItem) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionListResponseDataItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeFunctionVersionListResponseDataItemAPIKeyUnion contains all possible
// properties and values from
// [ComputeFunctionVersionListResponseDataItemAPIKeyPermissions],
// [ComputeFunctionVersionListResponseDataItemAPIKeyFull].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ComputeFunctionVersionListResponseDataItemAPIKeyUnion struct {
	// This field is from variant
	// [ComputeFunctionVersionListResponseDataItemAPIKeyPermissions].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant
	// [ComputeFunctionVersionListResponseDataItemAPIKeyFull].
	Full bool `json:"full"`
	JSON struct {
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u ComputeFunctionVersionListResponseDataItemAPIKeyUnion) AsComputeFunctionVersionListResponseDataItemAPIKeyPermissions() (v ComputeFunctionVersionListResponseDataItemAPIKeyPermissions) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeFunctionVersionListResponseDataItemAPIKeyUnion) AsComputeFunctionVersionListResponseDataItemAPIKeyFull() (v ComputeFunctionVersionListResponseDataItemAPIKeyFull) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeFunctionVersionListResponseDataItemAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeFunctionVersionListResponseDataItemAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionListResponseDataItemAPIKeyPermissions struct {
	Permissions map[string][]string `json:"permissions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionVersionListResponseDataItemAPIKeyPermissions) RawJSON() string {
	return r.JSON.raw
}
func (r *ComputeFunctionVersionListResponseDataItemAPIKeyPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionListResponseDataItemAPIKeyFull struct {
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
func (r ComputeFunctionVersionListResponseDataItemAPIKeyFull) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionListResponseDataItemAPIKeyFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionListResponseDataItemInspectError struct {
	Code    string                                                        `json:"code" api:"required"`
	Message string                                                        `json:"message" api:"required"`
	Errors  []ComputeFunctionVersionListResponseDataItemInspectErrorError `json:"errors"`
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
func (r ComputeFunctionVersionListResponseDataItemInspectError) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionListResponseDataItemInspectError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionListResponseDataItemInspectErrorError struct {
	Code    string                                                            `json:"code" api:"required"`
	Message string                                                            `json:"message" api:"required"`
	Line    int64                                                             `json:"line"`
	Repair  ComputeFunctionVersionListResponseDataItemInspectErrorErrorRepair `json:"repair"`
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
func (r ComputeFunctionVersionListResponseDataItemInspectErrorError) RawJSON() string {
	return r.JSON.raw
}
func (r *ComputeFunctionVersionListResponseDataItemInspectErrorError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionListResponseDataItemInspectErrorErrorRepair struct {
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
func (r ComputeFunctionVersionListResponseDataItemInspectErrorErrorRepair) RawJSON() string {
	return r.JSON.raw
}
func (r *ComputeFunctionVersionListResponseDataItemInspectErrorErrorRepair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionListResponseDataItemManifest struct {
	Credentials []ComputeFunctionVersionListResponseDataItemManifestCredential `json:"credentials"`
	Hosts       []string                                                       `json:"hosts"`
	Scopes      []string                                                       `json:"scopes"`
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
func (r ComputeFunctionVersionListResponseDataItemManifest) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionVersionListResponseDataItemManifest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionListResponseDataItemManifestCredential struct {
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
func (r ComputeFunctionVersionListResponseDataItemManifestCredential) RawJSON() string {
	return r.JSON.raw
}
func (r *ComputeFunctionVersionListResponseDataItemManifestCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionVersionNewParams struct {
	Source   string                                     `json:"source" api:"required"`
	APIKey   ComputeFunctionVersionNewParamsAPIKeyUnion `json:"apiKey,omitzero"`
	Manifest ComputeFunctionVersionNewParamsManifest    `json:"manifest,omitzero"`
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

type ComputeFunctionVersionNewParamsManifest struct {
	Credentials []ComputeFunctionVersionNewParamsManifestCredential `json:"credentials,omitzero"`
	Hosts       []string                                            `json:"hosts,omitzero"`
	Scopes      []string                                            `json:"scopes,omitzero"`
	paramObj
}

func (r ComputeFunctionVersionNewParamsManifest) MarshalJSON() (data []byte, err error) {
	type shadow ComputeFunctionVersionNewParamsManifest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeFunctionVersionNewParamsManifest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, PackageName are required.
type ComputeFunctionVersionNewParamsManifestCredential struct {
	Name        string `json:"name" api:"required"`
	PackageName string `json:"packageName" api:"required"`
	paramObj
}

func (r ComputeFunctionVersionNewParamsManifestCredential) MarshalJSON() (data []byte, err error) {
	type shadow ComputeFunctionVersionNewParamsManifestCredential
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeFunctionVersionNewParamsManifestCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
