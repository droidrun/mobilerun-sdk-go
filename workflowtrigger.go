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

// WorkflowTriggerService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowTriggerService] method instead.
type WorkflowTriggerService struct {
	options []option.RequestOption
}

// NewWorkflowTriggerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowTriggerService(opts ...option.RequestOption) (r WorkflowTriggerService) {
	r = WorkflowTriggerService{}
	r.options = opts
	return
}

// Create a trigger with an activation type of `event`, `schedule`, or `custom`.
// Each type requires its own fields (e.g. `eventType` and optional `conditions`
// for events, `scheduleRule` and `timezone` for schedules, `customPayloadSchema`
// for custom or schedule triggers); mismatched fields are rejected. Schedule
// triggers may store an optional JSON object `payload` (at most 1 MiB); it
// requires `customPayloadSchema`, is validated against it (422 on mismatch), and
// is available as `event.payload` on every scheduled, manual, or verification run.
// When a schema is configured, omitted payload values are validated as an empty
// object.
//
// Supports an optional `Idempotency-Key` header (1-255 printable ASCII
// characters). Replays with the same key and an identical body return the original
// 201; a changed body under the same key returns 422 `idempotency_key_reused`.
func (r *WorkflowTriggerService) New(ctx context.Context, body WorkflowTriggerNewParams, opts ...option.RequestOption) (res *WorkflowTriggerNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "triggers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a single trigger by its ID, including its activation type and
// type-specific configuration. Returns 404 if no trigger matches.
func (r *WorkflowTriggerService) Get(ctx context.Context, triggerID string, opts ...option.RequestOption) (res *WorkflowTriggerGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if triggerID == "" {
		err = errors.New("missing required triggerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("triggers/%s", triggerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a trigger; all fields are optional. When `activation` is
// changed, the type-specific field rules are re-validated. Schedule triggers
// accept stored `payload` values (at most 1 MiB); null clears them. The resulting
// payload must match customPayloadSchema (422 on mismatch, including schema
// changes that invalidate stored values). Absent or null values are validated as
// an empty object when a schema is configured. Changes apply from the next run.
// Returns 404 if the trigger does not exist.
func (r *WorkflowTriggerService) Update(ctx context.Context, triggerID string, body WorkflowTriggerUpdateParams, opts ...option.RequestOption) (res *WorkflowTriggerUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if triggerID == "" {
		err = errors.New("missing required triggerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("triggers/%s", triggerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Return a paginated list of triggers. Supports filtering by `activation` and
// `eventType`, free-text `search`, and ordering by name, createdAt, or updatedAt.
func (r *WorkflowTriggerService) List(ctx context.Context, query WorkflowTriggerListParams, opts ...option.RequestOption) (res *WorkflowTriggerListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "triggers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a trigger by its ID. Returns 404 if no trigger matches.
func (r *WorkflowTriggerService) Delete(ctx context.Context, triggerID string, opts ...option.RequestOption) (res *WorkflowTriggerDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if triggerID == "" {
		err = errors.New("missing required triggerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("triggers/%s", triggerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Invoke a custom trigger directly with an arbitrary JSON payload.
//
// Fan-out: a trigger may be referenced by multiple flows (workflows). Firing it
// enqueues one execution per enabled, non-deleted flow attached to this trigger,
// each receiving the same payload. The `enqueuedCount` in the response reports how
// many were enqueued (0 if no flows are attached, or if all matching flows are
// gated by a cooldown).
//
// Payload validation:
//
//   - If the trigger has a `customPayloadSchema`, the payload is validated against
//     it (JSON Schema).
//   - If no schema is configured, the payload only needs to be a JSON object — any
//     keys and values are accepted.
//
// Only triggers with `activation = "custom"` can be fired through this endpoint;
// event and schedule triggers return 409.
//
// Supports an optional `Idempotency-Key` header (1-255 printable ASCII
// characters), which maps to the derived `invocationId` used for fan-out
// deduplication. The payload is bound to the key on the first accepted fire,
// before fan-out. A repeat with the same key and an identical payload returns 202
// with `Idempotent-Replayed: true` (`deduplicated: true` when flows were skipped;
// `enqueuedCount` counts only executions enqueued by that call); a changed payload
// under the same key returns 422 `idempotency_key_reused`.
func (r *WorkflowTriggerService) Fire(ctx context.Context, triggerID string, body WorkflowTriggerFireParams, opts ...option.RequestOption) (res *WorkflowTriggerFireResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if triggerID == "" {
		err = errors.New("missing required triggerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("triggers/%s/fire", triggerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type WorkflowTriggerNewResponse struct {
	Data WorkflowTriggerNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerNewResponseData struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "event", "schedule", "custom".
	Activation          string                                     `json:"activation" api:"required"`
	CreatedAt           time.Time                                  `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy           string                                     `json:"createdBy" api:"required" format:"uuid"`
	CustomPayloadSchema map[string]any                             `json:"customPayloadSchema" api:"required"`
	Description         string                                     `json:"description" api:"required"`
	EventType           string                                     `json:"eventType" api:"required"`
	Name                string                                     `json:"name" api:"required"`
	OwnerID             string                                     `json:"ownerId" api:"required" format:"uuid"`
	Payload             map[string]any                             `json:"payload" api:"required"`
	ScheduleRule        WorkflowTriggerNewResponseDataScheduleRule `json:"scheduleRule" api:"required"`
	Timezone            string                                     `json:"timezone" api:"required"`
	UpdatedAt           time.Time                                  `json:"updatedAt" api:"required" format:"date-time"`
	// Deprecated: use ownerId (tenancy) / createdBy (actor).
	//
	// Deprecated: deprecated
	UserID       string `json:"userId" api:"required" format:"uuid"`
	Conditions   any    `json:"conditions"`
	NextFireTime string `json:"nextFireTime" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Activation          respjson.Field
		CreatedAt           respjson.Field
		CreatedBy           respjson.Field
		CustomPayloadSchema respjson.Field
		Description         respjson.Field
		EventType           respjson.Field
		Name                respjson.Field
		OwnerID             respjson.Field
		Payload             respjson.Field
		ScheduleRule        respjson.Field
		Timezone            respjson.Field
		UpdatedAt           respjson.Field
		UserID              respjson.Field
		Conditions          respjson.Field
		NextFireTime        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerNewResponseDataScheduleRule struct {
	// Any of "once", "cron", "recurring".
	Type string `json:"type" api:"required"`
	// ISO 8601 datetime (for type=once)
	DateTime string `json:"dateTime"`
	// Cron expression (for type=cron)
	Expression string `json:"expression"`
	// Optional per-occurrence random window around the nominal schedule time
	Jitter WorkflowTriggerNewResponseDataScheduleRuleJitter `json:"jitter"`
	// RRULE string (for type=recurring)
	Rrule string `json:"rrule"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		DateTime    respjson.Field
		Expression  respjson.Field
		Jitter      respjson.Field
		Rrule       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerNewResponseDataScheduleRule) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerNewResponseDataScheduleRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional per-occurrence random window around the nominal schedule time
type WorkflowTriggerNewResponseDataScheduleRuleJitter struct {
	AfterMinutes  int64 `json:"afterMinutes"`
	BeforeMinutes int64 `json:"beforeMinutes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AfterMinutes  respjson.Field
		BeforeMinutes respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerNewResponseDataScheduleRuleJitter) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerNewResponseDataScheduleRuleJitter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerGetResponse struct {
	Data WorkflowTriggerGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerGetResponseData struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "event", "schedule", "custom".
	Activation          string                                     `json:"activation" api:"required"`
	CreatedAt           time.Time                                  `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy           string                                     `json:"createdBy" api:"required" format:"uuid"`
	CustomPayloadSchema map[string]any                             `json:"customPayloadSchema" api:"required"`
	Description         string                                     `json:"description" api:"required"`
	EventType           string                                     `json:"eventType" api:"required"`
	Name                string                                     `json:"name" api:"required"`
	OwnerID             string                                     `json:"ownerId" api:"required" format:"uuid"`
	Payload             map[string]any                             `json:"payload" api:"required"`
	ScheduleRule        WorkflowTriggerGetResponseDataScheduleRule `json:"scheduleRule" api:"required"`
	Timezone            string                                     `json:"timezone" api:"required"`
	UpdatedAt           time.Time                                  `json:"updatedAt" api:"required" format:"date-time"`
	// Deprecated: use ownerId (tenancy) / createdBy (actor).
	//
	// Deprecated: deprecated
	UserID       string `json:"userId" api:"required" format:"uuid"`
	Conditions   any    `json:"conditions"`
	NextFireTime string `json:"nextFireTime" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Activation          respjson.Field
		CreatedAt           respjson.Field
		CreatedBy           respjson.Field
		CustomPayloadSchema respjson.Field
		Description         respjson.Field
		EventType           respjson.Field
		Name                respjson.Field
		OwnerID             respjson.Field
		Payload             respjson.Field
		ScheduleRule        respjson.Field
		Timezone            respjson.Field
		UpdatedAt           respjson.Field
		UserID              respjson.Field
		Conditions          respjson.Field
		NextFireTime        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerGetResponseDataScheduleRule struct {
	// Any of "once", "cron", "recurring".
	Type string `json:"type" api:"required"`
	// ISO 8601 datetime (for type=once)
	DateTime string `json:"dateTime"`
	// Cron expression (for type=cron)
	Expression string `json:"expression"`
	// Optional per-occurrence random window around the nominal schedule time
	Jitter WorkflowTriggerGetResponseDataScheduleRuleJitter `json:"jitter"`
	// RRULE string (for type=recurring)
	Rrule string `json:"rrule"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		DateTime    respjson.Field
		Expression  respjson.Field
		Jitter      respjson.Field
		Rrule       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerGetResponseDataScheduleRule) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerGetResponseDataScheduleRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional per-occurrence random window around the nominal schedule time
type WorkflowTriggerGetResponseDataScheduleRuleJitter struct {
	AfterMinutes  int64 `json:"afterMinutes"`
	BeforeMinutes int64 `json:"beforeMinutes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AfterMinutes  respjson.Field
		BeforeMinutes respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerGetResponseDataScheduleRuleJitter) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerGetResponseDataScheduleRuleJitter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerUpdateResponse struct {
	Data WorkflowTriggerUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerUpdateResponseData struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "event", "schedule", "custom".
	Activation          string                                        `json:"activation" api:"required"`
	CreatedAt           time.Time                                     `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy           string                                        `json:"createdBy" api:"required" format:"uuid"`
	CustomPayloadSchema map[string]any                                `json:"customPayloadSchema" api:"required"`
	Description         string                                        `json:"description" api:"required"`
	EventType           string                                        `json:"eventType" api:"required"`
	Name                string                                        `json:"name" api:"required"`
	OwnerID             string                                        `json:"ownerId" api:"required" format:"uuid"`
	Payload             map[string]any                                `json:"payload" api:"required"`
	ScheduleRule        WorkflowTriggerUpdateResponseDataScheduleRule `json:"scheduleRule" api:"required"`
	Timezone            string                                        `json:"timezone" api:"required"`
	UpdatedAt           time.Time                                     `json:"updatedAt" api:"required" format:"date-time"`
	// Deprecated: use ownerId (tenancy) / createdBy (actor).
	//
	// Deprecated: deprecated
	UserID       string `json:"userId" api:"required" format:"uuid"`
	Conditions   any    `json:"conditions"`
	NextFireTime string `json:"nextFireTime" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Activation          respjson.Field
		CreatedAt           respjson.Field
		CreatedBy           respjson.Field
		CustomPayloadSchema respjson.Field
		Description         respjson.Field
		EventType           respjson.Field
		Name                respjson.Field
		OwnerID             respjson.Field
		Payload             respjson.Field
		ScheduleRule        respjson.Field
		Timezone            respjson.Field
		UpdatedAt           respjson.Field
		UserID              respjson.Field
		Conditions          respjson.Field
		NextFireTime        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerUpdateResponseDataScheduleRule struct {
	// Any of "once", "cron", "recurring".
	Type string `json:"type" api:"required"`
	// ISO 8601 datetime (for type=once)
	DateTime string `json:"dateTime"`
	// Cron expression (for type=cron)
	Expression string `json:"expression"`
	// Optional per-occurrence random window around the nominal schedule time
	Jitter WorkflowTriggerUpdateResponseDataScheduleRuleJitter `json:"jitter"`
	// RRULE string (for type=recurring)
	Rrule string `json:"rrule"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		DateTime    respjson.Field
		Expression  respjson.Field
		Jitter      respjson.Field
		Rrule       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerUpdateResponseDataScheduleRule) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerUpdateResponseDataScheduleRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional per-occurrence random window around the nominal schedule time
type WorkflowTriggerUpdateResponseDataScheduleRuleJitter struct {
	AfterMinutes  int64 `json:"afterMinutes"`
	BeforeMinutes int64 `json:"beforeMinutes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AfterMinutes  respjson.Field
		BeforeMinutes respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerUpdateResponseDataScheduleRuleJitter) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerUpdateResponseDataScheduleRuleJitter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerListResponse struct {
	Items      []WorkflowTriggerListResponseItem `json:"items" api:"required"`
	Pagination shared.Pagination                 `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerListResponseItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "event", "schedule", "custom".
	Activation          string                                      `json:"activation" api:"required"`
	CreatedAt           time.Time                                   `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy           string                                      `json:"createdBy" api:"required" format:"uuid"`
	CustomPayloadSchema map[string]any                              `json:"customPayloadSchema" api:"required"`
	Description         string                                      `json:"description" api:"required"`
	EventType           string                                      `json:"eventType" api:"required"`
	Name                string                                      `json:"name" api:"required"`
	OwnerID             string                                      `json:"ownerId" api:"required" format:"uuid"`
	Payload             map[string]any                              `json:"payload" api:"required"`
	ScheduleRule        WorkflowTriggerListResponseItemScheduleRule `json:"scheduleRule" api:"required"`
	Timezone            string                                      `json:"timezone" api:"required"`
	UpdatedAt           time.Time                                   `json:"updatedAt" api:"required" format:"date-time"`
	// Deprecated: use ownerId (tenancy) / createdBy (actor).
	//
	// Deprecated: deprecated
	UserID       string `json:"userId" api:"required" format:"uuid"`
	Conditions   any    `json:"conditions"`
	NextFireTime string `json:"nextFireTime" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Activation          respjson.Field
		CreatedAt           respjson.Field
		CreatedBy           respjson.Field
		CustomPayloadSchema respjson.Field
		Description         respjson.Field
		EventType           respjson.Field
		Name                respjson.Field
		OwnerID             respjson.Field
		Payload             respjson.Field
		ScheduleRule        respjson.Field
		Timezone            respjson.Field
		UpdatedAt           respjson.Field
		UserID              respjson.Field
		Conditions          respjson.Field
		NextFireTime        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerListResponseItemScheduleRule struct {
	// Any of "once", "cron", "recurring".
	Type string `json:"type" api:"required"`
	// ISO 8601 datetime (for type=once)
	DateTime string `json:"dateTime"`
	// Cron expression (for type=cron)
	Expression string `json:"expression"`
	// Optional per-occurrence random window around the nominal schedule time
	Jitter WorkflowTriggerListResponseItemScheduleRuleJitter `json:"jitter"`
	// RRULE string (for type=recurring)
	Rrule string `json:"rrule"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		DateTime    respjson.Field
		Expression  respjson.Field
		Jitter      respjson.Field
		Rrule       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerListResponseItemScheduleRule) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerListResponseItemScheduleRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional per-occurrence random window around the nominal schedule time
type WorkflowTriggerListResponseItemScheduleRuleJitter struct {
	AfterMinutes  int64 `json:"afterMinutes"`
	BeforeMinutes int64 `json:"beforeMinutes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AfterMinutes  respjson.Field
		BeforeMinutes respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerListResponseItemScheduleRuleJitter) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerListResponseItemScheduleRuleJitter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerDeleteResponse struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerFireResponse struct {
	// True when at least one attached flow was skipped because it already had an
	// execution for this (flow, invocationId). Only ever true when the request carried
	// an `Idempotency-Key`.
	Deduplicated bool `json:"deduplicated" api:"required"`
	// Number of flow executions enqueued. May be 0 if no flows are attached to this
	// trigger, if all attached flows are currently in cooldown, or if every attached
	// flow was deduplicated.
	EnqueuedCount int64 `json:"enqueuedCount" api:"required"`
	// Unique ID for this fire invocation (derived from the `Idempotency-Key` header
	// when one is sent, otherwise generated).
	InvocationID string `json:"invocationId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deduplicated  respjson.Field
		EnqueuedCount respjson.Field
		InvocationID  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerFireResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerFireResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerNewParams struct {
	// Any of "event", "schedule", "custom".
	Activation  WorkflowTriggerNewParamsActivation `json:"activation,omitzero" api:"required"`
	Name        string                             `json:"name" api:"required"`
	Description param.Opt[string]                  `json:"description,omitzero"`
	EventType   param.Opt[string]                  `json:"eventType,omitzero"`
	Timezone    param.Opt[string]                  `json:"timezone,omitzero"`
	Conditions  WorkflowTriggerNewParamsConditions `json:"conditions,omitzero"`
	// Optional JSON Schema for validating payloads for custom or schedule triggers
	CustomPayloadSchema map[string]any `json:"customPayloadSchema,omitzero"`
	// Stored input values for schedule triggers (maximum 1 MiB of JSON). Requires
	// customPayloadSchema; available as event.payload on every run.
	Payload      map[string]any                       `json:"payload,omitzero"`
	ScheduleRule WorkflowTriggerNewParamsScheduleRule `json:"scheduleRule,omitzero"`
	paramObj
}

func (r WorkflowTriggerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTriggerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerNewParamsActivation string

const (
	WorkflowTriggerNewParamsActivationEvent    WorkflowTriggerNewParamsActivation = "event"
	WorkflowTriggerNewParamsActivationSchedule WorkflowTriggerNewParamsActivation = "schedule"
	WorkflowTriggerNewParamsActivationCustom   WorkflowTriggerNewParamsActivation = "custom"
)

type WorkflowTriggerNewParamsConditions struct {
	All []any `json:"all,omitzero"`
	Any []any `json:"any,omitzero"`
	paramObj
}

func (r WorkflowTriggerNewParamsConditions) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerNewParamsConditions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTriggerNewParamsConditions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type WorkflowTriggerNewParamsScheduleRule struct {
	// Any of "once", "cron", "recurring".
	Type string `json:"type,omitzero" api:"required"`
	// ISO 8601 datetime (for type=once)
	DateTime param.Opt[string] `json:"dateTime,omitzero"`
	// Cron expression (for type=cron)
	Expression param.Opt[string] `json:"expression,omitzero"`
	// RRULE string (for type=recurring)
	Rrule param.Opt[string] `json:"rrule,omitzero"`
	// Optional per-occurrence random window around the nominal schedule time
	Jitter WorkflowTriggerNewParamsScheduleRuleJitter `json:"jitter,omitzero"`
	paramObj
}

func (r WorkflowTriggerNewParamsScheduleRule) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerNewParamsScheduleRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTriggerNewParamsScheduleRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowTriggerNewParamsScheduleRule](
		"type", "once", "cron", "recurring",
	)
}

// Optional per-occurrence random window around the nominal schedule time
type WorkflowTriggerNewParamsScheduleRuleJitter struct {
	AfterMinutes  param.Opt[int64] `json:"afterMinutes,omitzero"`
	BeforeMinutes param.Opt[int64] `json:"beforeMinutes,omitzero"`
	paramObj
}

func (r WorkflowTriggerNewParamsScheduleRuleJitter) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerNewParamsScheduleRuleJitter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTriggerNewParamsScheduleRuleJitter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerUpdateParams struct {
	Timezone    param.Opt[string]                     `json:"timezone,omitzero"`
	Description param.Opt[string]                     `json:"description,omitzero"`
	EventType   param.Opt[string]                     `json:"eventType,omitzero"`
	Name        param.Opt[string]                     `json:"name,omitzero"`
	Conditions  WorkflowTriggerUpdateParamsConditions `json:"conditions,omitzero"`
	// Optional JSON Schema for validating payloads for custom or schedule triggers
	CustomPayloadSchema map[string]any `json:"customPayloadSchema,omitzero"`
	// Stored schedule input values (maximum 1 MiB of JSON); null clears them. Requires
	// customPayloadSchema.
	Payload      map[string]any                          `json:"payload,omitzero"`
	ScheduleRule WorkflowTriggerUpdateParamsScheduleRule `json:"scheduleRule,omitzero"`
	// Any of "event", "schedule", "custom".
	Activation WorkflowTriggerUpdateParamsActivation `json:"activation,omitzero"`
	paramObj
}

func (r WorkflowTriggerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTriggerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerUpdateParamsActivation string

const (
	WorkflowTriggerUpdateParamsActivationEvent    WorkflowTriggerUpdateParamsActivation = "event"
	WorkflowTriggerUpdateParamsActivationSchedule WorkflowTriggerUpdateParamsActivation = "schedule"
	WorkflowTriggerUpdateParamsActivationCustom   WorkflowTriggerUpdateParamsActivation = "custom"
)

type WorkflowTriggerUpdateParamsConditions struct {
	All []any `json:"all,omitzero"`
	Any []any `json:"any,omitzero"`
	paramObj
}

func (r WorkflowTriggerUpdateParamsConditions) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerUpdateParamsConditions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTriggerUpdateParamsConditions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type WorkflowTriggerUpdateParamsScheduleRule struct {
	// Any of "once", "cron", "recurring".
	Type string `json:"type,omitzero" api:"required"`
	// ISO 8601 datetime (for type=once)
	DateTime param.Opt[string] `json:"dateTime,omitzero"`
	// Cron expression (for type=cron)
	Expression param.Opt[string] `json:"expression,omitzero"`
	// RRULE string (for type=recurring)
	Rrule param.Opt[string] `json:"rrule,omitzero"`
	// Optional per-occurrence random window around the nominal schedule time
	Jitter WorkflowTriggerUpdateParamsScheduleRuleJitter `json:"jitter,omitzero"`
	paramObj
}

func (r WorkflowTriggerUpdateParamsScheduleRule) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerUpdateParamsScheduleRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTriggerUpdateParamsScheduleRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowTriggerUpdateParamsScheduleRule](
		"type", "once", "cron", "recurring",
	)
}

// Optional per-occurrence random window around the nominal schedule time
type WorkflowTriggerUpdateParamsScheduleRuleJitter struct {
	AfterMinutes  param.Opt[int64] `json:"afterMinutes,omitzero"`
	BeforeMinutes param.Opt[int64] `json:"beforeMinutes,omitzero"`
	paramObj
}

func (r WorkflowTriggerUpdateParamsScheduleRuleJitter) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerUpdateParamsScheduleRuleJitter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTriggerUpdateParamsScheduleRuleJitter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerListParams struct {
	EventType param.Opt[string] `query:"eventType,omitzero" json:"-"`
	Page      param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize  param.Opt[int64]  `query:"pageSize,omitzero" json:"-"`
	Search    param.Opt[string] `query:"search,omitzero" json:"-"`
	// Any of "event", "schedule", "custom".
	Activation WorkflowTriggerListParamsActivation `query:"activation,omitzero" json:"-"`
	// Any of "name", "createdAt", "updatedAt".
	OrderBy WorkflowTriggerListParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// Any of "asc", "desc".
	OrderByDirection WorkflowTriggerListParamsOrderByDirection `query:"orderByDirection,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowTriggerListParams]'s query parameters as
// `url.Values`.
func (r WorkflowTriggerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowTriggerListParamsActivation string

const (
	WorkflowTriggerListParamsActivationEvent    WorkflowTriggerListParamsActivation = "event"
	WorkflowTriggerListParamsActivationSchedule WorkflowTriggerListParamsActivation = "schedule"
	WorkflowTriggerListParamsActivationCustom   WorkflowTriggerListParamsActivation = "custom"
)

type WorkflowTriggerListParamsOrderBy string

const (
	WorkflowTriggerListParamsOrderByName      WorkflowTriggerListParamsOrderBy = "name"
	WorkflowTriggerListParamsOrderByCreatedAt WorkflowTriggerListParamsOrderBy = "createdAt"
	WorkflowTriggerListParamsOrderByUpdatedAt WorkflowTriggerListParamsOrderBy = "updatedAt"
)

type WorkflowTriggerListParamsOrderByDirection string

const (
	WorkflowTriggerListParamsOrderByDirectionAsc  WorkflowTriggerListParamsOrderByDirection = "asc"
	WorkflowTriggerListParamsOrderByDirectionDesc WorkflowTriggerListParamsOrderByDirection = "desc"
)

type WorkflowTriggerFireParams struct {
	// Arbitrary JSON object forwarded to every flow attached to this trigger.
	// Validated against the trigger's customPayloadSchema when one is configured;
	// otherwise only "must be a JSON object" is enforced.
	Payload map[string]any `json:"payload,omitzero" api:"required"`
	paramObj
}

func (r WorkflowTriggerFireParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerFireParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTriggerFireParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
