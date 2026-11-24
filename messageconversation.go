// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
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
)

// MessageConversationService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageConversationService] method instead.
type MessageConversationService struct {
	options []option.RequestOption
}

// NewMessageConversationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessageConversationService(opts ...option.RequestOption) (r MessageConversationService) {
	r = MessageConversationService{}
	r.options = opts
	return
}

// Lists SMS conversations by recent activity. Use both cursor fields from
// `nextCursor` to fetch the next page.
func (r *MessageConversationService) List(ctx context.Context, query MessageConversationListParams, opts ...option.RequestOption) (res *MessageConversationListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "numbers/messages/conversations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Marks inbound messages in a conversation as read through the supplied cursor.
// Repeating the request is safe.
func (r *MessageConversationService) MarkRead(ctx context.Context, body MessageConversationMarkReadParams, opts ...option.RequestOption) (res *MessageConversationMarkReadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "numbers/messages/conversations/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type MessageConversationListResponse struct {
	Items      []MessageConversationListResponseItem     `json:"items" api:"required"`
	NextCursor MessageConversationListResponseNextCursor `json:"nextCursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageConversationListResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageConversationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageConversationListResponseItem struct {
	EsimIDs     []string                                       `json:"esimIds" api:"required" format:"uuid"`
	LastMessage MessageConversationListResponseItemLastMessage `json:"lastMessage" api:"required"`
	PeerKey     string                                         `json:"peerKey" api:"required"`
	UnreadCount int64                                          `json:"unreadCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EsimIDs     respjson.Field
		LastMessage respjson.Field
		PeerKey     respjson.Field
		UnreadCount respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageConversationListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *MessageConversationListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageConversationListResponseItemLastMessage struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Body string `json:"body" api:"required"`
	// Any of "inbound", "outbound".
	Direction  string    `json:"direction" api:"required"`
	OccurredAt time.Time `json:"occurredAt" api:"required" format:"date-time"`
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
		ID          respjson.Field
		Body        respjson.Field
		Direction   respjson.Field
		OccurredAt  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageConversationListResponseItemLastMessage) RawJSON() string { return r.JSON.raw }
func (r *MessageConversationListResponseItemLastMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageConversationListResponseNextCursor struct {
	LastMessageID  string    `json:"lastMessageId" api:"required" format:"uuid"`
	LastOccurredAt time.Time `json:"lastOccurredAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastMessageID  respjson.Field
		LastOccurredAt respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageConversationListResponseNextCursor) RawJSON() string { return r.JSON.raw }
func (r *MessageConversationListResponseNextCursor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageConversationMarkReadResponse struct {
	Data MessageConversationMarkReadResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageConversationMarkReadResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageConversationMarkReadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageConversationMarkReadResponseData struct {
	Updated int64 `json:"updated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Updated     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageConversationMarkReadResponseData) RawJSON() string { return r.JSON.raw }
func (r *MessageConversationMarkReadResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageConversationListParams struct {
	CursorLastMessageID  param.Opt[string]    `query:"cursorLastMessageId,omitzero" format:"uuid" json:"-"`
	CursorLastOccurredAt param.Opt[time.Time] `query:"cursorLastOccurredAt,omitzero" format:"date-time" json:"-"`
	EsimID               param.Opt[string]    `query:"esimId,omitzero" format:"uuid" json:"-"`
	Limit                param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	NumberID             param.Opt[string]    `query:"numberId,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [MessageConversationListParams]'s query parameters as
// `url.Values`.
func (r MessageConversationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageConversationMarkReadParams struct {
	// The thread's canonical peer key (see GET .../conversations)
	PeerKey       string `json:"peerKey" api:"required"`
	UpToMessageID string `json:"upToMessageId" api:"required" format:"uuid"`
	// Mark inbound messages read up to (and including) this occurredAt/upToMessageId
	// cursor
	UpToOccurredAt time.Time `json:"upToOccurredAt" api:"required" format:"date-time"`
	paramObj
}

func (r MessageConversationMarkReadParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageConversationMarkReadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageConversationMarkReadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
