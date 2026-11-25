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

// DeviceTaskService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceTaskService] method instead.
type DeviceTaskService struct {
	options []option.RequestOption
}

// NewDeviceTaskService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceTaskService(opts ...option.RequestOption) (r DeviceTaskService) {
	r = DeviceTaskService{}
	r.options = opts
	return
}

// Returns a paginated list of tasks that have run on the device, along with
// pagination metadata.
func (r *DeviceTaskService) List(ctx context.Context, deviceID string, query DeviceTaskListParams, opts ...option.RequestOption) (res *pagination.Page[DeviceTaskListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/tasks", url.PathEscape(deviceID))
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

// Returns a paginated list of tasks that have run on the device, along with
// pagination metadata.
func (r *DeviceTaskService) ListAutoPaging(ctx context.Context, deviceID string, query DeviceTaskListParams, opts ...option.RequestOption) *pagination.PageAutoPager[DeviceTaskListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, deviceID, query, opts...))
}

type DeviceTaskListResponse struct {
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	TaskID    string    `json:"taskId" api:"required"`
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		TaskID      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceTaskListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceTaskListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceTaskListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Any of "id", "createdAt", "updatedAt", "assignedAt".
	OrderBy DeviceTaskListParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// Any of "asc", "desc".
	OrderByDirection DeviceTaskListParamsOrderByDirection `query:"orderByDirection,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceTaskListParams]'s query parameters as `url.Values`.
func (r DeviceTaskListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceTaskListParamsOrderBy string

const (
	DeviceTaskListParamsOrderByID         DeviceTaskListParamsOrderBy = "id"
	DeviceTaskListParamsOrderByCreatedAt  DeviceTaskListParamsOrderBy = "createdAt"
	DeviceTaskListParamsOrderByUpdatedAt  DeviceTaskListParamsOrderBy = "updatedAt"
	DeviceTaskListParamsOrderByAssignedAt DeviceTaskListParamsOrderBy = "assignedAt"
)

type DeviceTaskListParamsOrderByDirection string

const (
	DeviceTaskListParamsOrderByDirectionAsc  DeviceTaskListParamsOrderByDirection = "asc"
	DeviceTaskListParamsOrderByDirectionDesc DeviceTaskListParamsOrderByDirection = "desc"
)
