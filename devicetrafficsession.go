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

// DeviceTrafficSessionService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceTrafficSessionService] method instead.
type DeviceTrafficSessionService struct {
	options []option.RequestOption
}

// NewDeviceTrafficSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeviceTrafficSessionService(opts ...option.RequestOption) (r DeviceTrafficSessionService) {
	r = DeviceTrafficSessionService{}
	r.options = opts
	return
}

// Starts one live-only decoded HTTP/1.1, HTTP/2, HTTP/3 and WebSocket traffic
// session.
func (r *DeviceTrafficSessionService) New(ctx context.Context, deviceID string, body DeviceTrafficSessionNewParams, opts ...option.RequestOption) (res *DeviceTrafficSessionNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/traffic/sessions", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns status and the device stream credential for a live session.
func (r *DeviceTrafficSessionService) Get(ctx context.Context, sessionID string, query DeviceTrafficSessionGetParams, opts ...option.RequestOption) (res *DeviceTrafficSessionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required sessionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/traffic/sessions/%s", url.PathEscape(query.DeviceID), url.PathEscape(sessionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List device traffic sessions
func (r *DeviceTrafficSessionService) List(ctx context.Context, deviceID string, query DeviceTrafficSessionListParams, opts ...option.RequestOption) (res *pagination.Page[DeviceTrafficSessionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/traffic/sessions", url.PathEscape(deviceID))
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

// List device traffic sessions
func (r *DeviceTrafficSessionService) ListAutoPaging(ctx context.Context, deviceID string, query DeviceTrafficSessionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[DeviceTrafficSessionListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, deviceID, query, opts...))
}

// Stop device traffic inspection
func (r *DeviceTrafficSessionService) Delete(ctx context.Context, sessionID string, body DeviceTrafficSessionDeleteParams, opts ...option.RequestOption) (res *DeviceTrafficSessionDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required sessionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/traffic/sessions/%s", url.PathEscape(body.DeviceID), url.PathEscape(sessionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type StreamCredentials struct {
	Token    string `json:"token" api:"required"`
	Protocol string `json:"protocol" api:"required"`
	URL      string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Protocol    respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamCredentials) RawJSON() string { return r.JSON.raw }
func (r *StreamCredentials) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceTrafficSessionNewResponse struct {
	ID           string    `json:"id" api:"required"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	DeviceID     string    `json:"deviceId" api:"required"`
	ExpiresAt    time.Time `json:"expiresAt" api:"required" format:"date-time"`
	MaxBodyBytes int64     `json:"maxBodyBytes" api:"required"`
	// Any of "none".
	Retention DeviceTrafficSessionNewResponseRetention `json:"retention" api:"required"`
	// Any of "starting", "active", "stopping", "stopped", "failed", "expired".
	State DeviceTrafficSessionNewResponseState `json:"state" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema    string                               `json:"$schema" format:"uri"`
	Error     DeviceTrafficSessionNewResponseError `json:"error"`
	StartedAt time.Time                            `json:"startedAt" format:"date-time"`
	StoppedAt time.Time                            `json:"stoppedAt" format:"date-time"`
	Stream    StreamCredentials                    `json:"stream"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		DeviceID     respjson.Field
		ExpiresAt    respjson.Field
		MaxBodyBytes respjson.Field
		Retention    respjson.Field
		State        respjson.Field
		Schema       respjson.Field
		Error        respjson.Field
		StartedAt    respjson.Field
		StoppedAt    respjson.Field
		Stream       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceTrafficSessionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceTrafficSessionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceTrafficSessionNewResponseRetention string

const (
	DeviceTrafficSessionNewResponseRetentionNone DeviceTrafficSessionNewResponseRetention = "none"
)

type DeviceTrafficSessionNewResponseState string

const (
	DeviceTrafficSessionNewResponseStateStarting DeviceTrafficSessionNewResponseState = "starting"
	DeviceTrafficSessionNewResponseStateActive   DeviceTrafficSessionNewResponseState = "active"
	DeviceTrafficSessionNewResponseStateStopping DeviceTrafficSessionNewResponseState = "stopping"
	DeviceTrafficSessionNewResponseStateStopped  DeviceTrafficSessionNewResponseState = "stopped"
	DeviceTrafficSessionNewResponseStateFailed   DeviceTrafficSessionNewResponseState = "failed"
	DeviceTrafficSessionNewResponseStateExpired  DeviceTrafficSessionNewResponseState = "expired"
)

type DeviceTrafficSessionNewResponseError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceTrafficSessionNewResponseError) RawJSON() string { return r.JSON.raw }
func (r *DeviceTrafficSessionNewResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceTrafficSessionGetResponse struct {
	ID           string    `json:"id" api:"required"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	DeviceID     string    `json:"deviceId" api:"required"`
	ExpiresAt    time.Time `json:"expiresAt" api:"required" format:"date-time"`
	MaxBodyBytes int64     `json:"maxBodyBytes" api:"required"`
	// Any of "none".
	Retention DeviceTrafficSessionGetResponseRetention `json:"retention" api:"required"`
	// Any of "starting", "active", "stopping", "stopped", "failed", "expired".
	State DeviceTrafficSessionGetResponseState `json:"state" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema    string                               `json:"$schema" format:"uri"`
	Error     DeviceTrafficSessionGetResponseError `json:"error"`
	StartedAt time.Time                            `json:"startedAt" format:"date-time"`
	StoppedAt time.Time                            `json:"stoppedAt" format:"date-time"`
	Stream    StreamCredentials                    `json:"stream"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		DeviceID     respjson.Field
		ExpiresAt    respjson.Field
		MaxBodyBytes respjson.Field
		Retention    respjson.Field
		State        respjson.Field
		Schema       respjson.Field
		Error        respjson.Field
		StartedAt    respjson.Field
		StoppedAt    respjson.Field
		Stream       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceTrafficSessionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceTrafficSessionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceTrafficSessionGetResponseRetention string

const (
	DeviceTrafficSessionGetResponseRetentionNone DeviceTrafficSessionGetResponseRetention = "none"
)

type DeviceTrafficSessionGetResponseState string

const (
	DeviceTrafficSessionGetResponseStateStarting DeviceTrafficSessionGetResponseState = "starting"
	DeviceTrafficSessionGetResponseStateActive   DeviceTrafficSessionGetResponseState = "active"
	DeviceTrafficSessionGetResponseStateStopping DeviceTrafficSessionGetResponseState = "stopping"
	DeviceTrafficSessionGetResponseStateStopped  DeviceTrafficSessionGetResponseState = "stopped"
	DeviceTrafficSessionGetResponseStateFailed   DeviceTrafficSessionGetResponseState = "failed"
	DeviceTrafficSessionGetResponseStateExpired  DeviceTrafficSessionGetResponseState = "expired"
)

type DeviceTrafficSessionGetResponseError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceTrafficSessionGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *DeviceTrafficSessionGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceTrafficSessionListResponse struct {
	ID           string    `json:"id" api:"required"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	DeviceID     string    `json:"deviceId" api:"required"`
	ExpiresAt    time.Time `json:"expiresAt" api:"required" format:"date-time"`
	MaxBodyBytes int64     `json:"maxBodyBytes" api:"required"`
	// Any of "none".
	Retention DeviceTrafficSessionListResponseRetention `json:"retention" api:"required"`
	// Any of "starting", "active", "stopping", "stopped", "failed", "expired".
	State DeviceTrafficSessionListResponseState `json:"state" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema    string                                `json:"$schema" format:"uri"`
	Error     DeviceTrafficSessionListResponseError `json:"error"`
	StartedAt time.Time                             `json:"startedAt" format:"date-time"`
	StoppedAt time.Time                             `json:"stoppedAt" format:"date-time"`
	Stream    StreamCredentials                     `json:"stream"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		DeviceID     respjson.Field
		ExpiresAt    respjson.Field
		MaxBodyBytes respjson.Field
		Retention    respjson.Field
		State        respjson.Field
		Schema       respjson.Field
		Error        respjson.Field
		StartedAt    respjson.Field
		StoppedAt    respjson.Field
		Stream       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceTrafficSessionListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceTrafficSessionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceTrafficSessionListResponseRetention string

const (
	DeviceTrafficSessionListResponseRetentionNone DeviceTrafficSessionListResponseRetention = "none"
)

type DeviceTrafficSessionListResponseState string

const (
	DeviceTrafficSessionListResponseStateStarting DeviceTrafficSessionListResponseState = "starting"
	DeviceTrafficSessionListResponseStateActive   DeviceTrafficSessionListResponseState = "active"
	DeviceTrafficSessionListResponseStateStopping DeviceTrafficSessionListResponseState = "stopping"
	DeviceTrafficSessionListResponseStateStopped  DeviceTrafficSessionListResponseState = "stopped"
	DeviceTrafficSessionListResponseStateFailed   DeviceTrafficSessionListResponseState = "failed"
	DeviceTrafficSessionListResponseStateExpired  DeviceTrafficSessionListResponseState = "expired"
)

type DeviceTrafficSessionListResponseError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceTrafficSessionListResponseError) RawJSON() string { return r.JSON.raw }
func (r *DeviceTrafficSessionListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceTrafficSessionDeleteResponse struct {
	ID           string    `json:"id" api:"required"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	DeviceID     string    `json:"deviceId" api:"required"`
	ExpiresAt    time.Time `json:"expiresAt" api:"required" format:"date-time"`
	MaxBodyBytes int64     `json:"maxBodyBytes" api:"required"`
	// Any of "none".
	Retention DeviceTrafficSessionDeleteResponseRetention `json:"retention" api:"required"`
	// Any of "starting", "active", "stopping", "stopped", "failed", "expired".
	State DeviceTrafficSessionDeleteResponseState `json:"state" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema    string                                  `json:"$schema" format:"uri"`
	Error     DeviceTrafficSessionDeleteResponseError `json:"error"`
	StartedAt time.Time                               `json:"startedAt" format:"date-time"`
	StoppedAt time.Time                               `json:"stoppedAt" format:"date-time"`
	Stream    StreamCredentials                       `json:"stream"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		DeviceID     respjson.Field
		ExpiresAt    respjson.Field
		MaxBodyBytes respjson.Field
		Retention    respjson.Field
		State        respjson.Field
		Schema       respjson.Field
		Error        respjson.Field
		StartedAt    respjson.Field
		StoppedAt    respjson.Field
		Stream       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceTrafficSessionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceTrafficSessionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceTrafficSessionDeleteResponseRetention string

const (
	DeviceTrafficSessionDeleteResponseRetentionNone DeviceTrafficSessionDeleteResponseRetention = "none"
)

type DeviceTrafficSessionDeleteResponseState string

const (
	DeviceTrafficSessionDeleteResponseStateStarting DeviceTrafficSessionDeleteResponseState = "starting"
	DeviceTrafficSessionDeleteResponseStateActive   DeviceTrafficSessionDeleteResponseState = "active"
	DeviceTrafficSessionDeleteResponseStateStopping DeviceTrafficSessionDeleteResponseState = "stopping"
	DeviceTrafficSessionDeleteResponseStateStopped  DeviceTrafficSessionDeleteResponseState = "stopped"
	DeviceTrafficSessionDeleteResponseStateFailed   DeviceTrafficSessionDeleteResponseState = "failed"
	DeviceTrafficSessionDeleteResponseStateExpired  DeviceTrafficSessionDeleteResponseState = "expired"
)

type DeviceTrafficSessionDeleteResponseError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceTrafficSessionDeleteResponseError) RawJSON() string { return r.JSON.raw }
func (r *DeviceTrafficSessionDeleteResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceTrafficSessionNewParams struct {
	ExpiresInSeconds param.Opt[int64] `json:"expiresInSeconds,omitzero"`
	MaxBodyBytes     param.Opt[int64] `json:"maxBodyBytes,omitzero"`
	paramObj
}

func (r DeviceTrafficSessionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceTrafficSessionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceTrafficSessionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceTrafficSessionGetParams struct {
	DeviceID string `path:"deviceId" api:"required" json:"-"`
	paramObj
}

type DeviceTrafficSessionListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceTrafficSessionListParams]'s query parameters as
// `url.Values`.
func (r DeviceTrafficSessionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceTrafficSessionDeleteParams struct {
	DeviceID string `path:"deviceId" api:"required" json:"-"`
	paramObj
}
