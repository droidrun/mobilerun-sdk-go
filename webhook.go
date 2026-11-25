// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiquery"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/pagination"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// WebhookService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	options      []option.RequestOption
	Integrations WebhookIntegrationService
	Deliveries   WebhookDeliveryService
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.options = opts
	r.Integrations = NewWebhookIntegrationService(opts...)
	r.Deliveries = NewWebhookDeliveryService(opts...)
	return
}

// Creates a webhook subscription and an optional list of event types to subscribe
// to (defaults to all when omitted). `kind: "http"` (the default) delivers signed
// JSON to a URL; the response includes the generated signing secret, which is
// returned only once at creation time and cannot be retrieved later.
// `kind: "integration"` posts each event as a message into a connected integration
// target (e.g. a Slack channel): pick the `capability` from
// `GET /webhooks/integrations` and the `args` from
// `GET /webhooks/integrations/{capabilityId}/targets`. Integration webhooks have
// no signing secret. Send an `Idempotency-Key` header to make the create safe to
// retry: the same key with the same body returns the existing webhook (status 201,
// `Idempotent-Replayed: true`) instead of creating a second one, and the same key
// with a different body is a 422. A replayed create does not return the signing
// secret again; use rotate-secret if the original response was lost. If that
// webhook was deleted in the meantime, the retry is a 404.
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *WebhookNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single webhook subscription by id, including its URL (or integration
// target), subscribed event types, state, and system-observed delivery health. The
// signing secret is never included.
func (r *WebhookService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a webhook subscription. Any combination of the subscribed event types,
// state (ACTIVE or DISABLED), and description may be changed, and at least one
// field must be supplied. Setting state to ACTIVE re-enables a subscription that
// was auto-blocked after sustained delivery failures. The URL or integration
// target cannot be changed; delete and recreate the webhook instead.
func (r *WebhookService) Update(ctx context.Context, id string, body WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of your webhook subscriptions, optionally filtered by
// status (active, failing, blocked, or disabled), by `search` (a case-insensitive
// substring match against the URL or description), and/or by delivery `kind`. The
// response also includes per-status counts across all of your subscriptions (of
// the requested `kind`, if given; the other filters do not apply to the counts).
func (r *WebhookService) List(ctx context.Context, query WebhookListParams, opts ...option.RequestOption) (res *pagination.Page[WebhookListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "webhooks"
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

// Returns a paginated list of your webhook subscriptions, optionally filtered by
// status (active, failing, blocked, or disabled), by `search` (a case-insensitive
// substring match against the URL or description), and/or by delivery `kind`. The
// response also includes per-status counts across all of your subscriptions (of
// the requested `kind`, if given; the other filters do not apply to the counts).
func (r *WebhookService) ListAutoPaging(ctx context.Context, query WebhookListParams, opts ...option.RequestOption) *pagination.PageAutoPager[WebhookListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a webhook subscription so it stops receiving deliveries. Returns 204 No
// Content on success.
func (r *WebhookService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Generates a new signing secret for the webhook subscription and returns it once
// in the response. The previous secret is replaced immediately, so any signature
// verification on your endpoint must be updated to use the new value. Only `http`
// webhooks have a signing secret; for `integration` webhooks this returns 400.
func (r *WebhookService) RotateSecret(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookRotateSecretResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s/rotate-secret", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Sends a single test payload to the webhook subscription URL (or a test message
// to its integration target) to verify connectivity. The response reports whether
// the attempt succeeded along with the returned HTTP status code or error, if any.
func (r *WebhookService) TestDelivery(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookTestDeliveryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s/test", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type WebhookNewResponse struct {
	Data WebhookNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewResponseData struct {
	ID        string `json:"id" api:"required" format:"uuid"`
	BlockedAt string `json:"blockedAt" api:"required"`
	// Why the webhook was blocked, e.g. `integration_disconnected` (reconnect the
	// integration) or `integration_target_invalid` (the channel is gone or not
	// allowed).
	BlockedReason string `json:"blockedReason" api:"required"`
	CreatedAt     string `json:"createdAt" api:"required"`
	// Id of the actor who created this endpoint. Null when no creator was recorded.
	CreatedBy   string   `json:"createdBy" api:"required" format:"uuid"`
	Description string   `json:"description" api:"required"`
	EventTypes  []string `json:"eventTypes" api:"required"`
	// System-observed delivery health. `blocked` endpoints are auto-disabled after
	// sustained failure; PATCH state=ACTIVE to re-enable.
	//
	// Any of "healthy", "failing", "blocked".
	Health string `json:"health" api:"required"`
	// Integration target; null for `http` webhooks.
	Integration WebhookNewResponseDataIntegration `json:"integration" api:"required"`
	// `http` posts signed JSON to `url`; `integration` posts a message through a
	// connected integration (e.g. a Slack channel).
	//
	// Any of "http", "integration".
	Kind string `json:"kind" api:"required"`
	// Always false for `integration` webhooks.
	SigningEnabled bool `json:"signingEnabled" api:"required"`
	// Any of "ACTIVE", "DISABLED", "DELETED".
	State     string `json:"state" api:"required"`
	UpdatedAt string `json:"updatedAt" api:"required"`
	// Delivery URL; null for `integration` webhooks.
	URL string `json:"url" api:"required"`
	// Signing secret for `http` webhooks — shown only once. Store it now. Absent for
	// `integration` webhooks.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		BlockedAt      respjson.Field
		BlockedReason  respjson.Field
		CreatedAt      respjson.Field
		CreatedBy      respjson.Field
		Description    respjson.Field
		EventTypes     respjson.Field
		Health         respjson.Field
		Integration    respjson.Field
		Kind           respjson.Field
		SigningEnabled respjson.Field
		State          respjson.Field
		UpdatedAt      respjson.Field
		URL            respjson.Field
		Secret         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration target; null for `http` webhooks.
type WebhookNewResponseDataIntegration struct {
	Args       map[string]any                              `json:"args" api:"required"`
	Capability WebhookNewResponseDataIntegrationCapability `json:"capability" api:"required"`
	// Target name at creation time, e.g. `#ops`.
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args        respjson.Field
		Capability  respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponseDataIntegration) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponseDataIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewResponseDataIntegrationCapability struct {
	ID       string `json:"id" api:"required"`
	Revision int64  `json:"revision" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Revision    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponseDataIntegrationCapability) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponseDataIntegrationCapability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookGetResponse struct {
	Data WebhookGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookGetResponseData struct {
	ID        string `json:"id" api:"required" format:"uuid"`
	BlockedAt string `json:"blockedAt" api:"required"`
	// Why the webhook was blocked, e.g. `integration_disconnected` (reconnect the
	// integration) or `integration_target_invalid` (the channel is gone or not
	// allowed).
	BlockedReason string `json:"blockedReason" api:"required"`
	CreatedAt     string `json:"createdAt" api:"required"`
	// Id of the actor who created this endpoint. Null when no creator was recorded.
	CreatedBy   string   `json:"createdBy" api:"required" format:"uuid"`
	Description string   `json:"description" api:"required"`
	EventTypes  []string `json:"eventTypes" api:"required"`
	// System-observed delivery health. `blocked` endpoints are auto-disabled after
	// sustained failure; PATCH state=ACTIVE to re-enable.
	//
	// Any of "healthy", "failing", "blocked".
	Health string `json:"health" api:"required"`
	// Integration target; null for `http` webhooks.
	Integration WebhookGetResponseDataIntegration `json:"integration" api:"required"`
	// `http` posts signed JSON to `url`; `integration` posts a message through a
	// connected integration (e.g. a Slack channel).
	//
	// Any of "http", "integration".
	Kind string `json:"kind" api:"required"`
	// Always false for `integration` webhooks.
	SigningEnabled bool `json:"signingEnabled" api:"required"`
	// Any of "ACTIVE", "DISABLED", "DELETED".
	State     string `json:"state" api:"required"`
	UpdatedAt string `json:"updatedAt" api:"required"`
	// Delivery URL; null for `integration` webhooks.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		BlockedAt      respjson.Field
		BlockedReason  respjson.Field
		CreatedAt      respjson.Field
		CreatedBy      respjson.Field
		Description    respjson.Field
		EventTypes     respjson.Field
		Health         respjson.Field
		Integration    respjson.Field
		Kind           respjson.Field
		SigningEnabled respjson.Field
		State          respjson.Field
		UpdatedAt      respjson.Field
		URL            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration target; null for `http` webhooks.
type WebhookGetResponseDataIntegration struct {
	Args       map[string]any                              `json:"args" api:"required"`
	Capability WebhookGetResponseDataIntegrationCapability `json:"capability" api:"required"`
	// Target name at creation time, e.g. `#ops`.
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args        respjson.Field
		Capability  respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookGetResponseDataIntegration) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetResponseDataIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookGetResponseDataIntegrationCapability struct {
	ID       string `json:"id" api:"required"`
	Revision int64  `json:"revision" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Revision    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookGetResponseDataIntegrationCapability) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetResponseDataIntegrationCapability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateResponse struct {
	Data WebhookUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateResponseData struct {
	ID        string `json:"id" api:"required" format:"uuid"`
	BlockedAt string `json:"blockedAt" api:"required"`
	// Why the webhook was blocked, e.g. `integration_disconnected` (reconnect the
	// integration) or `integration_target_invalid` (the channel is gone or not
	// allowed).
	BlockedReason string `json:"blockedReason" api:"required"`
	CreatedAt     string `json:"createdAt" api:"required"`
	// Id of the actor who created this endpoint. Null when no creator was recorded.
	CreatedBy   string   `json:"createdBy" api:"required" format:"uuid"`
	Description string   `json:"description" api:"required"`
	EventTypes  []string `json:"eventTypes" api:"required"`
	// System-observed delivery health. `blocked` endpoints are auto-disabled after
	// sustained failure; PATCH state=ACTIVE to re-enable.
	//
	// Any of "healthy", "failing", "blocked".
	Health string `json:"health" api:"required"`
	// Integration target; null for `http` webhooks.
	Integration WebhookUpdateResponseDataIntegration `json:"integration" api:"required"`
	// `http` posts signed JSON to `url`; `integration` posts a message through a
	// connected integration (e.g. a Slack channel).
	//
	// Any of "http", "integration".
	Kind string `json:"kind" api:"required"`
	// Always false for `integration` webhooks.
	SigningEnabled bool `json:"signingEnabled" api:"required"`
	// Any of "ACTIVE", "DISABLED", "DELETED".
	State     string `json:"state" api:"required"`
	UpdatedAt string `json:"updatedAt" api:"required"`
	// Delivery URL; null for `integration` webhooks.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		BlockedAt      respjson.Field
		BlockedReason  respjson.Field
		CreatedAt      respjson.Field
		CreatedBy      respjson.Field
		Description    respjson.Field
		EventTypes     respjson.Field
		Health         respjson.Field
		Integration    respjson.Field
		Kind           respjson.Field
		SigningEnabled respjson.Field
		State          respjson.Field
		UpdatedAt      respjson.Field
		URL            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration target; null for `http` webhooks.
type WebhookUpdateResponseDataIntegration struct {
	Args       map[string]any                                 `json:"args" api:"required"`
	Capability WebhookUpdateResponseDataIntegrationCapability `json:"capability" api:"required"`
	// Target name at creation time, e.g. `#ops`.
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args        respjson.Field
		Capability  respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookUpdateResponseDataIntegration) RawJSON() string { return r.JSON.raw }
func (r *WebhookUpdateResponseDataIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateResponseDataIntegrationCapability struct {
	ID       string `json:"id" api:"required"`
	Revision int64  `json:"revision" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Revision    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookUpdateResponseDataIntegrationCapability) RawJSON() string { return r.JSON.raw }
func (r *WebhookUpdateResponseDataIntegrationCapability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponse struct {
	ID        string `json:"id" api:"required" format:"uuid"`
	BlockedAt string `json:"blockedAt" api:"required"`
	// Why the webhook was blocked, e.g. `integration_disconnected` (reconnect the
	// integration) or `integration_target_invalid` (the channel is gone or not
	// allowed).
	BlockedReason string `json:"blockedReason" api:"required"`
	CreatedAt     string `json:"createdAt" api:"required"`
	// Id of the actor who created this endpoint. Null when no creator was recorded.
	CreatedBy   string   `json:"createdBy" api:"required" format:"uuid"`
	Description string   `json:"description" api:"required"`
	EventTypes  []string `json:"eventTypes" api:"required"`
	// System-observed delivery health. `blocked` endpoints are auto-disabled after
	// sustained failure; PATCH state=ACTIVE to re-enable.
	//
	// Any of "healthy", "failing", "blocked".
	Health WebhookListResponseHealth `json:"health" api:"required"`
	// Integration target; null for `http` webhooks.
	Integration WebhookListResponseIntegration `json:"integration" api:"required"`
	// `http` posts signed JSON to `url`; `integration` posts a message through a
	// connected integration (e.g. a Slack channel).
	//
	// Any of "http", "integration".
	Kind WebhookListResponseKind `json:"kind" api:"required"`
	// Always false for `integration` webhooks.
	SigningEnabled bool `json:"signingEnabled" api:"required"`
	// Any of "ACTIVE", "DISABLED", "DELETED".
	State     WebhookListResponseState `json:"state" api:"required"`
	UpdatedAt string                   `json:"updatedAt" api:"required"`
	// Delivery URL; null for `integration` webhooks.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		BlockedAt      respjson.Field
		BlockedReason  respjson.Field
		CreatedAt      respjson.Field
		CreatedBy      respjson.Field
		Description    respjson.Field
		EventTypes     respjson.Field
		Health         respjson.Field
		Integration    respjson.Field
		Kind           respjson.Field
		SigningEnabled respjson.Field
		State          respjson.Field
		UpdatedAt      respjson.Field
		URL            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// System-observed delivery health. `blocked` endpoints are auto-disabled after
// sustained failure; PATCH state=ACTIVE to re-enable.
type WebhookListResponseHealth string

const (
	WebhookListResponseHealthHealthy WebhookListResponseHealth = "healthy"
	WebhookListResponseHealthFailing WebhookListResponseHealth = "failing"
	WebhookListResponseHealthBlocked WebhookListResponseHealth = "blocked"
)

// Integration target; null for `http` webhooks.
type WebhookListResponseIntegration struct {
	Args       map[string]any                           `json:"args" api:"required"`
	Capability WebhookListResponseIntegrationCapability `json:"capability" api:"required"`
	// Target name at creation time, e.g. `#ops`.
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args        respjson.Field
		Capability  respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponseIntegration) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponseIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponseIntegrationCapability struct {
	ID       string `json:"id" api:"required"`
	Revision int64  `json:"revision" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Revision    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponseIntegrationCapability) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponseIntegrationCapability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `http` posts signed JSON to `url`; `integration` posts a message through a
// connected integration (e.g. a Slack channel).
type WebhookListResponseKind string

const (
	WebhookListResponseKindHTTP        WebhookListResponseKind = "http"
	WebhookListResponseKindIntegration WebhookListResponseKind = "integration"
)

type WebhookListResponseState string

const (
	WebhookListResponseStateActive   WebhookListResponseState = "ACTIVE"
	WebhookListResponseStateDisabled WebhookListResponseState = "DISABLED"
	WebhookListResponseStateDeleted  WebhookListResponseState = "DELETED"
)

type WebhookRotateSecretResponse struct {
	Data WebhookRotateSecretResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookRotateSecretResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookRotateSecretResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookRotateSecretResponseData struct {
	ID        string `json:"id" api:"required" format:"uuid"`
	BlockedAt string `json:"blockedAt" api:"required"`
	// Why the webhook was blocked, e.g. `integration_disconnected` (reconnect the
	// integration) or `integration_target_invalid` (the channel is gone or not
	// allowed).
	BlockedReason string `json:"blockedReason" api:"required"`
	CreatedAt     string `json:"createdAt" api:"required"`
	// Id of the actor who created this endpoint. Null when no creator was recorded.
	CreatedBy   string   `json:"createdBy" api:"required" format:"uuid"`
	Description string   `json:"description" api:"required"`
	EventTypes  []string `json:"eventTypes" api:"required"`
	// System-observed delivery health. `blocked` endpoints are auto-disabled after
	// sustained failure; PATCH state=ACTIVE to re-enable.
	//
	// Any of "healthy", "failing", "blocked".
	Health string `json:"health" api:"required"`
	// Integration target; null for `http` webhooks.
	Integration WebhookRotateSecretResponseDataIntegration `json:"integration" api:"required"`
	// `http` posts signed JSON to `url`; `integration` posts a message through a
	// connected integration (e.g. a Slack channel).
	//
	// Any of "http", "integration".
	Kind string `json:"kind" api:"required"`
	// Signing secret — shown only once. Store it now.
	Secret string `json:"secret" api:"required"`
	// Always false for `integration` webhooks.
	SigningEnabled bool `json:"signingEnabled" api:"required"`
	// Any of "ACTIVE", "DISABLED", "DELETED".
	State     string `json:"state" api:"required"`
	UpdatedAt string `json:"updatedAt" api:"required"`
	// Delivery URL; null for `integration` webhooks.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		BlockedAt      respjson.Field
		BlockedReason  respjson.Field
		CreatedAt      respjson.Field
		CreatedBy      respjson.Field
		Description    respjson.Field
		EventTypes     respjson.Field
		Health         respjson.Field
		Integration    respjson.Field
		Kind           respjson.Field
		Secret         respjson.Field
		SigningEnabled respjson.Field
		State          respjson.Field
		UpdatedAt      respjson.Field
		URL            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookRotateSecretResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookRotateSecretResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration target; null for `http` webhooks.
type WebhookRotateSecretResponseDataIntegration struct {
	Args       map[string]any                                       `json:"args" api:"required"`
	Capability WebhookRotateSecretResponseDataIntegrationCapability `json:"capability" api:"required"`
	// Target name at creation time, e.g. `#ops`.
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args        respjson.Field
		Capability  respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookRotateSecretResponseDataIntegration) RawJSON() string { return r.JSON.raw }
func (r *WebhookRotateSecretResponseDataIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookRotateSecretResponseDataIntegrationCapability struct {
	ID       string `json:"id" api:"required"`
	Revision int64  `json:"revision" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Revision    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookRotateSecretResponseDataIntegrationCapability) RawJSON() string { return r.JSON.raw }
func (r *WebhookRotateSecretResponseDataIntegrationCapability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTestDeliveryResponse struct {
	Data WebhookTestDeliveryResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTestDeliveryResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookTestDeliveryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTestDeliveryResponseData struct {
	Error      string  `json:"error" api:"required"`
	StatusCode float64 `json:"statusCode" api:"required"`
	Success    bool    `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		StatusCode  respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTestDeliveryResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookTestDeliveryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfObject *WebhookNewParamsBodyObject `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfWebhookNewsBodyObject2 *WebhookNewParamsBodyObject2 `json:",inline"`

	paramObj
}

func (u WebhookNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfObject, u.OfWebhookNewsBodyObject2)
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type WebhookNewParamsBodyObject struct {
	URL         string            `json:"url" api:"required" format:"uri"`
	Description param.Opt[string] `json:"description,omitzero"`
	EventTypes  []string          `json:"eventTypes,omitzero"`
	// Delivery transport. Omitted ⇒ `http`.
	//
	// Any of "http".
	Kind string `json:"kind,omitzero"`
	paramObj
}

func (r WebhookNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParamsBodyObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParamsBodyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookNewParamsBodyObject](
		"kind", "http",
	)
}

// The properties Args, Capability, Kind are required.
type WebhookNewParamsBodyObject2 struct {
	// Target args exactly as returned by
	// `GET /webhooks/integrations/{capabilityId}/targets`.
	Args       map[string]any                        `json:"args,omitzero" api:"required"`
	Capability WebhookNewParamsBodyObject2Capability `json:"capability,omitzero" api:"required"`
	// Any of "integration".
	Kind        string            `json:"kind,omitzero" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	EventTypes  []string          `json:"eventTypes,omitzero"`
	paramObj
}

func (r WebhookNewParamsBodyObject2) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParamsBodyObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParamsBodyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookNewParamsBodyObject2](
		"kind", "integration",
	)
}

// The properties ID, Revision are required.
type WebhookNewParamsBodyObject2Capability struct {
	ID       string `json:"id" api:"required"`
	Revision int64  `json:"revision" api:"required"`
	paramObj
}

func (r WebhookNewParamsBodyObject2Capability) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParamsBodyObject2Capability
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParamsBodyObject2Capability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	EventTypes  []string          `json:"eventTypes,omitzero"`
	// Any of "ACTIVE", "DISABLED".
	State WebhookUpdateParamsState `json:"state,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateParamsState string

const (
	WebhookUpdateParamsStateActive   WebhookUpdateParamsState = "ACTIVE"
	WebhookUpdateParamsStateDisabled WebhookUpdateParamsState = "DISABLED"
)

type WebhookListParams struct {
	// Only include webhooks created by this actor id. Mutually exclusive with `mine`.
	CreatedBy param.Opt[string] `query:"createdBy,omitzero" format:"uuid" json:"-"`
	Page      param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize  param.Opt[int64]  `query:"pageSize,omitzero" json:"-"`
	// Case-insensitive substring match against the URL or description.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Only include webhooks of this delivery kind.
	//
	// Any of "http", "integration".
	Kind WebhookListParamsKind `query:"kind,omitzero" json:"-"`
	// When true, only include webhooks created by you (not just owned by your org).
	//
	// Any of "true", "false".
	Mine WebhookListParamsMine `query:"mine,omitzero" json:"-"`
	// Any of "active", "failing", "blocked", "disabled".
	Status WebhookListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookListParams]'s query parameters as `url.Values`.
func (r WebhookListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only include webhooks of this delivery kind.
type WebhookListParamsKind string

const (
	WebhookListParamsKindHTTP        WebhookListParamsKind = "http"
	WebhookListParamsKindIntegration WebhookListParamsKind = "integration"
)

// When true, only include webhooks created by you (not just owned by your org).
type WebhookListParamsMine string

const (
	WebhookListParamsMineTrue  WebhookListParamsMine = "true"
	WebhookListParamsMineFalse WebhookListParamsMine = "false"
)

type WebhookListParamsStatus string

const (
	WebhookListParamsStatusActive   WebhookListParamsStatus = "active"
	WebhookListParamsStatusFailing  WebhookListParamsStatus = "failing"
	WebhookListParamsStatusBlocked  WebhookListParamsStatus = "blocked"
	WebhookListParamsStatusDisabled WebhookListParamsStatus = "disabled"
)
