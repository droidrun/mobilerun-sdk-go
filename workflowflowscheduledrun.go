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

// WorkflowFlowScheduledRunService contains methods and other services that help
// with interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowFlowScheduledRunService] method instead.
type WorkflowFlowScheduledRunService struct {
	options []option.RequestOption
}

// NewWorkflowFlowScheduledRunService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWorkflowFlowScheduledRunService(opts ...option.RequestOption) (r WorkflowFlowScheduledRunService) {
	r = WorkflowFlowScheduledRunService{}
	r.options = opts
	return
}

// Plan Run now for a future time, using the same devices, payload rules and
// admission at fire time. A flow can have at most 100 scheduled runs. Supports an
// optional Idempotency-Key; a replay returns the current scheduled run with the
// original 201 status.
func (r *WorkflowFlowScheduledRunService) New(ctx context.Context, flowID string, body WorkflowFlowScheduledRunNewParams, opts ...option.RequestOption) (res *WorkflowFlowScheduledRunNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/scheduled-runs", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists upcoming runs by default, including claimed runs awaiting dispatch. Set
// status=all to include terminal rows. Ordered by runAt ascending.
func (r *WorkflowFlowScheduledRunService) List(ctx context.Context, flowID string, query WorkflowFlowScheduledRunListParams, opts ...option.RequestOption) (res *pagination.Page[ScheduledRun], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/scheduled-runs", flowID)
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

// Lists upcoming runs by default, including claimed runs awaiting dispatch. Set
// status=all to include terminal rows. Ordered by runAt ascending.
func (r *WorkflowFlowScheduledRunService) ListAutoPaging(ctx context.Context, flowID string, query WorkflowFlowScheduledRunListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ScheduledRun] {
	return pagination.NewPageAutoPager(r.List(ctx, flowID, query, opts...))
}

// Cancel a planned flow run
func (r *WorkflowFlowScheduledRunService) Cancel(ctx context.Context, scheduledRunID string, body WorkflowFlowScheduledRunCancelParams, opts ...option.RequestOption) (res *WorkflowFlowScheduledRunCancelResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.FlowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	if scheduledRunID == "" {
		err = errors.New("missing required scheduledRunId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/scheduled-runs/%s/cancel", body.FlowID, scheduledRunID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ScheduledRun struct {
	ID           string         `json:"id" api:"required" format:"uuid"`
	CancelledAt  time.Time      `json:"cancelledAt" api:"required" format:"date-time"`
	CreatedAt    time.Time      `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy    string         `json:"createdBy" api:"required" format:"uuid"`
	Error        string         `json:"error" api:"required"`
	ExecutionIDs []string       `json:"executionIds" api:"required" format:"uuid"`
	FinishedAt   time.Time      `json:"finishedAt" api:"required" format:"date-time"`
	FlowID       string         `json:"flowId" api:"required" format:"uuid"`
	Payload      map[string]any `json:"payload" api:"required"`
	RunAt        time.Time      `json:"runAt" api:"required" format:"date-time"`
	SkipReason   string         `json:"skipReason" api:"required"`
	// Any of "scheduled", "started", "skipped", "failed", "missed", "cancelled".
	Status ScheduledRunStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CancelledAt  respjson.Field
		CreatedAt    respjson.Field
		CreatedBy    respjson.Field
		Error        respjson.Field
		ExecutionIDs respjson.Field
		FinishedAt   respjson.Field
		FlowID       respjson.Field
		Payload      respjson.Field
		RunAt        respjson.Field
		SkipReason   respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduledRun) RawJSON() string { return r.JSON.raw }
func (r *ScheduledRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScheduledRunStatus string

const (
	ScheduledRunStatusScheduled ScheduledRunStatus = "scheduled"
	ScheduledRunStatusStarted   ScheduledRunStatus = "started"
	ScheduledRunStatusSkipped   ScheduledRunStatus = "skipped"
	ScheduledRunStatusFailed    ScheduledRunStatus = "failed"
	ScheduledRunStatusMissed    ScheduledRunStatus = "missed"
	ScheduledRunStatusCancelled ScheduledRunStatus = "cancelled"
)

type WorkflowFlowScheduledRunNewResponse struct {
	Data ScheduledRun `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowScheduledRunNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowScheduledRunNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowScheduledRunCancelResponse struct {
	Data ScheduledRun `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowScheduledRunCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowScheduledRunCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowScheduledRunNewParams struct {
	// RFC 3339 time with an offset; between one minute and 365 days from now.
	RunAt   time.Time      `json:"runAt" api:"required" format:"date-time"`
	Payload map[string]any `json:"payload,omitzero"`
	paramObj
}

func (r WorkflowFlowScheduledRunNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowScheduledRunNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowScheduledRunNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowScheduledRunListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Any of "scheduled", "all".
	Status WorkflowFlowScheduledRunListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowFlowScheduledRunListParams]'s query parameters as
// `url.Values`.
func (r WorkflowFlowScheduledRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowFlowScheduledRunListParamsStatus string

const (
	WorkflowFlowScheduledRunListParamsStatusScheduled WorkflowFlowScheduledRunListParamsStatus = "scheduled"
	WorkflowFlowScheduledRunListParamsStatusAll       WorkflowFlowScheduledRunListParamsStatus = "all"
)

type WorkflowFlowScheduledRunCancelParams struct {
	FlowID string `path:"flowId" api:"required" format:"uuid" json:"-"`
	paramObj
}
