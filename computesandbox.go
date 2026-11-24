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

// ComputeSandboxService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeSandboxService] method instead.
type ComputeSandboxService struct {
	options  []option.RequestOption
	Commands ComputeSandboxCommandService
	Files    ComputeSandboxFileService
}

// NewComputeSandboxService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputeSandboxService(opts ...option.RequestOption) (r ComputeSandboxService) {
	r = ComputeSandboxService{}
	r.options = opts
	r.Commands = NewComputeSandboxCommandService(opts...)
	r.Files = NewComputeSandboxFileService(opts...)
	return
}

// Create a sandbox
func (r *ComputeSandboxService) New(ctx context.Context, body ComputeSandboxNewParams, opts ...option.RequestOption) (res *ComputeSandboxNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "compute/sandboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a sandbox
func (r *ComputeSandboxService) Get(ctx context.Context, sandboxID string, opts ...option.RequestOption) (res *ComputeSandboxGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if sandboxID == "" {
		err = errors.New("missing required sandboxId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/sandboxes/%s", sandboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List sandboxes
func (r *ComputeSandboxService) List(ctx context.Context, query ComputeSandboxListParams, opts ...option.RequestOption) (res *ComputeSandboxListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "compute/sandboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Stop a sandbox
func (r *ComputeSandboxService) Delete(ctx context.Context, sandboxID string, opts ...option.RequestOption) (res *ComputeSandboxDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if sandboxID == "" {
		err = errors.New("missing required sandboxId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/sandboxes/%s", sandboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Execute a shell command
func (r *ComputeSandboxService) Exec(ctx context.Context, sandboxID string, body ComputeSandboxExecParams, opts ...option.RequestOption) (res *ComputeSandboxExecResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if sandboxID == "" {
		err = errors.New("missing required sandboxId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/sandboxes/%s/exec", sandboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ComputeSandboxNewResponse struct {
	Data ComputeSandboxNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxNewResponseData struct {
	ID           string    `json:"id" api:"required" format:"uuid"`
	AllowedHosts []string  `json:"allowedHosts" api:"required"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	ExpiresAt    time.Time `json:"expiresAt" api:"required" format:"date-time"`
	IdleTimeoutS int64     `json:"idleTimeoutS" api:"required"`
	Name         string    `json:"name" api:"required"`
	// Any of "starting", "running", "stopped".
	Status    string    `json:"status" api:"required"`
	StoppedAt time.Time `json:"stoppedAt" api:"required" format:"date-time"`
	// Any of "deleted", "idle", "timeout", "lost".
	StopReason  string                                   `json:"stopReason" api:"required"`
	UsedSeconds float64                                  `json:"usedSeconds" api:"required"`
	APIKey      ComputeSandboxNewResponseDataAPIKeyUnion `json:"apiKey"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AllowedHosts respjson.Field
		CreatedAt    respjson.Field
		ExpiresAt    respjson.Field
		IdleTimeoutS respjson.Field
		Name         respjson.Field
		Status       respjson.Field
		StoppedAt    respjson.Field
		StopReason   respjson.Field
		UsedSeconds  respjson.Field
		APIKey       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeSandboxNewResponseDataAPIKeyUnion contains all possible properties and
// values from [ComputeSandboxNewResponseDataAPIKeyObject],
// [ComputeSandboxNewResponseDataAPIKeyObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ComputeSandboxNewResponseDataAPIKeyUnion struct {
	ExpiresAt time.Time `json:"expiresAt"`
	// This field is from variant [ComputeSandboxNewResponseDataAPIKeyObject].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant [ComputeSandboxNewResponseDataAPIKeyObject2].
	Full bool `json:"full"`
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u ComputeSandboxNewResponseDataAPIKeyUnion) AsComputeSandboxNewResponseDataAPIKeyObject() (v ComputeSandboxNewResponseDataAPIKeyObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeSandboxNewResponseDataAPIKeyUnion) AsComputeSandboxNewResponseDataAPIKeyObject2() (v ComputeSandboxNewResponseDataAPIKeyObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeSandboxNewResponseDataAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeSandboxNewResponseDataAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxNewResponseDataAPIKeyObject struct {
	ExpiresAt   time.Time           `json:"expiresAt" api:"required" format:"date-time"`
	Permissions map[string][]string `json:"permissions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxNewResponseDataAPIKeyObject) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxNewResponseDataAPIKeyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxNewResponseDataAPIKeyObject2 struct {
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// Any of true.
	Full bool `json:"full" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Full        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxNewResponseDataAPIKeyObject2) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxNewResponseDataAPIKeyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxGetResponse struct {
	Data ComputeSandboxGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxGetResponseData struct {
	ID           string    `json:"id" api:"required" format:"uuid"`
	AllowedHosts []string  `json:"allowedHosts" api:"required"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	ExpiresAt    time.Time `json:"expiresAt" api:"required" format:"date-time"`
	IdleTimeoutS int64     `json:"idleTimeoutS" api:"required"`
	Name         string    `json:"name" api:"required"`
	// Any of "starting", "running", "stopped".
	Status    string    `json:"status" api:"required"`
	StoppedAt time.Time `json:"stoppedAt" api:"required" format:"date-time"`
	// Any of "deleted", "idle", "timeout", "lost".
	StopReason  string                                   `json:"stopReason" api:"required"`
	UsedSeconds float64                                  `json:"usedSeconds" api:"required"`
	APIKey      ComputeSandboxGetResponseDataAPIKeyUnion `json:"apiKey"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AllowedHosts respjson.Field
		CreatedAt    respjson.Field
		ExpiresAt    respjson.Field
		IdleTimeoutS respjson.Field
		Name         respjson.Field
		Status       respjson.Field
		StoppedAt    respjson.Field
		StopReason   respjson.Field
		UsedSeconds  respjson.Field
		APIKey       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeSandboxGetResponseDataAPIKeyUnion contains all possible properties and
// values from [ComputeSandboxGetResponseDataAPIKeyObject],
// [ComputeSandboxGetResponseDataAPIKeyObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ComputeSandboxGetResponseDataAPIKeyUnion struct {
	ExpiresAt time.Time `json:"expiresAt"`
	// This field is from variant [ComputeSandboxGetResponseDataAPIKeyObject].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant [ComputeSandboxGetResponseDataAPIKeyObject2].
	Full bool `json:"full"`
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u ComputeSandboxGetResponseDataAPIKeyUnion) AsComputeSandboxGetResponseDataAPIKeyObject() (v ComputeSandboxGetResponseDataAPIKeyObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeSandboxGetResponseDataAPIKeyUnion) AsComputeSandboxGetResponseDataAPIKeyObject2() (v ComputeSandboxGetResponseDataAPIKeyObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeSandboxGetResponseDataAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeSandboxGetResponseDataAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxGetResponseDataAPIKeyObject struct {
	ExpiresAt   time.Time           `json:"expiresAt" api:"required" format:"date-time"`
	Permissions map[string][]string `json:"permissions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxGetResponseDataAPIKeyObject) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxGetResponseDataAPIKeyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxGetResponseDataAPIKeyObject2 struct {
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// Any of true.
	Full bool `json:"full" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Full        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxGetResponseDataAPIKeyObject2) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxGetResponseDataAPIKeyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxListResponse struct {
	Data ComputeSandboxListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxListResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxListResponseData struct {
	Items      []ComputeSandboxListResponseDataItem `json:"items" api:"required"`
	Pagination shared.Pagination                    `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxListResponseDataItem struct {
	ID           string    `json:"id" api:"required" format:"uuid"`
	AllowedHosts []string  `json:"allowedHosts" api:"required"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	ExpiresAt    time.Time `json:"expiresAt" api:"required" format:"date-time"`
	IdleTimeoutS int64     `json:"idleTimeoutS" api:"required"`
	Name         string    `json:"name" api:"required"`
	// Any of "starting", "running", "stopped".
	Status    string    `json:"status" api:"required"`
	StoppedAt time.Time `json:"stoppedAt" api:"required" format:"date-time"`
	// Any of "deleted", "idle", "timeout", "lost".
	StopReason  string                                        `json:"stopReason" api:"required"`
	UsedSeconds float64                                       `json:"usedSeconds" api:"required"`
	APIKey      ComputeSandboxListResponseDataItemAPIKeyUnion `json:"apiKey"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AllowedHosts respjson.Field
		CreatedAt    respjson.Field
		ExpiresAt    respjson.Field
		IdleTimeoutS respjson.Field
		Name         respjson.Field
		Status       respjson.Field
		StoppedAt    respjson.Field
		StopReason   respjson.Field
		UsedSeconds  respjson.Field
		APIKey       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxListResponseDataItem) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxListResponseDataItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeSandboxListResponseDataItemAPIKeyUnion contains all possible properties
// and values from [ComputeSandboxListResponseDataItemAPIKeyObject],
// [ComputeSandboxListResponseDataItemAPIKeyObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ComputeSandboxListResponseDataItemAPIKeyUnion struct {
	ExpiresAt time.Time `json:"expiresAt"`
	// This field is from variant [ComputeSandboxListResponseDataItemAPIKeyObject].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant [ComputeSandboxListResponseDataItemAPIKeyObject2].
	Full bool `json:"full"`
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u ComputeSandboxListResponseDataItemAPIKeyUnion) AsComputeSandboxListResponseDataItemAPIKeyObject() (v ComputeSandboxListResponseDataItemAPIKeyObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeSandboxListResponseDataItemAPIKeyUnion) AsComputeSandboxListResponseDataItemAPIKeyObject2() (v ComputeSandboxListResponseDataItemAPIKeyObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeSandboxListResponseDataItemAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeSandboxListResponseDataItemAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxListResponseDataItemAPIKeyObject struct {
	ExpiresAt   time.Time           `json:"expiresAt" api:"required" format:"date-time"`
	Permissions map[string][]string `json:"permissions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxListResponseDataItemAPIKeyObject) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxListResponseDataItemAPIKeyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxListResponseDataItemAPIKeyObject2 struct {
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// Any of true.
	Full bool `json:"full" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Full        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxListResponseDataItemAPIKeyObject2) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxListResponseDataItemAPIKeyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxDeleteResponse struct {
	Data ComputeSandboxDeleteResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxDeleteResponseData struct {
	ID           string    `json:"id" api:"required" format:"uuid"`
	AllowedHosts []string  `json:"allowedHosts" api:"required"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	ExpiresAt    time.Time `json:"expiresAt" api:"required" format:"date-time"`
	IdleTimeoutS int64     `json:"idleTimeoutS" api:"required"`
	Name         string    `json:"name" api:"required"`
	// Any of "starting", "running", "stopped".
	Status    string    `json:"status" api:"required"`
	StoppedAt time.Time `json:"stoppedAt" api:"required" format:"date-time"`
	// Any of "deleted", "idle", "timeout", "lost".
	StopReason  string                                      `json:"stopReason" api:"required"`
	UsedSeconds float64                                     `json:"usedSeconds" api:"required"`
	APIKey      ComputeSandboxDeleteResponseDataAPIKeyUnion `json:"apiKey"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AllowedHosts respjson.Field
		CreatedAt    respjson.Field
		ExpiresAt    respjson.Field
		IdleTimeoutS respjson.Field
		Name         respjson.Field
		Status       respjson.Field
		StoppedAt    respjson.Field
		StopReason   respjson.Field
		UsedSeconds  respjson.Field
		APIKey       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeSandboxDeleteResponseDataAPIKeyUnion contains all possible properties and
// values from [ComputeSandboxDeleteResponseDataAPIKeyObject],
// [ComputeSandboxDeleteResponseDataAPIKeyObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ComputeSandboxDeleteResponseDataAPIKeyUnion struct {
	ExpiresAt time.Time `json:"expiresAt"`
	// This field is from variant [ComputeSandboxDeleteResponseDataAPIKeyObject].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant [ComputeSandboxDeleteResponseDataAPIKeyObject2].
	Full bool `json:"full"`
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u ComputeSandboxDeleteResponseDataAPIKeyUnion) AsComputeSandboxDeleteResponseDataAPIKeyObject() (v ComputeSandboxDeleteResponseDataAPIKeyObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeSandboxDeleteResponseDataAPIKeyUnion) AsComputeSandboxDeleteResponseDataAPIKeyObject2() (v ComputeSandboxDeleteResponseDataAPIKeyObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeSandboxDeleteResponseDataAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeSandboxDeleteResponseDataAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxDeleteResponseDataAPIKeyObject struct {
	ExpiresAt   time.Time           `json:"expiresAt" api:"required" format:"date-time"`
	Permissions map[string][]string `json:"permissions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxDeleteResponseDataAPIKeyObject) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxDeleteResponseDataAPIKeyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxDeleteResponseDataAPIKeyObject2 struct {
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// Any of true.
	Full bool `json:"full" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Full        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxDeleteResponseDataAPIKeyObject2) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxDeleteResponseDataAPIKeyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxExecResponse struct {
	Data ComputeSandboxExecResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxExecResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxExecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxExecResponseData struct {
	CommandID  string    `json:"commandId" api:"required" format:"uuid"`
	DurationMs float64   `json:"durationMs" api:"required"`
	ExitCode   int64     `json:"exitCode" api:"required"`
	StartedAt  time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Any of "running", "exited".
	Status    string `json:"status" api:"required"`
	Stderr    string `json:"stderr" api:"required"`
	Stdout    string `json:"stdout" api:"required"`
	TimedOut  bool   `json:"timedOut" api:"required"`
	Truncated bool   `json:"truncated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommandID   respjson.Field
		DurationMs  respjson.Field
		ExitCode    respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Stderr      respjson.Field
		Stdout      respjson.Field
		TimedOut    respjson.Field
		Truncated   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxExecResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxExecResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxNewParams struct {
	IdleTimeoutS param.Opt[int64]                   `json:"idleTimeoutS,omitzero"`
	Name         param.Opt[string]                  `json:"name,omitzero"`
	TimeoutS     param.Opt[int64]                   `json:"timeoutS,omitzero"`
	AllowedHosts []string                           `json:"allowedHosts,omitzero"`
	APIKey       ComputeSandboxNewParamsAPIKeyUnion `json:"apiKey,omitzero"`
	Env          map[string]string                  `json:"env,omitzero"`
	paramObj
}

func (r ComputeSandboxNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputeSandboxNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeSandboxNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ComputeSandboxNewParamsAPIKeyUnion struct {
	OfComputeSandboxNewsAPIKeyPermissions *ComputeSandboxNewParamsAPIKeyPermissions `json:",omitzero,inline"`
	OfComputeSandboxNewsAPIKeyFull        *ComputeSandboxNewParamsAPIKeyFull        `json:",omitzero,inline"`
	OfComputeSandboxNewsAPIKeyKey         *ComputeSandboxNewParamsAPIKeyKey         `json:",omitzero,inline"`
	paramUnion
}

func (u ComputeSandboxNewParamsAPIKeyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfComputeSandboxNewsAPIKeyPermissions, u.OfComputeSandboxNewsAPIKeyFull, u.OfComputeSandboxNewsAPIKeyKey)
}
func (u *ComputeSandboxNewParamsAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Permissions is required.
type ComputeSandboxNewParamsAPIKeyPermissions struct {
	Permissions map[string][]string `json:"permissions,omitzero" api:"required"`
	paramObj
}

func (r ComputeSandboxNewParamsAPIKeyPermissions) MarshalJSON() (data []byte, err error) {
	type shadow ComputeSandboxNewParamsAPIKeyPermissions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeSandboxNewParamsAPIKeyPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Full is required.
type ComputeSandboxNewParamsAPIKeyFull struct {
	// Any of true.
	Full bool `json:"full,omitzero" api:"required"`
	paramObj
}

func (r ComputeSandboxNewParamsAPIKeyFull) MarshalJSON() (data []byte, err error) {
	type shadow ComputeSandboxNewParamsAPIKeyFull
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeSandboxNewParamsAPIKeyFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ComputeSandboxNewParamsAPIKeyFull](
		"full", true,
	)
}

// The property Key is required.
type ComputeSandboxNewParamsAPIKeyKey struct {
	Key string `json:"key" api:"required"`
	paramObj
}

func (r ComputeSandboxNewParamsAPIKeyKey) MarshalJSON() (data []byte, err error) {
	type shadow ComputeSandboxNewParamsAPIKeyKey
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeSandboxNewParamsAPIKeyKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Any of "starting", "running", "stopped".
	Status ComputeSandboxListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeSandboxListParams]'s query parameters as
// `url.Values`.
func (r ComputeSandboxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ComputeSandboxListParamsStatus string

const (
	ComputeSandboxListParamsStatusStarting ComputeSandboxListParamsStatus = "starting"
	ComputeSandboxListParamsStatusRunning  ComputeSandboxListParamsStatus = "running"
	ComputeSandboxListParamsStatusStopped  ComputeSandboxListParamsStatus = "stopped"
)

type ComputeSandboxExecParams struct {
	Command  string            `json:"command" api:"required"`
	Cwd      param.Opt[string] `json:"cwd,omitzero"`
	TimeoutS param.Opt[int64]  `json:"timeoutS,omitzero"`
	WaitS    param.Opt[int64]  `json:"waitS,omitzero"`
	Env      map[string]string `json:"env,omitzero"`
	paramObj
}

func (r ComputeSandboxExecParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputeSandboxExecParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeSandboxExecParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
