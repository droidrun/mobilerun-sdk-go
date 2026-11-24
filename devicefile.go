// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiform"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiquery"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
	"github.com/droidrun/mobilerun-sdk-go/v6/shared"
)

// DeviceFileService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceFileService] method instead.
type DeviceFileService struct {
	options []option.RequestOption
}

// NewDeviceFileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceFileService(opts ...option.RequestOption) (r DeviceFileService) {
	r = DeviceFileService{}
	r.options = opts
	return
}

// Lists the files at the directory path given in the path query parameter,
// returning each entry's metadata along with the path and total count.
func (r *DeviceFileService) List(ctx context.Context, deviceID string, params DeviceFileListParams, opts ...option.RequestOption) (res *DeviceFileListResponse, err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/files", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Deletes the file at the path given in the path query parameter from the device.
func (r *DeviceFileService) Delete(ctx context.Context, deviceID string, params DeviceFileDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/files", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Pulls the file at the given path query parameter from the device and returns its
// raw bytes as an octet-stream.
func (r *DeviceFileService) Download(ctx context.Context, deviceID string, params DeviceFileDownloadParams, opts ...option.RequestOption) (res *string, err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/files/download", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Uploads a file to the device via multipart form data, writing it into the
// directory given by the path query parameter using the uploaded file's name.
func (r *DeviceFileService) Upload(ctx context.Context, deviceID string, params DeviceFileUploadParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/files", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

type DeviceFileListResponse struct {
	Files []DeviceFileListResponseFile `json:"files" api:"required"`
	Path  string                       `json:"path" api:"required"`
	Total int64                        `json:"total" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Files       respjson.Field
		Path        respjson.Field
		Total       respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceFileListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceFileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceFileListResponseFile struct {
	ExtendedAttributes bool                                  `json:"extendedAttributes" api:"required"`
	Group              string                                `json:"group" api:"required"`
	HardLinks          int64                                 `json:"hardLinks" api:"required"`
	ModifiedAt         time.Time                             `json:"modifiedAt" api:"required" format:"date-time"`
	Name               string                                `json:"name" api:"required"`
	Owner              string                                `json:"owner" api:"required"`
	Permissions        DeviceFileListResponseFilePermissions `json:"permissions" api:"required"`
	Size               int64                                 `json:"size" api:"required"`
	Type               string                                `json:"type" api:"required"`
	SymlinkTarget      string                                `json:"symlinkTarget"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtendedAttributes respjson.Field
		Group              respjson.Field
		HardLinks          respjson.Field
		ModifiedAt         respjson.Field
		Name               respjson.Field
		Owner              respjson.Field
		Permissions        respjson.Field
		Size               respjson.Field
		Type               respjson.Field
		SymlinkTarget      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceFileListResponseFile) RawJSON() string { return r.JSON.raw }
func (r *DeviceFileListResponseFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceFileListResponseFilePermissions struct {
	Group   shared.PermissionSet                         `json:"group" api:"required"`
	Others  shared.PermissionSet                         `json:"others" api:"required"`
	Owner   shared.PermissionSet                         `json:"owner" api:"required"`
	Special DeviceFileListResponseFilePermissionsSpecial `json:"special" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Group       respjson.Field
		Others      respjson.Field
		Owner       respjson.Field
		Special     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceFileListResponseFilePermissions) RawJSON() string { return r.JSON.raw }
func (r *DeviceFileListResponseFilePermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceFileListResponseFilePermissionsSpecial struct {
	SetGid bool `json:"setGid" api:"required"`
	SetUid bool `json:"setUid" api:"required"`
	Sticky bool `json:"sticky" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SetGid      respjson.Field
		SetUid      respjson.Field
		Sticky      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceFileListResponseFilePermissionsSpecial) RawJSON() string { return r.JSON.raw }
func (r *DeviceFileListResponseFilePermissionsSpecial) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceFileListParams struct {
	Path             string           `query:"path" api:"required" json:"-"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceFileListParams]'s query parameters as `url.Values`.
func (r DeviceFileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceFileDeleteParams struct {
	Path             string           `query:"path" api:"required" json:"-"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceFileDeleteParams]'s query parameters as `url.Values`.
func (r DeviceFileDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceFileDownloadParams struct {
	Path             string           `query:"path" api:"required" json:"-"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceFileDownloadParams]'s query parameters as
// `url.Values`.
func (r DeviceFileDownloadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceFileUploadParams struct {
	Path             string           `query:"path" api:"required" json:"-"`
	File             io.Reader        `json:"file,omitzero" api:"required" format:"binary"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceFileUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [DeviceFileUploadParams]'s query parameters as `url.Values`.
func (r DeviceFileUploadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
