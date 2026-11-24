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

// WebhookDeliveryService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookDeliveryService] method instead.
type WebhookDeliveryService struct {
	options []option.RequestOption
}

// NewWebhookDeliveryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookDeliveryService(opts ...option.RequestOption) (r WebhookDeliveryService) {
	r = WebhookDeliveryService{}
	r.options = opts
	return
}

// Returns a paginated feed of webhook deliveries across all of your subscriptions,
// with the originating endpoint URL included on each record. Results can be
// filtered by delivery status (pending, success, skipped, or dead), by a `since`
// timestamp, by `eventId` (exact match against the originating event id), and/or
// by endpoint delivery `kind`.
func (r *WebhookDeliveryService) List(ctx context.Context, query WebhookDeliveryListParams, opts ...option.RequestOption) (res *WebhookDeliveryListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks/deliveries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of deliveries for a single webhook subscription,
// identified by its id. Each record reports the event, delivery status, attempt
// count, and the last response code or error. Results can be filtered by `eventId`
// (exact match against the originating event id).
func (r *WebhookDeliveryService) ListForWebhook(ctx context.Context, id string, query WebhookDeliveryListForWebhookParams, opts ...option.RequestOption) (res *WebhookDeliveryListForWebhookResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s/deliveries", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a single delivery for a webhook subscription along with the full list of
// captured attempt records. Each attempt includes the request URL, method, headers
// and body, whether it was signed, and the response status, headers, and snippet.
func (r *WebhookDeliveryService) GetAttempts(ctx context.Context, deliveryID string, query WebhookDeliveryGetAttemptsParams, opts ...option.RequestOption) (res *WebhookDeliveryGetAttemptsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if deliveryID == "" {
		err = errors.New("missing required deliveryId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s/deliveries/%s", query.ID, deliveryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns aggregate delivery statistics across all of your webhooks, including the
// total count, a breakdown by status (pending, success, skipped, dead), and the
// overall success rate. Optional `since` and endpoint delivery `kind` filters
// narrow the reporting window.
func (r *WebhookDeliveryService) Stats(ctx context.Context, query WebhookDeliveryStatsParams, opts ...option.RequestOption) (res *WebhookDeliveryStatsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks/deliveries/stats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type WebhookDeliveryListResponse struct {
	Items      []WebhookDeliveryListResponseItem `json:"items" api:"required"`
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
func (r WebhookDeliveryListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListResponseItem struct {
	ID          string  `json:"id" api:"required" format:"uuid"`
	Attempts    float64 `json:"attempts" api:"required"`
	CompletedAt string  `json:"completedAt" api:"required"`
	CreatedAt   string  `json:"createdAt" api:"required"`
	// Id of the parent endpoint's creator. Null when the endpoint row is gone or its
	// creator was never recorded.
	CreatedBy  string  `json:"createdBy" api:"required" format:"uuid"`
	DurationMs float64 `json:"durationMs" api:"required"`
	EndpointID string  `json:"endpointId" api:"required" format:"uuid"`
	// Any of "http", "integration".
	EndpointKind string `json:"endpointKind" api:"required"`
	// Integration target label of the parent webhook (e.g. `#ops`); null for `http`
	// webhooks.
	EndpointLabel string `json:"endpointLabel" api:"required"`
	// Delivery URL of the parent webhook; null for `integration` webhooks.
	EndpointURL    string  `json:"endpointUrl" api:"required"`
	EventID        string  `json:"eventId" api:"required"`
	EventType      string  `json:"eventType" api:"required"`
	IsTest         bool    `json:"isTest" api:"required"`
	LastError      string  `json:"lastError" api:"required"`
	LastStatusCode float64 `json:"lastStatusCode" api:"required"`
	OccurredAt     string  `json:"occurredAt" api:"required"`
	Source         string  `json:"source" api:"required"`
	// Any of "pending", "success", "skipped", "dead".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Attempts       respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		CreatedBy      respjson.Field
		DurationMs     respjson.Field
		EndpointID     respjson.Field
		EndpointKind   respjson.Field
		EndpointLabel  respjson.Field
		EndpointURL    respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		IsTest         respjson.Field
		LastError      respjson.Field
		LastStatusCode respjson.Field
		OccurredAt     respjson.Field
		Source         respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListForWebhookResponse struct {
	Items      []WebhookDeliveryListForWebhookResponseItem `json:"items" api:"required"`
	Pagination shared.Pagination                           `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListForWebhookResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListForWebhookResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListForWebhookResponseItem struct {
	ID          string  `json:"id" api:"required" format:"uuid"`
	Attempts    float64 `json:"attempts" api:"required"`
	CompletedAt string  `json:"completedAt" api:"required"`
	CreatedAt   string  `json:"createdAt" api:"required"`
	// Id of the parent endpoint's creator. Null when the endpoint row is gone or its
	// creator was never recorded.
	CreatedBy      string  `json:"createdBy" api:"required" format:"uuid"`
	DurationMs     float64 `json:"durationMs" api:"required"`
	EndpointID     string  `json:"endpointId" api:"required" format:"uuid"`
	EventID        string  `json:"eventId" api:"required"`
	EventType      string  `json:"eventType" api:"required"`
	IsTest         bool    `json:"isTest" api:"required"`
	LastError      string  `json:"lastError" api:"required"`
	LastStatusCode float64 `json:"lastStatusCode" api:"required"`
	OccurredAt     string  `json:"occurredAt" api:"required"`
	Source         string  `json:"source" api:"required"`
	// Any of "pending", "success", "skipped", "dead".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Attempts       respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		CreatedBy      respjson.Field
		DurationMs     respjson.Field
		EndpointID     respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		IsTest         respjson.Field
		LastError      respjson.Field
		LastStatusCode respjson.Field
		OccurredAt     respjson.Field
		Source         respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryListForWebhookResponseItem) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryListForWebhookResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryGetAttemptsResponse struct {
	Data WebhookDeliveryGetAttemptsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryGetAttemptsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryGetAttemptsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryGetAttemptsResponseData struct {
	ID          string                                          `json:"id" api:"required" format:"uuid"`
	Attempts    []WebhookDeliveryGetAttemptsResponseDataAttempt `json:"attempts" api:"required"`
	CompletedAt string                                          `json:"completedAt" api:"required"`
	CreatedAt   string                                          `json:"createdAt" api:"required"`
	// Id of the parent endpoint's creator. Null when the endpoint row is gone or its
	// creator was never recorded.
	CreatedBy      string  `json:"createdBy" api:"required" format:"uuid"`
	DurationMs     float64 `json:"durationMs" api:"required"`
	EndpointID     string  `json:"endpointId" api:"required" format:"uuid"`
	EventID        string  `json:"eventId" api:"required"`
	EventType      string  `json:"eventType" api:"required"`
	IsTest         bool    `json:"isTest" api:"required"`
	LastError      string  `json:"lastError" api:"required"`
	LastStatusCode float64 `json:"lastStatusCode" api:"required"`
	OccurredAt     string  `json:"occurredAt" api:"required"`
	Source         string  `json:"source" api:"required"`
	// Any of "pending", "success", "skipped", "dead".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Attempts       respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		CreatedBy      respjson.Field
		DurationMs     respjson.Field
		EndpointID     respjson.Field
		EventID        respjson.Field
		EventType      respjson.Field
		IsTest         respjson.Field
		LastError      respjson.Field
		LastStatusCode respjson.Field
		OccurredAt     respjson.Field
		Source         respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryGetAttemptsResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryGetAttemptsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryGetAttemptsResponseDataAttempt struct {
	AttemptNo       float64           `json:"attemptNo" api:"required"`
	DurationMs      float64           `json:"durationMs" api:"required"`
	Error           string            `json:"error" api:"required"`
	RequestBody     string            `json:"requestBody" api:"required"`
	RequestHeaders  map[string]string `json:"requestHeaders" api:"required"`
	RequestMethod   string            `json:"requestMethod" api:"required"`
	RequestURL      string            `json:"requestUrl" api:"required"`
	ResponseHeaders map[string]string `json:"responseHeaders" api:"required"`
	ResponseSnippet string            `json:"responseSnippet" api:"required"`
	ResponseStatus  float64           `json:"responseStatus" api:"required"`
	SentAt          string            `json:"sentAt" api:"required"`
	Signed          bool              `json:"signed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttemptNo       respjson.Field
		DurationMs      respjson.Field
		Error           respjson.Field
		RequestBody     respjson.Field
		RequestHeaders  respjson.Field
		RequestMethod   respjson.Field
		RequestURL      respjson.Field
		ResponseHeaders respjson.Field
		ResponseSnippet respjson.Field
		ResponseStatus  respjson.Field
		SentAt          respjson.Field
		Signed          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryGetAttemptsResponseDataAttempt) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryGetAttemptsResponseDataAttempt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryStatsResponse struct {
	Data WebhookDeliveryStatsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryStatsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryStatsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryStatsResponseData struct {
	ByStatus    WebhookDeliveryStatsResponseDataByStatus `json:"byStatus" api:"required"`
	SuccessRate float64                                  `json:"successRate" api:"required"`
	Total       float64                                  `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByStatus    respjson.Field
		SuccessRate respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryStatsResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryStatsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryStatsResponseDataByStatus struct {
	Dead    float64 `json:"dead" api:"required"`
	Pending float64 `json:"pending" api:"required"`
	Skipped float64 `json:"skipped" api:"required"`
	Success float64 `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dead        respjson.Field
		Pending     respjson.Field
		Skipped     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryStatsResponseDataByStatus) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryStatsResponseDataByStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryListParams struct {
	// Exact text match against the originating event id.
	EventID  param.Opt[string]    `query:"eventId,omitzero" json:"-"`
	Page     param.Opt[int64]     `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64]     `query:"pageSize,omitzero" json:"-"`
	Since    param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// Only include deliveries to endpoints of this kind.
	//
	// Any of "http", "integration".
	Kind WebhookDeliveryListParamsKind `query:"kind,omitzero" json:"-"`
	// Any of "pending", "success", "skipped", "dead".
	Status WebhookDeliveryListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryListParams]'s query parameters as
// `url.Values`.
func (r WebhookDeliveryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only include deliveries to endpoints of this kind.
type WebhookDeliveryListParamsKind string

const (
	WebhookDeliveryListParamsKindHTTP        WebhookDeliveryListParamsKind = "http"
	WebhookDeliveryListParamsKindIntegration WebhookDeliveryListParamsKind = "integration"
)

type WebhookDeliveryListParamsStatus string

const (
	WebhookDeliveryListParamsStatusPending WebhookDeliveryListParamsStatus = "pending"
	WebhookDeliveryListParamsStatusSuccess WebhookDeliveryListParamsStatus = "success"
	WebhookDeliveryListParamsStatusSkipped WebhookDeliveryListParamsStatus = "skipped"
	WebhookDeliveryListParamsStatusDead    WebhookDeliveryListParamsStatus = "dead"
)

type WebhookDeliveryListForWebhookParams struct {
	// Exact text match against the originating event id.
	EventID  param.Opt[string] `query:"eventId,omitzero" json:"-"`
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64]  `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryListForWebhookParams]'s query parameters as
// `url.Values`.
func (r WebhookDeliveryListForWebhookParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookDeliveryGetAttemptsParams struct {
	ID string `path:"id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type WebhookDeliveryStatsParams struct {
	Since param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// Only include deliveries to endpoints of this kind.
	//
	// Any of "http", "integration".
	Kind WebhookDeliveryStatsParamsKind `query:"kind,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeliveryStatsParams]'s query parameters as
// `url.Values`.
func (r WebhookDeliveryStatsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only include deliveries to endpoints of this kind.
type WebhookDeliveryStatsParamsKind string

const (
	WebhookDeliveryStatsParamsKindHTTP        WebhookDeliveryStatsParamsKind = "http"
	WebhookDeliveryStatsParamsKindIntegration WebhookDeliveryStatsParamsKind = "integration"
)
