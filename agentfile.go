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
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// AgentFileService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentFileService] method instead.
type AgentFileService struct {
	options []option.RequestOption
}

// NewAgentFileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentFileService(opts ...option.RequestOption) (r AgentFileService) {
	r = AgentFileService{}
	r.options = opts
	return
}

// List the ready files of one agent, optionally filtered by zone.
func (r *AgentFileService) List(ctx context.Context, agentID string, query AgentFileListParams, opts ...option.RequestOption) (res *AgentFileListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/files", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AgentFileListResponse struct {
	Files []AgentFileListResponseFile `json:"files" api:"required"`
	Quota AgentFileListResponseQuota  `json:"quota" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Files       respjson.Field
		Quota       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentFileListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentFileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentFileListResponseFile struct {
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "user", "agent", "workflow".
	CreatedBy   string  `json:"createdBy" api:"required"`
	DisplayName string  `json:"displayName" api:"required"`
	Enabled     bool    `json:"enabled" api:"required"`
	FileID      string  `json:"fileId" api:"required" format:"uuid"`
	Filename    string  `json:"filename" api:"required"`
	MimeType    string  `json:"mimeType" api:"required"`
	SizeBytes   float64 `json:"sizeBytes" api:"required"`
	// Any of "user", "agent", "workflow", "skills".
	Zone             string                                         `json:"zone" api:"required"`
	AgentID          string                                         `json:"agentId" format:"uuid"`
	CollectionStatus AgentFileListResponseFileCollectionStatusUnion `json:"collectionStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		CreatedBy        respjson.Field
		DisplayName      respjson.Field
		Enabled          respjson.Field
		FileID           respjson.Field
		Filename         respjson.Field
		MimeType         respjson.Field
		SizeBytes        respjson.Field
		Zone             respjson.Field
		AgentID          respjson.Field
		CollectionStatus respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentFileListResponseFile) RawJSON() string { return r.JSON.raw }
func (r *AgentFileListResponseFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentFileListResponseFileCollectionStatusUnion contains all possible properties
// and values from [AgentFileListResponseFileCollectionStatusObject],
// [AgentFileListResponseFileCollectionStatusObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AgentFileListResponseFileCollectionStatusUnion struct {
	Completeness string `json:"completeness"`
	Coverage     string `json:"coverage"`
	EndReached   bool   `json:"endReached"`
	// This field is from variant [AgentFileListResponseFileCollectionStatusObject].
	Records   AgentFileListResponseFileCollectionStatusObjectRecords `json:"records"`
	Resumable bool                                                   `json:"resumable"`
	Version   float64                                                `json:"version"`
	JSON      struct {
		Completeness respjson.Field
		Coverage     respjson.Field
		EndReached   respjson.Field
		Records      respjson.Field
		Resumable    respjson.Field
		Version      respjson.Field
		raw          string
	} `json:"-"`
}

func (u AgentFileListResponseFileCollectionStatusUnion) AsAgentFileListResponseFileCollectionStatusObject() (v AgentFileListResponseFileCollectionStatusObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentFileListResponseFileCollectionStatusUnion) AsAgentFileListResponseFileCollectionStatusObject2() (v AgentFileListResponseFileCollectionStatusObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentFileListResponseFileCollectionStatusUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentFileListResponseFileCollectionStatusUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentFileListResponseFileCollectionStatusObject struct {
	// Any of "complete", "partial", "none", "unknown".
	Completeness string `json:"completeness" api:"required"`
	// Any of "exhaustive", "bounded", "unknown".
	Coverage   string                                                 `json:"coverage" api:"required"`
	EndReached bool                                                   `json:"endReached" api:"required"`
	Records    AgentFileListResponseFileCollectionStatusObjectRecords `json:"records" api:"required"`
	Resumable  bool                                                   `json:"resumable" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completeness respjson.Field
		Coverage     respjson.Field
		EndReached   respjson.Field
		Records      respjson.Field
		Resumable    respjson.Field
		Version      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentFileListResponseFileCollectionStatusObject) RawJSON() string { return r.JSON.raw }
func (r *AgentFileListResponseFileCollectionStatusObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentFileListResponseFileCollectionStatusObjectRecords struct {
	Reported int64  `json:"reported" api:"required"`
	Unit     string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reported    respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentFileListResponseFileCollectionStatusObjectRecords) RawJSON() string { return r.JSON.raw }
func (r *AgentFileListResponseFileCollectionStatusObjectRecords) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentFileListResponseFileCollectionStatusObject2 struct {
	// Any of "unknown".
	Completeness string `json:"completeness" api:"required"`
	// Any of "unknown".
	Coverage string `json:"coverage" api:"required"`
	// Any of false.
	EndReached bool `json:"endReached" api:"required"`
	// Any of false.
	Resumable bool `json:"resumable" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completeness respjson.Field
		Coverage     respjson.Field
		EndReached   respjson.Field
		Resumable    respjson.Field
		Version      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentFileListResponseFileCollectionStatusObject2) RawJSON() string { return r.JSON.raw }
func (r *AgentFileListResponseFileCollectionStatusObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentFileListResponseQuota struct {
	IncludedBytes  int64 `json:"includedBytes" api:"required"`
	MaxBytes       int64 `json:"maxBytes" api:"required"`
	OverageAllowed bool  `json:"overageAllowed" api:"required"`
	UsedBytes      int64 `json:"usedBytes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludedBytes  respjson.Field
		MaxBytes       respjson.Field
		OverageAllowed respjson.Field
		UsedBytes      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentFileListResponseQuota) RawJSON() string { return r.JSON.raw }
func (r *AgentFileListResponseQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentFileListParams struct {
	// Any of "user", "agent", "workflow", "skills".
	Zone AgentFileListParamsZone `query:"zone,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentFileListParams]'s query parameters as `url.Values`.
func (r AgentFileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentFileListParamsZone string

const (
	AgentFileListParamsZoneUser     AgentFileListParamsZone = "user"
	AgentFileListParamsZoneAgent    AgentFileListParamsZone = "agent"
	AgentFileListParamsZoneWorkflow AgentFileListParamsZone = "workflow"
	AgentFileListParamsZoneSkills   AgentFileListParamsZone = "skills"
)
