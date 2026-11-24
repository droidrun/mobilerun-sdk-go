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
)

// FileService contains methods and other services that help with interacting with
// the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.options = opts
	return
}

// Partial update of `displayName` and/or `enabled`. Only files with `zone=skills`
// are mutable; other zones return 422 `unsupported_zone`.
func (r *FileService) Update(ctx context.Context, fileID string, body FileUpdateParams, opts ...option.RequestOption) (res *FileUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List the user's ready files, optionally filtered by zone.
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *FileListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "agents/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Hard-delete a file.
func (r *FileService) Delete(ctx context.Context, fileID string, opts ...option.RequestOption) (res *FileDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Soft-cancels an in-flight upload before confirm, transitioning it to expired and
// deleting the underlying object. Only acts on `pending` uploads — refuses to
// touch `ready` files, so confirmed files are never removed. Idempotent:
// `{ cancelled: false }` if the file exists but is no longer pending.
func (r *FileService) CancelPending(ctx context.Context, fileID string, opts ...option.RequestOption) (res *FileCancelPendingResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/files/%s/pending", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Confirm a file upload after validating the stored object with a server-side HEAD
// request.
func (r *FileService) Confirm(ctx context.Context, fileID string, opts ...option.RequestOption) (res *FileConfirmResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/files/%s/confirm", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Redirects to a short-lived presigned download URL. Missing, non-ready, and files
// owned by another tenant all return the same 404 response.
func (r *FileService) Download(ctx context.Context, fileID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return err
	}
	path := fmt.Sprintf("agents/files/%s/download", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Mint a presigned PUT URL and a fileId for uploading a user file.
func (r *FileService) UploadURL(ctx context.Context, body FileUploadURLParams, opts ...option.RequestOption) (res *FileUploadURLResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "agents/files/upload-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type FileUpdateResponse struct {
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "user", "agent", "workflow".
	CreatedBy   FileUpdateResponseCreatedBy `json:"createdBy" api:"required"`
	DisplayName string                      `json:"displayName" api:"required"`
	Enabled     bool                        `json:"enabled" api:"required"`
	FileID      string                      `json:"fileId" api:"required" format:"uuid"`
	Filename    string                      `json:"filename" api:"required"`
	MimeType    string                      `json:"mimeType" api:"required"`
	SizeBytes   float64                     `json:"sizeBytes" api:"required"`
	// Any of "user", "agent", "workflow", "skills".
	Zone             FileUpdateResponseZone                  `json:"zone" api:"required"`
	AgentID          string                                  `json:"agentId" format:"uuid"`
	CollectionStatus FileUpdateResponseCollectionStatusUnion `json:"collectionStatus"`
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
func (r FileUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FileUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUpdateResponseCreatedBy string

const (
	FileUpdateResponseCreatedByUser     FileUpdateResponseCreatedBy = "user"
	FileUpdateResponseCreatedByAgent    FileUpdateResponseCreatedBy = "agent"
	FileUpdateResponseCreatedByWorkflow FileUpdateResponseCreatedBy = "workflow"
)

type FileUpdateResponseZone string

const (
	FileUpdateResponseZoneUser     FileUpdateResponseZone = "user"
	FileUpdateResponseZoneAgent    FileUpdateResponseZone = "agent"
	FileUpdateResponseZoneWorkflow FileUpdateResponseZone = "workflow"
	FileUpdateResponseZoneSkills   FileUpdateResponseZone = "skills"
)

// FileUpdateResponseCollectionStatusUnion contains all possible properties and
// values from [FileUpdateResponseCollectionStatusObject],
// [FileUpdateResponseCollectionStatusObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FileUpdateResponseCollectionStatusUnion struct {
	Completeness string `json:"completeness"`
	Coverage     string `json:"coverage"`
	EndReached   bool   `json:"endReached"`
	// This field is from variant [FileUpdateResponseCollectionStatusObject].
	Records   FileUpdateResponseCollectionStatusObjectRecords `json:"records"`
	Resumable bool                                            `json:"resumable"`
	Version   float64                                         `json:"version"`
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

func (u FileUpdateResponseCollectionStatusUnion) AsFileUpdateResponseCollectionStatusObject() (v FileUpdateResponseCollectionStatusObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileUpdateResponseCollectionStatusUnion) AsFileUpdateResponseCollectionStatusObject2() (v FileUpdateResponseCollectionStatusObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FileUpdateResponseCollectionStatusUnion) RawJSON() string { return u.JSON.raw }

func (r *FileUpdateResponseCollectionStatusUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUpdateResponseCollectionStatusObject struct {
	// Any of "complete", "partial", "none", "unknown".
	Completeness string `json:"completeness" api:"required"`
	// Any of "exhaustive", "bounded", "unknown".
	Coverage   string                                          `json:"coverage" api:"required"`
	EndReached bool                                            `json:"endReached" api:"required"`
	Records    FileUpdateResponseCollectionStatusObjectRecords `json:"records" api:"required"`
	Resumable  bool                                            `json:"resumable" api:"required"`
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
func (r FileUpdateResponseCollectionStatusObject) RawJSON() string { return r.JSON.raw }
func (r *FileUpdateResponseCollectionStatusObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUpdateResponseCollectionStatusObjectRecords struct {
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
func (r FileUpdateResponseCollectionStatusObjectRecords) RawJSON() string { return r.JSON.raw }
func (r *FileUpdateResponseCollectionStatusObjectRecords) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUpdateResponseCollectionStatusObject2 struct {
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
func (r FileUpdateResponseCollectionStatusObject2) RawJSON() string { return r.JSON.raw }
func (r *FileUpdateResponseCollectionStatusObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListResponse struct {
	Files []FileListResponseFile `json:"files" api:"required"`
	Quota FileListResponseQuota  `json:"quota" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Files       respjson.Field
		Quota       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileListResponse) RawJSON() string { return r.JSON.raw }
func (r *FileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListResponseFile struct {
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
	Zone             string                                    `json:"zone" api:"required"`
	AgentID          string                                    `json:"agentId" format:"uuid"`
	CollectionStatus FileListResponseFileCollectionStatusUnion `json:"collectionStatus"`
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
func (r FileListResponseFile) RawJSON() string { return r.JSON.raw }
func (r *FileListResponseFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FileListResponseFileCollectionStatusUnion contains all possible properties and
// values from [FileListResponseFileCollectionStatusObject],
// [FileListResponseFileCollectionStatusObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FileListResponseFileCollectionStatusUnion struct {
	Completeness string `json:"completeness"`
	Coverage     string `json:"coverage"`
	EndReached   bool   `json:"endReached"`
	// This field is from variant [FileListResponseFileCollectionStatusObject].
	Records   FileListResponseFileCollectionStatusObjectRecords `json:"records"`
	Resumable bool                                              `json:"resumable"`
	Version   float64                                           `json:"version"`
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

func (u FileListResponseFileCollectionStatusUnion) AsFileListResponseFileCollectionStatusObject() (v FileListResponseFileCollectionStatusObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileListResponseFileCollectionStatusUnion) AsFileListResponseFileCollectionStatusObject2() (v FileListResponseFileCollectionStatusObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FileListResponseFileCollectionStatusUnion) RawJSON() string { return u.JSON.raw }

func (r *FileListResponseFileCollectionStatusUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListResponseFileCollectionStatusObject struct {
	// Any of "complete", "partial", "none", "unknown".
	Completeness string `json:"completeness" api:"required"`
	// Any of "exhaustive", "bounded", "unknown".
	Coverage   string                                            `json:"coverage" api:"required"`
	EndReached bool                                              `json:"endReached" api:"required"`
	Records    FileListResponseFileCollectionStatusObjectRecords `json:"records" api:"required"`
	Resumable  bool                                              `json:"resumable" api:"required"`
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
func (r FileListResponseFileCollectionStatusObject) RawJSON() string { return r.JSON.raw }
func (r *FileListResponseFileCollectionStatusObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListResponseFileCollectionStatusObjectRecords struct {
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
func (r FileListResponseFileCollectionStatusObjectRecords) RawJSON() string { return r.JSON.raw }
func (r *FileListResponseFileCollectionStatusObjectRecords) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListResponseFileCollectionStatusObject2 struct {
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
func (r FileListResponseFileCollectionStatusObject2) RawJSON() string { return r.JSON.raw }
func (r *FileListResponseFileCollectionStatusObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListResponseQuota struct {
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
func (r FileListResponseQuota) RawJSON() string { return r.JSON.raw }
func (r *FileListResponseQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileDeleteResponse struct {
	// Any of true.
	Ok bool `json:"ok" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ok          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FileDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileCancelPendingResponse struct {
	Cancelled bool `json:"cancelled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cancelled   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileCancelPendingResponse) RawJSON() string { return r.JSON.raw }
func (r *FileCancelPendingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileConfirmResponse struct {
	ActualSizeBytes float64   `json:"actualSizeBytes" api:"required"`
	CreatedAt       time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "user", "agent", "workflow".
	CreatedBy   FileConfirmResponseCreatedBy `json:"createdBy" api:"required"`
	DisplayName string                       `json:"displayName" api:"required"`
	Enabled     bool                         `json:"enabled" api:"required"`
	FileID      string                       `json:"fileId" api:"required" format:"uuid"`
	Filename    string                       `json:"filename" api:"required"`
	MimeType    string                       `json:"mimeType" api:"required"`
	SizeBytes   float64                      `json:"sizeBytes" api:"required"`
	// Any of "ready".
	State FileConfirmResponseState `json:"state" api:"required"`
	// Any of "user", "agent", "workflow", "skills".
	Zone             FileConfirmResponseZone                  `json:"zone" api:"required"`
	AgentID          string                                   `json:"agentId" format:"uuid"`
	CollectionStatus FileConfirmResponseCollectionStatusUnion `json:"collectionStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActualSizeBytes  respjson.Field
		CreatedAt        respjson.Field
		CreatedBy        respjson.Field
		DisplayName      respjson.Field
		Enabled          respjson.Field
		FileID           respjson.Field
		Filename         respjson.Field
		MimeType         respjson.Field
		SizeBytes        respjson.Field
		State            respjson.Field
		Zone             respjson.Field
		AgentID          respjson.Field
		CollectionStatus respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileConfirmResponse) RawJSON() string { return r.JSON.raw }
func (r *FileConfirmResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileConfirmResponseCreatedBy string

const (
	FileConfirmResponseCreatedByUser     FileConfirmResponseCreatedBy = "user"
	FileConfirmResponseCreatedByAgent    FileConfirmResponseCreatedBy = "agent"
	FileConfirmResponseCreatedByWorkflow FileConfirmResponseCreatedBy = "workflow"
)

type FileConfirmResponseState string

const (
	FileConfirmResponseStateReady FileConfirmResponseState = "ready"
)

type FileConfirmResponseZone string

const (
	FileConfirmResponseZoneUser     FileConfirmResponseZone = "user"
	FileConfirmResponseZoneAgent    FileConfirmResponseZone = "agent"
	FileConfirmResponseZoneWorkflow FileConfirmResponseZone = "workflow"
	FileConfirmResponseZoneSkills   FileConfirmResponseZone = "skills"
)

// FileConfirmResponseCollectionStatusUnion contains all possible properties and
// values from [FileConfirmResponseCollectionStatusObject],
// [FileConfirmResponseCollectionStatusObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FileConfirmResponseCollectionStatusUnion struct {
	Completeness string `json:"completeness"`
	Coverage     string `json:"coverage"`
	EndReached   bool   `json:"endReached"`
	// This field is from variant [FileConfirmResponseCollectionStatusObject].
	Records   FileConfirmResponseCollectionStatusObjectRecords `json:"records"`
	Resumable bool                                             `json:"resumable"`
	Version   float64                                          `json:"version"`
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

func (u FileConfirmResponseCollectionStatusUnion) AsFileConfirmResponseCollectionStatusObject() (v FileConfirmResponseCollectionStatusObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileConfirmResponseCollectionStatusUnion) AsFileConfirmResponseCollectionStatusObject2() (v FileConfirmResponseCollectionStatusObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FileConfirmResponseCollectionStatusUnion) RawJSON() string { return u.JSON.raw }

func (r *FileConfirmResponseCollectionStatusUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileConfirmResponseCollectionStatusObject struct {
	// Any of "complete", "partial", "none", "unknown".
	Completeness string `json:"completeness" api:"required"`
	// Any of "exhaustive", "bounded", "unknown".
	Coverage   string                                           `json:"coverage" api:"required"`
	EndReached bool                                             `json:"endReached" api:"required"`
	Records    FileConfirmResponseCollectionStatusObjectRecords `json:"records" api:"required"`
	Resumable  bool                                             `json:"resumable" api:"required"`
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
func (r FileConfirmResponseCollectionStatusObject) RawJSON() string { return r.JSON.raw }
func (r *FileConfirmResponseCollectionStatusObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileConfirmResponseCollectionStatusObjectRecords struct {
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
func (r FileConfirmResponseCollectionStatusObjectRecords) RawJSON() string { return r.JSON.raw }
func (r *FileConfirmResponseCollectionStatusObjectRecords) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileConfirmResponseCollectionStatusObject2 struct {
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
func (r FileConfirmResponseCollectionStatusObject2) RawJSON() string { return r.JSON.raw }
func (r *FileConfirmResponseCollectionStatusObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUploadURLResponse struct {
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	FileID    string    `json:"fileId" api:"required" format:"uuid"`
	PutURL    string    `json:"putUrl" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		FileID      respjson.Field
		PutURL      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileUploadURLResponse) RawJSON() string { return r.JSON.raw }
func (r *FileUploadURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUpdateParams struct {
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	Enabled     param.Opt[bool]   `json:"enabled,omitzero"`
	paramObj
}

func (r FileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListParams struct {
	// Any of "user", "agent", "workflow", "skills".
	Zone FileListParamsZone `query:"zone,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileListParamsZone string

const (
	FileListParamsZoneUser     FileListParamsZone = "user"
	FileListParamsZoneAgent    FileListParamsZone = "agent"
	FileListParamsZoneWorkflow FileListParamsZone = "workflow"
	FileListParamsZoneSkills   FileListParamsZone = "skills"
)

type FileUploadURLParams struct {
	Filename      string            `json:"filename" api:"required"`
	MimeType      string            `json:"mimeType" api:"required"`
	SizeBytes     int64             `json:"sizeBytes" api:"required"`
	ChatSessionID param.Opt[string] `json:"chatSessionId,omitzero" format:"uuid"`
	// Any of "user", "skills".
	Zone FileUploadURLParamsZone `json:"zone,omitzero"`
	paramObj
}

func (r FileUploadURLParams) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadURLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUploadURLParamsZone string

const (
	FileUploadURLParamsZoneUser   FileUploadURLParamsZone = "user"
	FileUploadURLParamsZoneSkills FileUploadURLParamsZone = "skills"
)
