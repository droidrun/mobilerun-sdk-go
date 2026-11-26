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

// NumberMessageService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumberMessageService] method instead.
type NumberMessageService struct {
	options []option.RequestOption
}

// NewNumberMessageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNumberMessageService(opts ...option.RequestOption) (r NumberMessageService) {
	r = NumberMessageService{}
	r.options = opts
	return
}

// Returns SMS messages on the number, inbound and outbound, scoped to the
// authenticated user. Newest first. Messages stay with the number across device
// switches.
func (r *NumberMessageService) List(ctx context.Context, id string, query NumberMessageListParams, opts ...option.RequestOption) (res *pagination.Page[NumberMessageListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("numbers/phones/%s/messages", id)
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

// Returns SMS messages on the number, inbound and outbound, scoped to the
// authenticated user. Newest first. Messages stay with the number across device
// switches.
func (r *NumberMessageService) ListAutoPaging(ctx context.Context, id string, query NumberMessageListParams, opts ...option.RequestOption) *pagination.PageAutoPager[NumberMessageListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, id, query, opts...))
}

// Queues an SMS from one of the caller's numbers that can send. Replaying an
// Idempotency-Key with the same payload returns the original message, reusing it
// with a different payload returns 422. 409 capability_unavailable: this number
// cannot send right now. 422 invalid_recipient / unsupported_destination /
// body_too_long: the recipient or text does not fit the number. 429
// daily_limit_reached: the rolling 24 h limit for the account is used up; 429
// rate_limited: too many sends by the owner in a short window; 429
// number_rate_limited: too many outbound sends from this number in the configured
// window (all sends from this number count, including those made on your behalf).
// 403 send_disabled: self-service SMS send is switched off;
// outbound_sms_not_granted: outbound SMS requires an admin grant.
func (r *NumberMessageService) Send(ctx context.Context, id string, body NumberMessageSendParams, opts ...option.RequestOption) (res *NumberMessageSendResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("numbers/phones/%s/messages", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type NumberMessageListResponse struct {
	ID             string                                  `json:"id" api:"required" format:"uuid"`
	Body           string                                  `json:"body" api:"required"`
	Classification NumberMessageListResponseClassification `json:"classification" api:"required"`
	CreatedAt      time.Time                               `json:"createdAt" api:"required" format:"date-time"`
	DeliveryStatus string                                  `json:"deliveryStatus" api:"required"`
	DetectedSender string                                  `json:"detectedSender" api:"required"`
	// Any of "inbound", "outbound".
	Direction  NumberMessageListResponseDirection `json:"direction" api:"required"`
	EsimID     string                             `json:"esimId" api:"required" format:"uuid"`
	OccurredAt time.Time                          `json:"occurredAt" api:"required" format:"date-time"`
	PeerKey    string                             `json:"peerKey" api:"required"`
	PeerNumber string                             `json:"peerNumber" api:"required"`
	// Message status. Outbound messages move queued → sending → sent, or end at
	// failed. `sent` means the mobile network accepted the message and is the final
	// success status. `delivered` and `sent_unconfirmed` are kept for compatibility;
	// current lines return no carrier delivery report, so neither is emitted.
	//
	// Any of "received", "queued", "claimed", "sending", "sent", "sent_unconfirmed",
	// "delivered", "failed".
	Status NumberMessageListResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Body           respjson.Field
		Classification respjson.Field
		CreatedAt      respjson.Field
		DeliveryStatus respjson.Field
		DetectedSender respjson.Field
		Direction      respjson.Field
		EsimID         respjson.Field
		OccurredAt     respjson.Field
		PeerKey        respjson.Field
		PeerNumber     respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberMessageListResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberMessageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberMessageListResponseClassification struct {
	Spam bool `json:"spam" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Spam        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberMessageListResponseClassification) RawJSON() string { return r.JSON.raw }
func (r *NumberMessageListResponseClassification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberMessageListResponseDirection string

const (
	NumberMessageListResponseDirectionInbound  NumberMessageListResponseDirection = "inbound"
	NumberMessageListResponseDirectionOutbound NumberMessageListResponseDirection = "outbound"
)

// Message status. Outbound messages move queued → sending → sent, or end at
// failed. `sent` means the mobile network accepted the message and is the final
// success status. `delivered` and `sent_unconfirmed` are kept for compatibility;
// current lines return no carrier delivery report, so neither is emitted.
type NumberMessageListResponseStatus string

const (
	NumberMessageListResponseStatusReceived        NumberMessageListResponseStatus = "received"
	NumberMessageListResponseStatusQueued          NumberMessageListResponseStatus = "queued"
	NumberMessageListResponseStatusClaimed         NumberMessageListResponseStatus = "claimed"
	NumberMessageListResponseStatusSending         NumberMessageListResponseStatus = "sending"
	NumberMessageListResponseStatusSent            NumberMessageListResponseStatus = "sent"
	NumberMessageListResponseStatusSentUnconfirmed NumberMessageListResponseStatus = "sent_unconfirmed"
	NumberMessageListResponseStatusDelivered       NumberMessageListResponseStatus = "delivered"
	NumberMessageListResponseStatusFailed          NumberMessageListResponseStatus = "failed"
)

type NumberMessageSendResponse struct {
	Data NumberMessageSendResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberMessageSendResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberMessageSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberMessageSendResponseData struct {
	ID             string                                      `json:"id" api:"required" format:"uuid"`
	Body           string                                      `json:"body" api:"required"`
	Classification NumberMessageSendResponseDataClassification `json:"classification" api:"required"`
	CreatedAt      time.Time                                   `json:"createdAt" api:"required" format:"date-time"`
	DeliveryStatus string                                      `json:"deliveryStatus" api:"required"`
	DetectedSender string                                      `json:"detectedSender" api:"required"`
	// Any of "inbound", "outbound".
	Direction  string    `json:"direction" api:"required"`
	EsimID     string    `json:"esimId" api:"required" format:"uuid"`
	OccurredAt time.Time `json:"occurredAt" api:"required" format:"date-time"`
	PeerKey    string    `json:"peerKey" api:"required"`
	PeerNumber string    `json:"peerNumber" api:"required"`
	// Message status. Outbound messages move queued → sending → sent, or end at
	// failed. `sent` means the mobile network accepted the message and is the final
	// success status. `delivered` and `sent_unconfirmed` are kept for compatibility;
	// current lines return no carrier delivery report, so neither is emitted.
	//
	// Any of "received", "queued", "claimed", "sending", "sent", "sent_unconfirmed",
	// "delivered", "failed".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Body           respjson.Field
		Classification respjson.Field
		CreatedAt      respjson.Field
		DeliveryStatus respjson.Field
		DetectedSender respjson.Field
		Direction      respjson.Field
		EsimID         respjson.Field
		OccurredAt     respjson.Field
		PeerKey        respjson.Field
		PeerNumber     respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberMessageSendResponseData) RawJSON() string { return r.JSON.raw }
func (r *NumberMessageSendResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberMessageSendResponseDataClassification struct {
	Spam bool `json:"spam" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Spam        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberMessageSendResponseDataClassification) RawJSON() string { return r.JSON.raw }
func (r *NumberMessageSendResponseDataClassification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberMessageListParams struct {
	Page     param.Opt[int64]     `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64]     `query:"pageSize,omitzero" json:"-"`
	Since    param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// Any of "all", "inbound", "outbound".
	Direction NumberMessageListParamsDirection `query:"direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberMessageListParams]'s query parameters as
// `url.Values`.
func (r NumberMessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NumberMessageListParamsDirection string

const (
	NumberMessageListParamsDirectionAll      NumberMessageListParamsDirection = "all"
	NumberMessageListParamsDirectionInbound  NumberMessageListParamsDirection = "inbound"
	NumberMessageListParamsDirectionOutbound NumberMessageListParamsDirection = "outbound"
)

type NumberMessageSendParams struct {
	// SMS body text, up to 1600 characters (longer bodies are rejected with 400). A
	// number's own limit may be lower and is answered 422 body_too_long.
	Body string `json:"body" api:"required"`
	// Recipient phone number, as E.164 or a US 10/11-digit number. Destinations the
	// number cannot send to are rejected with 422 unsupported_destination, non-numbers
	// with 422 invalid_recipient.
	To string `json:"to" api:"required"`
	// Deprecated: best effort and ignored on most lines. Hosted SIM numbers return no
	// carrier delivery report, so their messages end at `sent` and no
	// `sms.send.delivered` event follows. Defaults to false.
	DeliveryReport param.Opt[bool] `json:"deliveryReport,omitzero"`
	paramObj
}

func (r NumberMessageSendParams) MarshalJSON() (data []byte, err error) {
	type shadow NumberMessageSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberMessageSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
