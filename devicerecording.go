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

// DeviceRecordingService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceRecordingService] method instead.
type DeviceRecordingService struct {
	options []option.RequestOption
}

// NewDeviceRecordingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceRecordingService(opts ...option.RequestOption) (r DeviceRecordingService) {
	r = DeviceRecordingService{}
	r.options = opts
	return
}

// List device recordings
func (r *DeviceRecordingService) List(ctx context.Context, deviceID string, query DeviceRecordingListParams, opts ...option.RequestOption) (res *[]DeviceRecordingListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/recordings", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a device recording
func (r *DeviceRecordingService) Delete(ctx context.Context, recordingID string, body DeviceRecordingDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	if recordingID == "" {
		err = errors.New("missing required recordingId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/recordings/%s", url.PathEscape(body.DeviceID), url.PathEscape(recordingID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Screenshots saved into the recording with
// `GET /devices/{deviceId}/screenshot?save=true`, oldest first, each with a
// presigned download URL.
func (r *DeviceRecordingService) Screenshots(ctx context.Context, recordingID string, query DeviceRecordingScreenshotsParams, opts ...option.RequestOption) (res *[]DeviceRecordingScreenshotsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	if recordingID == "" {
		err = errors.New("missing required recordingId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/recordings/%s/screenshots", url.PathEscape(query.DeviceID), url.PathEscape(recordingID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Start a device recording
func (r *DeviceRecordingService) Start(ctx context.Context, deviceID string, body DeviceRecordingStartParams, opts ...option.RequestOption) (res *DeviceRecordingStartResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/recordings", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a device recording
func (r *DeviceRecordingService) Status(ctx context.Context, recordingID string, query DeviceRecordingStatusParams, opts ...option.RequestOption) (res *DeviceRecordingStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	if recordingID == "" {
		err = errors.New("missing required recordingId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/recordings/%s", url.PathEscape(query.DeviceID), url.PathEscape(recordingID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Stop a device recording
func (r *DeviceRecordingService) Stop(ctx context.Context, recordingID string, body DeviceRecordingStopParams, opts ...option.RequestOption) (res *DeviceRecordingStopResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	if recordingID == "" {
		err = errors.New("missing required recordingId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/recordings/%s", url.PathEscape(body.DeviceID), url.PathEscape(recordingID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get a device recording trajectory
func (r *DeviceRecordingService) Trajectory(ctx context.Context, recordingID string, query DeviceRecordingTrajectoryParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if query.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	if recordingID == "" {
		err = errors.New("missing required recordingId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/recordings/%s/trajectory", url.PathEscape(query.DeviceID), url.PathEscape(recordingID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Get a device recording video
func (r *DeviceRecordingService) Video(ctx context.Context, recordingID string, query DeviceRecordingVideoParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if query.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	if recordingID == "" {
		err = errors.New("missing required recordingId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/recordings/%s/video", url.PathEscape(query.DeviceID), url.PathEscape(recordingID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

type DeviceRecordingListResponse struct {
	ID        string                             `json:"id" api:"required"`
	Actions   int64                              `json:"actions" api:"required"`
	DeviceID  string                             `json:"deviceId" api:"required"`
	Display   DeviceRecordingListResponseDisplay `json:"display" api:"required"`
	ExpiresAt time.Time                          `json:"expiresAt" api:"required" format:"date-time"`
	Name      string                             `json:"name" api:"required"`
	StartedAt time.Time                          `json:"startedAt" api:"required" format:"date-time"`
	Status    string                             `json:"status" api:"required"`
	Types     []string                           `json:"types" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema  string                           `json:"$schema" format:"uri"`
	EndedAt time.Time                        `json:"endedAt" format:"date-time"`
	Error   string                           `json:"error"`
	Video   DeviceRecordingListResponseVideo `json:"video"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Actions     respjson.Field
		DeviceID    respjson.Field
		Display     respjson.Field
		ExpiresAt   respjson.Field
		Name        respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Types       respjson.Field
		Schema      respjson.Field
		EndedAt     respjson.Field
		Error       respjson.Field
		Video       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceRecordingListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceRecordingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRecordingListResponseDisplay struct {
	Height   int64 `json:"height" api:"required"`
	Rotation int64 `json:"rotation" api:"required"`
	Width    int64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Rotation    respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceRecordingListResponseDisplay) RawJSON() string { return r.JSON.raw }
func (r *DeviceRecordingListResponseDisplay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRecordingListResponseVideo struct {
	DurationMs  int64  `json:"durationMs" api:"required"`
	Format      string `json:"format" api:"required"`
	SizeBytes   int64  `json:"sizeBytes" api:"required"`
	Limited     bool   `json:"limited"`
	LimitReason string `json:"limitReason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationMs  respjson.Field
		Format      respjson.Field
		SizeBytes   respjson.Field
		Limited     respjson.Field
		LimitReason respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceRecordingListResponseVideo) RawJSON() string { return r.JSON.raw }
func (r *DeviceRecordingListResponseVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRecordingScreenshotsResponse struct {
	ID         string    `json:"id" api:"required"`
	CapturedAt time.Time `json:"capturedAt" api:"required" format:"date-time"`
	SizeBytes  int64     `json:"sizeBytes" api:"required"`
	// Presigned download URL, valid for 15 minutes.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CapturedAt  respjson.Field
		SizeBytes   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceRecordingScreenshotsResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceRecordingScreenshotsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRecordingStartResponse struct {
	ID        string                              `json:"id" api:"required"`
	Actions   int64                               `json:"actions" api:"required"`
	DeviceID  string                              `json:"deviceId" api:"required"`
	Display   DeviceRecordingStartResponseDisplay `json:"display" api:"required"`
	ExpiresAt time.Time                           `json:"expiresAt" api:"required" format:"date-time"`
	Name      string                              `json:"name" api:"required"`
	StartedAt time.Time                           `json:"startedAt" api:"required" format:"date-time"`
	Status    string                              `json:"status" api:"required"`
	Types     []string                            `json:"types" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema  string                            `json:"$schema" format:"uri"`
	EndedAt time.Time                         `json:"endedAt" format:"date-time"`
	Error   string                            `json:"error"`
	Video   DeviceRecordingStartResponseVideo `json:"video"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Actions     respjson.Field
		DeviceID    respjson.Field
		Display     respjson.Field
		ExpiresAt   respjson.Field
		Name        respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Types       respjson.Field
		Schema      respjson.Field
		EndedAt     respjson.Field
		Error       respjson.Field
		Video       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceRecordingStartResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceRecordingStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRecordingStartResponseDisplay struct {
	Height   int64 `json:"height" api:"required"`
	Rotation int64 `json:"rotation" api:"required"`
	Width    int64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Rotation    respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceRecordingStartResponseDisplay) RawJSON() string { return r.JSON.raw }
func (r *DeviceRecordingStartResponseDisplay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRecordingStartResponseVideo struct {
	DurationMs  int64  `json:"durationMs" api:"required"`
	Format      string `json:"format" api:"required"`
	SizeBytes   int64  `json:"sizeBytes" api:"required"`
	Limited     bool   `json:"limited"`
	LimitReason string `json:"limitReason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationMs  respjson.Field
		Format      respjson.Field
		SizeBytes   respjson.Field
		Limited     respjson.Field
		LimitReason respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceRecordingStartResponseVideo) RawJSON() string { return r.JSON.raw }
func (r *DeviceRecordingStartResponseVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRecordingStatusResponse struct {
	ID        string                               `json:"id" api:"required"`
	Actions   int64                                `json:"actions" api:"required"`
	DeviceID  string                               `json:"deviceId" api:"required"`
	Display   DeviceRecordingStatusResponseDisplay `json:"display" api:"required"`
	ExpiresAt time.Time                            `json:"expiresAt" api:"required" format:"date-time"`
	Name      string                               `json:"name" api:"required"`
	StartedAt time.Time                            `json:"startedAt" api:"required" format:"date-time"`
	Status    string                               `json:"status" api:"required"`
	Types     []string                             `json:"types" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema  string                             `json:"$schema" format:"uri"`
	EndedAt time.Time                          `json:"endedAt" format:"date-time"`
	Error   string                             `json:"error"`
	Video   DeviceRecordingStatusResponseVideo `json:"video"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Actions     respjson.Field
		DeviceID    respjson.Field
		Display     respjson.Field
		ExpiresAt   respjson.Field
		Name        respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Types       respjson.Field
		Schema      respjson.Field
		EndedAt     respjson.Field
		Error       respjson.Field
		Video       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceRecordingStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceRecordingStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRecordingStatusResponseDisplay struct {
	Height   int64 `json:"height" api:"required"`
	Rotation int64 `json:"rotation" api:"required"`
	Width    int64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Rotation    respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceRecordingStatusResponseDisplay) RawJSON() string { return r.JSON.raw }
func (r *DeviceRecordingStatusResponseDisplay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRecordingStatusResponseVideo struct {
	DurationMs  int64  `json:"durationMs" api:"required"`
	Format      string `json:"format" api:"required"`
	SizeBytes   int64  `json:"sizeBytes" api:"required"`
	Limited     bool   `json:"limited"`
	LimitReason string `json:"limitReason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationMs  respjson.Field
		Format      respjson.Field
		SizeBytes   respjson.Field
		Limited     respjson.Field
		LimitReason respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceRecordingStatusResponseVideo) RawJSON() string { return r.JSON.raw }
func (r *DeviceRecordingStatusResponseVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRecordingStopResponse struct {
	ID        string                             `json:"id" api:"required"`
	Actions   int64                              `json:"actions" api:"required"`
	DeviceID  string                             `json:"deviceId" api:"required"`
	Display   DeviceRecordingStopResponseDisplay `json:"display" api:"required"`
	ExpiresAt time.Time                          `json:"expiresAt" api:"required" format:"date-time"`
	Name      string                             `json:"name" api:"required"`
	StartedAt time.Time                          `json:"startedAt" api:"required" format:"date-time"`
	Status    string                             `json:"status" api:"required"`
	Types     []string                           `json:"types" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema  string                           `json:"$schema" format:"uri"`
	EndedAt time.Time                        `json:"endedAt" format:"date-time"`
	Error   string                           `json:"error"`
	Video   DeviceRecordingStopResponseVideo `json:"video"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Actions     respjson.Field
		DeviceID    respjson.Field
		Display     respjson.Field
		ExpiresAt   respjson.Field
		Name        respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Types       respjson.Field
		Schema      respjson.Field
		EndedAt     respjson.Field
		Error       respjson.Field
		Video       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceRecordingStopResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceRecordingStopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRecordingStopResponseDisplay struct {
	Height   int64 `json:"height" api:"required"`
	Rotation int64 `json:"rotation" api:"required"`
	Width    int64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Rotation    respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceRecordingStopResponseDisplay) RawJSON() string { return r.JSON.raw }
func (r *DeviceRecordingStopResponseDisplay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRecordingStopResponseVideo struct {
	DurationMs  int64  `json:"durationMs" api:"required"`
	Format      string `json:"format" api:"required"`
	SizeBytes   int64  `json:"sizeBytes" api:"required"`
	Limited     bool   `json:"limited"`
	LimitReason string `json:"limitReason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationMs  respjson.Field
		Format      respjson.Field
		SizeBytes   respjson.Field
		Limited     respjson.Field
		LimitReason respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceRecordingStopResponseVideo) RawJSON() string { return r.JSON.raw }
func (r *DeviceRecordingStopResponseVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRecordingListParams struct {
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	Type   param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceRecordingListParams]'s query parameters as
// `url.Values`.
func (r DeviceRecordingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceRecordingDeleteParams struct {
	DeviceID string `path:"deviceId" api:"required" json:"-"`
	paramObj
}

type DeviceRecordingScreenshotsParams struct {
	DeviceID string `path:"deviceId" api:"required" json:"-"`
	paramObj
}

type DeviceRecordingStartParams struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// Capture quality from 1 (lowest) to 10 (full stream quality). Defaults to the
	// device's full quality. Honored only by devices that support adjustable capture
	// quality.
	Quality       param.Opt[int64] `json:"quality,omitzero"`
	RetentionDays param.Opt[int64] `json:"retentionDays,omitzero"`
	// Artifacts to capture: trajectory (input actions; on some devices only when the
	// device announces trajectory capture), video, and audio (captured into the video
	// artifact, so it requires video; honored only by devices that support audio
	// capture). Defaults to trajectory and video, narrowed to what the device
	// produces.
	Types []string `json:"types,omitzero"`
	paramObj
}

func (r DeviceRecordingStartParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceRecordingStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceRecordingStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRecordingStatusParams struct {
	DeviceID string `path:"deviceId" api:"required" json:"-"`
	paramObj
}

type DeviceRecordingStopParams struct {
	DeviceID string `path:"deviceId" api:"required" json:"-"`
	paramObj
}

type DeviceRecordingTrajectoryParams struct {
	DeviceID string `path:"deviceId" api:"required" json:"-"`
	paramObj
}

type DeviceRecordingVideoParams struct {
	DeviceID string `path:"deviceId" api:"required" json:"-"`
	paramObj
}
