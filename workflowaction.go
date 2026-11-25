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

// WorkflowActionService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowActionService] method instead.
type WorkflowActionService struct {
	options  []option.RequestOption
	Services WorkflowActionServiceService
}

// NewWorkflowActionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowActionService(opts ...option.RequestOption) (r WorkflowActionService) {
	r = WorkflowActionService{}
	r.options = opts
	r.Services = NewWorkflowActionServiceService(opts...)
	return
}

// Create a reusable action from a catalog entry (`catalogEntryId`), with an
// optional `params` object supplying the values for that entry's service method.
// Returns 400 if the params are invalid for the chosen catalog entry.
//
// Supports an optional `Idempotency-Key` header (1-255 printable ASCII
// characters). Replays with the same key and an identical body return the original
// 201; a changed body under the same key returns 422 `idempotency_key_reused`.
func (r *WorkflowActionService) New(ctx context.Context, body WorkflowActionNewParams, opts ...option.RequestOption) (res *WorkflowActionNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "actions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a single action by its ID, including its configured service, method, and
// params. Supports `profile=agent_routine`; returns 404 if no action matches or
// the action is outside that profile.
func (r *WorkflowActionService) Get(ctx context.Context, actionID string, query WorkflowActionGetParams, opts ...option.RequestOption) (res *WorkflowActionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("actions/%s", actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially update an action's name, description, or params; all fields are
// optional. Returns 404 if the action does not exist.
func (r *WorkflowActionService) Update(ctx context.Context, actionID string, body WorkflowActionUpdateParams, opts ...option.RequestOption) (res *WorkflowActionUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("actions/%s", actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Return a paginated list of actions. Supports filtering by
// `profile=agent_routine`, `service`, free-text `search`, and ordering by name,
// createdAt, or updatedAt.
func (r *WorkflowActionService) List(ctx context.Context, query WorkflowActionListParams, opts ...option.RequestOption) (res *pagination.Page[Action], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "actions"
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

// Return a paginated list of actions. Supports filtering by
// `profile=agent_routine`, `service`, free-text `search`, and ordering by name,
// createdAt, or updatedAt.
func (r *WorkflowActionService) ListAutoPaging(ctx context.Context, query WorkflowActionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Action] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete an action by its ID. Returns 404 if no action matches.
func (r *WorkflowActionService) Delete(ctx context.Context, actionID string, opts ...option.RequestOption) (res *WorkflowActionDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("actions/%s", actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type Action struct {
	ID             string    `json:"id" api:"required" format:"uuid"`
	CatalogEntryID string    `json:"catalogEntryId" api:"required" format:"uuid"`
	CreatedAt      time.Time `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy      string    `json:"createdBy" api:"required" format:"uuid"`
	Description    string    `json:"description" api:"required"`
	// Deprecated: use uses.
	//
	// Deprecated: deprecated
	Method  string `json:"method" api:"required"`
	Name    string `json:"name" api:"required"`
	OwnerID string `json:"ownerId" api:"required" format:"uuid"`
	// Deprecated: use uses.
	//
	// Any of "tasks_api", "devices_api", "agents_api", "webhooks", "integrations_api",
	// "triggers_api", "sandbox_api".
	//
	// Deprecated: deprecated
	Service   ActionService `json:"service" api:"required"`
	UpdatedAt time.Time     `json:"updatedAt" api:"required" format:"date-time"`
	// Deprecated: use ownerId (tenancy) / createdBy (actor).
	//
	// Deprecated: deprecated
	UserID string `json:"userId" api:"required" format:"uuid"`
	// Authored step name; null when this step has no name. Includes hidden executable
	// steps.
	Uses         string `json:"uses" api:"required"`
	Params       any    `json:"params"`
	ParamsSchema any    `json:"paramsSchema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CatalogEntryID respjson.Field
		CreatedAt      respjson.Field
		CreatedBy      respjson.Field
		Description    respjson.Field
		Method         respjson.Field
		Name           respjson.Field
		OwnerID        respjson.Field
		Service        respjson.Field
		UpdatedAt      respjson.Field
		UserID         respjson.Field
		Uses           respjson.Field
		Params         respjson.Field
		ParamsSchema   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Action) RawJSON() string { return r.JSON.raw }
func (r *Action) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Deprecated: use uses.
type ActionService string

const (
	ActionServiceTasksAPI        ActionService = "tasks_api"
	ActionServiceDevicesAPI      ActionService = "devices_api"
	ActionServiceAgentsAPI       ActionService = "agents_api"
	ActionServiceWebhooks        ActionService = "webhooks"
	ActionServiceIntegrationsAPI ActionService = "integrations_api"
	ActionServiceTriggersAPI     ActionService = "triggers_api"
	ActionServiceSandboxAPI      ActionService = "sandbox_api"
)

type WorkflowActionNewResponse struct {
	Data Action `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionGetResponse struct {
	Data Action `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionUpdateResponse struct {
	Data Action `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionDeleteResponse struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionNewParams struct {
	CatalogEntryID string            `json:"catalogEntryId" api:"required" format:"uuid"`
	Name           string            `json:"name" api:"required"`
	Description    param.Opt[string] `json:"description,omitzero"`
	Params         map[string]any    `json:"params,omitzero"`
	paramObj
}

func (r WorkflowActionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowActionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowActionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionGetParams struct {
	// Restrict reads to the agent routine step surface.
	//
	// Any of "agent_routine".
	Profile WorkflowActionGetParamsProfile `query:"profile,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowActionGetParams]'s query parameters as
// `url.Values`.
func (r WorkflowActionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Restrict reads to the agent routine step surface.
type WorkflowActionGetParamsProfile string

const (
	WorkflowActionGetParamsProfileAgentRoutine WorkflowActionGetParamsProfile = "agent_routine"
)

type WorkflowActionUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	Params      map[string]any    `json:"params,omitzero"`
	paramObj
}

func (r WorkflowActionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowActionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowActionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionListParams struct {
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64]  `query:"pageSize,omitzero" json:"-"`
	Search   param.Opt[string] `query:"search,omitzero" json:"-"`
	// Any of "name", "createdAt", "updatedAt".
	OrderBy WorkflowActionListParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// Any of "asc", "desc".
	OrderByDirection WorkflowActionListParamsOrderByDirection `query:"orderByDirection,omitzero" json:"-"`
	// Restrict reads to the agent routine step surface.
	//
	// Any of "agent_routine".
	Profile WorkflowActionListParamsProfile `query:"profile,omitzero" json:"-"`
	// Any of "tasks_api", "devices_api", "agents_api", "webhooks", "integrations_api",
	// "triggers_api", "sandbox_api".
	Service WorkflowActionListParamsService `query:"service,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowActionListParams]'s query parameters as
// `url.Values`.
func (r WorkflowActionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowActionListParamsOrderBy string

const (
	WorkflowActionListParamsOrderByName      WorkflowActionListParamsOrderBy = "name"
	WorkflowActionListParamsOrderByCreatedAt WorkflowActionListParamsOrderBy = "createdAt"
	WorkflowActionListParamsOrderByUpdatedAt WorkflowActionListParamsOrderBy = "updatedAt"
)

type WorkflowActionListParamsOrderByDirection string

const (
	WorkflowActionListParamsOrderByDirectionAsc  WorkflowActionListParamsOrderByDirection = "asc"
	WorkflowActionListParamsOrderByDirectionDesc WorkflowActionListParamsOrderByDirection = "desc"
)

// Restrict reads to the agent routine step surface.
type WorkflowActionListParamsProfile string

const (
	WorkflowActionListParamsProfileAgentRoutine WorkflowActionListParamsProfile = "agent_routine"
)

type WorkflowActionListParamsService string

const (
	WorkflowActionListParamsServiceTasksAPI        WorkflowActionListParamsService = "tasks_api"
	WorkflowActionListParamsServiceDevicesAPI      WorkflowActionListParamsService = "devices_api"
	WorkflowActionListParamsServiceAgentsAPI       WorkflowActionListParamsService = "agents_api"
	WorkflowActionListParamsServiceWebhooks        WorkflowActionListParamsService = "webhooks"
	WorkflowActionListParamsServiceIntegrationsAPI WorkflowActionListParamsService = "integrations_api"
	WorkflowActionListParamsServiceTriggersAPI     WorkflowActionListParamsService = "triggers_api"
	WorkflowActionListParamsServiceSandboxAPI      WorkflowActionListParamsService = "sandbox_api"
)
