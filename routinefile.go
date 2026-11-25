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
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiquery"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/pagination"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// RoutineFileService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoutineFileService] method instead.
type RoutineFileService struct {
	options []option.RequestOption
}

// NewRoutineFileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoutineFileService(opts ...option.RequestOption) (r RoutineFileService) {
	r = RoutineFileService{}
	r.options = opts
	return
}

// Lists the files saved in Mobilerun storage for routines whose delivery
// destination is Mobilerun with a folder: the run recording, screenshots and
// files.upload outputs. Newest first; filter by `folder`, `flowId` or
// `executionId`. Files stay until you delete them or delete their routine.
func (r *RoutineFileService) List(ctx context.Context, query RoutineFileListParams, opts ...option.RequestOption) (res *pagination.Page[RoutineFileListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "routine-files"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists the files saved in Mobilerun storage for routines whose delivery
// destination is Mobilerun with a folder: the run recording, screenshots and
// files.upload outputs. Newest first; filter by `folder`, `flowId` or
// `executionId`. Files stay until you delete them or delete their routine.
func (r *RoutineFileService) ListAutoPaging(ctx context.Context, query RoutineFileListParams, opts ...option.RequestOption) *pagination.PageAutoPager[RoutineFileListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes one saved file permanently and frees its storage (the file disappears at
// once; the stored object is removed right after). Idempotent: an already deleted
// file returns 204 too.
func (r *RoutineFileService) Delete(ctx context.Context, fileID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return err
	}
	path := fmt.Sprintf("routine-files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Deletes every saved file of exactly one folder (`folder`) or one run
// (`executionId`) permanently. Pass exactly one of the two. Returns how many files
// were deleted.
func (r *RoutineFileService) DeleteMany(ctx context.Context, body RoutineFileDeleteManyParams, opts ...option.RequestOption) (res *RoutineFileDeleteManyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "routine-files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Redirects (302) to a time-limited (at most 5 minutes) download URL for the file.
// 404 when the file does not exist or belongs to someone else.
func (r *RoutineFileService) Download(ctx context.Context, fileID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return err
	}
	path := fmt.Sprintf("routine-files/%s/download", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Lists the folders that hold saved files with their file count, total size and
// newest save time, newest first. Independent of whether the routine or run still
// exists; `flows` names the routines that saved into each folder.
func (r *RoutineFileService) ListFolders(ctx context.Context, query RoutineFileListFoldersParams, opts ...option.RequestOption) (res *pagination.Page[RoutineFileListFoldersResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "routine-files/folders"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists the folders that hold saved files with their file count, total size and
// newest save time, newest first. Independent of whether the routine or run still
// exists; `flows` names the routines that saved into each folder.
func (r *RoutineFileService) ListFoldersAutoPaging(ctx context.Context, query RoutineFileListFoldersParams, opts ...option.RequestOption) *pagination.PageAutoPager[RoutineFileListFoldersResponse] {
	return pagination.NewPageAutoPager(r.ListFolders(ctx, query, opts...))
}

type RoutineFileListResponse struct {
	ID          string    `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time `json:"createdAt" api:"required" format:"date-time"`
	ExecutionID string    `json:"executionId" api:"required" format:"uuid"`
	Filename    string    `json:"filename" api:"required"`
	// True when the routine no longer exists (replaced by a deploy, or deleted before
	// its files were deleted with it). Its files stay until you delete them.
	FlowDeleted bool   `json:"flowDeleted" api:"required"`
	FlowID      string `json:"flowId" api:"required" format:"uuid"`
	// Null when the routine no longer exists.
	FlowName string `json:"flowName" api:"required"`
	// Canonical folder the file was saved under, for example "/Exports/241109".
	Folder string `json:"folder" api:"required"`
	// Any of "recording", "screenshot", "file".
	Kind      RoutineFileListResponseKind `json:"kind" api:"required"`
	MimeType  string                      `json:"mimeType" api:"required"`
	SizeBytes int64                       `json:"sizeBytes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ExecutionID respjson.Field
		Filename    respjson.Field
		FlowDeleted respjson.Field
		FlowID      respjson.Field
		FlowName    respjson.Field
		Folder      respjson.Field
		Kind        respjson.Field
		MimeType    respjson.Field
		SizeBytes   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoutineFileListResponse) RawJSON() string { return r.JSON.raw }
func (r *RoutineFileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoutineFileListResponseKind string

const (
	RoutineFileListResponseKindRecording  RoutineFileListResponseKind = "recording"
	RoutineFileListResponseKindScreenshot RoutineFileListResponseKind = "screenshot"
	RoutineFileListResponseKindFile       RoutineFileListResponseKind = "file"
)

type RoutineFileDeleteManyResponse struct {
	Deleted int64 `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoutineFileDeleteManyResponse) RawJSON() string { return r.JSON.raw }
func (r *RoutineFileDeleteManyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoutineFileListFoldersResponse struct {
	FileCount     int64                                `json:"fileCount" api:"required"`
	Flows         []RoutineFileListFoldersResponseFlow `json:"flows" api:"required"`
	Folder        string                               `json:"folder" api:"required"`
	LastCreatedAt time.Time                            `json:"lastCreatedAt" api:"required" format:"date-time"`
	SizeBytes     int64                                `json:"sizeBytes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileCount     respjson.Field
		Flows         respjson.Field
		Folder        respjson.Field
		LastCreatedAt respjson.Field
		SizeBytes     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoutineFileListFoldersResponse) RawJSON() string { return r.JSON.raw }
func (r *RoutineFileListFoldersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoutineFileListFoldersResponseFlow struct {
	FlowDeleted bool   `json:"flowDeleted" api:"required"`
	FlowID      string `json:"flowId" api:"required" format:"uuid"`
	FlowName    string `json:"flowName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FlowDeleted respjson.Field
		FlowID      respjson.Field
		FlowName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoutineFileListFoldersResponseFlow) RawJSON() string { return r.JSON.raw }
func (r *RoutineFileListFoldersResponseFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoutineFileListParams struct {
	ExecutionID param.Opt[string] `query:"executionId,omitzero" format:"uuid" json:"-"`
	FlowID      param.Opt[string] `query:"flowId,omitzero" format:"uuid" json:"-"`
	// Exact folder (canonicalized: "//" collapsed, trailing "/" stripped).
	Folder   param.Opt[string] `query:"folder,omitzero" json:"-"`
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64]  `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoutineFileListParams]'s query parameters as `url.Values`.
func (r RoutineFileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoutineFileDeleteManyParams struct {
	ExecutionID param.Opt[string] `query:"executionId,omitzero" format:"uuid" json:"-"`
	Folder      param.Opt[string] `query:"folder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoutineFileDeleteManyParams]'s query parameters as
// `url.Values`.
func (r RoutineFileDeleteManyParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoutineFileListFoldersParams struct {
	FlowID   param.Opt[string] `query:"flowId,omitzero" format:"uuid" json:"-"`
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64]  `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoutineFileListFoldersParams]'s query parameters as
// `url.Values`.
func (r RoutineFileListFoldersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
