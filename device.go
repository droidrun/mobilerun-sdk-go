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
	shimjson "github.com/droidrun/mobilerun-sdk-go/v6/internal/encoding/json"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
	"github.com/droidrun/mobilerun-sdk-go/v6/shared"
)

// DeviceService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceService] method instead.
type DeviceService struct {
	options         []option.RequestOption
	Actions         DeviceActionService
	Apps            DeviceAppService
	Esim            DeviceEsimService
	Files           DeviceFileService
	Keyboard        DeviceKeyboardService
	Location        DeviceLocationService
	Packages        DevicePackageService
	Profile         DeviceProfileService
	Proxy           DeviceProxyService
	Screenshots     DeviceScreenshotService
	State           DeviceStateService
	Tasks           DeviceTaskService
	Clipboard       DeviceClipboardService
	Timezone        DeviceTimezoneService
	Language        DeviceLanguageService
	DeepLink        DeviceDeepLinkService
	Browser         DeviceBrowserService
	Kiosk           DeviceKioskService
	Recordings      DeviceRecordingService
	TrafficSessions DeviceTrafficSessionService
	Number          DeviceNumberService
	NumberBindings  DeviceNumberBindingService
}

// NewDeviceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDeviceService(opts ...option.RequestOption) (r DeviceService) {
	r = DeviceService{}
	r.options = opts
	r.Actions = NewDeviceActionService(opts...)
	r.Apps = NewDeviceAppService(opts...)
	r.Esim = NewDeviceEsimService(opts...)
	r.Files = NewDeviceFileService(opts...)
	r.Keyboard = NewDeviceKeyboardService(opts...)
	r.Location = NewDeviceLocationService(opts...)
	r.Packages = NewDevicePackageService(opts...)
	r.Profile = NewDeviceProfileService(opts...)
	r.Proxy = NewDeviceProxyService(opts...)
	r.Screenshots = NewDeviceScreenshotService(opts...)
	r.State = NewDeviceStateService(opts...)
	r.Tasks = NewDeviceTaskService(opts...)
	r.Clipboard = NewDeviceClipboardService(opts...)
	r.Timezone = NewDeviceTimezoneService(opts...)
	r.Language = NewDeviceLanguageService(opts...)
	r.DeepLink = NewDeviceDeepLinkService(opts...)
	r.Browser = NewDeviceBrowserService(opts...)
	r.Kiosk = NewDeviceKioskService(opts...)
	r.Recordings = NewDeviceRecordingService(opts...)
	r.TrafficSessions = NewDeviceTrafficSessionService(opts...)
	r.Number = NewDeviceNumberService(opts...)
	r.NumberBindings = NewDeviceNumberBindingService(opts...)
	return
}

