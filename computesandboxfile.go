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

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiform"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiquery"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
)

// ComputeSandboxFileService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeSandboxFileService] method instead.
type ComputeSandboxFileService struct {
	options []option.RequestOption
}

// NewComputeSandboxFileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewComputeSandboxFileService(opts ...option.RequestOption) (r ComputeSandboxFileService) {
	r = ComputeSandboxFileService{}
	r.options = opts
	return
}

// Download a file
func (r *ComputeSandboxFileService) Get(ctx context.Context, sandboxID string, query ComputeSandboxFileGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if sandboxID == "" {
		err = errors.New("missing required sandboxId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/sandboxes/%s/files", sandboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Upload a file
func (r *ComputeSandboxFileService) Upload(ctx context.Context, sandboxID string, body io.Reader, params ComputeSandboxFileUploadParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*"), option.WithRequestBody("application/octet-stream", body)}, opts...)
	if sandboxID == "" {
		err = errors.New("missing required sandboxId parameter")
		return err
	}
	path := fmt.Sprintf("compute/sandboxes/%s/files", sandboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return err
}

type ComputeSandboxFileGetParams struct {
	Path string `query:"path" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeSandboxFileGetParams]'s query parameters as
// `url.Values`.
func (r ComputeSandboxFileGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ComputeSandboxFileUploadParams struct {
	Path string `query:"path" api:"required" json:"-"`
	paramObj
}

func (r ComputeSandboxFileUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [ComputeSandboxFileUploadParams]'s query parameters as
// `url.Values`.
func (r ComputeSandboxFileUploadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
