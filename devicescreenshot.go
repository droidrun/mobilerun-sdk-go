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
)

// DeviceScreenshotService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceScreenshotService] method instead.
type DeviceScreenshotService struct {
	options []option.RequestOption
}

// NewDeviceScreenshotService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeviceScreenshotService(opts ...option.RequestOption) (r DeviceScreenshotService) {
	r = DeviceScreenshotService{}
	r.options = opts
	return
}

// Saved screenshots, newest first. Pass nextCursor as cursor to retrieve the next
// page.
func (r *DeviceScreenshotService) List(ctx context.Context, deviceID string, query DeviceScreenshotListParams, opts ...option.RequestOption) (res *DeviceScreenshotListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/screenshots", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a device screenshot
func (r *DeviceScreenshotService) Delete(ctx context.Context, screenshotID string, body DeviceScreenshotDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	if screenshotID == "" {
		err = errors.New("missing required screenshotId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/screenshots/%s", url.PathEscape(body.DeviceID), url.PathEscape(screenshotID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Captures the device screen and returns it as a PNG image. Pass save=true to
// store it as device media.
func (r *DeviceScreenshotService) Capture(ctx context.Context, deviceID string, params DeviceScreenshotCaptureParams, opts ...option.RequestOption) (res *string, err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/screenshots", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Redirects to a presigned PNG URL valid for 15 minutes.
func (r *DeviceScreenshotService) Image(ctx context.Context, screenshotID string, query DeviceScreenshotImageParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if query.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	if screenshotID == "" {
		err = errors.New("missing required screenshotId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/screenshots/%s/image", url.PathEscape(query.DeviceID), url.PathEscape(screenshotID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

type DeviceScreenshotListResponse struct {
	Items      []DeviceScreenshotListResponseItem `json:"items" api:"required"`
	NextCursor string                             `json:"nextCursor" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceScreenshotListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceScreenshotListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceScreenshotListResponseItem struct {
	ID          string    `json:"id" api:"required"`
	CapturedAt  time.Time `json:"capturedAt" api:"required" format:"date-time"`
	DeviceID    string    `json:"deviceId" api:"required"`
	ExpiresAt   time.Time `json:"expiresAt" api:"required" format:"date-time"`
	RecordingID string    `json:"recordingId" api:"required"`
	SizeBytes   int64     `json:"sizeBytes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CapturedAt  respjson.Field
		DeviceID    respjson.Field
		ExpiresAt   respjson.Field
		RecordingID respjson.Field
		SizeBytes   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceScreenshotListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *DeviceScreenshotListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceScreenshotListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceScreenshotListParams]'s query parameters as
// `url.Values`.
func (r DeviceScreenshotListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceScreenshotDeleteParams struct {
	DeviceID string `path:"deviceId" api:"required" json:"-"`
	paramObj
}

type DeviceScreenshotCaptureParams struct {
	HideOverlay param.Opt[bool] `query:"hideOverlay,omitzero" json:"-"`
	// Store the screenshot as device media: in the owner's active recording, or
	// standalone for 7 days. Without save=true nothing is stored. The PNG response is
	// unchanged; X-Screenshot-Saved reports whether it was stored.
	Save             param.Opt[bool]  `query:"save,omitzero" json:"-"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceScreenshotCaptureParams]'s query parameters as
// `url.Values`.
func (r DeviceScreenshotCaptureParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceScreenshotImageParams struct {
	DeviceID string `path:"deviceId" api:"required" json:"-"`
	paramObj
}
