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
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// WebhookIntegrationService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookIntegrationService] method instead.
type WebhookIntegrationService struct {
	options []option.RequestOption
}

// NewWebhookIntegrationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWebhookIntegrationService(opts ...option.RequestOption) (r WebhookIntegrationService) {
	r = WebhookIntegrationService{}
	r.options = opts
	return
}

// Returns the integrations that can receive webhook events as channel messages
// (e.g. Slack), with your org's connection state for each. Connect an integration
// in the Cloud app before creating a webhook for it.
func (r *WebhookIntegrationService) List(ctx context.Context, opts ...option.RequestOption) (res *WebhookIntegrationListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "webhooks/integrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the live list of targets (e.g. Slack channels the connected account is a
// member of) for an integration. Pass a target's `args` together with the
// integration `capability` to `POST /webhooks` with `kind: "integration"`.
func (r *WebhookIntegrationService) ListTargets(ctx context.Context, capabilityID string, query WebhookIntegrationListTargetsParams, opts ...option.RequestOption) (res *WebhookIntegrationListTargetsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if capabilityID == "" {
		err = errors.New("missing required capabilityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/integrations/%s/targets", url.PathEscape(capabilityID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type WebhookIntegrationListResponse struct {
	Data []WebhookIntegrationListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookIntegrationListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookIntegrationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookIntegrationListResponseData struct {
	Capability WebhookIntegrationListResponseDataCapability `json:"capability" api:"required"`
	// Connection state for your org. Only `connected` integrations can be used as a
	// destination.
	//
	// Any of "connected", "not_connected", "expired".
	Connection string `json:"connection" api:"required"`
	// Opaque key for the Cloud connect dialog.
	ConnectKey string `json:"connectKey" api:"required"`
	// True for the previous revision during a switch; prefer the non-deprecated one.
	Deprecated bool `json:"deprecated" api:"required"`
	// Server-resolved brand logo URL; null when unknown. Show this instead of deriving
	// a logo from connectKey.
	IconURL string `json:"iconUrl" api:"required" format:"uri"`
	Label   string `json:"label" api:"required"`
	// True when targets (e.g. channels) can be listed via
	// `GET /webhooks/integrations/{capabilityId}/targets`.
	SupportsTargets bool `json:"supportsTargets" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capability      respjson.Field
		Connection      respjson.Field
		ConnectKey      respjson.Field
		Deprecated      respjson.Field
		IconURL         respjson.Field
		Label           respjson.Field
		SupportsTargets respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookIntegrationListResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookIntegrationListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookIntegrationListResponseDataCapability struct {
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
func (r WebhookIntegrationListResponseDataCapability) RawJSON() string { return r.JSON.raw }
func (r *WebhookIntegrationListResponseDataCapability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookIntegrationListTargetsResponse struct {
	Data []WebhookIntegrationListTargetsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookIntegrationListTargetsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookIntegrationListTargetsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookIntegrationListTargetsResponseData struct {
	// Pass as `args` when creating an `integration` webhook.
	Args      map[string]any `json:"args" api:"required"`
	IsPrivate bool           `json:"isPrivate" api:"required"`
	Label     string         `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args        respjson.Field
		IsPrivate   respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookIntegrationListTargetsResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookIntegrationListTargetsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookIntegrationListTargetsParams struct {
	Revision int64 `query:"revision" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookIntegrationListTargetsParams]'s query parameters as
// `url.Values`.
func (r WebhookIntegrationListTargetsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