// Requests a new device from the specification in the request body. Optional query
// parameters select the canonical device type, country, billing mode, and base
// profile. Returns the device, its resolved billing strategy, and its stream
// token.
func (r *DeviceService) New(ctx context.Context, params DeviceNewParams, opts ...option.RequestOption) (res *DeviceNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the current state and metadata for a single device, including its
// lifecycle state, type, stream URL, billing strategy, and timestamps. A stream
// token is included while the device is active, but may be omitted when it cannot
// be read at that moment. Refetch the device to obtain the stream token.
func (r *DeviceService) Get(ctx context.Context, deviceID string, opts ...option.RequestOption) (res *DeviceGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a paginated list of the user's devices along with pagination metadata.
// By default every device except terminated and unknown-state ones is listed,
// including stopped and maintenance devices; narrow with e.g. `state=ready` or
// `state=stopped,maintenance`, and pass `state=terminated` or `state=unknown` to
// see those.
func (r *DeviceService) List(ctx context.Context, query DeviceListParams, opts ...option.RequestOption) (res *DeviceListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deprecated: use GET /devices/summary instead. Returns the number of active
// claimed devices for the user, broken down by device type, in the legacy response
// shape.
//
// Deprecated: deprecated
func (r *DeviceService) Count(ctx context.Context, opts ...option.RequestOption) (res *DeviceCountResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "devices/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a live snapshot of the device's spoofed identity, including model,
// display, identifiers, and carrier. Devices without fingerprint support return an
// unsupported-feature error.
func (r *DeviceService) Fingerprint(ctx context.Context, deviceID string, query DeviceFingerprintParams, opts ...option.RequestOption) (res *DeviceFingerprintResponse, err error) {
	if !param.IsOmitted(query.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", query.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/fingerprint", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Triggers a reboot of the device. The device transitions through its reboot
// lifecycle and becomes ready again once the restart completes. Only devices whose
// capabilities report `reboot=true` can be rebooted. Personal (BYOD) devices
// cannot be rebooted remotely: the request is refused with 400
// `DEVICE_TOOL_UNSUPPORTED` and nothing happens on the device.
func (r *DeviceService) Reboot(ctx context.Context, deviceID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/reboot", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Resets the device. `mode` selects what is erased: `full` (factory reset: apps,
// accounts, storage, new identity) or `identity` (new fingerprint, apps and data
// kept). Without `mode` the reset is `full`, or `identity` on devices with
// `factoryReset=false`. A mode or `preserve` value the device does not support is
// refused with 400 `DEVICE_TOOL_UNSUPPORTED`. `preserve` restores the latest
// proxy, location, locale or timezone set through the API after the reset; omitted
// or empty preserves nothing. The device becomes ready again after the reset.
func (r *DeviceService) Reset(ctx context.Context, deviceID string, body DeviceResetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/reset", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Wakes a parked device: it first checks that the device can run again, then the
// same device starts and per-minute billing resumes. On a device that is not
// parked, the request does nothing. Unsupported intents return 400
// DEVICE_TOOL_UNSUPPORTED.
func (r *DeviceService) Resume(ctx context.Context, deviceID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/resume", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Returns the set of capabilities supported by this device. Used to determine
// which tools and features are available for the device.
func (r *DeviceService) GetCapabilities(ctx context.Context, deviceID string, opts ...option.RequestOption) (res *DeviceGetCapabilitiesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/capabilities", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Sets the display name for a device from the name in the request body and returns
// the updated device.
func (r *DeviceService) SetName(ctx context.Context, deviceID string, body DeviceSetNameParams, opts ...option.RequestOption) (res *DeviceSetNameResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/name", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Parks the device: its data, apps and identity are kept, but nothing runs and
// nothing is billed until it is resumed. Hosted subscription devices refuse a new
// park with 400 DEVICE_TOOL_UNSUPPORTED. Stop remains available for cleanup when
// the device already targets stopped. Other devices require stop support;
// unsupported intents return 400 DEVICE_TOOL_UNSUPPORTED.
func (r *DeviceService) Stop(ctx context.Context, deviceID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/stop", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Returns the total number of the user's devices and counts grouped by current
// state and device type.
func (r *DeviceService) Summary(ctx context.Context, query DeviceSummaryParams, opts ...option.RequestOption) (res *DeviceSummaryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "devices/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Terminates the device and releases its resources. For a personal (BYOD) device
// the owner (API key or user) can instead remove the phone from the account so
// anyone can pair it again; the phone must be disconnected first, otherwise the
// request is rejected with 409 BYOD_DEVICE_CONNECTED. Other subscription-billed
// devices cannot be terminated with this request; it is rejected with 401. The
// `terminateAt` and `previousDeviceId` body fields are reserved for internal use;
// requests that set them are rejected with 401.
func (r *DeviceService) Terminate(ctx context.Context, deviceID string, body DeviceTerminateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Blocks until the device reaches the ready state, then returns the same payload
// as Get device info. The call returns early with an error if the wait is
// cancelled or times out.
func (r *DeviceService) WaitReady(ctx context.Context, deviceID string, opts ...option.RequestOption) (res *DeviceWaitReadyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/wait", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type DeviceNewResponse struct {
	ID           string    `json:"id" api:"required"`
	ActiveTaskID string    `json:"activeTaskId" api:"required"`
	AssignedAt   time.Time `json:"assignedAt" api:"required" format:"date-time"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	Name         string    `json:"name" api:"required"`
	// Operating system the device runs.
	//
	// Any of "android", "ios".
	Platform     DeviceNewResponsePlatform `json:"platform" api:"required"`
	State        string                    `json:"state" api:"required"`
	StateMessage string                    `json:"stateMessage" api:"required"`
	StreamURL    string                    `json:"streamUrl" api:"required"`
	TaskCount    int64                     `json:"taskCount" api:"required"`
	TerminatesAt time.Time                 `json:"terminatesAt" api:"required" format:"date-time"`
	Type         string                    `json:"type" api:"required"`
	UpdatedAt    time.Time                 `json:"updatedAt" api:"required" format:"date-time"`
	// A URL to the JSON Schema for this object.
	Schema          string `json:"$schema" format:"uri"`
	BillingStrategy string `json:"billingStrategy"`
	CreatedBy       string `json:"createdBy"`
	// The platform default idle timeout this device uses when its own setting is
	// cleared (PUT idle-timeout with minutes null). Omitted when no default applies to
	// this device.
	IdleTimeoutDefaultMinutes int64 `json:"idleTimeoutDefaultMinutes"`
	// The effective idle timeout: minutes without client activity after which the
	// device is stopped automatically (state stopped, stopReason IDLE_TIMEOUT). It is
	// the device's own setting when one exists, otherwise the platform default for
	// pay-as-you-go (minute-billed) devices that can be stopped. Omitted while no idle
	// timer applies (explicitly turned off, or no default for this device). Change it
	// with PUT /devices/{deviceId}/idle-timeout.
	IdleTimeoutMinutes int64 `json:"idleTimeoutMinutes"`
	// Where idleTimeoutMinutes comes from: device = the device's own setting (also
	// when it is off), default = the platform default. Omitted when no idle timer can
	// apply.
	//
	// Any of "device", "default".
	IdleTimeoutSource DeviceNewResponseIdleTimeoutSource `json:"idleTimeoutSource"`
	// When the device last had client activity (RFC 3339). Tracked only while the idle
	// timer is on, to about one minute of precision; omitted otherwise.
	LastActivityAt time.Time `json:"lastActivityAt" format:"date-time"`
	OwnerID        string    `json:"ownerId"`
	ProviderID     string    `json:"providerId"`
	// Why the device was last stopped. Present only while state is stopped; clients
	// should use this value instead of parsing stateMessage.
	//
	// Any of "STOP_REQUESTED", "MINUTE_BALANCE_EXHAUSTED", "RESOURCE_EXHAUSTED",
	// "SUBSCRIPTION_REQUIRED", "PAYMENT_REQUIRED", "DEVICE_OPERATION_FAILED",
	// "DEVICE_LEFT", "IDLE_TIMEOUT".
	StopReason DeviceNewResponseStopReason `json:"stopReason"`
	// When devices-api recorded the stop reason (RFC 3339). Present only while state
	// is stopped and a reason is recorded.
	StopReasonAt time.Time `json:"stopReasonAt" format:"date-time"`
	StreamToken  string    `json:"streamToken"`
	// Deprecated: use ownerId/createdBy.
	//
	// Deprecated: deprecated
	UserID string `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		ActiveTaskID              respjson.Field
		AssignedAt                respjson.Field
		CreatedAt                 respjson.Field
		Name                      respjson.Field
		Platform                  respjson.Field
		State                     respjson.Field
		StateMessage              respjson.Field
		StreamURL                 respjson.Field
		TaskCount                 respjson.Field
		TerminatesAt              respjson.Field
		Type                      respjson.Field
		UpdatedAt                 respjson.Field
		Schema                    respjson.Field
		BillingStrategy           respjson.Field
		CreatedBy                 respjson.Field
		IdleTimeoutDefaultMinutes respjson.Field
		IdleTimeoutMinutes        respjson.Field
		IdleTimeoutSource         respjson.Field
		LastActivityAt            respjson.Field
		OwnerID                   respjson.Field
		ProviderID                respjson.Field
		StopReason                respjson.Field
		StopReasonAt              respjson.Field
		StreamToken               respjson.Field
		UserID                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Operating system the device runs.
type DeviceNewResponsePlatform string

const (
	DeviceNewResponsePlatformAndroid DeviceNewResponsePlatform = "android"
	DeviceNewResponsePlatformIos     DeviceNewResponsePlatform = "ios"
)

// Where idleTimeoutMinutes comes from: device = the device's own setting (also
// when it is off), default = the platform default. Omitted when no idle timer can
// apply.
type DeviceNewResponseIdleTimeoutSource string

const (
	DeviceNewResponseIdleTimeoutSourceDevice  DeviceNewResponseIdleTimeoutSource = "device"
	DeviceNewResponseIdleTimeoutSourceDefault DeviceNewResponseIdleTimeoutSource = "default"
)

// Why the device was last stopped. Present only while state is stopped; clients
// should use this value instead of parsing stateMessage.
type DeviceNewResponseStopReason string

const (
	DeviceNewResponseStopReasonStopRequested          DeviceNewResponseStopReason = "STOP_REQUESTED"
	DeviceNewResponseStopReasonMinuteBalanceExhausted DeviceNewResponseStopReason = "MINUTE_BALANCE_EXHAUSTED"
	DeviceNewResponseStopReasonResourceExhausted      DeviceNewResponseStopReason = "RESOURCE_EXHAUSTED"
	DeviceNewResponseStopReasonSubscriptionRequired   DeviceNewResponseStopReason = "SUBSCRIPTION_REQUIRED"
	DeviceNewResponseStopReasonPaymentRequired        DeviceNewResponseStopReason = "PAYMENT_REQUIRED"
	DeviceNewResponseStopReasonDeviceOperationFailed  DeviceNewResponseStopReason = "DEVICE_OPERATION_FAILED"
	DeviceNewResponseStopReasonDeviceLeft             DeviceNewResponseStopReason = "DEVICE_LEFT"
	DeviceNewResponseStopReasonIdleTimeout            DeviceNewResponseStopReason = "IDLE_TIMEOUT"
)

type DeviceGetResponse struct {
	ID           string    `json:"id" api:"required"`
	ActiveTaskID string    `json:"activeTaskId" api:"required"`
	AssignedAt   time.Time `json:"assignedAt" api:"required" format:"date-time"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	Name         string    `json:"name" api:"required"`
	// Operating system the device runs.
	//
	// Any of "android", "ios".
	Platform     DeviceGetResponsePlatform `json:"platform" api:"required"`
	State        string                    `json:"state" api:"required"`
	StateMessage string                    `json:"stateMessage" api:"required"`
	StreamURL    string                    `json:"streamUrl" api:"required"`
	TaskCount    int64                     `json:"taskCount" api:"required"`
	TerminatesAt time.Time                 `json:"terminatesAt" api:"required" format:"date-time"`
	Type         string                    `json:"type" api:"required"`
	UpdatedAt    time.Time                 `json:"updatedAt" api:"required" format:"date-time"`
	// A URL to the JSON Schema for this object.
	Schema          string `json:"$schema" format:"uri"`
	BillingStrategy string `json:"billingStrategy"`
	CreatedBy       string `json:"createdBy"`
	// The platform default idle timeout this device uses when its own setting is
	// cleared (PUT idle-timeout with minutes null). Omitted when no default applies to
	// this device.
	IdleTimeoutDefaultMinutes int64 `json:"idleTimeoutDefaultMinutes"`
	// The effective idle timeout: minutes without client activity after which the
	// device is stopped automatically (state stopped, stopReason IDLE_TIMEOUT). It is
	// the device's own setting when one exists, otherwise the platform default for
	// pay-as-you-go (minute-billed) devices that can be stopped. Omitted while no idle
	// timer applies (explicitly turned off, or no default for this device). Change it
	// with PUT /devices/{deviceId}/idle-timeout.
	IdleTimeoutMinutes int64 `json:"idleTimeoutMinutes"`
	// Where idleTimeoutMinutes comes from: device = the device's own setting (also
	// when it is off), default = the platform default. Omitted when no idle timer can
	// apply.
	//
	// Any of "device", "default".
	IdleTimeoutSource DeviceGetResponseIdleTimeoutSource `json:"idleTimeoutSource"`
	// When the device last had client activity (RFC 3339). Tracked only while the idle
	// timer is on, to about one minute of precision; omitted otherwise.
	LastActivityAt time.Time `json:"lastActivityAt" format:"date-time"`
	OwnerID        string    `json:"ownerId"`
	ProviderID     string    `json:"providerId"`
	// Why the device was last stopped. Present only while state is stopped; clients
	// should use this value instead of parsing stateMessage.
	//
	// Any of "STOP_REQUESTED", "MINUTE_BALANCE_EXHAUSTED", "RESOURCE_EXHAUSTED",
	// "SUBSCRIPTION_REQUIRED", "PAYMENT_REQUIRED", "DEVICE_OPERATION_FAILED",
	// "DEVICE_LEFT", "IDLE_TIMEOUT".
	StopReason DeviceGetResponseStopReason `json:"stopReason"`
	// When devices-api recorded the stop reason (RFC 3339). Present only while state
	// is stopped and a reason is recorded.
	StopReasonAt time.Time `json:"stopReasonAt" format:"date-time"`
	StreamToken  string    `json:"streamToken"`
	// Deprecated: use ownerId/createdBy.
	//
	// Deprecated: deprecated
	UserID string `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		ActiveTaskID              respjson.Field
		AssignedAt                respjson.Field
		CreatedAt                 respjson.Field
		Name                      respjson.Field
		Platform                  respjson.Field
		State                     respjson.Field
		StateMessage              respjson.Field
		StreamURL                 respjson.Field
		TaskCount                 respjson.Field
		TerminatesAt              respjson.Field
		Type                      respjson.Field
		UpdatedAt                 respjson.Field
		Schema                    respjson.Field
		BillingStrategy           respjson.Field
		CreatedBy                 respjson.Field
		IdleTimeoutDefaultMinutes respjson.Field
		IdleTimeoutMinutes        respjson.Field
		IdleTimeoutSource         respjson.Field
		LastActivityAt            respjson.Field
		OwnerID                   respjson.Field
		ProviderID                respjson.Field
		StopReason                respjson.Field
		StopReasonAt              respjson.Field
		StreamToken               respjson.Field
		UserID                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Operating system the device runs.
type DeviceGetResponsePlatform string

const (
	DeviceGetResponsePlatformAndroid DeviceGetResponsePlatform = "android"
	DeviceGetResponsePlatformIos     DeviceGetResponsePlatform = "ios"
)

// Where idleTimeoutMinutes comes from: device = the device's own setting (also
// when it is off), default = the platform default. Omitted when no idle timer can
// apply.
type DeviceGetResponseIdleTimeoutSource string

const (
	DeviceGetResponseIdleTimeoutSourceDevice  DeviceGetResponseIdleTimeoutSource = "device"
	DeviceGetResponseIdleTimeoutSourceDefault DeviceGetResponseIdleTimeoutSource = "default"
)

// Why the device was last stopped. Present only while state is stopped; clients
// should use this value instead of parsing stateMessage.
type DeviceGetResponseStopReason string

const (
	DeviceGetResponseStopReasonStopRequested          DeviceGetResponseStopReason = "STOP_REQUESTED"
	DeviceGetResponseStopReasonMinuteBalanceExhausted DeviceGetResponseStopReason = "MINUTE_BALANCE_EXHAUSTED"
	DeviceGetResponseStopReasonResourceExhausted      DeviceGetResponseStopReason = "RESOURCE_EXHAUSTED"
	DeviceGetResponseStopReasonSubscriptionRequired   DeviceGetResponseStopReason = "SUBSCRIPTION_REQUIRED"
	DeviceGetResponseStopReasonPaymentRequired        DeviceGetResponseStopReason = "PAYMENT_REQUIRED"
	DeviceGetResponseStopReasonDeviceOperationFailed  DeviceGetResponseStopReason = "DEVICE_OPERATION_FAILED"
	DeviceGetResponseStopReasonDeviceLeft             DeviceGetResponseStopReason = "DEVICE_LEFT"
	DeviceGetResponseStopReasonIdleTimeout            DeviceGetResponseStopReason = "IDLE_TIMEOUT"
)

type DeviceListResponse struct {
	Items      []DeviceListResponseItem `json:"items" api:"required"`
	Pagination shared.Meta              `json:"pagination" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceListResponseItem struct {
	ID           string    `json:"id" api:"required"`
	ActiveTaskID string    `json:"activeTaskId" api:"required"`
	AssignedAt   time.Time `json:"assignedAt" api:"required" format:"date-time"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	Name         string    `json:"name" api:"required"`
	// Operating system the device runs.
	//
	// Any of "android", "ios".
	Platform     string    `json:"platform" api:"required"`
	State        string    `json:"state" api:"required"`
	StateMessage string    `json:"stateMessage" api:"required"`
	StreamURL    string    `json:"streamUrl" api:"required"`
	TaskCount    int64     `json:"taskCount" api:"required"`
	TerminatesAt time.Time `json:"terminatesAt" api:"required" format:"date-time"`
	Type         string    `json:"type" api:"required"`
	UpdatedAt    time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// A URL to the JSON Schema for this object.
	Schema          string `json:"$schema" format:"uri"`
	BillingStrategy string `json:"billingStrategy"`
	CreatedBy       string `json:"createdBy"`
	// The platform default idle timeout this device uses when its own setting is
	// cleared (PUT idle-timeout with minutes null). Omitted when no default applies to
	// this device.
	IdleTimeoutDefaultMinutes int64 `json:"idleTimeoutDefaultMinutes"`
	// The effective idle timeout: minutes without client activity after which the
	// device is stopped automatically (state stopped, stopReason IDLE_TIMEOUT). It is
	// the device's own setting when one exists, otherwise the platform default for
	// pay-as-you-go (minute-billed) devices that can be stopped. Omitted while no idle
	// timer applies (explicitly turned off, or no default for this device). Change it
	// with PUT /devices/{deviceId}/idle-timeout.
	IdleTimeoutMinutes int64 `json:"idleTimeoutMinutes"`
	// Where idleTimeoutMinutes comes from: device = the device's own setting (also
	// when it is off), default = the platform default. Omitted when no idle timer can
	// apply.
	//
	// Any of "device", "default".
	IdleTimeoutSource string `json:"idleTimeoutSource"`
	// When the device last had client activity (RFC 3339). Tracked only while the idle
	// timer is on, to about one minute of precision; omitted otherwise.
	LastActivityAt time.Time `json:"lastActivityAt" format:"date-time"`
	OwnerID        string    `json:"ownerId"`
	ProviderID     string    `json:"providerId"`
	// Why the device was last stopped. Present only while state is stopped; clients
	// should use this value instead of parsing stateMessage.
	//
	// Any of "STOP_REQUESTED", "MINUTE_BALANCE_EXHAUSTED", "RESOURCE_EXHAUSTED",
	// "SUBSCRIPTION_REQUIRED", "PAYMENT_REQUIRED", "DEVICE_OPERATION_FAILED",
	// "DEVICE_LEFT", "IDLE_TIMEOUT".
	StopReason string `json:"stopReason"`
	// When devices-api recorded the stop reason (RFC 3339). Present only while state
	// is stopped and a reason is recorded.
	StopReasonAt time.Time `json:"stopReasonAt" format:"date-time"`
	StreamToken  string    `json:"streamToken"`
	// Deprecated: use ownerId/createdBy.
	//
	// Deprecated: deprecated
	UserID string `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		ActiveTaskID              respjson.Field
		AssignedAt                respjson.Field
		CreatedAt                 respjson.Field
		Name                      respjson.Field
		Platform                  respjson.Field
		State                     respjson.Field
		StateMessage              respjson.Field
		StreamURL                 respjson.Field
		TaskCount                 respjson.Field
		TerminatesAt              respjson.Field
		Type                      respjson.Field
		UpdatedAt                 respjson.Field
		Schema                    respjson.Field
		BillingStrategy           respjson.Field
		CreatedBy                 respjson.Field
		IdleTimeoutDefaultMinutes respjson.Field
		IdleTimeoutMinutes        respjson.Field
		IdleTimeoutSource         respjson.Field
		LastActivityAt            respjson.Field
		OwnerID                   respjson.Field
		ProviderID                respjson.Field
		StopReason                respjson.Field
		StopReasonAt              respjson.Field
		StreamToken               respjson.Field
		UserID                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *DeviceListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceCountResponse map[string]int64

type DeviceFingerprintResponse struct {
	Carrier     shared.DeviceCarrier             `json:"carrier" api:"required"`
	Display     DeviceFingerprintResponseDisplay `json:"display" api:"required"`
	Identifiers shared.DeviceIdentifiers         `json:"identifiers" api:"required"`
	Model       DeviceFingerprintResponseModel   `json:"model" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier     respjson.Field
		Display     respjson.Field
		Identifiers respjson.Field
		Model       respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceFingerprintResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceFingerprintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceFingerprintResponseDisplay struct {
	DensityDpi int64 `json:"densityDpi"`
	Height     int64 `json:"height"`
	Width      int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DensityDpi  respjson.Field
		Height      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceFingerprintResponseDisplay) RawJSON() string { return r.JSON.raw }
func (r *DeviceFingerprintResponseDisplay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceFingerprintResponseModel struct {
	AospVersion  string `json:"aospVersion"`
	Brand        string `json:"brand"`
	Device       string `json:"device"`
	Hardware     string `json:"hardware"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AospVersion  respjson.Field
		Brand        respjson.Field
		Device       respjson.Field
		Hardware     respjson.Field
		Manufacturer respjson.Field
		Model        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceFingerprintResponseModel) RawJSON() string { return r.JSON.raw }
func (r *DeviceFingerprintResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceGetCapabilitiesResponse struct {
	Capabilities DeviceGetCapabilitiesResponseCapabilities `json:"capabilities" api:"required"`
	DeviceType   string                                    `json:"deviceType" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capabilities respjson.Field
		DeviceType   respjson.Field
		Schema       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceGetCapabilitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceGetCapabilitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceGetCapabilitiesResponseCapabilities struct {
	Accessibility       bool `json:"accessibility" api:"required"`
	Agent               bool `json:"agent" api:"required"`
	Apps                bool `json:"apps" api:"required"`
	Browser             bool `json:"browser" api:"required"`
	CameraInjection     bool `json:"cameraInjection" api:"required"`
	Clipboard           bool `json:"clipboard" api:"required"`
	Esim                bool `json:"esim" api:"required"`
	FactoryReset        bool `json:"factoryReset" api:"required"`
	Files               bool `json:"files" api:"required"`
	Fingerprint         bool `json:"fingerprint" api:"required"`
	Frida               bool `json:"frida" api:"required"`
	Geo                 bool `json:"geo" api:"required"`
	HumanTouch          bool `json:"humanTouch" api:"required"`
	IdentityReset       bool `json:"identityReset" api:"required"`
	Language            bool `json:"language" api:"required"`
	Logcat              bool `json:"logcat" api:"required"`
	MicrophoneInjection bool `json:"microphoneInjection" api:"required"`
	Proxy               bool `json:"proxy" api:"required"`
	Reboot              bool `json:"reboot" api:"required"`
	Recording           bool `json:"recording" api:"required"`
	Reset               bool `json:"reset" api:"required"`
	// Any of "proxy", "location", "locale", "timezone".
	ResetPreserve     []string `json:"resetPreserve" api:"required"`
	SetGeo            bool     `json:"setGeo" api:"required"`
	SetTime           bool     `json:"setTime" api:"required"`
	Shell             bool     `json:"shell" api:"required"`
	Spoofing          bool     `json:"spoofing" api:"required"`
	Stop              bool     `json:"stop" api:"required"`
	Stream            bool     `json:"stream" api:"required"`
	Telephony         bool     `json:"telephony" api:"required"`
	Time              bool     `json:"time" api:"required"`
	TrafficInspection bool     `json:"trafficInspection" api:"required"`
	// Any of "direct".
	MediaInjectionTransport string `json:"mediaInjectionTransport"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accessibility           respjson.Field
		Agent                   respjson.Field
		Apps                    respjson.Field
		Browser                 respjson.Field
		CameraInjection         respjson.Field
		Clipboard               respjson.Field
		Esim                    respjson.Field
		FactoryReset            respjson.Field
		Files                   respjson.Field
		Fingerprint             respjson.Field
		Frida                   respjson.Field
		Geo                     respjson.Field
		HumanTouch              respjson.Field
		IdentityReset           respjson.Field
		Language                respjson.Field
		Logcat                  respjson.Field
		MicrophoneInjection     respjson.Field
		Proxy                   respjson.Field
		Reboot                  respjson.Field
		Recording               respjson.Field
		Reset                   respjson.Field
		ResetPreserve           respjson.Field
		SetGeo                  respjson.Field
		SetTime                 respjson.Field
		Shell                   respjson.Field
		Spoofing                respjson.Field
		Stop                    respjson.Field
		Stream                  respjson.Field
		Telephony               respjson.Field
		Time                    respjson.Field
		TrafficInspection       respjson.Field
		MediaInjectionTransport respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceGetCapabilitiesResponseCapabilities) RawJSON() string { return r.JSON.raw }
func (r *DeviceGetCapabilitiesResponseCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSetNameResponse struct {
	ID           string    `json:"id" api:"required"`
	ActiveTaskID string    `json:"activeTaskId" api:"required"`
	AssignedAt   time.Time `json:"assignedAt" api:"required" format:"date-time"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	Name         string    `json:"name" api:"required"`
	// Operating system the device runs.
	//
	// Any of "android", "ios".
	Platform     DeviceSetNameResponsePlatform `json:"platform" api:"required"`
	State        string                        `json:"state" api:"required"`
	StateMessage string                        `json:"stateMessage" api:"required"`
	StreamURL    string                        `json:"streamUrl" api:"required"`
	TaskCount    int64                         `json:"taskCount" api:"required"`
	TerminatesAt time.Time                     `json:"terminatesAt" api:"required" format:"date-time"`
	Type         string                        `json:"type" api:"required"`
	UpdatedAt    time.Time                     `json:"updatedAt" api:"required" format:"date-time"`
	// A URL to the JSON Schema for this object.
	Schema          string `json:"$schema" format:"uri"`
	BillingStrategy string `json:"billingStrategy"`
	CreatedBy       string `json:"createdBy"`
	// The platform default idle timeout this device uses when its own setting is
	// cleared (PUT idle-timeout with minutes null). Omitted when no default applies to
	// this device.
	IdleTimeoutDefaultMinutes int64 `json:"idleTimeoutDefaultMinutes"`
	// The effective idle timeout: minutes without client activity after which the
	// device is stopped automatically (state stopped, stopReason IDLE_TIMEOUT). It is
	// the device's own setting when one exists, otherwise the platform default for
	// pay-as-you-go (minute-billed) devices that can be stopped. Omitted while no idle
	// timer applies (explicitly turned off, or no default for this device). Change it
	// with PUT /devices/{deviceId}/idle-timeout.
	IdleTimeoutMinutes int64 `json:"idleTimeoutMinutes"`
	// Where idleTimeoutMinutes comes from: device = the device's own setting (also
	// when it is off), default = the platform default. Omitted when no idle timer can
	// apply.
	//
	// Any of "device", "default".
	IdleTimeoutSource DeviceSetNameResponseIdleTimeoutSource `json:"idleTimeoutSource"`
	// When the device last had client activity (RFC 3339). Tracked only while the idle
	// timer is on, to about one minute of precision; omitted otherwise.
	LastActivityAt time.Time `json:"lastActivityAt" format:"date-time"`
	OwnerID        string    `json:"ownerId"`
	ProviderID     string    `json:"providerId"`
	// Why the device was last stopped. Present only while state is stopped; clients
	// should use this value instead of parsing stateMessage.
	//
	// Any of "STOP_REQUESTED", "MINUTE_BALANCE_EXHAUSTED", "RESOURCE_EXHAUSTED",
	// "SUBSCRIPTION_REQUIRED", "PAYMENT_REQUIRED", "DEVICE_OPERATION_FAILED",
	// "DEVICE_LEFT", "IDLE_TIMEOUT".
	StopReason DeviceSetNameResponseStopReason `json:"stopReason"`
	// When devices-api recorded the stop reason (RFC 3339). Present only while state
	// is stopped and a reason is recorded.
	StopReasonAt time.Time `json:"stopReasonAt" format:"date-time"`
	StreamToken  string    `json:"streamToken"`
	// Deprecated: use ownerId/createdBy.
	//
	// Deprecated: deprecated
	UserID string `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		ActiveTaskID              respjson.Field
		AssignedAt                respjson.Field
		CreatedAt                 respjson.Field
		Name                      respjson.Field
		Platform                  respjson.Field
		State                     respjson.Field
		StateMessage              respjson.Field
		StreamURL                 respjson.Field
		TaskCount                 respjson.Field
		TerminatesAt              respjson.Field
		Type                      respjson.Field
		UpdatedAt                 respjson.Field
		Schema                    respjson.Field
		BillingStrategy           respjson.Field
		CreatedBy                 respjson.Field
		IdleTimeoutDefaultMinutes respjson.Field
		IdleTimeoutMinutes        respjson.Field
		IdleTimeoutSource         respjson.Field
		LastActivityAt            respjson.Field
		OwnerID                   respjson.Field
		ProviderID                respjson.Field
		StopReason                respjson.Field
		StopReasonAt              respjson.Field
		StreamToken               respjson.Field
		UserID                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceSetNameResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceSetNameResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Operating system the device runs.
type DeviceSetNameResponsePlatform string

const (
	DeviceSetNameResponsePlatformAndroid DeviceSetNameResponsePlatform = "android"
	DeviceSetNameResponsePlatformIos     DeviceSetNameResponsePlatform = "ios"
)

// Where idleTimeoutMinutes comes from: device = the device's own setting (also
// when it is off), default = the platform default. Omitted when no idle timer can
// apply.
type DeviceSetNameResponseIdleTimeoutSource string

const (
	DeviceSetNameResponseIdleTimeoutSourceDevice  DeviceSetNameResponseIdleTimeoutSource = "device"
	DeviceSetNameResponseIdleTimeoutSourceDefault DeviceSetNameResponseIdleTimeoutSource = "default"
)

// Why the device was last stopped. Present only while state is stopped; clients
// should use this value instead of parsing stateMessage.
type DeviceSetNameResponseStopReason string

const (
	DeviceSetNameResponseStopReasonStopRequested          DeviceSetNameResponseStopReason = "STOP_REQUESTED"
	DeviceSetNameResponseStopReasonMinuteBalanceExhausted DeviceSetNameResponseStopReason = "MINUTE_BALANCE_EXHAUSTED"
	DeviceSetNameResponseStopReasonResourceExhausted      DeviceSetNameResponseStopReason = "RESOURCE_EXHAUSTED"
	DeviceSetNameResponseStopReasonSubscriptionRequired   DeviceSetNameResponseStopReason = "SUBSCRIPTION_REQUIRED"
	DeviceSetNameResponseStopReasonPaymentRequired        DeviceSetNameResponseStopReason = "PAYMENT_REQUIRED"
	DeviceSetNameResponseStopReasonDeviceOperationFailed  DeviceSetNameResponseStopReason = "DEVICE_OPERATION_FAILED"
	DeviceSetNameResponseStopReasonDeviceLeft             DeviceSetNameResponseStopReason = "DEVICE_LEFT"
	DeviceSetNameResponseStopReasonIdleTimeout            DeviceSetNameResponseStopReason = "IDLE_TIMEOUT"
)

type DeviceSummaryResponse struct {
	ByState map[string]int64 `json:"byState" api:"required"`
	ByType  map[string]int64 `json:"byType" api:"required"`
	Total   int64            `json:"total" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByState     respjson.Field
		ByType      respjson.Field
		Total       respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceWaitReadyResponse struct {
	ID           string    `json:"id" api:"required"`
	ActiveTaskID string    `json:"activeTaskId" api:"required"`
	AssignedAt   time.Time `json:"assignedAt" api:"required" format:"date-time"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	Name         string    `json:"name" api:"required"`
	// Operating system the device runs.
	//
	// Any of "android", "ios".
	Platform     DeviceWaitReadyResponsePlatform `json:"platform" api:"required"`
	State        string                          `json:"state" api:"required"`
	StateMessage string                          `json:"stateMessage" api:"required"`
	StreamURL    string                          `json:"streamUrl" api:"required"`
	TaskCount    int64                           `json:"taskCount" api:"required"`
	TerminatesAt time.Time                       `json:"terminatesAt" api:"required" format:"date-time"`
	Type         string                          `json:"type" api:"required"`
	UpdatedAt    time.Time                       `json:"updatedAt" api:"required" format:"date-time"`
	// A URL to the JSON Schema for this object.
	Schema          string `json:"$schema" format:"uri"`
	BillingStrategy string `json:"billingStrategy"`
	CreatedBy       string `json:"createdBy"`
	// The platform default idle timeout this device uses when its own setting is
	// cleared (PUT idle-timeout with minutes null). Omitted when no default applies to
	// this device.
	IdleTimeoutDefaultMinutes int64 `json:"idleTimeoutDefaultMinutes"`
	// The effective idle timeout: minutes without client activity after which the
	// device is stopped automatically (state stopped, stopReason IDLE_TIMEOUT). It is
	// the device's own setting when one exists, otherwise the platform default for
	// pay-as-you-go (minute-billed) devices that can be stopped. Omitted while no idle
	// timer applies (explicitly turned off, or no default for this device). Change it
	// with PUT /devices/{deviceId}/idle-timeout.
	IdleTimeoutMinutes int64 `json:"idleTimeoutMinutes"`
	// Where idleTimeoutMinutes comes from: device = the device's own setting (also
	// when it is off), default = the platform default. Omitted when no idle timer can
	// apply.
	//
	// Any of "device", "default".
	IdleTimeoutSource DeviceWaitReadyResponseIdleTimeoutSource `json:"idleTimeoutSource"`
	// When the device last had client activity (RFC 3339). Tracked only while the idle
	// timer is on, to about one minute of precision; omitted otherwise.
	LastActivityAt time.Time `json:"lastActivityAt" format:"date-time"`
	OwnerID        string    `json:"ownerId"`
	ProviderID     string    `json:"providerId"`
	// Why the device was last stopped. Present only while state is stopped; clients
	// should use this value instead of parsing stateMessage.
	//
	// Any of "STOP_REQUESTED", "MINUTE_BALANCE_EXHAUSTED", "RESOURCE_EXHAUSTED",
	// "SUBSCRIPTION_REQUIRED", "PAYMENT_REQUIRED", "DEVICE_OPERATION_FAILED",
	// "DEVICE_LEFT", "IDLE_TIMEOUT".
	StopReason DeviceWaitReadyResponseStopReason `json:"stopReason"`
	// When devices-api recorded the stop reason (RFC 3339). Present only while state
	// is stopped and a reason is recorded.
	StopReasonAt time.Time `json:"stopReasonAt" format:"date-time"`
	StreamToken  string    `json:"streamToken"`
	// Deprecated: use ownerId/createdBy.
	//
	// Deprecated: deprecated
	UserID string `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		ActiveTaskID              respjson.Field
		AssignedAt                respjson.Field
		CreatedAt                 respjson.Field
		Name                      respjson.Field
		Platform                  respjson.Field
		State                     respjson.Field
		StateMessage              respjson.Field
		StreamURL                 respjson.Field
		TaskCount                 respjson.Field
		TerminatesAt              respjson.Field
		Type                      respjson.Field
		UpdatedAt                 respjson.Field
		Schema                    respjson.Field
		BillingStrategy           respjson.Field
		CreatedBy                 respjson.Field
		IdleTimeoutDefaultMinutes respjson.Field
		IdleTimeoutMinutes        respjson.Field
		IdleTimeoutSource         respjson.Field
		LastActivityAt            respjson.Field
		OwnerID                   respjson.Field
		ProviderID                respjson.Field
		StopReason                respjson.Field
		StopReasonAt              respjson.Field
		StreamToken               respjson.Field
		UserID                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceWaitReadyResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceWaitReadyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Operating system the device runs.
type DeviceWaitReadyResponsePlatform string

const (
	DeviceWaitReadyResponsePlatformAndroid DeviceWaitReadyResponsePlatform = "android"
	DeviceWaitReadyResponsePlatformIos     DeviceWaitReadyResponsePlatform = "ios"
)

// Where idleTimeoutMinutes comes from: device = the device's own setting (also
// when it is off), default = the platform default. Omitted when no idle timer can
// apply.
type DeviceWaitReadyResponseIdleTimeoutSource string

const (
	DeviceWaitReadyResponseIdleTimeoutSourceDevice  DeviceWaitReadyResponseIdleTimeoutSource = "device"
	DeviceWaitReadyResponseIdleTimeoutSourceDefault DeviceWaitReadyResponseIdleTimeoutSource = "default"
)

// Why the device was last stopped. Present only while state is stopped; clients
// should use this value instead of parsing stateMessage.
type DeviceWaitReadyResponseStopReason string

const (
	DeviceWaitReadyResponseStopReasonStopRequested          DeviceWaitReadyResponseStopReason = "STOP_REQUESTED"
	DeviceWaitReadyResponseStopReasonMinuteBalanceExhausted DeviceWaitReadyResponseStopReason = "MINUTE_BALANCE_EXHAUSTED"
	DeviceWaitReadyResponseStopReasonResourceExhausted      DeviceWaitReadyResponseStopReason = "RESOURCE_EXHAUSTED"
	DeviceWaitReadyResponseStopReasonSubscriptionRequired   DeviceWaitReadyResponseStopReason = "SUBSCRIPTION_REQUIRED"
	DeviceWaitReadyResponseStopReasonPaymentRequired        DeviceWaitReadyResponseStopReason = "PAYMENT_REQUIRED"
	DeviceWaitReadyResponseStopReasonDeviceOperationFailed  DeviceWaitReadyResponseStopReason = "DEVICE_OPERATION_FAILED"
	DeviceWaitReadyResponseStopReasonDeviceLeft             DeviceWaitReadyResponseStopReason = "DEVICE_LEFT"
	DeviceWaitReadyResponseStopReasonIdleTimeout            DeviceWaitReadyResponseStopReason = "IDLE_TIMEOUT"
)

type DeviceNewParams struct {
	DeviceSpec shared.DeviceSpecParam
	// ISO 3166-1 alpha-2 country code. If omitted the system picks the country with
	// the most availability.
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// Profile ID to use as device spec
	ProfileID param.Opt[string] `query:"profileId,omitzero" json:"-"`
	// Optional device template, e.g. onboarding. The device gets the template's
	// Android version and preinstalled apps and is ready in seconds when a prepared
	// phone is available; otherwise it boots with the same setup. Country, name and
	// proxy still come from this request. Cannot be combined with profileId, apps or
	// androidVersion. Omit it to create from the request alone.
	Template param.Opt[string] `query:"template,omitzero" json:"-"`
	// Billing mode. 'auto' tries subscription first, then minute billing if no
	// subscription entitlement exists or all subscription slots are in use, provided
	// minute billing is enabled. When subscription billing is disabled, auto uses
	// minutes directly. Billing-service failures never trigger fallback.
	// 'subscription' requires an available subscription slot and never falls back.
	// 'minute' uses minute billing only, subject to balance and concurrency checks.
	// Modes depend on the device type's billing configuration.
	//
	// Any of "auto", "subscription", "minute".
	Billing DeviceNewParamsBilling `query:"billing,omitzero" json:"-"`
	// Device type: android_cloud_phone (cloud Android phone), ios_phone or
	// device_slot. Any other value is rejected with 422. Not every device type is
	// available for creation; an unavailable type returns DEVICE_TYPE_UNAVAILABLE
	// (422).
	//
	// Any of "android_cloud_phone", "ios_phone", "device_slot".
	DeviceType DeviceNewParamsDeviceType `query:"deviceType,omitzero" json:"-"`
	paramObj
}

func (r DeviceNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DeviceSpec)
}
func (r *DeviceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [DeviceNewParams]'s query parameters as `url.Values`.
func (r DeviceNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Billing mode. 'auto' tries subscription first, then minute billing if no
// subscription entitlement exists or all subscription slots are in use, provided
// minute billing is enabled. When subscription billing is disabled, auto uses
// minutes directly. Billing-service failures never trigger fallback.
// 'subscription' requires an available subscription slot and never falls back.
// 'minute' uses minute billing only, subject to balance and concurrency checks.
// Modes depend on the device type's billing configuration.
type DeviceNewParamsBilling string

const (
	DeviceNewParamsBillingAuto         DeviceNewParamsBilling = "auto"
	DeviceNewParamsBillingSubscription DeviceNewParamsBilling = "subscription"
	DeviceNewParamsBillingMinute       DeviceNewParamsBilling = "minute"
)

// Device type: android_cloud_phone (cloud Android phone), ios_phone or
// device_slot. Any other value is rejected with 422. Not every device type is
// available for creation; an unavailable type returns DEVICE_TYPE_UNAVAILABLE
// (422).
type DeviceNewParamsDeviceType string

const (
	DeviceNewParamsDeviceTypeAndroidCloudPhone DeviceNewParamsDeviceType = "android_cloud_phone"
	DeviceNewParamsDeviceTypeIosPhone          DeviceNewParamsDeviceType = "ios_phone"
	DeviceNewParamsDeviceTypeDeviceSlot        DeviceNewParamsDeviceType = "device_slot"
)

type DeviceListParams struct {
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// Include devices created at or after this RFC3339 timestamp.
	CreatedAfter param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	// Include devices created before this RFC3339 timestamp (exclusive). Must be later
	// than createdAfter when both are set.
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// Filter to devices created by this user id. Mutually exclusive with mine.
	CreatedBy param.Opt[string] `query:"createdBy,omitzero" json:"-"`
	// When true, only return devices created by the authenticated caller.
	Mine       param.Opt[bool]   `query:"mine,omitzero" json:"-"`
	Name       param.Opt[string] `query:"name,omitzero" json:"-"`
	Page       param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize   param.Opt[int64]  `query:"pageSize,omitzero" json:"-"`
	ProviderID param.Opt[string] `query:"providerId,omitzero" json:"-"`
	// Case-insensitive search: matches devices whose name contains the term or whose
	// ID starts with it. Combined with all other filters.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Filter by current device state. Comma-separated (state=ready,stopped). Defaults
	// to every state except terminated and unknown, so stopped and maintenance devices
	// are listed; pass state=terminated or state=unknown to see those.
	//
	// Any of "creating", "assigned", "ready", "rebooting", "migrating", "resetting",
	// "terminated", "maintenance", "stopped", "unknown".
	State []string `query:"state,omitzero" json:"-"`
	// Any of "id", "createdAt", "updatedAt", "assignedAt", "gridPosition".
	OrderBy DeviceListParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// Any of "asc", "desc".
	OrderByDirection DeviceListParamsOrderByDirection `query:"orderByDirection,omitzero" json:"-"`
	// Filter by the device's platform as served in the platform field.
	//
	// Any of "android", "ios".
	Platform DeviceListParamsPlatform `query:"platform,omitzero" json:"-"`
	// Filter by device type.
	//
	// Any of "android_cloud_phone", "ios_phone", "device_slot".
	Type DeviceListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceListParams]'s query parameters as `url.Values`.
func (r DeviceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceListParamsOrderBy string

const (
	DeviceListParamsOrderByID           DeviceListParamsOrderBy = "id"
	DeviceListParamsOrderByCreatedAt    DeviceListParamsOrderBy = "createdAt"
	DeviceListParamsOrderByUpdatedAt    DeviceListParamsOrderBy = "updatedAt"
	DeviceListParamsOrderByAssignedAt   DeviceListParamsOrderBy = "assignedAt"
	DeviceListParamsOrderByGridPosition DeviceListParamsOrderBy = "gridPosition"
)

type DeviceListParamsOrderByDirection string

const (
	DeviceListParamsOrderByDirectionAsc  DeviceListParamsOrderByDirection = "asc"
	DeviceListParamsOrderByDirectionDesc DeviceListParamsOrderByDirection = "desc"
)

// Filter by the device's platform as served in the platform field.
type DeviceListParamsPlatform string

const (
	DeviceListParamsPlatformAndroid DeviceListParamsPlatform = "android"
	DeviceListParamsPlatformIos     DeviceListParamsPlatform = "ios"
)

// Filter by device type.
type DeviceListParamsType string

const (
	DeviceListParamsTypeAndroidCloudPhone DeviceListParamsType = "android_cloud_phone"
	DeviceListParamsTypeIosPhone          DeviceListParamsType = "ios_phone"
	DeviceListParamsTypeDeviceSlot        DeviceListParamsType = "device_slot"
)

type DeviceFingerprintParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

type DeviceResetParams struct {
	// Device configuration to restore after the reset: the latest value set through
	// the API. Omitted, null or empty preserves nothing.
	//
	// Any of "proxy", "location", "locale", "timezone".
	Preserve []string `json:"preserve,omitzero"`
	// What to erase. full: factory reset — erases the user data (apps, accounts,
	// storage) and gives the device a new identity. identity: new device fingerprint,
	// apps and data are kept. When omitted: full, or identity on devices whose
	// capabilities report factoryReset=false.
	//
	// Any of "full", "identity".
	Mode DeviceResetParamsMode `json:"mode,omitzero"`
	paramObj
}

func (r DeviceResetParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceResetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceResetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What to erase. full: factory reset — erases the user data (apps, accounts,
// storage) and gives the device a new identity. identity: new device fingerprint,
// apps and data are kept. When omitted: full, or identity on devices whose
// capabilities report factoryReset=false.
type DeviceResetParamsMode string

const (
	DeviceResetParamsModeFull     DeviceResetParamsMode = "full"
	DeviceResetParamsModeIdentity DeviceResetParamsMode = "identity"
)

type DeviceSetNameParams struct {
	Name string `json:"name" api:"required"`
	paramObj
}

func (r DeviceSetNameParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceSetNameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceSetNameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSummaryParams struct {
	// Any of "creating", "assigned", "ready", "rebooting", "migrating", "resetting",
	// "terminated", "maintenance", "stopped", "unknown".
	State []string `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceSummaryParams]'s query parameters as `url.Values`.
func (r DeviceSummaryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceTerminateParams struct {
	PreviousDeviceID param.Opt[string]    `json:"previousDeviceId,omitzero"`
	TerminateAt      param.Opt[time.Time] `json:"terminateAt,omitzero" format:"date-time"`
	paramObj
}

func (r DeviceTerminateParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceTerminateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceTerminateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
