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
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
	"github.com/droidrun/mobilerun-sdk-go/v6/shared"
)

// WorkflowActionCatalogService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowActionCatalogService] method instead.
type WorkflowActionCatalogService struct {
	options []option.RequestOption
}

// NewWorkflowActionCatalogService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkflowActionCatalogService(opts ...option.RequestOption) (r WorkflowActionCatalogService) {
	r = WorkflowActionCatalogService{}
	r.options = opts
	return
}

// Fetch a single action catalog entry by its ID, including its public `stepName`
// for routine documents, service, method, parameter schema, and `deviceFree`
// eligibility. Supports `profile=agent_routine`; returns 404 if no entry matches
// or the entry is outside that profile.
func (r *WorkflowActionCatalogService) Get(ctx context.Context, catalogEntryID string, query WorkflowActionCatalogGetParams, opts ...option.RequestOption) (res *WorkflowActionCatalogGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if catalogEntryID == "" {
		err = errors.New("missing required catalogEntryId parameter")
		return nil, err
	}
	path := fmt.Sprintf("action-catalog/%s", catalogEntryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return a paginated list of catalog entries — the service/method templates that
// actions are created from, each carrying its public `stepName` for routine
// documents, parameter schema and `deviceFree` eligibility. Supports filtering by
// `service` and `profile=agent_routine`.
func (r *WorkflowActionCatalogService) List(ctx context.Context, query WorkflowActionCatalogListParams, opts ...option.RequestOption) (res *WorkflowActionCatalogListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "action-catalog"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type WorkflowActionCatalogGetResponse struct {
	Data WorkflowActionCatalogGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionCatalogGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionCatalogGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionCatalogGetResponseData struct {
	ID          string    `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time `json:"createdAt" api:"required" format:"date-time"`
	Description string    `json:"description" api:"required"`
	// True when this step itself runs without a device; only such steps may appear in
	// a device-free flow. Not the same as cloudRoutine: a device-free step can still
	// be device-routine-only (see cloudRoutineReason).
	DeviceFree bool `json:"deviceFree" api:"required"`
	// Deprecated: use stepName.
	//
	// Deprecated: deprecated
	Method string `json:"method" api:"required"`
	Name   string `json:"name" api:"required"`
	// Any of "mobilerun", "oneDrive", "googleDrive".
	Provider string `json:"provider" api:"required"`
	// Deprecated: use stepName.
	//
	// Any of "tasks_api", "devices_api", "agents_api", "webhooks", "integrations_api",
	// "triggers_api", "sandbox_api".
	//
	// Deprecated: deprecated
	Service string `json:"service" api:"required"`
	// Public step name for routine document uses, e.g. files.upload or scripts.run.
	StepName  string    `json:"stepName" api:"required"`
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// True when this step is allowed in a cloud agent routine. Present only for
	// profile=agent_routine. False does not mean the step needs a device; read
	// cloudRoutineReason.
	CloudRoutine bool `json:"cloudRoutine"`
	// Why cloudRoutine is false: needs_device (the step needs a device),
	// device_routine_only (allowed only in device routines; the step itself does not
	// need a device). Present only when cloudRoutine is false.
	//
	// Any of "needs_device", "device_routine_only".
	CloudRoutineReason string `json:"cloudRoutineReason"`
	OutputSchema       any    `json:"outputSchema"`
	ParamsSchema       any    `json:"paramsSchema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		DeviceFree         respjson.Field
		Method             respjson.Field
		Name               respjson.Field
		Provider           respjson.Field
		Service            respjson.Field
		StepName           respjson.Field
		UpdatedAt          respjson.Field
		CloudRoutine       respjson.Field
		CloudRoutineReason respjson.Field
		OutputSchema       respjson.Field
		ParamsSchema       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionCatalogGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionCatalogGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionCatalogListResponse struct {
	Items      []WorkflowActionCatalogListResponseItem `json:"items" api:"required"`
	Pagination shared.Pagination                       `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionCatalogListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionCatalogListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionCatalogListResponseItem struct {
	ID          string    `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time `json:"createdAt" api:"required" format:"date-time"`
	Description string    `json:"description" api:"required"`
	// True when this step itself runs without a device; only such steps may appear in
	// a device-free flow. Not the same as cloudRoutine: a device-free step can still
	// be device-routine-only (see cloudRoutineReason).
	DeviceFree bool `json:"deviceFree" api:"required"`
	// Deprecated: use stepName.
	//
	// Deprecated: deprecated
	Method string `json:"method" api:"required"`
	Name   string `json:"name" api:"required"`
	// Any of "mobilerun", "oneDrive", "googleDrive".
	Provider string `json:"provider" api:"required"`
	// Deprecated: use stepName.
	//
	// Any of "tasks_api", "devices_api", "agents_api", "webhooks", "integrations_api",
	// "triggers_api", "sandbox_api".
	//
	// Deprecated: deprecated
	Service string `json:"service" api:"required"`
	// Public step name for routine document uses, e.g. files.upload or scripts.run.
	StepName  string    `json:"stepName" api:"required"`
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// True when this step is allowed in a cloud agent routine. Present only for
	// profile=agent_routine. False does not mean the step needs a device; read
	// cloudRoutineReason.
	CloudRoutine bool `json:"cloudRoutine"`
	// Why cloudRoutine is false: needs_device (the step needs a device),
	// device_routine_only (allowed only in device routines; the step itself does not
	// need a device). Present only when cloudRoutine is false.
	//
	// Any of "needs_device", "device_routine_only".
	CloudRoutineReason string `json:"cloudRoutineReason"`
	OutputSchema       any    `json:"outputSchema"`
	ParamsSchema       any    `json:"paramsSchema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		DeviceFree         respjson.Field
		Method             respjson.Field
		Name               respjson.Field
		Provider           respjson.Field
		Service            respjson.Field
		StepName           respjson.Field
		UpdatedAt          respjson.Field
		CloudRoutine       respjson.Field
		CloudRoutineReason respjson.Field
		OutputSchema       respjson.Field
		ParamsSchema       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionCatalogListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionCatalogListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionCatalogGetParams struct {
	// Restrict reads to the agent routine step surface.
	//
	// Any of "agent_routine".
	Profile WorkflowActionCatalogGetParamsProfile `query:"profile,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowActionCatalogGetParams]'s query parameters as
// `url.Values`.
func (r WorkflowActionCatalogGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Restrict reads to the agent routine step surface.
type WorkflowActionCatalogGetParamsProfile string

const (
	WorkflowActionCatalogGetParamsProfileAgentRoutine WorkflowActionCatalogGetParamsProfile = "agent_routine"
)

type WorkflowActionCatalogListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Restrict reads to the agent routine step surface.
	//
	// Any of "agent_routine".
	Profile WorkflowActionCatalogListParamsProfile `query:"profile,omitzero" json:"-"`
	// Any of "tasks_api", "devices_api", "agents_api", "webhooks", "integrations_api",
	// "triggers_api", "sandbox_api".
	Service WorkflowActionCatalogListParamsService `query:"service,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowActionCatalogListParams]'s query parameters as
// `url.Values`.
func (r WorkflowActionCatalogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Restrict reads to the agent routine step surface.
type WorkflowActionCatalogListParamsProfile string

const (
	WorkflowActionCatalogListParamsProfileAgentRoutine WorkflowActionCatalogListParamsProfile = "agent_routine"
)

type WorkflowActionCatalogListParamsService string

const (
	WorkflowActionCatalogListParamsServiceTasksAPI        WorkflowActionCatalogListParamsService = "tasks_api"
	WorkflowActionCatalogListParamsServiceDevicesAPI      WorkflowActionCatalogListParamsService = "devices_api"
	WorkflowActionCatalogListParamsServiceAgentsAPI       WorkflowActionCatalogListParamsService = "agents_api"
	WorkflowActionCatalogListParamsServiceWebhooks        WorkflowActionCatalogListParamsService = "webhooks"
	WorkflowActionCatalogListParamsServiceIntegrationsAPI WorkflowActionCatalogListParamsService = "integrations_api"
	WorkflowActionCatalogListParamsServiceTriggersAPI     WorkflowActionCatalogListParamsService = "triggers_api"
	WorkflowActionCatalogListParamsServiceSandboxAPI      WorkflowActionCatalogListParamsService = "sandbox_api"
)
