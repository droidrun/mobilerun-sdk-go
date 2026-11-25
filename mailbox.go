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

// MailboxService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMailboxService] method instead.
type MailboxService struct {
	options     []option.RequestOption
	Connections MailboxConnectionService
	Messages    MailboxMessageService
}

// NewMailboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMailboxService(opts ...option.RequestOption) (r MailboxService) {
	r = MailboxService{}
	r.options = opts
	r.Connections = NewMailboxConnectionService(opts...)
	r.Messages = NewMailboxMessageService(opts...)
	return
}

// Creates a mailbox on the default domain or a connected custom domain. An
// optional `localPart` selects the address. Replaying the same Idempotency-Key and
// payload returns the original mailbox. Fresh requests and replays return 202 with
// the current mailbox state. Poll the mailbox when the response does not yet
// include a checkout URL.
func (r *MailboxService) New(ctx context.Context, body MailboxNewParams, opts ...option.RequestOption) (res *MailboxNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "mailboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a mailbox
func (r *MailboxService) Get(ctx context.Context, mailboxID string, opts ...option.RequestOption) (res *MailboxGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if mailboxID == "" {
		err = errors.New("missing required mailboxId parameter")
		return nil, err
	}
	path := fmt.Sprintf("mailboxes/%s", mailboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the label of a mailbox.
func (r *MailboxService) Update(ctx context.Context, mailboxID string, body MailboxUpdateParams, opts ...option.RequestOption) (res *MailboxUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if mailboxID == "" {
		err = errors.New("missing required mailboxId parameter")
		return nil, err
	}
	path := fmt.Sprintf("mailboxes/%s", mailboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists the caller-owned mailboxes with page-based pagination.
func (r *MailboxService) List(ctx context.Context, query MailboxListParams, opts ...option.RequestOption) (res *pagination.Page[Mailbox], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "mailboxes"
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

// Lists the caller-owned mailboxes with page-based pagination.
func (r *MailboxService) ListAutoPaging(ctx context.Context, query MailboxListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Mailbox] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Cancels a pending mailbox or schedules an active paid mailbox for cancellation.
// Existing addresses and messages are retained. Repeating the request is safe. An
// external inbox (Gmail) cannot be cancelled here; disconnect the link instead.
func (r *MailboxService) Delete(ctx context.Context, mailboxID string, opts ...option.RequestOption) (res *MailboxDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if mailboxID == "" {
		err = errors.New("missing required mailboxId parameter")
		return nil, err
	}
	path := fmt.Sprintf("mailboxes/%s", mailboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns the number of mailboxes currently available through included capacity.
func (r *MailboxService) Capacity(ctx context.Context, opts ...option.RequestOption) (res *MailboxCapacityResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "mailboxes/capacity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Removes only this mailbox link. The Gmail connection used by your agents stays
// in place.
func (r *MailboxService) Disconnect(ctx context.Context, mailboxID string, opts ...option.RequestOption) (res *MailboxDisconnectResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if mailboxID == "" {
		err = errors.New("missing required mailboxId parameter")
		return nil, err
	}
	path := fmt.Sprintf("mailboxes/%s/disconnect", mailboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns the most likely recent OTP for the mailbox. Returns 204 when no matching
// code is available.
func (r *MailboxService) Otp(ctx context.Context, mailboxID string, query MailboxOtpParams, opts ...option.RequestOption) (res *MailboxOtpResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if mailboxID == "" {
		err = errors.New("missing required mailboxId parameter")
		return nil, err
	}
	path := fmt.Sprintf("mailboxes/%s/otp", mailboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Restarts an archived mailbox with the same address. Uses included capacity when
// available unless paid service is requested.
func (r *MailboxService) Restart(ctx context.Context, mailboxID string, body MailboxRestartParams, opts ...option.RequestOption) (res *MailboxRestartResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if mailboxID == "" {
		err = errors.New("missing required mailboxId parameter")
		return nil, err
	}
	path := fmt.Sprintf("mailboxes/%s/restart", mailboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Withdraws a scheduled cancellation. Only available while cancellation is
// pending.
func (r *MailboxService) Uncancel(ctx context.Context, mailboxID string, opts ...option.RequestOption) (res *MailboxUncancelResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if mailboxID == "" {
		err = errors.New("missing required mailboxId parameter")
		return nil, err
	}
	path := fmt.Sprintf("mailboxes/%s/uncancel", mailboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type Mailbox struct {
	ID      string `json:"id" api:"required" format:"uuid"`
	Address string `json:"address" api:"required"`
	// Any of "rent", "included", "domain", "external".
	BillingMode       MailboxBillingMode     `json:"billingMode" api:"required"`
	CancelAtPeriodEnd bool                   `json:"cancelAtPeriodEnd" api:"required"`
	CheckoutExpiresAt time.Time              `json:"checkoutExpiresAt" api:"required" format:"date-time"`
	CheckoutURL       string                 `json:"checkoutUrl" api:"required"`
	CreatedAt         time.Time              `json:"createdAt" api:"required" format:"date-time"`
	CurrentPeriodEnd  time.Time              `json:"currentPeriodEnd" api:"required" format:"date-time"`
	DomainID          string                 `json:"domainId" api:"required" format:"uuid"`
	InboundMessages   MailboxInboundMessages `json:"inboundMessages" api:"required"`
	Label             string                 `json:"label" api:"required"`
	// Any of "matix", "gmail".
	Provider MailboxProvider `json:"provider" api:"required"`
	// Any of "provisioning", "awaiting_payment", "active", "cancel_scheduled",
	// "archived", "billing_error", "pending_connection", "connection_expired",
	// "connection_removed".
	Status MailboxStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Address           respjson.Field
		BillingMode       respjson.Field
		CancelAtPeriodEnd respjson.Field
		CheckoutExpiresAt respjson.Field
		CheckoutURL       respjson.Field
		CreatedAt         respjson.Field
		CurrentPeriodEnd  respjson.Field
		DomainID          respjson.Field
		InboundMessages   respjson.Field
		Label             respjson.Field
		Provider          respjson.Field
		Status            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Mailbox) RawJSON() string { return r.JSON.raw }
func (r *Mailbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxBillingMode string

const (
	MailboxBillingModeRent     MailboxBillingMode = "rent"
	MailboxBillingModeIncluded MailboxBillingMode = "included"
	MailboxBillingModeDomain   MailboxBillingMode = "domain"
	MailboxBillingModeExternal MailboxBillingMode = "external"
)

type MailboxInboundMessages struct {
	Exhausted bool      `json:"exhausted" api:"required"`
	Included  int64     `json:"included" api:"required"`
	ResetsAt  time.Time `json:"resetsAt" api:"required" format:"date-time"`
	Used      int64     `json:"used" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exhausted   respjson.Field
		Included    respjson.Field
		ResetsAt    respjson.Field
		Used        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxInboundMessages) RawJSON() string { return r.JSON.raw }
func (r *MailboxInboundMessages) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxProvider string

const (
	MailboxProviderMatix MailboxProvider = "matix"
	MailboxProviderGmail MailboxProvider = "gmail"
)

type MailboxStatus string

const (
	MailboxStatusProvisioning      MailboxStatus = "provisioning"
	MailboxStatusAwaitingPayment   MailboxStatus = "awaiting_payment"
	MailboxStatusActive            MailboxStatus = "active"
	MailboxStatusCancelScheduled   MailboxStatus = "cancel_scheduled"
	MailboxStatusArchived          MailboxStatus = "archived"
	MailboxStatusBillingError      MailboxStatus = "billing_error"
	MailboxStatusPendingConnection MailboxStatus = "pending_connection"
	MailboxStatusConnectionExpired MailboxStatus = "connection_expired"
	MailboxStatusConnectionRemoved MailboxStatus = "connection_removed"
)

type MailboxNewResponse struct {
	Data Mailbox `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxGetResponse struct {
	Data Mailbox `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxUpdateResponse struct {
	Data Mailbox `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxDeleteResponse struct {
	Data Mailbox `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxCapacityResponse struct {
	Data MailboxCapacityResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxCapacityResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxCapacityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxCapacityResponseData struct {
	// null with status unlimited
	Included int64 `json:"included" api:"required"`
	// null with status unlimited
	IncludedRemaining int64 `json:"includedRemaining" api:"required"`
	// null with status unlimited
	Remaining int64 `json:"remaining" api:"required"`
	// Any of "available", "exhausted", "not_included", "unlimited".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Included          respjson.Field
		IncludedRemaining respjson.Field
		Remaining         respjson.Field
		Status            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxCapacityResponseData) RawJSON() string { return r.JSON.raw }
func (r *MailboxCapacityResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxDisconnectResponse struct {
	Data Mailbox `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxDisconnectResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxDisconnectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxOtpResponse struct {
	Data MailboxOtpResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxOtpResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxOtpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxOtpResponseData struct {
	// OTP code as text to preserve leading zeros.
	Code string `json:"code" api:"required"`
	// Any of "high", "medium", "low".
	Confidence string    `json:"confidence" api:"required"`
	MessageID  string    `json:"messageId" api:"required"`
	ReceivedAt time.Time `json:"receivedAt" api:"required" format:"date-time"`
	Sender     string    `json:"sender" api:"required"`
	Subject    string    `json:"subject" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Confidence  respjson.Field
		MessageID   respjson.Field
		ReceivedAt  respjson.Field
		Sender      respjson.Field
		Subject     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxOtpResponseData) RawJSON() string { return r.JSON.raw }
func (r *MailboxOtpResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxRestartResponse struct {
	Data Mailbox `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxRestartResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxRestartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxUncancelResponse struct {
	Data Mailbox `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxUncancelResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxUncancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxNewParams struct {
	// Optional active custom mailbox domain owned by the caller. Omit to use the
	// system domain.
	DomainID param.Opt[string] `json:"domainId,omitzero" format:"uuid"`
	Label    param.Opt[string] `json:"label,omitzero"`
	// Optional mailbox name before the "@". Omit to generate a random address.
	LocalPart param.Opt[string] `json:"localPart,omitzero"`
	// included uses package capacity when available and otherwise starts paid
	// checkout; included_only fails without creating a paid reservation when no
	// included slot remains; rent always starts paid checkout.
	//
	// Any of "included", "included_only", "rent".
	BillingPreference MailboxNewParamsBillingPreference `json:"billingPreference,omitzero"`
	paramObj
}

func (r MailboxNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MailboxNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailboxNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// included uses package capacity when available and otherwise starts paid
// checkout; included_only fails without creating a paid reservation when no
// included slot remains; rent always starts paid checkout.
type MailboxNewParamsBillingPreference string

const (
	MailboxNewParamsBillingPreferenceIncluded     MailboxNewParamsBillingPreference = "included"
	MailboxNewParamsBillingPreferenceIncludedOnly MailboxNewParamsBillingPreference = "included_only"
	MailboxNewParamsBillingPreferenceRent         MailboxNewParamsBillingPreference = "rent"
)

type MailboxUpdateParams struct {
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	paramObj
}

func (r MailboxUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MailboxUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailboxUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MailboxListParams]'s query parameters as `url.Values`.
func (r MailboxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MailboxOtpParams struct {
	After     param.Opt[time.Time] `query:"after,omitzero" format:"date-time" json:"-"`
	MaxLength param.Opt[int64]     `query:"maxLength,omitzero" json:"-"`
	MinLength param.Opt[int64]     `query:"minLength,omitzero" json:"-"`
	Sender    param.Opt[string]    `query:"sender,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MailboxOtpParams]'s query parameters as `url.Values`.
func (r MailboxOtpParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MailboxRestartParams struct {
	// included uses package capacity when available and otherwise starts paid
	// checkout; included_only fails without creating a paid reservation when no
	// included slot remains; rent always starts paid checkout.
	//
	// Any of "included", "included_only", "rent".
	BillingPreference MailboxRestartParamsBillingPreference `json:"billingPreference,omitzero"`
	paramObj
}

func (r MailboxRestartParams) MarshalJSON() (data []byte, err error) {
	type shadow MailboxRestartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailboxRestartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// included uses package capacity when available and otherwise starts paid
// checkout; included_only fails without creating a paid reservation when no
// included slot remains; rent always starts paid checkout.
type MailboxRestartParamsBillingPreference string

const (
	MailboxRestartParamsBillingPreferenceIncluded     MailboxRestartParamsBillingPreference = "included"
	MailboxRestartParamsBillingPreferenceIncludedOnly MailboxRestartParamsBillingPreference = "included_only"
	MailboxRestartParamsBillingPreferenceRent         MailboxRestartParamsBillingPreference = "rent"
)
