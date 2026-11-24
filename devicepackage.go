// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiquery"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
)

// DevicePackageService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePackageService] method instead.
type DevicePackageService struct {
	options []option.RequestOption
}

// NewDevicePackageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDevicePackageService(opts ...option.RequestOption) (r DevicePackageService) {
	r = DevicePackageService{}
	r.options = opts
	return
}

// Returns the package names of apps installed on the device. System and protected
// packages are excluded unless the corresponding query parameters are set.
func (r *DevicePackageService) List(ctx context.Context, deviceID string, params DevicePackageListParams, opts ...option.RequestOption) (res *[]string, err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/packages", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type DevicePackageListParams struct {
	IncludeProtectedPackages param.Opt[bool]  `query:"includeProtectedPackages,omitzero" json:"-"`
	IncludeSystemPackages    param.Opt[bool]  `query:"includeSystemPackages,omitzero" json:"-"`
	XDeviceDisplayID         param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DevicePackageListParams]'s query parameters as
// `url.Values`.
func (r DevicePackageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
