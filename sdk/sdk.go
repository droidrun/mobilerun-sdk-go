// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sdk

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/pagination"
	"github.com/droidrun/mobilerun-sdk-go/v6/shared"
)

// CredentialsNamespace adds object managers to the raw CredentialService; raw methods stay available.
type CredentialsNamespace struct {
	*mobilerun.CredentialService
	Packages *CredentialPackageManager
}

// ComputeNamespace adds object managers to the raw ComputeService; raw methods stay available.
type ComputeNamespace struct {
	*mobilerun.ComputeService
	Functions *ComputeFunctionManager
	Runs      *RunManager
	Sandboxes *SandboxManager
}

// ConnectNamespace adds object managers to the raw ConnectService; raw methods stay available.
type ConnectNamespace struct {
	*mobilerun.ConnectService
	Proxies *ConnectProxyManager
	Users   *ConnectUserManager
}

// WorkflowsNamespace adds object managers to the raw WorkflowService; raw methods stay available.
type WorkflowsNamespace struct {
	*mobilerun.WorkflowService
	Triggers      *TriggerManager
	ActionCatalog *CatalogEntryManager
	Actions       *WorkflowActionManager
	Flows         *FlowManager
	Executions    *ExecutionManager
}

// MobilerunSDK is the object API entry point. API is the complete raw client; the
// other fields return object managers (or raw services without objects).
type MobilerunSDK struct {
	API          *mobilerun.Client
	Apps         *AppManager
	Credentials  *CredentialsNamespace
	Devices      *DeviceManager
	StreamTokens *mobilerun.StreamTokenService
	Identity     *mobilerun.IdentityService
	Models       *mobilerun.ModelService
	Proxies      *ProxyConfigurationManager
	Compute      *ComputeNamespace
	Connect      *ConnectNamespace
	Tasks        *TaskManager
	Workflows    *WorkflowsNamespace
	Webhooks     *WebhookManager
	Mailboxes    *MailboxManager
	RoutineFiles *RoutineFileManager
	Files        *AgentFileManager
	Agents       *AgentManager
	Assistant    *mobilerun.AssistantService
	Messages     *NumberMessageManager
	Numbers      *PhoneNumberManager
	Store        *mobilerun.StoreService
}

// NewMobilerunSDK creates the object API and its raw client. Options are the raw
// client's options, read from the environment as for mobilerun.NewClient.
func NewMobilerunSDK(opts ...option.RequestOption) *MobilerunSDK {
	raw := mobilerun.NewClient(opts...)
	api := &raw
	c := &clientRef{api: api, options: append(mobilerun.DefaultClientOptions(), opts...)}
	s := &MobilerunSDK{API: api}
	s.Apps = newAppManager(c)
	s.Credentials = &CredentialsNamespace{CredentialService: &api.Credentials, Packages: newCredentialPackageManager(c)}
	s.Devices = newDeviceManager(c)
	s.StreamTokens = &api.StreamTokens
	s.Identity = &api.Identity
	s.Models = &api.Models
	s.Proxies = newProxyConfigurationManager(c)
	s.Compute = &ComputeNamespace{ComputeService: &api.Compute, Functions: newComputeFunctionManager(c), Runs: newRunManager(c), Sandboxes: newSandboxManager(c)}
	s.Connect = &ConnectNamespace{ConnectService: &api.Connect, Proxies: newConnectProxyManager(c), Users: newConnectUserManager(c)}
	s.Tasks = newTaskManager(c)
	s.Workflows = &WorkflowsNamespace{WorkflowService: &api.Workflows, Triggers: newTriggerManager(c), ActionCatalog: newCatalogEntryManager(c), Actions: newWorkflowActionManager(c), Flows: newFlowManager(c), Executions: newExecutionManager(c)}
	s.Webhooks = newWebhookManager(c)
	s.Mailboxes = newMailboxManager(c)
	s.RoutineFiles = newRoutineFileManager(c)
	s.Files = newAgentFileManager(c)
	s.Agents = newAgentManager(c)
	s.Assistant = &api.Assistant
	s.Messages = newNumberMessageManager(c)
	s.Numbers = newPhoneNumberManager(c)
	s.Store = &api.Store
	_ = c
	return s
}

func (s *MobilerunSDK) String() string { return objectString("MobilerunSDK", nil) }

// VerifyWebhook verifies a signed delivery and returns its parsed JSON body.
// Pass the exact raw body bytes before JSON parsing. No API client or API key is
// needed. Deduplicate replays inside the tolerance window by an ID in the signed
// body.
func VerifyWebhook(body []byte, headers http.Header, secret string, opts *SignatureVerifyOptions) (any, error) {
	return verifySignature(body, headers, secret, opts, signatureScheme{
		signatureHeader: "x-mobilerun-signature",
		timestampHeader: "x-mobilerun-timestamp",
		prefix:          "sha256=",
		unit:            "ms",
		tolerance:       300000 * time.Millisecond,
	})
}

// Screenshot is a binary response with declared headers.
type Screenshot struct {
	BinaryResult
	// ID is the X-Screenshot-Id response header.
	ID string
	// Saved is the X-Screenshot-Saved response header (nil when absent).
	Saved *bool
}

// AppManager creates, fetches and lists App handles.
type AppManager struct {
	c *clientRef
}

func newAppManager(c *clientRef) *AppManager {
	m := &AppManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *AppManager) FromID(id string) *App { return newApp(m.c, id) }

func (m *AppManager) String() string { return objectString("AppManager", nil) }

// Retrieves an app by its ID
//
// Get fetches one App and returns its handle.
func (m *AppManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*App, error) {
	res, err := m.c.api.Apps.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newApp(m.c, id).withInfo(res, "data"), nil
}

// Retrieves a paginated list of apps with filtering and search capabilities
//
// List returns the first page of App handles; ListAutoPaging iterates every page.
func (m *AppManager) List(ctx context.Context, params mobilerun.AppListParams, opts ...option.RequestOption) (*ObjectPage[*App, *pagination.Page[mobilerun.AppListResponse]], error) {
	var wrap func(page *pagination.Page[mobilerun.AppListResponse]) *ObjectPage[*App, *pagination.Page[mobilerun.AppListResponse]]
	wrap = func(page *pagination.Page[mobilerun.AppListResponse]) *ObjectPage[*App, *pagination.Page[mobilerun.AppListResponse]] {
		result := &ObjectPage[*App, *pagination.Page[mobilerun.AppListResponse]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newApp(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*App, *pagination.Page[mobilerun.AppListResponse]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Apps.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the App handles of every page.
func (m *AppManager) ListAutoPaging(ctx context.Context, params mobilerun.AppListParams, opts ...option.RequestOption) *ObjectAutoPager[*App] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// Upload uploads files to presigned storage, confirms the upload and returns the
// App. Uploading uses storage quota and may incur charges; the returned handle
// does not own cleanup. Failures before confirmation clean up the pending upload
// once, best-effort, and return an UploadError.
func (m *AppManager) Upload(ctx context.Context, files []UploadFile, params mobilerun.AppNewSignedUploadURLParams, opts ...option.RequestOption) (*App, error) {
	prepared, err := prepareUploads(files, []string{"application/vnd.android.package-archive", "application/octet-stream", "application/zip"})
	if err != nil {
		return nil, err
	}
	rows := make([]map[string]any, len(prepared))
	for i, file := range prepared {
		rows[i] = map[string]any{"fileName": file.filename, "contentType": file.contentType}
	}
	requestOpts := bindOptions(opts, binding{"files", "body", rows})
	issued, err := m.c.api.Apps.NewSignedUploadURL(ctx, params, requestOpts...)
	if err != nil {
		return nil, err
	}
	uploadID := objectID(issued, "versionId")
	confirming, confirmedOK := false, false
	cleanup := func(cause error) error {
		if !confirmedOK && (!confirming || definiteConfirmFailure(cause)) {
			cleanupCtx, cancel := cleanupContext(ctx)
			_, _ = m.c.api.Apps.MarkFailed(cleanupCtx, uploadID, opts...)
			cancel()
		}
		if ctx.Err() != nil {
			return cause
		}
		return &UploadError{Cause: cause}
	}
	urls, err := uploadURLs(issued, prepared, "", "r2UploadUrls", "fileName", "r2UploadUrl")
	if err != nil {
		return nil, cleanup(err)
	}
	for i, file := range prepared {
		if err := putUpload(ctx, m.c, file, urls[i], opts); err != nil {
			return nil, cleanup(err)
		}
	}
	confirming = true
	_, err = m.c.api.Apps.ConfirmUpload(ctx, uploadID, opts...)
	if err != nil {
		return nil, cleanup(err)
	}
	confirmedOK = true
	targetID := objectID(issued, "appId")
	info, err := m.c.api.Apps.Get(ctx, targetID, opts...)
	if err != nil {
		return nil, cleanup(err)
	}
	return newApp(m.c, objectID(info, "data.id")).withInfo(info, "data"), nil
}

// Creates or updates an app and returns pre-signed upload URLs for each file.
// Honours an optional `Idempotency-Key` header: a retry with the same key and body
// returns the same app and version with freshly signed URLs and
// `Idempotent-Replayed: true`, and never creates a second version. An
// Idempotency-Key is honoured for 7 days; after that it behaves like a fresh
// request. A replay whose version was deleted is a 404; a replay whose version
// already left the queued state is a 409 `idempotency_key_consumed`; a replay that
// cannot get the version lock within 5 s is a 409
// `idempotency_request_in_progress`. Every response produced by a replay,
// including those 404/409 errors, carries `Idempotent-Replayed: true`. A same-key
// request still in flight after a 5 s wait is a 409
// `idempotency_request_in_progress` (retry after `Retry-After`). If signing fails,
// nothing is persisted (with or without a key) and the same request, key included,
// can be retried.
func (m *AppManager) NewSignedUploadURL(ctx context.Context, params mobilerun.AppNewSignedUploadURLParams, opts ...option.RequestOption) (*mobilerun.AppNewSignedUploadURLResponse, error) {
	return m.c.api.Apps.NewSignedUploadURL(ctx, params, opts...)
}

// Returns the user’s storage allowance in bytes: bytes used, bytes included in the
// plan, the hard maximum (included + max purchasable overage; null = unlimited)
// and whether usage above the included amount is billed as overage. With storage
// billing off there is no limit: only usedBytes is set.
func (m *AppManager) StorageUsage(ctx context.Context, opts ...option.RequestOption) (*mobilerun.AppStorageUsageResponse, error) {
	return m.c.api.Apps.StorageUsage(ctx, opts...)
}

// CredentialPackageManager creates, fetches and lists CredentialPackage handles.
type CredentialPackageManager struct {
	c *clientRef
}

func newCredentialPackageManager(c *clientRef) *CredentialPackageManager {
	m := &CredentialPackageManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *CredentialPackageManager) FromID(id string) *CredentialPackage {
	return newCredentialPackage(m.c, id)
}

func (m *CredentialPackageManager) String() string {
	return objectString("CredentialPackageManager", nil)
}

// Returns all credentials stored under the given `packageName`. Each credential
// includes its name, secret path, and the list of fields it holds.
//
// Get fetches one CredentialPackage and returns its handle.
func (m *CredentialPackageManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*CredentialPackage, error) {
	res, err := m.c.api.Credentials.Packages.List(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newCredentialPackage(m.c, firstRowID(res, "data", "packageName", id)).withInfo(res, ""), nil
}

// Creates a new package (identified by `packageName`) under which credentials can
// be grouped. Returns a conflict if a package with the same name already exists
// for the user.
//
// New creates a CredentialPackage. The returned handle owns cleanup: Close
// disposes it.
func (m *CredentialPackageManager) New(ctx context.Context, params mobilerun.CredentialPackageNewParams, opts ...option.RequestOption) (*CredentialPackage, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Credentials.Packages.New(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newCredentialPackage(m.c, objectID(res, "data.packageName")).withInfo(res, "data")
	o.state.setOwned(true)
	return o, nil
}

// Returns the names of all packages (apps) the authenticated owner has credentials
// grouped under. Use this to discover which `packageName` values are valid for the
// per-package credential routes.
//
// List returns the first page of CredentialPackage handles; ListAutoPaging
// iterates every page.
func (m *CredentialPackageManager) List(ctx context.Context, opts ...option.RequestOption) (*ObjectPage[*CredentialPackage, *mobilerun.CredentialPackageListAllResponse], error) {
	load := func() (*ObjectPage[*CredentialPackage, *mobilerun.CredentialPackageListAllResponse], error) {
		pageOpts := opts
		res, err := m.c.api.Credentials.Packages.ListAll(ctx, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*CredentialPackage, *mobilerun.CredentialPackageListAllResponse]{Response: res}
		for _, info := range objectRows(res, "data") {
			result.Items = append(result.Items, newCredentialPackage(m.c, objectID(info, "packageName")).withInfo(info, ""))
		}
		return result, nil
	}
	return load()
}

// ListAutoPaging iterates the CredentialPackage handles of every page.
func (m *CredentialPackageManager) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *ObjectAutoPager[*CredentialPackage] {
	page, err := m.List(ctx, opts...)
	return newObjectAutoPager(page, err)
}

// CredentialManager creates, fetches and lists Credential handles.
type CredentialManager struct {
	c           *clientRef
	packageName string
}

func newCredentialManager(c *clientRef, packageName string) *CredentialManager {
	m := &CredentialManager{c: c, packageName: packageName}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *CredentialManager) FromID(id string) *Credential {
	return newCredential(m.c, m.packageName, id)
}

func (m *CredentialManager) String() string { return objectString("CredentialManager", nil) }

// Creates a credential under the given package with a `credentialName` and at
// least one field. Field types must be unique. Each field has a `fieldType`
// (email, username, password, api_token, phone_number, two_factor_secret, or
// backup_codes) and a value.
//
// New creates a Credential. The returned handle owns cleanup: Close disposes it.
func (m *CredentialManager) New(ctx context.Context, params mobilerun.CredentialPackageCredentialNewParams, opts ...option.RequestOption) (*Credential, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Credentials.Packages.Credentials.New(ctx, m.packageName, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newCredential(m.c, m.packageName, objectID(res, "data.credentialName")).withInfo(res, "data")
	o.state.setOwned(true)
	return o, nil
}

// Returns all credentials stored under the given `packageName`. Each credential
// includes its name, secret path, and the list of fields it holds.
//
// List returns the first page of Credential handles; ListAutoPaging iterates every
// page.
func (m *CredentialManager) List(ctx context.Context, opts ...option.RequestOption) (*ObjectPage[*Credential, *mobilerun.CredentialPackageListResponse], error) {
	load := func() (*ObjectPage[*Credential, *mobilerun.CredentialPackageListResponse], error) {
		pageOpts := opts
		res, err := m.c.api.Credentials.Packages.List(ctx, m.packageName, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*Credential, *mobilerun.CredentialPackageListResponse]{Response: res}
		for _, info := range objectRows(res, "data") {
			result.Items = append(result.Items, newCredential(m.c, m.packageName, objectID(info, "credentialName")).withInfo(info, ""))
		}
		return result, nil
	}
	return load()
}

// ListAutoPaging iterates the Credential handles of every page.
func (m *CredentialManager) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *ObjectAutoPager[*Credential] {
	page, err := m.List(ctx, opts...)
	return newObjectAutoPager(page, err)
}

// PackageName returns the bound CredentialPackage ID.
func (m *CredentialManager) PackageName() string { return m.packageName }

// DeviceManager creates, fetches and lists Device handles.
type DeviceManager struct {
	c *clientRef

	NumberBindings *DeviceManagerNumberBindings
}

func newDeviceManager(c *clientRef) *DeviceManager {
	m := &DeviceManager{c: c}
	m.NumberBindings = newDeviceManagerNumberBindings(m)
	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *DeviceManager) FromID(id string) *Device { return newDevice(m.c, id) }

func (m *DeviceManager) String() string { return objectString("DeviceManager", nil) }

// Returns the current state and metadata for a single device, including its
// lifecycle state, type, stream URL, billing strategy, and timestamps. A stream
// token is included while the device is active, but may be omitted when it cannot
// be read at that moment. Refetch the device to obtain the stream token.
//
// Get fetches one Device and returns its handle.
func (m *DeviceManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*Device, error) {
	res, err := m.c.api.Devices.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newDevice(m.c, id).withInfo(res, ""), nil
}

// Starts billing for the device according to the selected billing mode.
//
// Requests a new device from the specification in the request body. Optional query
// parameters select the canonical device type, country, billing mode, and base
// profile. Returns the device, its resolved billing strategy, and its stream
// token.
//
// New creates a Device. The returned handle owns cleanup: Close disposes it. It
// waits until the object is ready (WithoutWaitReady, WithReadyTimeout); on failure
// it returns the handle with an ObjectReadyError after disposing the object once
// (WithoutDisposeOnReadyFailure keeps it).
func (m *DeviceManager) New(ctx context.Context, params mobilerun.DeviceNewParams, opts ...option.RequestOption) (*Device, error) {
	controls, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Devices.New(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newDevice(m.c, objectID(res, "id")).withInfo(res, "")
	o.state.setOwned(true)
	if !controls.noWaitReady {
		if _, err := o.WaitReady(ctx, readyOptions(controls, opts)...); err != nil {
			readyErr := &ObjectReadyError{Object: o, ID: o.id, Cause: err}
			if !controls.keepOnReadyFailure {
				cleanupCtx, cancel := cleanupContext(ctx)
				cleanupErr := o.Close(cleanupCtx, opts...)
				cancel()
				if cleanupErr != nil {
					readyErr.CleanupErr = cleanupErr
				} else {
					readyErr.Disposed = true
				}
			}
			return o, readyErr
		}
	}
	return o, nil
}

// Returns a paginated list of the user's devices along with pagination metadata.
// By default every device except terminated and unknown-state ones is listed,
// including stopped and maintenance devices; narrow with e.g. `state=ready` or
// `state=stopped,maintenance`, and pass `state=terminated` or `state=unknown` to
// see those.
//
// List returns the first page of Device handles; ListAutoPaging iterates every
// page.
func (m *DeviceManager) List(ctx context.Context, params mobilerun.DeviceListParams, opts ...option.RequestOption) (*ObjectPage[*Device, *pagination.Page[mobilerun.Device]], error) {
	var wrap func(page *pagination.Page[mobilerun.Device]) *ObjectPage[*Device, *pagination.Page[mobilerun.Device]]
	wrap = func(page *pagination.Page[mobilerun.Device]) *ObjectPage[*Device, *pagination.Page[mobilerun.Device]] {
		result := &ObjectPage[*Device, *pagination.Page[mobilerun.Device]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newDevice(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*Device, *pagination.Page[mobilerun.Device]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Devices.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the Device handles of every page.
func (m *DeviceManager) ListAutoPaging(ctx context.Context, params mobilerun.DeviceListParams, opts ...option.RequestOption) *ObjectAutoPager[*Device] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// Returns the total number of the user's devices and counts grouped by current
// state and device type.
func (m *DeviceManager) Summary(ctx context.Context, params mobilerun.DeviceSummaryParams, opts ...option.RequestOption) (*mobilerun.DeviceSummaryResponse, error) {
	return m.c.api.Devices.Summary(ctx, params, opts...)
}

// DeviceManagerNumberBindings groups the number_bindings methods of a DeviceManager.
type DeviceManagerNumberBindings struct {
	m *DeviceManager
}

func newDeviceManagerNumberBindings(m *DeviceManager) *DeviceManagerNumberBindings {
	n := &DeviceManagerNumberBindings{m: m}

	return n
}

func (n *DeviceManagerNumberBindings) String() string {
	return objectString("DeviceManagerNumberBindings", nil)
}

func (n *DeviceManagerNumberBindings) List(ctx context.Context, opts ...option.RequestOption) (*mobilerun.DeviceNumberBindingListResponse, error) {
	m := n.m
	return m.c.api.Devices.NumberBindings.List(ctx, opts...)
}

func (n *DeviceManagerNumberBindings) ListCandidates(ctx context.Context, opts ...option.RequestOption) (*mobilerun.DeviceNumberBindingListCandidatesResponse, error) {
	m := n.m
	return m.c.api.Devices.NumberBindings.ListCandidates(ctx, opts...)
}

// SavedScreenshotManager creates, fetches and lists SavedScreenshot handles.
type SavedScreenshotManager struct {
	c        *clientRef
	deviceID string
}

func newSavedScreenshotManager(c *clientRef, deviceID string) *SavedScreenshotManager {
	m := &SavedScreenshotManager{c: c, deviceID: deviceID}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *SavedScreenshotManager) FromID(id string) *SavedScreenshot {
	return newSavedScreenshot(m.c, m.deviceID, id)
}

func (m *SavedScreenshotManager) String() string { return objectString("SavedScreenshotManager", nil) }

// Saved screenshots, newest first. Pass nextCursor as cursor to retrieve the next
// page.
//
// List returns the first page of SavedScreenshot handles; ListAutoPaging iterates
// every page.
func (m *SavedScreenshotManager) List(ctx context.Context, params mobilerun.DeviceScreenshotListParams, opts ...option.RequestOption) (*ObjectPage[*SavedScreenshot, *mobilerun.DeviceScreenshotListResponse], error) {
	seen := map[string]bool{}
	var load func(page string) (*ObjectPage[*SavedScreenshot, *mobilerun.DeviceScreenshotListResponse], error)
	load = func(page string) (*ObjectPage[*SavedScreenshot, *mobilerun.DeviceScreenshotListResponse], error) {
		pageOpts := opts
		if page != "" {
			seen[page] = true
			pageOpts = bindOptions(opts, binding{"cursor", "query", page})
		}
		res, err := m.c.api.Devices.Screenshots.List(ctx, m.deviceID, params, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*SavedScreenshot, *mobilerun.DeviceScreenshotListResponse]{Response: res}
		for _, info := range objectRows(res, "items") {
			result.Items = append(result.Items, newSavedScreenshot(m.c, m.deviceID, objectID(info, "id")).withInfo(info, ""))
		}
		next, ok, err := nextObjectCursor(res, "nextCursor", seen)
		if err != nil {
			return nil, err
		}
		if ok {
			result.next = func() (*ObjectPage[*SavedScreenshot, *mobilerun.DeviceScreenshotListResponse], error) {
				return load(next)
			}
		}
		return result, nil
	}
	return load("")
}

// ListAutoPaging iterates the SavedScreenshot handles of every page.
func (m *SavedScreenshotManager) ListAutoPaging(ctx context.Context, params mobilerun.DeviceScreenshotListParams, opts ...option.RequestOption) *ObjectAutoPager[*SavedScreenshot] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// DeviceID returns the bound Device ID.
func (m *SavedScreenshotManager) DeviceID() string { return m.deviceID }

// RecordingManager creates, fetches and lists Recording handles.
type RecordingManager struct {
	c        *clientRef
	deviceID string
}

func newRecordingManager(c *clientRef, deviceID string) *RecordingManager {
	m := &RecordingManager{c: c, deviceID: deviceID}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *RecordingManager) FromID(id string) *Recording { return newRecording(m.c, m.deviceID, id) }

func (m *RecordingManager) String() string { return objectString("RecordingManager", nil) }

// Get a device recording
//
// Get fetches one Recording and returns its handle.
//
// The object layer sets deviceId; values for it in params are replaced.
func (m *RecordingManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*Recording, error) {
	res, err := m.c.api.Devices.Recordings.Status(ctx, id, mobilerun.DeviceRecordingStatusParams{DeviceID: m.deviceID}, opts...)
	if err != nil {
		return nil, err
	}
	return newRecording(m.c, m.deviceID, id).withInfo(res, ""), nil
}

// Start a device recording
//
// Start creates a Recording. The returned handle owns cleanup: Close disposes it.
func (m *RecordingManager) Start(ctx context.Context, params mobilerun.DeviceRecordingStartParams, opts ...option.RequestOption) (*Recording, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Devices.Recordings.Start(ctx, m.deviceID, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newRecording(m.c, m.deviceID, objectID(res, "id")).withInfo(res, "")
	o.state.setOwned(true)
	return o, nil
}

// List device recordings
//
// List returns the first page of Recording handles; ListAutoPaging iterates every
// page.
func (m *RecordingManager) List(ctx context.Context, params mobilerun.DeviceRecordingListParams, opts ...option.RequestOption) (*ObjectPage[*Recording, *[]mobilerun.DeviceRecordingListResponse], error) {
	load := func() (*ObjectPage[*Recording, *[]mobilerun.DeviceRecordingListResponse], error) {
		pageOpts := opts
		res, err := m.c.api.Devices.Recordings.List(ctx, m.deviceID, params, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*Recording, *[]mobilerun.DeviceRecordingListResponse]{Response: res}
		for _, info := range objectRows(res, "") {
			result.Items = append(result.Items, newRecording(m.c, m.deviceID, objectID(info, "id")).withInfo(info, ""))
		}
		return result, nil
	}
	return load()
}

// ListAutoPaging iterates the Recording handles of every page.
func (m *RecordingManager) ListAutoPaging(ctx context.Context, params mobilerun.DeviceRecordingListParams, opts ...option.RequestOption) *ObjectAutoPager[*Recording] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// DeviceID returns the bound Device ID.
func (m *RecordingManager) DeviceID() string { return m.deviceID }

// TrafficSessionManager creates, fetches and lists TrafficSession handles.
type TrafficSessionManager struct {
	c        *clientRef
	deviceID string
}

func newTrafficSessionManager(c *clientRef, deviceID string) *TrafficSessionManager {
	m := &TrafficSessionManager{c: c, deviceID: deviceID}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *TrafficSessionManager) FromID(id string) *TrafficSession {
	return newTrafficSession(m.c, m.deviceID, id)
}

func (m *TrafficSessionManager) String() string { return objectString("TrafficSessionManager", nil) }

// Returns status and the device stream credential for a live session.
//
// Get fetches one TrafficSession and returns its handle.
//
// The object layer sets deviceId; values for it in params are replaced.
func (m *TrafficSessionManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*TrafficSession, error) {
	res, err := m.c.api.Devices.TrafficSessions.Get(ctx, id, mobilerun.DeviceTrafficSessionGetParams{DeviceID: m.deviceID}, opts...)
	if err != nil {
		return nil, err
	}
	return newTrafficSession(m.c, m.deviceID, id).withInfo(res, ""), nil
}

// Starts one live-only decoded HTTP/1.1, HTTP/2, HTTP/3 and WebSocket traffic
// session.
//
// New creates a TrafficSession. The returned handle owns cleanup: Close disposes
// it.
func (m *TrafficSessionManager) New(ctx context.Context, params mobilerun.DeviceTrafficSessionNewParams, opts ...option.RequestOption) (*TrafficSession, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Devices.TrafficSessions.New(ctx, m.deviceID, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newTrafficSession(m.c, m.deviceID, objectID(res, "id")).withInfo(res, "")
	o.state.setOwned(true)
	return o, nil
}

// List device traffic sessions
//
// List returns the first page of TrafficSession handles; ListAutoPaging iterates
// every page.
func (m *TrafficSessionManager) List(ctx context.Context, params mobilerun.DeviceTrafficSessionListParams, opts ...option.RequestOption) (*ObjectPage[*TrafficSession, *pagination.Page[mobilerun.DeviceTrafficSessionListResponse]], error) {
	var wrap func(page *pagination.Page[mobilerun.DeviceTrafficSessionListResponse]) *ObjectPage[*TrafficSession, *pagination.Page[mobilerun.DeviceTrafficSessionListResponse]]
	wrap = func(page *pagination.Page[mobilerun.DeviceTrafficSessionListResponse]) *ObjectPage[*TrafficSession, *pagination.Page[mobilerun.DeviceTrafficSessionListResponse]] {
		result := &ObjectPage[*TrafficSession, *pagination.Page[mobilerun.DeviceTrafficSessionListResponse]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newTrafficSession(m.c, m.deviceID, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*TrafficSession, *pagination.Page[mobilerun.DeviceTrafficSessionListResponse]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Devices.TrafficSessions.List(ctx, m.deviceID, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the TrafficSession handles of every page.
func (m *TrafficSessionManager) ListAutoPaging(ctx context.Context, params mobilerun.DeviceTrafficSessionListParams, opts ...option.RequestOption) *ObjectAutoPager[*TrafficSession] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// DeviceID returns the bound Device ID.
func (m *TrafficSessionManager) DeviceID() string { return m.deviceID }

// ProxyConfigurationManager creates, fetches and lists ProxyConfiguration handles.
type ProxyConfigurationManager struct {
	c *clientRef
}

func newProxyConfigurationManager(c *clientRef) *ProxyConfigurationManager {
	m := &ProxyConfigurationManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *ProxyConfigurationManager) FromID(id string) *ProxyConfiguration {
	return newProxyConfiguration(m.c, id)
}

func (m *ProxyConfigurationManager) String() string {
	return objectString("ProxyConfigurationManager", nil)
}

// Fetches a single proxy config by its `proxyId`. The response shape depends on
// the proxy's `protocol` (socks5 or wireguard). Returns not found if no matching
// config exists.
//
// Get fetches one ProxyConfiguration and returns its handle.
func (m *ProxyConfigurationManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*ProxyConfiguration, error) {
	res, err := m.c.api.Proxies.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newProxyConfiguration(m.c, id).withInfo(res, "data"), nil
}

// Creates a proxy config. The body is a discriminated union on `protocol`:
// `socks5` requires name, host, port, user, and password, while `wireguard`
// requires name and a config string. Returns the created config with its generated
// `proxyId`.
//
// New creates a ProxyConfiguration. The returned handle owns cleanup: Close
// disposes it.
func (m *ProxyConfigurationManager) New(ctx context.Context, params mobilerun.ProxyNewParams, opts ...option.RequestOption) (*ProxyConfiguration, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Proxies.New(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newProxyConfiguration(m.c, objectID(res, "data.proxyId")).withInfo(res, "data")
	o.state.setOwned(true)
	return o, nil
}

// Returns all proxy configs for the authenticated user. An optional `protocol`
// query parameter (`socks5` or `wireguard`) filters the results by proxy type.
//
// List returns the first page of ProxyConfiguration handles; ListAutoPaging
// iterates every page.
func (m *ProxyConfigurationManager) List(ctx context.Context, params mobilerun.ProxyListParams, opts ...option.RequestOption) (*ObjectPage[*ProxyConfiguration, *mobilerun.ProxyListResponse], error) {
	load := func() (*ObjectPage[*ProxyConfiguration, *mobilerun.ProxyListResponse], error) {
		pageOpts := opts
		res, err := m.c.api.Proxies.List(ctx, params, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*ProxyConfiguration, *mobilerun.ProxyListResponse]{Response: res}
		for _, info := range objectRows(res, "data") {
			result.Items = append(result.Items, newProxyConfiguration(m.c, objectID(info, "proxyId")).withInfo(info, ""))
		}
		return result, nil
	}
	return load()
}

// ListAutoPaging iterates the ProxyConfiguration handles of every page.
func (m *ProxyConfigurationManager) ListAutoPaging(ctx context.Context, params mobilerun.ProxyListParams, opts ...option.RequestOption) *ObjectAutoPager[*ProxyConfiguration] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// ComputeFunctionManager creates, fetches and lists ComputeFunction handles.
type ComputeFunctionManager struct {
	c *clientRef

	Validations *ComputeFunctionManagerValidations
}

func newComputeFunctionManager(c *clientRef) *ComputeFunctionManager {
	m := &ComputeFunctionManager{c: c}
	m.Validations = newComputeFunctionManagerValidations(m)
	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *ComputeFunctionManager) FromID(id string) *ComputeFunction {
	return newComputeFunction(m.c, id)
}

func (m *ComputeFunctionManager) String() string { return objectString("ComputeFunctionManager", nil) }

// Get function metadata
//
// Get fetches one ComputeFunction and returns its handle.
func (m *ComputeFunctionManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*ComputeFunction, error) {
	res, err := m.c.api.Compute.Functions.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newComputeFunction(m.c, id).withInfo(res, "data"), nil
}

// Create a function with its first immutable version
//
// New creates a ComputeFunction. The returned handle owns cleanup: Close disposes
// it.
func (m *ComputeFunctionManager) New(ctx context.Context, params mobilerun.ComputeFunctionNewParams, opts ...option.RequestOption) (*ComputeFunction, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Compute.Functions.New(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newComputeFunction(m.c, objectID(res, "data.id")).withInfo(res, "data")
	o.state.setOwned(true)
	return o, nil
}

// List functions owned by this organisation
//
// List returns the first page of ComputeFunction handles; ListAutoPaging iterates
// every page.
func (m *ComputeFunctionManager) List(ctx context.Context, params mobilerun.ComputeFunctionListParams, opts ...option.RequestOption) (*ObjectPage[*ComputeFunction, *mobilerun.ComputeFunctionListResponse], error) {
	var load func(page int64) (*ObjectPage[*ComputeFunction, *mobilerun.ComputeFunctionListResponse], error)
	load = func(page int64) (*ObjectPage[*ComputeFunction, *mobilerun.ComputeFunctionListResponse], error) {
		pageOpts := opts
		if page > 0 {
			pageOpts = bindOptions(opts, binding{"page", "query", page})
		}
		res, err := m.c.api.Compute.Functions.List(ctx, params, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*ComputeFunction, *mobilerun.ComputeFunctionListResponse]{Response: res}
		for _, info := range objectRows(res, "data.items") {
			result.Items = append(result.Items, newComputeFunction(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		next, ok, err := nextPageNumber(res, "data.pagination.page", "data.pagination.hasNext", "data.pagination.pages", page)
		if err != nil {
			return nil, err
		}
		if ok {
			result.next = func() (*ObjectPage[*ComputeFunction, *mobilerun.ComputeFunctionListResponse], error) {
				return load(next)
			}
		}
		return result, nil
	}
	return load(0)
}

// ListAutoPaging iterates the ComputeFunction handles of every page.
func (m *ComputeFunctionManager) ListAutoPaging(ctx context.Context, params mobilerun.ComputeFunctionListParams, opts ...option.RequestOption) *ObjectAutoPager[*ComputeFunction] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// ComputeFunctionManagerValidations groups the validations methods of a ComputeFunctionManager.
type ComputeFunctionManagerValidations struct {
	m *ComputeFunctionManager
}

func newComputeFunctionManagerValidations(m *ComputeFunctionManager) *ComputeFunctionManagerValidations {
	n := &ComputeFunctionManagerValidations{m: m}

	return n
}

func (n *ComputeFunctionManagerValidations) String() string {
	return objectString("ComputeFunctionManagerValidations", nil)
}

// Validate Python function source without uploading it
func (n *ComputeFunctionManagerValidations) New(ctx context.Context, params mobilerun.ComputeFunctionValidationNewParams, opts ...option.RequestOption) (*mobilerun.ComputeFunctionValidationNewResponse, error) {
	m := n.m
	return m.c.api.Compute.Functions.Validations.New(ctx, params, opts...)
}

// Retrieve a function validation owned by this organisation
func (n *ComputeFunctionManagerValidations) Get(ctx context.Context, validationID string, params mobilerun.ComputeFunctionValidationGetParams, opts ...option.RequestOption) (*mobilerun.ComputeFunctionValidationGetResponse, error) {
	m := n.m
	return m.c.api.Compute.Functions.Validations.Get(ctx, validationID, params, opts...)
}

// RunManager creates, fetches and lists Run handles.
type RunManager struct {
	c *clientRef
}

func newRunManager(c *clientRef) *RunManager {
	m := &RunManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *RunManager) FromID(id string) *Run { return newRun(m.c, id) }

func (m *RunManager) String() string { return objectString("RunManager", nil) }

// Get a run
//
// Get fetches one Run and returns its handle.
func (m *RunManager) Get(ctx context.Context, id string, params mobilerun.ComputeRunGetParams, opts ...option.RequestOption) (*Run, error) {
	res, err := m.c.api.Compute.Runs.Get(ctx, id, params, opts...)
	if err != nil {
		return nil, err
	}
	return newRun(m.c, id).withInfo(res, "data"), nil
}

// List the runs started on the public routes, newest first
//
// List returns the first page of Run handles; ListAutoPaging iterates every page.
func (m *RunManager) List(ctx context.Context, params mobilerun.ComputeRunListParams, opts ...option.RequestOption) (*ObjectPage[*Run, *mobilerun.ComputeRunListResponse], error) {
	var load func(page int64) (*ObjectPage[*Run, *mobilerun.ComputeRunListResponse], error)
	load = func(page int64) (*ObjectPage[*Run, *mobilerun.ComputeRunListResponse], error) {
		pageOpts := opts
		if page > 0 {
			pageOpts = bindOptions(opts, binding{"page", "query", page})
		}
		res, err := m.c.api.Compute.Runs.List(ctx, params, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*Run, *mobilerun.ComputeRunListResponse]{Response: res}
		for _, info := range objectRows(res, "data.items") {
			result.Items = append(result.Items, newRun(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		next, ok, err := nextPageNumber(res, "data.pagination.page", "data.pagination.hasNext", "data.pagination.pages", page)
		if err != nil {
			return nil, err
		}
		if ok {
			result.next = func() (*ObjectPage[*Run, *mobilerun.ComputeRunListResponse], error) { return load(next) }
		}
		return result, nil
	}
	return load(0)
}

// ListAutoPaging iterates the Run handles of every page.
func (m *RunManager) ListAutoPaging(ctx context.Context, params mobilerun.ComputeRunListParams, opts ...option.RequestOption) *ObjectAutoPager[*Run] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// SandboxManager creates, fetches and lists Sandbox handles.
type SandboxManager struct {
	c *clientRef
}

func newSandboxManager(c *clientRef) *SandboxManager {
	m := &SandboxManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *SandboxManager) FromID(id string) *Sandbox { return newSandbox(m.c, id) }

func (m *SandboxManager) String() string { return objectString("SandboxManager", nil) }

// Get a sandbox
//
// Get fetches one Sandbox and returns its handle.
func (m *SandboxManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*Sandbox, error) {
	res, err := m.c.api.Compute.Sandboxes.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newSandbox(m.c, id).withInfo(res, "data"), nil
}

// Starts a compute sandbox; compute usage may incur charges until it stops.
//
// # Create a sandbox
//
// New creates a Sandbox. The returned handle owns cleanup: Close disposes it.
func (m *SandboxManager) New(ctx context.Context, params mobilerun.ComputeSandboxNewParams, opts ...option.RequestOption) (*Sandbox, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Compute.Sandboxes.New(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newSandbox(m.c, objectID(res, "data.id")).withInfo(res, "data")
	o.state.setOwned(true)
	return o, nil
}

// List sandboxes
//
// List returns the first page of Sandbox handles; ListAutoPaging iterates every
// page.
func (m *SandboxManager) List(ctx context.Context, params mobilerun.ComputeSandboxListParams, opts ...option.RequestOption) (*ObjectPage[*Sandbox, *mobilerun.ComputeSandboxListResponse], error) {
	var load func(page int64) (*ObjectPage[*Sandbox, *mobilerun.ComputeSandboxListResponse], error)
	load = func(page int64) (*ObjectPage[*Sandbox, *mobilerun.ComputeSandboxListResponse], error) {
		pageOpts := opts
		if page > 0 {
			pageOpts = bindOptions(opts, binding{"page", "query", page})
		}
		res, err := m.c.api.Compute.Sandboxes.List(ctx, params, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*Sandbox, *mobilerun.ComputeSandboxListResponse]{Response: res}
		for _, info := range objectRows(res, "data.items") {
			result.Items = append(result.Items, newSandbox(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		next, ok, err := nextPageNumber(res, "data.pagination.page", "data.pagination.hasNext", "data.pagination.pages", page)
		if err != nil {
			return nil, err
		}
		if ok {
			result.next = func() (*ObjectPage[*Sandbox, *mobilerun.ComputeSandboxListResponse], error) { return load(next) }
		}
		return result, nil
	}
	return load(0)
}

// ListAutoPaging iterates the Sandbox handles of every page.
func (m *SandboxManager) ListAutoPaging(ctx context.Context, params mobilerun.ComputeSandboxListParams, opts ...option.RequestOption) *ObjectAutoPager[*Sandbox] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// ConnectProxyManager creates, fetches and lists ConnectProxy handles.
type ConnectProxyManager struct {
	c *clientRef
}

func newConnectProxyManager(c *clientRef) *ConnectProxyManager {
	m := &ConnectProxyManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *ConnectProxyManager) FromID(id string) *ConnectProxy { return newConnectProxy(m.c, id) }

func (m *ConnectProxyManager) String() string { return objectString("ConnectProxyManager", nil) }

// Returns the proxy identified by the path ID. The response includes the proxy's
// password.
//
// Get fetches one ConnectProxy and returns its handle.
func (m *ConnectProxyManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*ConnectProxy, error) {
	res, err := m.c.api.Connect.Proxies.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newConnectProxy(m.c, id).withInfo(res, ""), nil
}

// Returns proxies owned by the caller. Credentials are omitted from the list.
//
// List returns the first page of ConnectProxy handles; ListAutoPaging iterates
// every page.
func (m *ConnectProxyManager) List(ctx context.Context, params mobilerun.ConnectProxyListParams, opts ...option.RequestOption) (*ObjectPage[*ConnectProxy, *pagination.Page[mobilerun.ConnectProxyListResponse]], error) {
	var wrap func(page *pagination.Page[mobilerun.ConnectProxyListResponse]) *ObjectPage[*ConnectProxy, *pagination.Page[mobilerun.ConnectProxyListResponse]]
	wrap = func(page *pagination.Page[mobilerun.ConnectProxyListResponse]) *ObjectPage[*ConnectProxy, *pagination.Page[mobilerun.ConnectProxyListResponse]] {
		result := &ObjectPage[*ConnectProxy, *pagination.Page[mobilerun.ConnectProxyListResponse]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newConnectProxy(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*ConnectProxy, *pagination.Page[mobilerun.ConnectProxyListResponse]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Connect.Proxies.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the ConnectProxy handles of every page.
func (m *ConnectProxyManager) ListAutoPaging(ctx context.Context, params mobilerun.ConnectProxyListParams, opts ...option.RequestOption) *ObjectAutoPager[*ConnectProxy] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// Purchases a proxy; may incur charges or require checkout.
//
// Provisions a proxy of the requested type for the caller in the selected country.
//
// Buy returns a ConnectProxy handle; the handle does not own cleanup.
func (m *ConnectProxyManager) Buy(ctx context.Context, params mobilerun.ConnectProxyBuyParams, opts ...option.RequestOption) (*ConnectProxy, error) {
	res, err := m.c.api.Connect.Proxies.Buy(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return newConnectProxy(m.c, objectID(res, "id")).withInfo(res, ""), nil
}

// ConnectUserManager creates, fetches and lists ConnectUser handles.
type ConnectUserManager struct {
	c *clientRef
}

func newConnectUserManager(c *clientRef) *ConnectUserManager {
	m := &ConnectUserManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *ConnectUserManager) FromID(id string) *ConnectUser { return newConnectUser(m.c, id) }

func (m *ConnectUserManager) String() string { return objectString("ConnectUserManager", nil) }

// Returns the SOCKS5 user identified by the path ID. The response includes the
// user's password.
//
// Get fetches one ConnectUser and returns its handle.
func (m *ConnectUserManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*ConnectUser, error) {
	res, err := m.c.api.Connect.Users.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newConnectUser(m.c, id).withInfo(res, ""), nil
}

// Creates a SOCKS5 credential, optionally bound to a proxy for dedicated routing.
// Username and password are generated when omitted.
//
// New creates a ConnectUser. The returned handle owns cleanup: Close disposes it.
func (m *ConnectUserManager) New(ctx context.Context, params mobilerun.ConnectUserNewParams, opts ...option.RequestOption) (*ConnectUser, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Connect.Users.New(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newConnectUser(m.c, objectID(res, "id")).withInfo(res, "")
	o.state.setOwned(true)
	return o, nil
}

// Returns SOCKS5 users owned by the caller. Passwords are omitted from the list.
//
// List returns the first page of ConnectUser handles; ListAutoPaging iterates
// every page.
func (m *ConnectUserManager) List(ctx context.Context, params mobilerun.ConnectUserListParams, opts ...option.RequestOption) (*ObjectPage[*ConnectUser, *pagination.Page[mobilerun.ConnectUserListResponse]], error) {
	var wrap func(page *pagination.Page[mobilerun.ConnectUserListResponse]) *ObjectPage[*ConnectUser, *pagination.Page[mobilerun.ConnectUserListResponse]]
	wrap = func(page *pagination.Page[mobilerun.ConnectUserListResponse]) *ObjectPage[*ConnectUser, *pagination.Page[mobilerun.ConnectUserListResponse]] {
		result := &ObjectPage[*ConnectUser, *pagination.Page[mobilerun.ConnectUserListResponse]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newConnectUser(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*ConnectUser, *pagination.Page[mobilerun.ConnectUserListResponse]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Connect.Users.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the ConnectUser handles of every page.
func (m *ConnectUserManager) ListAutoPaging(ctx context.Context, params mobilerun.ConnectUserListParams, opts ...option.RequestOption) *ObjectAutoPager[*ConnectUser] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// TaskManager creates, fetches and lists Task handles.
type TaskManager struct {
	c *clientRef
}

func newTaskManager(c *clientRef) *TaskManager {
	m := &TaskManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *TaskManager) FromID(id string) *Task { return newTask(m.c, id) }

func (m *TaskManager) String() string { return objectString("TaskManager", nil) }

// Get full details of a task by ID.
//
// Get fetches one Task and returns its handle.
func (m *TaskManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*Task, error) {
	res, err := m.c.api.Tasks.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newTask(m.c, id).withInfo(res, "task"), nil
}

// List tasks with optional filtering, sorting, and pagination.
//
// List returns the first page of Task handles; ListAutoPaging iterates every page.
func (m *TaskManager) List(ctx context.Context, params mobilerun.TaskListParams, opts ...option.RequestOption) (*ObjectPage[*Task, *pagination.Page[mobilerun.TaskListItem]], error) {
	var wrap func(page *pagination.Page[mobilerun.TaskListItem]) *ObjectPage[*Task, *pagination.Page[mobilerun.TaskListItem]]
	wrap = func(page *pagination.Page[mobilerun.TaskListItem]) *ObjectPage[*Task, *pagination.Page[mobilerun.TaskListItem]] {
		result := &ObjectPage[*Task, *pagination.Page[mobilerun.TaskListItem]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newTask(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*Task, *pagination.Page[mobilerun.TaskListItem]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Tasks.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the Task handles of every page.
func (m *TaskManager) ListAutoPaging(ctx context.Context, params mobilerun.TaskListParams, opts ...option.RequestOption) *ObjectAutoPager[*Task] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// Create and dispatch a new agent task, returning an SSE stream of task events.
// Cancels the task if the client disconnects.
func (m *TaskManager) RunStreamed(ctx context.Context, params mobilerun.TaskRunStreamedParams, opts ...option.RequestOption) (*mobilerun.TaskRunStreamedResponse, error) {
	return m.c.api.Tasks.RunStreamed(ctx, params, opts...)
}

// Task outcome counts per time bucket, success rate, duration percentiles and top
// failure reasons for a window. Tasks are bucketed by creation time.
func (m *TaskManager) Stats(ctx context.Context, params mobilerun.TaskStatsParams, opts ...option.RequestOption) (*mobilerun.TaskStatsResponse, error) {
	return m.c.api.Tasks.Stats(ctx, params, opts...)
}

// Create and dispatch a new agent task. Returns the task ID and device stream
// details.
//
// Run returns a Task handle; the handle does not own cleanup.
func (m *TaskManager) Run(ctx context.Context, params mobilerun.TaskRunParams, opts ...option.RequestOption) (*Task, error) {
	res, err := m.c.api.Tasks.Run(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return newTask(m.c, objectID(res, "id")).withInfo(res, ""), nil
}

// TriggerManager creates, fetches and lists Trigger handles.
type TriggerManager struct {
	c *clientRef
}

func newTriggerManager(c *clientRef) *TriggerManager {
	m := &TriggerManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *TriggerManager) FromID(id string) *Trigger { return newTrigger(m.c, id) }

func (m *TriggerManager) String() string { return objectString("TriggerManager", nil) }

// Fetch a single trigger by its ID, including its activation type and
// type-specific configuration. Returns 404 if no trigger matches.
//
// Get fetches one Trigger and returns its handle.
func (m *TriggerManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*Trigger, error) {
	res, err := m.c.api.Workflows.Triggers.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newTrigger(m.c, id).withInfo(res, "data"), nil
}

// Create a trigger with an activation type of `event`, `schedule`, or `custom`.
// Each type requires its own fields (e.g. `eventType` and optional `conditions`
// for events, `scheduleRule` and `timezone` for schedules, `customPayloadSchema`
// for custom or schedule triggers); mismatched fields are rejected. Schedule
// triggers may store an optional JSON object `payload` (at most 1 MiB); it
// requires `customPayloadSchema`, is validated against it (422 on mismatch), and
// is available as `event.payload` on every scheduled, manual, or verification run.
// When a schema is configured, omitted payload values are validated as an empty
// object.
//
// Supports an optional `Idempotency-Key` header (1-255 printable ASCII
// characters). Replays with the same key and an identical body return the original
// 201; a changed body under the same key returns 422 `idempotency_key_reused`.
//
// New creates a Trigger. The returned handle owns cleanup: Close disposes it.
func (m *TriggerManager) New(ctx context.Context, params mobilerun.WorkflowTriggerNewParams, opts ...option.RequestOption) (*Trigger, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Workflows.Triggers.New(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newTrigger(m.c, objectID(res, "data.id")).withInfo(res, "data")
	o.state.setOwned(true)
	return o, nil
}

// Return a paginated list of triggers. Supports filtering by `activation` and
// `eventType`, free-text `search`, and ordering by name, createdAt, or updatedAt.
//
// List returns the first page of Trigger handles; ListAutoPaging iterates every
// page.
func (m *TriggerManager) List(ctx context.Context, params mobilerun.WorkflowTriggerListParams, opts ...option.RequestOption) (*ObjectPage[*Trigger, *pagination.Page[mobilerun.Trigger]], error) {
	var wrap func(page *pagination.Page[mobilerun.Trigger]) *ObjectPage[*Trigger, *pagination.Page[mobilerun.Trigger]]
	wrap = func(page *pagination.Page[mobilerun.Trigger]) *ObjectPage[*Trigger, *pagination.Page[mobilerun.Trigger]] {
		result := &ObjectPage[*Trigger, *pagination.Page[mobilerun.Trigger]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newTrigger(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*Trigger, *pagination.Page[mobilerun.Trigger]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Workflows.Triggers.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the Trigger handles of every page.
func (m *TriggerManager) ListAutoPaging(ctx context.Context, params mobilerun.WorkflowTriggerListParams, opts ...option.RequestOption) *ObjectAutoPager[*Trigger] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// CatalogEntryManager creates, fetches and lists CatalogEntry handles.
type CatalogEntryManager struct {
	c *clientRef
}

func newCatalogEntryManager(c *clientRef) *CatalogEntryManager {
	m := &CatalogEntryManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *CatalogEntryManager) FromID(id string) *CatalogEntry { return newCatalogEntry(m.c, id) }

func (m *CatalogEntryManager) String() string { return objectString("CatalogEntryManager", nil) }

// Fetch a single action catalog entry by its ID, including its public `stepName`
// for routine documents, service, method, parameter schema, and `deviceFree`
// eligibility. Supports `profile=agent_routine`; returns 404 if no entry matches
// or the entry is outside that profile.
//
// Get fetches one CatalogEntry and returns its handle.
func (m *CatalogEntryManager) Get(ctx context.Context, id string, params mobilerun.WorkflowActionCatalogGetParams, opts ...option.RequestOption) (*CatalogEntry, error) {
	res, err := m.c.api.Workflows.ActionCatalog.Get(ctx, id, params, opts...)
	if err != nil {
		return nil, err
	}
	return newCatalogEntry(m.c, id).withInfo(res, "data"), nil
}

// Return a paginated list of catalog entries — the service/method templates that
// actions are created from, each carrying its public `stepName` for routine
// documents, parameter schema and `deviceFree` eligibility. Supports filtering by
// `service` and `profile=agent_routine`.
//
// List returns the first page of CatalogEntry handles; ListAutoPaging iterates
// every page.
func (m *CatalogEntryManager) List(ctx context.Context, params mobilerun.WorkflowActionCatalogListParams, opts ...option.RequestOption) (*ObjectPage[*CatalogEntry, *pagination.Page[mobilerun.ActionCatalogEntry]], error) {
	var wrap func(page *pagination.Page[mobilerun.ActionCatalogEntry]) *ObjectPage[*CatalogEntry, *pagination.Page[mobilerun.ActionCatalogEntry]]
	wrap = func(page *pagination.Page[mobilerun.ActionCatalogEntry]) *ObjectPage[*CatalogEntry, *pagination.Page[mobilerun.ActionCatalogEntry]] {
		result := &ObjectPage[*CatalogEntry, *pagination.Page[mobilerun.ActionCatalogEntry]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newCatalogEntry(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*CatalogEntry, *pagination.Page[mobilerun.ActionCatalogEntry]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Workflows.ActionCatalog.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the CatalogEntry handles of every page.
func (m *CatalogEntryManager) ListAutoPaging(ctx context.Context, params mobilerun.WorkflowActionCatalogListParams, opts ...option.RequestOption) *ObjectAutoPager[*CatalogEntry] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// WorkflowActionManager creates, fetches and lists WorkflowAction handles.
type WorkflowActionManager struct {
	c *clientRef

	Services *WorkflowActionManagerServices
}

func newWorkflowActionManager(c *clientRef) *WorkflowActionManager {
	m := &WorkflowActionManager{c: c}
	m.Services = newWorkflowActionManagerServices(m)
	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *WorkflowActionManager) FromID(id string) *WorkflowAction { return newWorkflowAction(m.c, id) }

func (m *WorkflowActionManager) String() string { return objectString("WorkflowActionManager", nil) }

// Fetch a single action by its ID, including its configured service, method, and
// params. Supports `profile=agent_routine`; returns 404 if no action matches or
// the action is outside that profile.
//
// Get fetches one WorkflowAction and returns its handle.
func (m *WorkflowActionManager) Get(ctx context.Context, id string, params mobilerun.WorkflowActionGetParams, opts ...option.RequestOption) (*WorkflowAction, error) {
	res, err := m.c.api.Workflows.Actions.Get(ctx, id, params, opts...)
	if err != nil {
		return nil, err
	}
	return newWorkflowAction(m.c, id).withInfo(res, "data"), nil
}

// Create a reusable action from a catalog entry (`catalogEntryId`), with an
// optional `params` object supplying the values for that entry's service method.
// Returns 400 if the params are invalid for the chosen catalog entry.
//
// Supports an optional `Idempotency-Key` header (1-255 printable ASCII
// characters). Replays with the same key and an identical body return the original
// 201; a changed body under the same key returns 422 `idempotency_key_reused`.
//
// New creates a WorkflowAction. The returned handle owns cleanup: Close disposes
// it.
func (m *WorkflowActionManager) New(ctx context.Context, params mobilerun.WorkflowActionNewParams, opts ...option.RequestOption) (*WorkflowAction, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Workflows.Actions.New(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newWorkflowAction(m.c, objectID(res, "data.id")).withInfo(res, "data")
	o.state.setOwned(true)
	return o, nil
}

// Return a paginated list of actions. Supports filtering by
// `profile=agent_routine`, `service`, free-text `search`, and ordering by name,
// createdAt, or updatedAt.
//
// List returns the first page of WorkflowAction handles; ListAutoPaging iterates
// every page.
func (m *WorkflowActionManager) List(ctx context.Context, params mobilerun.WorkflowActionListParams, opts ...option.RequestOption) (*ObjectPage[*WorkflowAction, *pagination.Page[mobilerun.Action]], error) {
	var wrap func(page *pagination.Page[mobilerun.Action]) *ObjectPage[*WorkflowAction, *pagination.Page[mobilerun.Action]]
	wrap = func(page *pagination.Page[mobilerun.Action]) *ObjectPage[*WorkflowAction, *pagination.Page[mobilerun.Action]] {
		result := &ObjectPage[*WorkflowAction, *pagination.Page[mobilerun.Action]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newWorkflowAction(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*WorkflowAction, *pagination.Page[mobilerun.Action]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Workflows.Actions.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the WorkflowAction handles of every page.
func (m *WorkflowActionManager) ListAutoPaging(ctx context.Context, params mobilerun.WorkflowActionListParams, opts ...option.RequestOption) *ObjectAutoPager[*WorkflowAction] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// WorkflowActionManagerServices groups the services methods of a WorkflowActionManager.
type WorkflowActionManagerServices struct {
	m *WorkflowActionManager
}

func newWorkflowActionManagerServices(m *WorkflowActionManager) *WorkflowActionManagerServices {
	n := &WorkflowActionManagerServices{m: m}

	return n
}

func (n *WorkflowActionManagerServices) String() string {
	return objectString("WorkflowActionManagerServices", nil)
}

// Return the names of the services that actions can be built against. Use these
// values to look up each service's allowed methods.
func (n *WorkflowActionManagerServices) List(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WorkflowActionServiceListResponse, error) {
	m := n.m
	return m.c.api.Workflows.Actions.Services.List(ctx, opts...)
}

// Return the methods allowed for the given service, each with its parameter
// definitions (name, type, whether required, description, and optional
// default/example). Returns 404 if the service is unknown.
func (n *WorkflowActionManagerServices) ListMethods(ctx context.Context, service string, opts ...option.RequestOption) (*mobilerun.WorkflowActionServiceListMethodsResponse, error) {
	m := n.m
	return m.c.api.Workflows.Actions.Services.ListMethods(ctx, service, opts...)
}

// FlowManager creates, fetches and lists Flow handles.
type FlowManager struct {
	c *clientRef
}

func newFlowManager(c *clientRef) *FlowManager {
	m := &FlowManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *FlowManager) FromID(id string) *Flow { return newFlow(m.c, id) }

func (m *FlowManager) String() string { return objectString("FlowManager", nil) }

// Fetch a single flow by its ID, including its trigger binding, configuration, and
// current status. Returns 404 if no flow matches.
//
// Get fetches one Flow and returns its handle.
func (m *FlowManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*Flow, error) {
	res, err := m.c.api.Workflows.Flows.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newFlow(m.c, id).withInfo(res, "data"), nil
}

// Create a flow that binds a trigger (`triggerId`) to an ordered list of actions.
// The list may be empty or omitted: a routine without steps is stored disabled and
// cannot be enabled, run or verified until it has a step. Enabling it returns 400
// with "Add a step before enabling this routine". Optional settings include target
// `deviceIds`, a cooldown (`cooldownSeconds`/`cooldownScope`), and webhook
// notifications on success or failure.
//
// Supports an optional `Idempotency-Key` header (1-255 printable ASCII
// characters). Replays with the same key and an identical body return the original
// 201; a changed body under the same key returns 422 `idempotency_key_reused`.
//
// New creates a Flow. The returned handle owns cleanup: Close disposes it.
func (m *FlowManager) New(ctx context.Context, params mobilerun.WorkflowFlowNewParams, opts ...option.RequestOption) (*Flow, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Workflows.Flows.New(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newFlow(m.c, objectID(res, "data.id")).withInfo(res, "data")
	o.state.setOwned(true)
	return o, nil
}

// Return a paginated list of flows. Supports filtering by `agentId` (UUID or
// `unassigned`; omit for all buckets), `triggerId`, `enabled`, one or more health
// `status` values (healthy, failing, blocked), `mine` (flows created by the
// calling actor), `createdBy` (flows created by a given actor id — mutually
// exclusive with `mine`), plus free-text `search` and ordering.
//
// List returns the first page of Flow handles; ListAutoPaging iterates every page.
func (m *FlowManager) List(ctx context.Context, params mobilerun.WorkflowFlowListParams, opts ...option.RequestOption) (*ObjectPage[*Flow, *pagination.Page[mobilerun.WorkflowFlowListResponse]], error) {
	var wrap func(page *pagination.Page[mobilerun.WorkflowFlowListResponse]) *ObjectPage[*Flow, *pagination.Page[mobilerun.WorkflowFlowListResponse]]
	wrap = func(page *pagination.Page[mobilerun.WorkflowFlowListResponse]) *ObjectPage[*Flow, *pagination.Page[mobilerun.WorkflowFlowListResponse]] {
		result := &ObjectPage[*Flow, *pagination.Page[mobilerun.WorkflowFlowListResponse]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newFlow(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*Flow, *pagination.Page[mobilerun.WorkflowFlowListResponse]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Workflows.Flows.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the Flow handles of every page.
func (m *FlowManager) ListAutoPaging(ctx context.Context, params mobilerun.WorkflowFlowListParams, opts ...option.RequestOption) *ObjectAutoPager[*Flow] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// Returns the environment routine limit and remaining capacity for this owner and
// agent bucket, counting enabled and disabled non-deleted flows. Omit agentId (or
// use unassigned) for the unassigned bucket. This is advisory; create and clone
// enforce admission under a bucket lock.
func (m *FlowManager) Capacity(ctx context.Context, params mobilerun.WorkflowFlowCapacityParams, opts ...option.RequestOption) (*mobilerun.WorkflowFlowCapacityResponse, error) {
	return m.c.api.Workflows.Flows.Capacity(ctx, params, opts...)
}

// Atomically creates the trigger, actions and disabled flow. Body: {document}.
// Optional Idempotency-Key fingerprints the canonical document; identical retries
// reload the resource, changed documents return idempotency_key_reused. The owner
// can enable a routine with at least one step directly through PATCH
// /flows/{flowId} using enabled: true or lifecycleStatus: "enabled", without
// verification. For the verified path used by agent deploys, call POST
// /flows/{flowId}/verify then POST /flows/{flowId}/activate; activation requires a
// successful verification of the current graph.
func (m *FlowManager) NewFromDocument(ctx context.Context, params mobilerun.WorkflowFlowNewFromDocumentParams, opts ...option.RequestOption) (*mobilerun.WorkflowFlowNewFromDocumentResponse, error) {
	return m.c.api.Workflows.Flows.NewFromDocument(ctx, params, opts...)
}

// The optional agent_routine profile lists only agent-eligible steps. This is
// authoring guidance; backend validation remains authoritative.
func (m *FlowManager) DocumentSchema(ctx context.Context, params mobilerun.WorkflowFlowDocumentSchemaParams, opts ...option.RequestOption) (*mobilerun.WorkflowFlowDocumentSchemaResponse, error) {
	return m.c.api.Workflows.Flows.DocumentSchema(ctx, params, opts...)
}

// Get flow template context
func (m *FlowManager) TemplateContext(ctx context.Context, params mobilerun.WorkflowFlowTemplateContextParams, opts ...option.RequestOption) (*mobilerun.WorkflowFlowTemplateContextResponse, error) {
	return m.c.api.Workflows.Flows.TemplateContext(ctx, params, opts...)
}

// Validate a flow
func (m *FlowManager) Validate(ctx context.Context, params mobilerun.WorkflowFlowValidateParams, opts ...option.RequestOption) (*mobilerun.WorkflowFlowValidateResponse, error) {
	return m.c.api.Workflows.Flows.Validate(ctx, params, opts...)
}

// Validates graph, settings and scheduled inputs without writing. Disconnected
// delivery destinations produce warnings; create and replace enforce the existing
// delivery connection checks.
func (m *FlowManager) ValidateDocument(ctx context.Context, params mobilerun.WorkflowFlowValidateDocumentParams, opts ...option.RequestOption) (*mobilerun.WorkflowFlowValidateDocumentResponse, error) {
	return m.c.api.Workflows.Flows.ValidateDocument(ctx, params, opts...)
}

// ScheduledRunManager creates, fetches and lists ScheduledRun handles.
type ScheduledRunManager struct {
	c      *clientRef
	flowID string
}

func newScheduledRunManager(c *clientRef, flowID string) *ScheduledRunManager {
	m := &ScheduledRunManager{c: c, flowID: flowID}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *ScheduledRunManager) FromID(id string) *ScheduledRun {
	return newScheduledRun(m.c, m.flowID, id)
}

func (m *ScheduledRunManager) String() string { return objectString("ScheduledRunManager", nil) }

// Plan Run now for a future time, using the same devices, payload rules and
// admission at fire time. A flow can have at most 100 scheduled runs. Supports an
// optional Idempotency-Key; a replay returns the current scheduled run with the
// original 201 status.
//
// New creates a ScheduledRun. The returned handle owns cleanup: Close disposes it.
func (m *ScheduledRunManager) New(ctx context.Context, params mobilerun.WorkflowFlowScheduledRunNewParams, opts ...option.RequestOption) (*ScheduledRun, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Workflows.Flows.ScheduledRuns.New(ctx, m.flowID, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newScheduledRun(m.c, m.flowID, objectID(res, "data.id")).withInfo(res, "data")
	o.state.setOwned(true)
	return o, nil
}

// Lists upcoming runs by default, including claimed runs awaiting dispatch. Set
// status=all to include terminal rows. Ordered by runAt ascending.
//
// List returns the first page of ScheduledRun handles; ListAutoPaging iterates
// every page.
func (m *ScheduledRunManager) List(ctx context.Context, params mobilerun.WorkflowFlowScheduledRunListParams, opts ...option.RequestOption) (*ObjectPage[*ScheduledRun, *pagination.Page[mobilerun.ScheduledRun]], error) {
	var wrap func(page *pagination.Page[mobilerun.ScheduledRun]) *ObjectPage[*ScheduledRun, *pagination.Page[mobilerun.ScheduledRun]]
	wrap = func(page *pagination.Page[mobilerun.ScheduledRun]) *ObjectPage[*ScheduledRun, *pagination.Page[mobilerun.ScheduledRun]] {
		result := &ObjectPage[*ScheduledRun, *pagination.Page[mobilerun.ScheduledRun]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newScheduledRun(m.c, m.flowID, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*ScheduledRun, *pagination.Page[mobilerun.ScheduledRun]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Workflows.Flows.ScheduledRuns.List(ctx, m.flowID, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the ScheduledRun handles of every page.
func (m *ScheduledRunManager) ListAutoPaging(ctx context.Context, params mobilerun.WorkflowFlowScheduledRunListParams, opts ...option.RequestOption) *ObjectAutoPager[*ScheduledRun] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// FlowID returns the bound Flow ID.
func (m *ScheduledRunManager) FlowID() string { return m.flowID }

// ExecutionManager creates, fetches and lists Execution handles.
type ExecutionManager struct {
	c *clientRef
}

func newExecutionManager(c *clientRef) *ExecutionManager {
	m := &ExecutionManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *ExecutionManager) FromID(id string) *Execution { return newExecution(m.c, id) }

func (m *ExecutionManager) String() string { return objectString("ExecutionManager", nil) }

// Fetch a single flow execution by its ID, including its status, kind, result or
// error, and start/finish timestamps. Add `include=input` to get what each step
// ran with. Returns 404 if no execution matches.
//
// Get fetches one Execution and returns its handle.
func (m *ExecutionManager) Get(ctx context.Context, id string, params mobilerun.WorkflowExecutionGetParams, opts ...option.RequestOption) (*Execution, error) {
	res, err := m.c.api.Workflows.Executions.Get(ctx, id, params, opts...)
	if err != nil {
		return nil, err
	}
	return newExecution(m.c, id).withInfo(res, "data"), nil
}

// Return a paginated history of flow executions. Supports filtering by `flowId`,
// `triggerId`, `invocationId` (the client/verify key), `status`, `kind` (one or
// more execution kinds), and a `from`/`to` time range, plus free-text `search` and
// ordering by startedAt, finishedAt, or status.
//
// List returns the first page of Execution handles; ListAutoPaging iterates every
// page.
func (m *ExecutionManager) List(ctx context.Context, params mobilerun.WorkflowExecutionListParams, opts ...option.RequestOption) (*ObjectPage[*Execution, *pagination.Page[mobilerun.WorkflowExecutionListResponse]], error) {
	var wrap func(page *pagination.Page[mobilerun.WorkflowExecutionListResponse]) *ObjectPage[*Execution, *pagination.Page[mobilerun.WorkflowExecutionListResponse]]
	wrap = func(page *pagination.Page[mobilerun.WorkflowExecutionListResponse]) *ObjectPage[*Execution, *pagination.Page[mobilerun.WorkflowExecutionListResponse]] {
		result := &ObjectPage[*Execution, *pagination.Page[mobilerun.WorkflowExecutionListResponse]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newExecution(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*Execution, *pagination.Page[mobilerun.WorkflowExecutionListResponse]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Workflows.Executions.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the Execution handles of every page.
func (m *ExecutionManager) ListAutoPaging(ctx context.Context, params mobilerun.WorkflowExecutionListParams, opts ...option.RequestOption) *ObjectAutoPager[*Execution] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// Return aggregate execution metrics — total count, counts by status, average
// duration, and the last execution time. Can be scoped by `flowId`, `triggerId`,
// and a `from`/`to` time range.
func (m *ExecutionManager) GetMetrics(ctx context.Context, params mobilerun.WorkflowExecutionGetMetricsParams, opts ...option.RequestOption) (*mobilerun.WorkflowExecutionGetMetricsResponse, error) {
	return m.c.api.Workflows.Executions.GetMetrics(ctx, params, opts...)
}

// WebhookManager creates, fetches and lists Webhook handles.
type WebhookManager struct {
	c *clientRef

	Integrations *WebhookManagerIntegrations
}

func newWebhookManager(c *clientRef) *WebhookManager {
	m := &WebhookManager{c: c}
	m.Integrations = newWebhookManagerIntegrations(m)
	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *WebhookManager) FromID(id string) *Webhook { return newWebhook(m.c, id) }

func (m *WebhookManager) String() string { return objectString("WebhookManager", nil) }

// Returns a single webhook subscription by id, including its URL (or integration
// target), subscribed event types, state, and system-observed delivery health. The
// signing secret is never included.
//
// Get fetches one Webhook and returns its handle.
func (m *WebhookManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*Webhook, error) {
	res, err := m.c.api.Webhooks.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newWebhook(m.c, id).withInfo(res, "data"), nil
}

// Creates a webhook subscription and an optional list of event types to subscribe
// to (defaults to all when omitted). `kind: "http"` (the default) delivers signed
// JSON to a URL; the response includes the generated signing secret, which is
// returned only once at creation time and cannot be retrieved later.
// `kind: "integration"` posts each event as a message into a connected integration
// target (e.g. a Slack channel): pick the `capability` from
// `GET /webhooks/integrations` and the `args` from
// `GET /webhooks/integrations/{capabilityId}/targets`. Integration webhooks have
// no signing secret. Send an `Idempotency-Key` header to make the create safe to
// retry: the same key with the same body returns the existing webhook (status 201,
// `Idempotent-Replayed: true`) instead of creating a second one, and the same key
// with a different body is a 422. A replayed create does not return the signing
// secret again; use rotate-secret if the original response was lost. If that
// webhook was deleted in the meantime, the retry is a 404.
//
// New creates a Webhook. The returned handle owns cleanup: Close disposes it.
func (m *WebhookManager) New(ctx context.Context, params mobilerun.WebhookNewParams, opts ...option.RequestOption) (*Webhook, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Webhooks.New(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newWebhook(m.c, objectID(res, "data.id")).withInfo(res, "data")
	o.state.setOwned(true)
	return o, nil
}

// Returns a paginated list of your webhook subscriptions, optionally filtered by
// status (active, failing, blocked, or disabled), by `search` (a case-insensitive
// substring match against the URL or description), and/or by delivery `kind`. The
// response also includes per-status counts across all of your subscriptions (of
// the requested `kind`, if given; the other filters do not apply to the counts).
//
// List returns the first page of Webhook handles; ListAutoPaging iterates every
// page.
func (m *WebhookManager) List(ctx context.Context, params mobilerun.WebhookListParams, opts ...option.RequestOption) (*ObjectPage[*Webhook, *pagination.Page[mobilerun.WebhookListResponse]], error) {
	var wrap func(page *pagination.Page[mobilerun.WebhookListResponse]) *ObjectPage[*Webhook, *pagination.Page[mobilerun.WebhookListResponse]]
	wrap = func(page *pagination.Page[mobilerun.WebhookListResponse]) *ObjectPage[*Webhook, *pagination.Page[mobilerun.WebhookListResponse]] {
		result := &ObjectPage[*Webhook, *pagination.Page[mobilerun.WebhookListResponse]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newWebhook(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*Webhook, *pagination.Page[mobilerun.WebhookListResponse]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Webhooks.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the Webhook handles of every page.
func (m *WebhookManager) ListAutoPaging(ctx context.Context, params mobilerun.WebhookListParams, opts ...option.RequestOption) *ObjectAutoPager[*Webhook] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// Verify verifies a signed delivery; see [VerifyWebhook].
func (m *WebhookManager) Verify(body []byte, headers http.Header, secret string, opts *SignatureVerifyOptions) (any, error) {
	return VerifyWebhook(body, headers, secret, opts)
}

// WebhookManagerIntegrations groups the integrations methods of a WebhookManager.
type WebhookManagerIntegrations struct {
	m *WebhookManager
}

func newWebhookManagerIntegrations(m *WebhookManager) *WebhookManagerIntegrations {
	n := &WebhookManagerIntegrations{m: m}

	return n
}

func (n *WebhookManagerIntegrations) String() string {
	return objectString("WebhookManagerIntegrations", nil)
}

// Returns the integrations that can receive webhook events as channel messages
// (e.g. Slack), with your org's connection state for each. Connect an integration
// in the Cloud app before creating a webhook for it.
func (n *WebhookManagerIntegrations) List(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WebhookIntegrationListResponse, error) {
	m := n.m
	return m.c.api.Webhooks.Integrations.List(ctx, opts...)
}

// Returns the live list of targets (e.g. Slack channels the connected account is a
// member of) for an integration. Pass a target's `args` together with the
// integration `capability` to `POST /webhooks` with `kind: "integration"`.
func (n *WebhookManagerIntegrations) ListTargets(ctx context.Context, capabilityID string, params mobilerun.WebhookIntegrationListTargetsParams, opts ...option.RequestOption) (*mobilerun.WebhookIntegrationListTargetsResponse, error) {
	m := n.m
	return m.c.api.Webhooks.Integrations.ListTargets(ctx, capabilityID, params, opts...)
}

// WebhookDeliveryManager creates, fetches and lists WebhookDelivery handles.
type WebhookDeliveryManager struct {
	c         *clientRef
	webhookID string
}

func newWebhookDeliveryManager(c *clientRef, webhookID string) *WebhookDeliveryManager {
	m := &WebhookDeliveryManager{c: c, webhookID: webhookID}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *WebhookDeliveryManager) FromID(id string) *WebhookDelivery {
	return newWebhookDelivery(m.c, m.webhookID, id)
}

func (m *WebhookDeliveryManager) String() string { return objectString("WebhookDeliveryManager", nil) }

// Returns a single delivery for a webhook subscription along with the full list of
// captured attempt records. Each attempt includes the request URL, method, headers
// and body, whether it was signed, and the response status, headers, and snippet.
//
// Get fetches one WebhookDelivery and returns its handle.
//
// The object layer sets id; values for it in params are replaced.
func (m *WebhookDeliveryManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*WebhookDelivery, error) {
	res, err := m.c.api.Webhooks.Deliveries.GetAttempts(ctx, id, mobilerun.WebhookDeliveryGetAttemptsParams{ID: m.webhookID}, opts...)
	if err != nil {
		return nil, err
	}
	return newWebhookDelivery(m.c, m.webhookID, id).withInfo(res, "data"), nil
}

// Returns a paginated list of deliveries for a single webhook subscription,
// identified by its id. Each record reports the event, delivery status, attempt
// count, and the last response code or error. Results can be filtered by `eventId`
// (exact match against the originating event id).
//
// List returns the first page of WebhookDelivery handles; ListAutoPaging iterates
// every page.
func (m *WebhookDeliveryManager) List(ctx context.Context, params mobilerun.WebhookDeliveryListForWebhookParams, opts ...option.RequestOption) (*ObjectPage[*WebhookDelivery, *pagination.Page[mobilerun.WebhookDeliveryListForWebhookResponse]], error) {
	var wrap func(page *pagination.Page[mobilerun.WebhookDeliveryListForWebhookResponse]) *ObjectPage[*WebhookDelivery, *pagination.Page[mobilerun.WebhookDeliveryListForWebhookResponse]]
	wrap = func(page *pagination.Page[mobilerun.WebhookDeliveryListForWebhookResponse]) *ObjectPage[*WebhookDelivery, *pagination.Page[mobilerun.WebhookDeliveryListForWebhookResponse]] {
		result := &ObjectPage[*WebhookDelivery, *pagination.Page[mobilerun.WebhookDeliveryListForWebhookResponse]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newWebhookDelivery(m.c, m.webhookID, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*WebhookDelivery, *pagination.Page[mobilerun.WebhookDeliveryListForWebhookResponse]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Webhooks.Deliveries.ListForWebhook(ctx, m.webhookID, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the WebhookDelivery handles of every page.
func (m *WebhookDeliveryManager) ListAutoPaging(ctx context.Context, params mobilerun.WebhookDeliveryListForWebhookParams, opts ...option.RequestOption) *ObjectAutoPager[*WebhookDelivery] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// WebhookID returns the bound Webhook ID.
func (m *WebhookDeliveryManager) WebhookID() string { return m.webhookID }

// MailboxManager creates, fetches and lists Mailbox handles.
type MailboxManager struct {
	c *clientRef

	Connections *MailboxConnectionManager
}

func newMailboxManager(c *clientRef) *MailboxManager {
	m := &MailboxManager{c: c}
	m.Connections = newMailboxConnectionManager(m.c)
	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *MailboxManager) FromID(id string) *Mailbox { return newMailbox(m.c, id) }

func (m *MailboxManager) String() string { return objectString("MailboxManager", nil) }

// Get a mailbox
//
// Get fetches one Mailbox and returns its handle.
func (m *MailboxManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*Mailbox, error) {
	res, err := m.c.api.Mailboxes.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newMailbox(m.c, id).withInfo(res, "data"), nil
}

// Claims or rents a mailbox. Included capacity may apply; paid rentals may require
// checkout.
//
// Creates a mailbox on the default domain or a connected custom domain. An
// optional `localPart` selects the address. Replaying the same Idempotency-Key and
// payload returns the original mailbox. Fresh requests and replays return 202 with
// the current mailbox state. Poll the mailbox when the response does not yet
// include a checkout URL.
//
// New creates a Mailbox. The returned handle owns cleanup: Close disposes it.
func (m *MailboxManager) New(ctx context.Context, params mobilerun.MailboxNewParams, opts ...option.RequestOption) (*Mailbox, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Mailboxes.New(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newMailbox(m.c, objectID(res, "data.id")).withInfo(res, "data")
	o.state.setOwned(true)
	return o, nil
}

// Lists the caller-owned mailboxes with page-based pagination.
//
// List returns the first page of Mailbox handles; ListAutoPaging iterates every
// page.
func (m *MailboxManager) List(ctx context.Context, params mobilerun.MailboxListParams, opts ...option.RequestOption) (*ObjectPage[*Mailbox, *pagination.Page[mobilerun.Mailbox]], error) {
	var wrap func(page *pagination.Page[mobilerun.Mailbox]) *ObjectPage[*Mailbox, *pagination.Page[mobilerun.Mailbox]]
	wrap = func(page *pagination.Page[mobilerun.Mailbox]) *ObjectPage[*Mailbox, *pagination.Page[mobilerun.Mailbox]] {
		result := &ObjectPage[*Mailbox, *pagination.Page[mobilerun.Mailbox]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newMailbox(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*Mailbox, *pagination.Page[mobilerun.Mailbox]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Mailboxes.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the Mailbox handles of every page.
func (m *MailboxManager) ListAutoPaging(ctx context.Context, params mobilerun.MailboxListParams, opts ...option.RequestOption) *ObjectAutoPager[*Mailbox] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// Returns the number of mailboxes currently available through included capacity.
func (m *MailboxManager) Capacity(ctx context.Context, opts ...option.RequestOption) (*mobilerun.MailboxCapacityResponse, error) {
	return m.c.api.Mailboxes.Capacity(ctx, opts...)
}

// MailboxConnectionManager creates, fetches and lists MailboxConnection handles.
type MailboxConnectionManager struct {
	c *clientRef
}

func newMailboxConnectionManager(c *clientRef) *MailboxConnectionManager {
	m := &MailboxConnectionManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *MailboxConnectionManager) FromID(id string) *MailboxConnection {
	return newMailboxConnection(m.c, id)
}

func (m *MailboxConnectionManager) String() string {
	return objectString("MailboxConnectionManager", nil)
}

// Reports pending, active (with the Gmail address), expired, or removed.
// connectionId is the mailbox id.
//
// Get fetches one MailboxConnection and returns its handle.
func (m *MailboxConnectionManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*MailboxConnection, error) {
	res, err := m.c.api.Mailboxes.Connections.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newMailboxConnection(m.c, id).withInfo(res, ""), nil
}

// Links the caller’s Gmail account. Replaying the same Idempotency-Key returns the
// same mailbox. When Gmail is already active, redirectUrl is null and the mailbox
// is usable immediately.
//
// New creates a MailboxConnection. The returned handle owns cleanup: Close
// disposes it.
func (m *MailboxConnectionManager) New(ctx context.Context, params mobilerun.MailboxConnectionNewParams, opts ...option.RequestOption) (*MailboxConnection, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Mailboxes.Connections.New(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newMailboxConnection(m.c, objectID(res, "connectionId")).withInfo(res, "")
	o.state.setOwned(true)
	return o, nil
}

// MailboxMessageManager creates, fetches and lists MailboxMessage handles.
type MailboxMessageManager struct {
	c         *clientRef
	mailboxID string
}

func newMailboxMessageManager(c *clientRef, mailboxID string) *MailboxMessageManager {
	m := &MailboxMessageManager{c: c, mailboxID: mailboxID}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *MailboxMessageManager) FromID(id string) *MailboxMessage {
	return newMailboxMessage(m.c, m.mailboxID, id)
}

func (m *MailboxMessageManager) String() string { return objectString("MailboxMessageManager", nil) }

// Get a mailbox message
//
// Get fetches one MailboxMessage and returns its handle.
//
// The object layer sets mailboxId; values for it in params are replaced.
func (m *MailboxMessageManager) Get(ctx context.Context, id string, params mobilerun.MailboxMessageGetParams, opts ...option.RequestOption) (*MailboxMessage, error) {
	params.MailboxID = m.mailboxID
	res, err := m.c.api.Mailboxes.Messages.Get(ctx, id, params, opts...)
	if err != nil {
		return nil, err
	}
	return newMailboxMessage(m.c, m.mailboxID, id).withInfo(res, "data"), nil
}

// Lists mailbox messages with cursor pagination and optional time, sender, and OTP
// filters.
//
// List returns the first page of MailboxMessage handles; ListAutoPaging iterates
// every page.
func (m *MailboxMessageManager) List(ctx context.Context, params mobilerun.MailboxMessageListParams, opts ...option.RequestOption) (*ObjectPage[*MailboxMessage, *mobilerun.MailboxMessageListResponse], error) {
	seen := map[string]bool{}
	var load func(page string) (*ObjectPage[*MailboxMessage, *mobilerun.MailboxMessageListResponse], error)
	load = func(page string) (*ObjectPage[*MailboxMessage, *mobilerun.MailboxMessageListResponse], error) {
		pageOpts := opts
		if page != "" {
			seen[page] = true
			pageOpts = bindOptions(opts, binding{"cursor", "query", page})
		}
		res, err := m.c.api.Mailboxes.Messages.List(ctx, m.mailboxID, params, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*MailboxMessage, *mobilerun.MailboxMessageListResponse]{Response: res}
		for _, info := range objectRows(res, "items") {
			result.Items = append(result.Items, newMailboxMessage(m.c, m.mailboxID, objectID(info, "id")).withInfo(info, ""))
		}
		next, ok, err := nextObjectCursor(res, "nextCursor", seen)
		if err != nil {
			return nil, err
		}
		if ok {
			result.next = func() (*ObjectPage[*MailboxMessage, *mobilerun.MailboxMessageListResponse], error) { return load(next) }
		}
		return result, nil
	}
	return load("")
}

// ListAutoPaging iterates the MailboxMessage handles of every page.
func (m *MailboxMessageManager) ListAutoPaging(ctx context.Context, params mobilerun.MailboxMessageListParams, opts ...option.RequestOption) *ObjectAutoPager[*MailboxMessage] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// MailboxID returns the bound Mailbox ID.
func (m *MailboxMessageManager) MailboxID() string { return m.mailboxID }

// RoutineFileManager creates, fetches and lists RoutineFile handles.
type RoutineFileManager struct {
	c *clientRef
}

func newRoutineFileManager(c *clientRef) *RoutineFileManager {
	m := &RoutineFileManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *RoutineFileManager) FromID(id string) *RoutineFile { return newRoutineFile(m.c, id) }

func (m *RoutineFileManager) String() string { return objectString("RoutineFileManager", nil) }

// Lists the files saved in Mobilerun storage for routines whose delivery
// destination is Mobilerun with a folder: the run recording, screenshots and
// files.upload outputs. Newest first; filter by `folder`, `flowId` or
// `executionId`. Files stay until you delete them or delete their routine.
//
// List returns the first page of RoutineFile handles; ListAutoPaging iterates
// every page.
func (m *RoutineFileManager) List(ctx context.Context, params mobilerun.RoutineFileListParams, opts ...option.RequestOption) (*ObjectPage[*RoutineFile, *pagination.Page[mobilerun.RoutineFileListResponse]], error) {
	var wrap func(page *pagination.Page[mobilerun.RoutineFileListResponse]) *ObjectPage[*RoutineFile, *pagination.Page[mobilerun.RoutineFileListResponse]]
	wrap = func(page *pagination.Page[mobilerun.RoutineFileListResponse]) *ObjectPage[*RoutineFile, *pagination.Page[mobilerun.RoutineFileListResponse]] {
		result := &ObjectPage[*RoutineFile, *pagination.Page[mobilerun.RoutineFileListResponse]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newRoutineFile(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*RoutineFile, *pagination.Page[mobilerun.RoutineFileListResponse]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.RoutineFiles.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the RoutineFile handles of every page.
func (m *RoutineFileManager) ListAutoPaging(ctx context.Context, params mobilerun.RoutineFileListParams, opts ...option.RequestOption) *ObjectAutoPager[*RoutineFile] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// Permanently deletes files for exactly one folder or executionId.
//
// Deletes every saved file of exactly one folder (`folder`) or one run
// (`executionId`) permanently. Pass exactly one of the two. Returns how many files
// were deleted.
func (m *RoutineFileManager) DeleteMany(ctx context.Context, params mobilerun.RoutineFileDeleteManyParams, opts ...option.RequestOption) (*mobilerun.RoutineFileDeleteManyResponse, error) {
	return m.c.api.RoutineFiles.DeleteMany(ctx, params, opts...)
}

// Lists the folders that hold saved files with their file count, total size and
// newest save time, newest first. Independent of whether the routine or run still
// exists; `flows` names the routines that saved into each folder.
func (m *RoutineFileManager) ListFolders(ctx context.Context, params mobilerun.RoutineFileListFoldersParams, opts ...option.RequestOption) (*pagination.Page[mobilerun.RoutineFileListFoldersResponse], error) {
	return m.c.api.RoutineFiles.ListFolders(ctx, params, append(append([]option.RequestOption{}, opts...), callerContext(ctx))...)
}

// Lists the folders that hold saved files with their file count, total size and
// newest save time, newest first. Independent of whether the routine or run still
// exists; `flows` names the routines that saved into each folder.
func (m *RoutineFileManager) ListFoldersAutoPaging(ctx context.Context, params mobilerun.RoutineFileListFoldersParams, opts ...option.RequestOption) *pagination.PageAutoPager[mobilerun.RoutineFileListFoldersResponse] {
	return m.c.api.RoutineFiles.ListFoldersAutoPaging(ctx, params, append(append([]option.RequestOption{}, opts...), callerContext(ctx))...)
}

// AgentFileManager creates, fetches and lists AgentFile handles.
type AgentFileManager struct {
	c *clientRef
}

func newAgentFileManager(c *clientRef) *AgentFileManager {
	m := &AgentFileManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *AgentFileManager) FromID(id string) *AgentFile { return newAgentFile(m.c, id) }

func (m *AgentFileManager) String() string { return objectString("AgentFileManager", nil) }

// List the user's ready files, optionally filtered by zone.
//
// List returns the first page of AgentFile handles; ListAutoPaging iterates every
// page.
func (m *AgentFileManager) List(ctx context.Context, params mobilerun.FileListParams, opts ...option.RequestOption) (*ObjectPage[*AgentFile, *mobilerun.FileListResponse], error) {
	load := func() (*ObjectPage[*AgentFile, *mobilerun.FileListResponse], error) {
		pageOpts := opts
		res, err := m.c.api.Files.List(ctx, params, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*AgentFile, *mobilerun.FileListResponse]{Response: res}
		for _, info := range objectRows(res, "files") {
			result.Items = append(result.Items, newAgentFile(m.c, objectID(info, "fileId")).withInfo(info, ""))
		}
		return result, nil
	}
	return load()
}

// ListAutoPaging iterates the AgentFile handles of every page.
func (m *AgentFileManager) ListAutoPaging(ctx context.Context, params mobilerun.FileListParams, opts ...option.RequestOption) *ObjectAutoPager[*AgentFile] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// Upload uploads a file to presigned storage, confirms the upload and returns the
// AgentFile. Uploading uses storage quota and may incur charges; the returned
// handle does not own cleanup. Failures before confirmation clean up the pending
// upload once, best-effort, and return an UploadError.
func (m *AgentFileManager) Upload(ctx context.Context, file UploadFile, params mobilerun.FileUploadURLParams, opts ...option.RequestOption) (*AgentFile, error) {
	prepared, err := prepareUploads([]UploadFile{file}, nil)
	if err != nil {
		return nil, err
	}
	requestOpts := bindOptions(opts, binding{"filename", "body", prepared[0].filename}, binding{"mimeType", "body", prepared[0].contentType}, binding{"sizeBytes", "body", prepared[0].size})
	issued, err := m.c.api.Files.UploadURL(ctx, params, requestOpts...)
	if err != nil {
		return nil, err
	}
	uploadID := objectID(issued, "fileId")
	confirming, confirmedOK := false, false
	cleanup := func(cause error) error {
		if !confirmedOK && (!confirming || definiteConfirmFailure(cause)) {
			cleanupCtx, cancel := cleanupContext(ctx)
			_, _ = m.c.api.Files.CancelPending(cleanupCtx, uploadID, opts...)
			cancel()
		}
		if ctx.Err() != nil {
			return cause
		}
		return &UploadError{Cause: cause}
	}
	urls, err := uploadURLs(issued, prepared, "putUrl", "", "", "")
	if err != nil {
		return nil, cleanup(err)
	}
	for i, file := range prepared {
		if err := putUpload(ctx, m.c, file, urls[i], opts); err != nil {
			return nil, cleanup(err)
		}
	}
	confirming = true
	confirmed, err := m.c.api.Files.Confirm(ctx, uploadID, opts...)
	if err != nil {
		return nil, cleanup(err)
	}
	confirmedOK = true
	return newAgentFile(m.c, objectID(confirmed, "fileId")).withInfo(confirmed, ""), nil
}

// Mint a presigned PUT URL and a fileId for uploading a user file.
func (m *AgentFileManager) UploadURL(ctx context.Context, params mobilerun.FileUploadURLParams, opts ...option.RequestOption) (*mobilerun.FileUploadURLResponse, error) {
	return m.c.api.Files.UploadURL(ctx, params, opts...)
}

// AgentManager creates, fetches and lists Agent handles.
type AgentManager struct {
	c *clientRef

	Requests *AgentRequestManager
}

func newAgentManager(c *clientRef) *AgentManager {
	m := &AgentManager{c: c}
	m.Requests = newAgentRequestManager(m.c)
	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *AgentManager) FromID(id string) *Agent { return newAgent(m.c, id) }

func (m *AgentManager) String() string { return objectString("AgentManager", nil) }

// Creates an Agent directly on the owner runtime, without a chat message or
// approval card. Provisioning runs in the background; the new Agent has no chat
// history. Requires Idempotency-Key; a replay returns the current Agent with the
// original 201 status.
//
// New creates a Agent. The returned handle owns cleanup: Close disposes it.
func (m *AgentManager) New(ctx context.Context, params mobilerun.AgentNewParams, opts ...option.RequestOption) (*Agent, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Agents.New(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newAgent(m.c, objectID(res, "agent.id")).withInfo(res, "agent")
	o.state.setOwned(true)
	return o, nil
}

// Owner-scoped. Returns every Agent of the owner that is not on a retired machine,
// archived and failed ones included, oldest first.
//
// List returns the first page of Agent handles; ListAutoPaging iterates every
// page.
func (m *AgentManager) List(ctx context.Context, opts ...option.RequestOption) (*ObjectPage[*Agent, *mobilerun.AgentListResponse], error) {
	load := func() (*ObjectPage[*Agent, *mobilerun.AgentListResponse], error) {
		pageOpts := opts
		res, err := m.c.api.Agents.List(ctx, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*Agent, *mobilerun.AgentListResponse]{Response: res}
		for _, info := range objectRows(res, "agents") {
			result.Items = append(result.Items, newAgent(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		return result, nil
	}
	return load()
}

// ListAutoPaging iterates the Agent handles of every page.
func (m *AgentManager) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *ObjectAutoPager[*Agent] {
	page, err := m.List(ctx, opts...)
	return newObjectAutoPager(page, err)
}

// List the environment model profiles and their default. Model routing is managed
// by the server.
func (m *AgentManager) ModelProfiles(ctx context.Context, opts ...option.RequestOption) (*mobilerun.AgentModelProfilesResponse, error) {
	return m.c.api.Agents.ModelProfiles(ctx, opts...)
}

// AgentRequestManager creates, fetches and lists AgentRequest handles.
type AgentRequestManager struct {
	c *clientRef
}

func newAgentRequestManager(c *clientRef) *AgentRequestManager {
	m := &AgentRequestManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *AgentRequestManager) FromID(id string) *AgentRequest { return newAgentRequest(m.c, id) }

func (m *AgentRequestManager) String() string { return objectString("AgentRequestManager", nil) }

// Owner-scoped. Returns the request with its payloads and event timeline; an
// unknown id or another owner's request is a 404.
//
// Get fetches one AgentRequest and returns its handle.
func (m *AgentRequestManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*AgentRequest, error) {
	res, err := m.c.api.Agents.Requests.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newAgentRequest(m.c, id).withInfo(res, ""), nil
}

// Owner-scoped. Returns newest requests first, ordered by creation time and id.
// Use nextCursor as before to read older pages. Invalid cursors return 400
// invalid_cursor. Use the single read for payloads and events.
//
// List returns the first page of AgentRequest handles; ListAutoPaging iterates
// every page.
func (m *AgentRequestManager) List(ctx context.Context, params mobilerun.AgentRequestListParams, opts ...option.RequestOption) (*ObjectPage[*AgentRequest, *mobilerun.AgentRequestListResponse], error) {
	load := func() (*ObjectPage[*AgentRequest, *mobilerun.AgentRequestListResponse], error) {
		pageOpts := opts
		res, err := m.c.api.Agents.Requests.List(ctx, params, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*AgentRequest, *mobilerun.AgentRequestListResponse]{Response: res}
		for _, info := range objectRows(res, "requests") {
			result.Items = append(result.Items, newAgentRequest(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		return result, nil
	}
	return load()
}

// ListAutoPaging iterates the AgentRequest handles of every page.
func (m *AgentRequestManager) ListAutoPaging(ctx context.Context, params mobilerun.AgentRequestListParams, opts ...option.RequestOption) *ObjectAutoPager[*AgentRequest] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// Owner-scoped. Returns newest requests where the Agent is source or target. Use
// nextCursor as before to read older pages. Invalid cursors return 400
// invalid_cursor.
func (m *AgentRequestManager) ListForAgent(ctx context.Context, agentID string, params mobilerun.AgentRequestListForAgentParams, opts ...option.RequestOption) (*mobilerun.AgentRequestListForAgentResponse, error) {
	return m.c.api.Agents.Requests.ListForAgent(ctx, agentID, params, opts...)
}

// Owner-scoped. Returns the Agent-to-Agent request whose target session is
// {sessionId}, including its payloads and event timeline. Unknown sessions,
// source-only sessions, sessions without a request, and another owner's sessions
// return 404 agent_request_not_found.
//
// GetBySession returns a AgentRequest handle; the handle does not own cleanup.
func (m *AgentRequestManager) GetBySession(ctx context.Context, sessionID string, opts ...option.RequestOption) (*AgentRequest, error) {
	res, err := m.c.api.Agents.Requests.GetBySession(ctx, sessionID, opts...)
	if err != nil {
		return nil, err
	}
	return newAgentRequest(m.c, objectID(res, "id")).withInfo(res, ""), nil
}

// NumberMessageManager creates, fetches and lists NumberMessage handles.
type NumberMessageManager struct {
	c *clientRef

	Conversations *NumberMessageManagerConversations
}

func newNumberMessageManager(c *clientRef) *NumberMessageManager {
	m := &NumberMessageManager{c: c}
	m.Conversations = newNumberMessageManagerConversations(m)
	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *NumberMessageManager) FromID(id string) *NumberMessage { return newNumberMessage(m.c, id) }

func (m *NumberMessageManager) String() string { return objectString("NumberMessageManager", nil) }

// Returns one message. For an outbound message that is not yet final, the status
// may be refreshed once (at most every 30 s) before it is returned; a failed
// refresh returns the stored message.
//
// Get fetches one NumberMessage and returns its handle.
func (m *NumberMessageManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*NumberMessage, error) {
	res, err := m.c.api.Messages.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newNumberMessage(m.c, id).withInfo(res, "data"), nil
}

// Lists SMS messages newest first, with filters for direction, eSIM, phone number,
// status, and conversation.
//
// List returns the first page of NumberMessage handles; ListAutoPaging iterates
// every page.
func (m *NumberMessageManager) List(ctx context.Context, params mobilerun.MessageListParams, opts ...option.RequestOption) (*ObjectPage[*NumberMessage, *pagination.Page[mobilerun.PublicMessageRow]], error) {
	var wrap func(page *pagination.Page[mobilerun.PublicMessageRow]) *ObjectPage[*NumberMessage, *pagination.Page[mobilerun.PublicMessageRow]]
	wrap = func(page *pagination.Page[mobilerun.PublicMessageRow]) *ObjectPage[*NumberMessage, *pagination.Page[mobilerun.PublicMessageRow]] {
		result := &ObjectPage[*NumberMessage, *pagination.Page[mobilerun.PublicMessageRow]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newNumberMessage(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*NumberMessage, *pagination.Page[mobilerun.PublicMessageRow]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Messages.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the NumberMessage handles of every page.
func (m *NumberMessageManager) ListAutoPaging(ctx context.Context, params mobilerun.MessageListParams, opts ...option.RequestOption) *ObjectAutoPager[*NumberMessage] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// NumberMessageManagerConversations groups the conversations methods of a NumberMessageManager.
type NumberMessageManagerConversations struct {
	m *NumberMessageManager
}

func newNumberMessageManagerConversations(m *NumberMessageManager) *NumberMessageManagerConversations {
	n := &NumberMessageManagerConversations{m: m}

	return n
}

func (n *NumberMessageManagerConversations) String() string {
	return objectString("NumberMessageManagerConversations", nil)
}

// Lists SMS conversations by recent activity. Use both cursor fields from
// `nextCursor` to fetch the next page.
func (n *NumberMessageManagerConversations) List(ctx context.Context, params mobilerun.MessageConversationListParams, opts ...option.RequestOption) (*mobilerun.MessageConversationListResponse, error) {
	m := n.m
	return m.c.api.Messages.Conversations.List(ctx, params, opts...)
}

// Marks inbound messages in a conversation as read through the supplied cursor.
// Repeating the request is safe.
func (n *NumberMessageManagerConversations) MarkRead(ctx context.Context, params mobilerun.MessageConversationMarkReadParams, opts ...option.RequestOption) (*mobilerun.MessageConversationMarkReadResponse, error) {
	m := n.m
	return m.c.api.Messages.Conversations.MarkRead(ctx, params, opts...)
}

// PhoneNumberManager creates, fetches and lists PhoneNumber handles.
type PhoneNumberManager struct {
	c *clientRef
}

func newPhoneNumberManager(c *clientRef) *PhoneNumberManager {
	m := &PhoneNumberManager{c: c}

	return m
}

// FromID returns an unfetched handle; it makes no request and has no Info.
func (m *PhoneNumberManager) FromID(id string) *PhoneNumber { return newPhoneNumber(m.c, id) }

func (m *PhoneNumberManager) String() string { return objectString("PhoneNumberManager", nil) }

// Retrieves a single phone number.
//
// Get fetches one PhoneNumber and returns its handle.
func (m *PhoneNumberManager) Get(ctx context.Context, id string, opts ...option.RequestOption) (*PhoneNumber, error) {
	res, err := m.c.api.Numbers.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return newPhoneNumber(m.c, id).withInfo(res, "data"), nil
}

// Purchases a phone number. Poll the returned number for status updates.
//
// Starts a phone-number purchase. Poll the returned phone number for status
// updates.
//
// New creates a PhoneNumber. The returned handle owns cleanup: Close disposes it.
func (m *PhoneNumberManager) New(ctx context.Context, params mobilerun.NumberNewParams, opts ...option.RequestOption) (*PhoneNumber, error) {
	_, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	res, err := m.c.api.Numbers.New(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	o := newPhoneNumber(m.c, objectID(res, "data.numberId")).withInfo(res, "data")
	o.state.setOwned(true)
	return o, nil
}

// Lists the caller's phone numbers.
//
// List returns the first page of PhoneNumber handles; ListAutoPaging iterates
// every page.
func (m *PhoneNumberManager) List(ctx context.Context, params mobilerun.NumberListParams, opts ...option.RequestOption) (*ObjectPage[*PhoneNumber, *pagination.Page[mobilerun.PublicNumber]], error) {
	var wrap func(page *pagination.Page[mobilerun.PublicNumber]) *ObjectPage[*PhoneNumber, *pagination.Page[mobilerun.PublicNumber]]
	wrap = func(page *pagination.Page[mobilerun.PublicNumber]) *ObjectPage[*PhoneNumber, *pagination.Page[mobilerun.PublicNumber]] {
		result := &ObjectPage[*PhoneNumber, *pagination.Page[mobilerun.PublicNumber]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newPhoneNumber(m.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*PhoneNumber, *pagination.Page[mobilerun.PublicNumber]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := m.c.api.Numbers.List(ctx, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the PhoneNumber handles of every page.
func (m *PhoneNumberManager) ListAutoPaging(ctx context.Context, params mobilerun.NumberListParams, opts ...option.RequestOption) *ObjectAutoPager[*PhoneNumber] {
	page, err := m.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// Returns included phone capacity for a country. Creating a phone is
// authoritative.
func (m *PhoneNumberManager) Capacity(ctx context.Context, params mobilerun.NumberCapacityParams, opts ...option.RequestOption) (*mobilerun.NumberCapacityResponse, error) {
	return m.c.api.Numbers.Capacity(ctx, params, opts...)
}

// Lists available countries and current phone-number availability.
func (m *PhoneNumberManager) Countries(ctx context.Context, opts ...option.RequestOption) (*mobilerun.NumberCountriesResponse, error) {
	return m.c.api.Numbers.Countries(ctx, opts...)
}

// App is a handle to one App.
//
// Field accessors read the last response (Info) without making a request.
type App struct {
	c *clientRef

	id    string
	state objectState
}

func newApp(c *clientRef, id string) *App {
	o := &App{c: c, id: id}

	return o
}

func (o *App) withInfo(info any, root string) *App {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *App) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *App) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *App) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *App) String() string {
	if o.Info() == nil {
		return objectString("App", &o.id)
	}
	return objectString("App", &o.id, "displayName", o.DisplayName(), "platform", o.Platform())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *App) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *App) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.AppGetResponse, error) {
	res, err := o.c.api.Apps.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// DisplayName reads displayName. It reads the last response without making a
// request.
func (o *App) DisplayName() string { return resultString(o.state.field("displayName")) }

// Platform reads platform. It reads the last response without making a request.
func (o *App) Platform() string { return resultString(o.state.field("platform")) }

// Deletes an uploaded app by ID. Removes its files and metadata.
func (o *App) Delete(ctx context.Context, opts ...option.RequestOption) (*mobilerun.AppDeleteResponse, error) {
	return o.c.api.Apps.Delete(ctx, o.id, opts...)
}

// Retrieves all versions of an app visible to the user (own uploads + system
// versions)
func (o *App) ListVersions(ctx context.Context, opts ...option.RequestOption) (*mobilerun.AppListVersionsResponse, error) {
	return o.c.api.Apps.ListVersions(ctx, o.id, opts...)
}

// A package identified by its name. get preserves the credentials-array response
// in info.
//
// Field accessors read the last response (Info) without making a request.
type CredentialPackage struct {
	c *clientRef

	id          string
	state       objectState
	Credentials *CredentialManager
	Links       *CredentialPackageLinks
}

func newCredentialPackage(c *clientRef, id string) *CredentialPackage {
	o := &CredentialPackage{c: c, id: id}
	o.Credentials = newCredentialManager(o.c, o.id)
	o.Links = newCredentialPackageLinks(o)
	return o
}

func (o *CredentialPackage) withInfo(info any, root string) *CredentialPackage {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *CredentialPackage) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *CredentialPackage) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *CredentialPackage) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *CredentialPackage) String() string {
	if o.Info() == nil {
		return objectString("CredentialPackage", &o.id)
	}
	return objectString("CredentialPackage", &o.id)
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *CredentialPackage) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *CredentialPackage) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.CredentialPackageListResponse, error) {
	res, err := o.c.api.Credentials.Packages.List(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "")
	return res, nil
}

// CredentialPackageLinks groups the links methods of a CredentialPackage.
type CredentialPackageLinks struct {
	o *CredentialPackage
}

func newCredentialPackageLinks(o *CredentialPackage) *CredentialPackageLinks {
	n := &CredentialPackageLinks{o: o}

	return n
}

func (n *CredentialPackageLinks) String() string { return objectString("CredentialPackageLinks", nil) }

// Creates a read-only credential that resolves its values live from the provider.
// `provider` defaults to `onepassword` (vaultId, itemId, fieldMap). `wallet` takes
// a cardId, is only allowed in the reserved payment_cards package, and only the
// cardholder may link the card (404 otherwise). Never replaces an existing
// credential of the same name.
func (n *CredentialPackageLinks) New(ctx context.Context, params mobilerun.CredentialPackageLinkNewParams, opts ...option.RequestOption) (*mobilerun.CredentialPackageLinkNewResponse, error) {
	o := n.o
	return o.c.api.Credentials.Packages.Links.New(ctx, o.id, params, opts...)
}

// Credential is a handle to one Credential.
//
// Field accessors read the last response (Info) without making a request.
type Credential struct {
	c           *clientRef
	packageName string
	id          string
	state       objectState
	Fields      *CredentialFields
}

func newCredential(c *clientRef, packageName string, id string) *Credential {
	o := &Credential{c: c, packageName: packageName, id: id}
	o.Fields = newCredentialFields(o)
	return o
}

func (o *Credential) withInfo(info any, root string) *Credential {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *Credential) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *Credential) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *Credential) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *Credential) String() string {
	if o.Info() == nil {
		return objectString("Credential", &o.id)
	}
	return objectString("Credential", &o.id, "source", o.Source())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *Credential) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Source reads source. It reads only the last New/List snapshot without making a
// request; FromID handles have no data.
func (o *Credential) Source() string { return resultString(o.state.field("source")) }

// PackageName returns the bound CredentialPackage ID.
func (o *Credential) PackageName() string { return o.packageName }

// Sets or clears the optional `label` (human-readable name, free text, max 200
// characters) of an existing credential. `null` or an empty text clears it. Works
// for linked credentials too: the label is metadata of the credential, not a field
// of the linked item. Returns the updated credential.
//
// The object layer sets packageName; values for it in params are replaced.
func (o *Credential) Update(ctx context.Context, params mobilerun.CredentialPackageCredentialUpdateParams, opts ...option.RequestOption) (*mobilerun.CredentialPackageCredentialUpdateResponse, error) {
	params.PackageName = o.packageName
	return o.c.api.Credentials.Packages.Credentials.Update(ctx, o.id, params, opts...)
}

// Permanently deletes the credential identified by `packageName` and
// `credentialName`, removing all of its fields. Returns the deleted credential.
//
// The object layer sets packageName; values for it in params are replaced.
func (o *Credential) Delete(ctx context.Context, opts ...option.RequestOption) (*mobilerun.CredentialPackageCredentialDeleteResponse, error) {
	return o.c.api.Credentials.Packages.Credentials.Delete(ctx, o.id, mobilerun.CredentialPackageCredentialDeleteParams{PackageName: o.packageName}, opts...)
}

// CredentialFields groups the fields methods of a Credential.
type CredentialFields struct {
	o *Credential
}

func newCredentialFields(o *Credential) *CredentialFields {
	n := &CredentialFields{o: o}

	return n
}

func (n *CredentialFields) String() string { return objectString("CredentialFields", nil) }

// Adds a single field to an existing credential. The body specifies a `fieldType`
// (one of the supported field types) and its value. Returns a conflict if a field
// of that type already exists on the credential.
//
// The object layer sets packageName; values for it in params are replaced.
func (n *CredentialFields) New(ctx context.Context, params mobilerun.CredentialPackageCredentialFieldNewParams, opts ...option.RequestOption) (*mobilerun.CredentialPackageCredentialFieldNewResponse, error) {
	o := n.o
	params.PackageName = o.packageName
	return o.c.api.Credentials.Packages.Credentials.Fields.New(ctx, o.id, params, opts...)
}

// Updates the value of an existing field on a credential, identified by
// `packageName`, `credentialName`, and `fieldType` in the path. The body carries
// the new value and returns the updated credential.
//
// The object layer sets credentialName, packageName; values for them in params are
// replaced.
func (n *CredentialFields) Update(ctx context.Context, fieldType mobilerun.CredentialPackageCredentialFieldUpdateParamsFieldType, params mobilerun.CredentialPackageCredentialFieldUpdateParams, opts ...option.RequestOption) (*mobilerun.CredentialPackageCredentialFieldUpdateResponse, error) {
	o := n.o
	params.CredentialName = o.id
	params.PackageName = o.packageName
	return o.c.api.Credentials.Packages.Credentials.Fields.Update(ctx, fieldType, params, opts...)
}

// Removes a single field of the given `fieldType` from the specified credential
// while leaving the credential itself intact. Returns the updated credential.
//
// The object layer sets credentialName, packageName; values for them in params are
// replaced.
func (n *CredentialFields) Delete(ctx context.Context, fieldType mobilerun.CredentialPackageCredentialFieldDeleteParamsFieldType, opts ...option.RequestOption) (*mobilerun.CredentialPackageCredentialFieldDeleteResponse, error) {
	o := n.o
	return o.c.api.Credentials.Packages.Credentials.Fields.Delete(ctx, fieldType, mobilerun.CredentialPackageCredentialFieldDeleteParams{CredentialName: o.id, PackageName: o.packageName}, opts...)
}

// Device is a handle to one Device.
//
// Field accessors read the last response (Info) without making a request.
type Device struct {
	c *clientRef

	id              string
	state           objectState
	Actions         *DeviceActions
	Apps            *DeviceApps
	Esim            *DeviceEsim
	Files           *DeviceFiles
	Keyboard        *DeviceKeyboard
	Location        *DeviceLocation
	Packages        *DevicePackages
	Profile         *DeviceProfile
	Proxy           *DeviceProxy
	Screenshots     *SavedScreenshotManager
	Tasks           *DeviceTasks
	Clipboard       *DeviceClipboard
	Timezone        *DeviceTimezone
	Language        *DeviceLanguage
	DeepLink        *DeviceDeepLink
	Browser         *DeviceBrowser
	Kiosk           *DeviceKiosk
	Recordings      *RecordingManager
	TrafficSessions *TrafficSessionManager
	Number          *DeviceNumber
}

func newDevice(c *clientRef, id string) *Device {
	o := &Device{c: c, id: id}
	o.Actions = newDeviceActions(o)
	o.Apps = newDeviceApps(o)
	o.Esim = newDeviceEsim(o)
	o.Files = newDeviceFiles(o)
	o.Keyboard = newDeviceKeyboard(o)
	o.Location = newDeviceLocation(o)
	o.Packages = newDevicePackages(o)
	o.Profile = newDeviceProfile(o)
	o.Proxy = newDeviceProxy(o)
	o.Screenshots = newSavedScreenshotManager(o.c, o.id)
	o.Tasks = newDeviceTasks(o)
	o.Clipboard = newDeviceClipboard(o)
	o.Timezone = newDeviceTimezone(o)
	o.Language = newDeviceLanguage(o)
	o.DeepLink = newDeviceDeepLink(o)
	o.Browser = newDeviceBrowser(o)
	o.Kiosk = newDeviceKiosk(o)
	o.Recordings = newRecordingManager(o.c, o.id)
	o.TrafficSessions = newTrafficSessionManager(o.c, o.id)
	o.Number = newDeviceNumber(o)
	return o
}

func (o *Device) withInfo(info any, root string) *Device {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *Device) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *Device) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *Device) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *Device) String() string {
	if o.Info() == nil {
		return objectString("Device", &o.id)
	}
	return objectString("Device", &o.id, "name", o.Name(), "state", o.State())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *Device) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	err := o.c.api.Devices.Terminate(ctx, o.id, mobilerun.DeviceTerminateParams{}, opts...)
	return err
}

// Refresh fetches the object and updates Info.
func (o *Device) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.Device, error) {
	res, err := o.c.api.Devices.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "")
	return res, nil
}

// WaitReady waits until the object is ready, for up to 300000 ms
// (WithReadyTimeout). Request timeouts and 408 are retried immediately; 409, 429
// and server errors with backoff. It updates Info.
func (o *Device) WaitReady(ctx context.Context, opts ...option.RequestOption) (*mobilerun.Device, error) {
	controls, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	return waitFor(ctx, durationOr(controls.readyTimeout, 300000*time.Millisecond), singleCallInterval, 0, func(ctx context.Context, pollOpts []option.RequestOption) (*mobilerun.Device, bool, error) {
		requestOpts := append(append([]option.RequestOption{}, opts...), pollOpts...)
		res, err := o.c.api.Devices.WaitReady(ctx, o.id, requestOpts...)
		if err != nil {
			return nil, false, err
		}
		o.state.setInfo(res, "")
		return res, true, nil
	}, retryReadiness)
}

// Name reads name. It reads the last response without making a request.
func (o *Device) Name() string { return resultString(o.state.field("name")) }

// State reads state. It reads the last response without making a request.
func (o *Device) State() string { return resultString(o.state.field("state")) }

// Wait for the declared success condition; failure raises with the last response.
//
// WaitStopped polls every 1000 ms for up to 300000 ms (WithPollInterval,
// WithWaitTimeout); failure states return a WaitFailureError and the deadline a
// WaitTimeoutError.
func (o *Device) WaitStopped(ctx context.Context, opts ...option.RequestOption) (*mobilerun.Device, error) {
	controls, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	poll := func(ctx context.Context, pollOpts []option.RequestOption) (*mobilerun.Device, bool, error) {
		requestOpts := append(append([]option.RequestOption{}, opts...), pollOpts...)
		res, err := o.c.api.Devices.Get(ctx, o.id, requestOpts...)
		if err != nil {
			return nil, false, err
		}
		o.state.setInfo(res, "")
		done, err := enumState(res, "state", []string{"stopped"}, []string{"terminate", "terminated"}, []string{"unknown", "creating", "assigned", "ready", "reboot", "rebooting", "reset", "resetting", "migrate", "migrating", "maintain", "maintenance"}, true)
		return res, done, err
	}
	return waitFor(ctx, durationOr(controls.waitTimeout, 300000*time.Millisecond), durationOr(controls.pollInterval, 1000*time.Millisecond), 30000*time.Millisecond, poll, retryPoll)
}

// Returns a live snapshot of the device's spoofed identity, including model,
// display, identifiers, and carrier. Devices without fingerprint support return an
// unsupported-feature error.
func (o *Device) Fingerprint(ctx context.Context, params mobilerun.DeviceFingerprintParams, opts ...option.RequestOption) (*mobilerun.DeviceFingerprintResponse, error) {
	return o.c.api.Devices.Fingerprint(ctx, o.id, params, opts...)
}

// Triggers a reboot of the device. The device transitions through its reboot
// lifecycle and becomes ready again once the restart completes. Only devices whose
// capabilities report `reboot=true` can be rebooted. Personal (BYOD) devices
// cannot be rebooted remotely: the request is refused with 400
// `DEVICE_TOOL_UNSUPPORTED` and nothing happens on the device.
func (o *Device) Reboot(ctx context.Context, opts ...option.RequestOption) error {
	return o.c.api.Devices.Reboot(ctx, o.id, opts...)
}

// Resets the device. `mode` selects what is erased: `full` (factory reset: apps,
// accounts, storage, new identity) or `identity` (new fingerprint, apps and data
// kept). Without `mode` the reset is `full`, or `identity` on devices with
// `factoryReset=false`. A mode or `preserve` value the device does not support is
// refused with 400 `DEVICE_TOOL_UNSUPPORTED`. `preserve` restores the latest
// proxy, location, locale or timezone set through the API after the reset; omitted
// or empty preserves nothing. The device becomes ready again after the reset.
func (o *Device) Reset(ctx context.Context, params mobilerun.DeviceResetParams, opts ...option.RequestOption) error {
	return o.c.api.Devices.Reset(ctx, o.id, params, opts...)
}

// Returns the set of capabilities supported by this device. Used to determine
// which tools and features are available for the device.
func (o *Device) GetCapabilities(ctx context.Context, opts ...option.RequestOption) (*mobilerun.DeviceGetCapabilitiesResponse, error) {
	return o.c.api.Devices.GetCapabilities(ctx, o.id, opts...)
}

// Sets the display name for a device from the name in the request body and returns
// the updated device.
func (o *Device) SetName(ctx context.Context, params mobilerun.DeviceSetNameParams, opts ...option.RequestOption) (*mobilerun.Device, error) {
	return o.c.api.Devices.SetName(ctx, o.id, params, opts...)
}

// Terminates the device and releases its resources. For a personal (BYOD) device
// the owner (API key or user) can instead remove the phone from the account so
// anyone can pair it again; the phone must be disconnected first, otherwise the
// request is rejected with 409 BYOD_DEVICE_CONNECTED. Other subscription-billed
// devices cannot be terminated with this request; it is rejected with 403
// FORBIDDEN: Subscription devices are managed through subscription cancellation.
// This includes owner requests that set `terminateAt` or `previousDeviceId`. Those
// body fields are reserved for internal use; owner requests on other devices that
// set them are rejected with 401.
func (o *Device) Terminate(ctx context.Context, params mobilerun.DeviceTerminateParams, opts ...option.RequestOption) error {
	return o.c.api.Devices.Terminate(ctx, o.id, params, opts...)
}

// Mints a short-lived single-use stream token for this device. Returns the token
// response; does not update the device snapshot.
//
// Mint a short-lived, single-use stream token for a device. The token permits
// streaming and control.
//
// The object layer sets deviceId; values for it in params are replaced.
func (o *Device) NewStreamToken(ctx context.Context, params mobilerun.StreamTokenNewParams, opts ...option.RequestOption) (*mobilerun.StreamTokenCreateResponse, error) {
	return o.c.api.StreamTokens.New(ctx, params, bindOptions(opts, binding{"deviceId", "body", o.id})...)
}

// Sets the idle timeout to 5–240 minutes, 0 to turn it off, or null to use the
// platform default. Returns the updated device as a new handle; call refresh() on
// an existing handle to update it. A custom timeout returns 409 when the device
// cannot be stopped; other values return 422.
//
// Sets the device's own idle stop setting and returns the updated device.
// `minutes` between 5 and 240 sets a custom timeout; `0` turns the idle timer off
// for this device; `null` removes the device's own setting, so the platform
// default applies again (a pay-as-you-go device that can be stopped stops after
// the default; other devices have no timer). Other values return 422. While a
// timer is on, the device is stopped automatically (state `stopped`, `stopReason`
// `IDLE_TIMEOUT`) once it has had no client activity for longer than that: API
// tool calls, open stream/adb/browser/frida/port connections and a held lease all
// count as activity, reading device state does not. Setting a custom timeout
// requires a device whose capabilities report stop=true, otherwise 409; `0` and
// `null` are always allowed. The response's `idleTimeoutMinutes` is the effective
// timeout.
//
// SetIdleTimeout returns a Device handle; the handle does not own cleanup.
func (o *Device) SetIdleTimeout(ctx context.Context, minutes int64, opts ...option.RequestOption) (*Device, error) {
	res, err := o.c.api.Devices.UpdateIdleTimeout(ctx, o.id, mobilerun.DeviceUpdateIdleTimeoutParams{}, bindOptions(opts, binding{"minutes", "body", minutes})...)
	if err != nil {
		return nil, err
	}
	return newDevice(o.c, objectID(res, "id")).withInfo(res, ""), nil
}

// Wakes a parked device: it first checks that the device can run again, then the
// same device starts and per-minute billing resumes. On a device that is not
// parked, the request does nothing. Unsupported intents return 400
// DEVICE_TOOL_UNSUPPORTED.
//
// After success Resume waits until the object is ready again (opt out with
// WithoutWaitReady, set the timeout with WithReadyTimeout) and returns the same
// handle.
func (o *Device) Resume(ctx context.Context, opts ...option.RequestOption) (*Device, error) {
	controls, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	err = o.c.api.Devices.Resume(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	if !controls.noWaitReady {
		if _, err := o.WaitReady(ctx, readyOptions(controls, opts)...); err != nil {
			return o, &ObjectReadyError{Object: o, ID: o.id, Cause: err}
		}
	}
	return o, nil
}

// Parks the device: its data, apps and identity are kept, but nothing runs and
// nothing is billed until it is resumed. Hosted subscription devices refuse a new
// park with 400 DEVICE_TOOL_UNSUPPORTED. Stop remains available for cleanup when
// the device already targets stopped. Other devices require stop support;
// unsupported intents return 400 DEVICE_TOOL_UNSUPPORTED.
//
// After success Stop runs WaitStopped (opt out with WithoutWait, set the timeout
// with WithWaitTimeout) and returns the same handle.
func (o *Device) Stop(ctx context.Context, opts ...option.RequestOption) (*Device, error) {
	controls, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	err = o.c.api.Devices.Stop(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	if !controls.noWait {
		if _, err := o.WaitStopped(ctx, waitOptions(controls, opts)...); err != nil {
			return nil, err
		}
	}
	return o, nil
}

// Taps the device screen at the given x/y coordinates. An optional humanize flag
// routes the tap through human-like input on devices that support it.
func (o *Device) Tap(ctx context.Context, x int64, y int64, params mobilerun.DeviceActionTapParams, opts ...option.RequestOption) error {
	return o.c.api.Devices.Actions.Tap(ctx, o.id, params, bindOptions(opts, binding{"x", "body", x}, binding{"y", "body", y})...)
}

// Types text into the focused input field. The optional completionMode defaults to
// accepted for backwards-compatible low latency; committed additionally waits for
// the complete text or a quiescent UI state.
func (o *Device) Type(ctx context.Context, text string, params mobilerun.DeviceKeyboardWriteParams, opts ...option.RequestOption) error {
	return o.c.api.Devices.Keyboard.Write(ctx, o.id, params, bindOptions(opts, binding{"text", "body", text})...)
}

// Captures the device screen and returns it as a PNG image. Pass save=true to
// store it as device media.
//
// Screenshot returns the raw response bytes and the declared headers.
func (o *Device) Screenshot(ctx context.Context, params mobilerun.DeviceScreenshotCaptureParams, opts ...option.RequestOption) (*Screenshot, error) {
	var body []byte
	var response *http.Response
	_, err := o.c.api.Devices.Screenshots.Capture(ctx, o.id, params, append(opts, option.WithResponseBodyInto(&body), option.WithResponseInto(&response))...)
	if err != nil {
		return nil, err
	}
	result := &Screenshot{BinaryResult: BinaryResult{Bytes: body, ContentType: response.Header.Get("Content-Type")}}
	result.ID = response.Header.Get("X-Screenshot-Id")
	result.Saved = headerBool(response.Header.Get("X-Screenshot-Saved"))
	return result, nil
}

// Create and dispatch a new agent task. Returns the task ID and device stream
// details.
//
// Run returns a Task handle; the handle does not own cleanup.
//
// The object layer sets deviceId; values for it in params are replaced.
func (o *Device) Run(ctx context.Context, params mobilerun.TaskRunParams, opts ...option.RequestOption) (*Task, error) {
	res, err := o.c.api.Tasks.Run(ctx, params, bindOptions(opts, binding{"deviceId", "body", o.id})...)
	if err != nil {
		return nil, err
	}
	return newTask(o.c, objectID(res, "id")).withInfo(res, ""), nil
}

// Swipes from a start coordinate to an end coordinate over the given duration in
// milliseconds. The optional humanize flag (best-effort, also with holdMs) adds a
// human-like path and timing. With holdMs > 0 the request returns 422 when
// duration + holdMs exceeds 9000 ms, or on devices that do not support it.
func (o *Device) Swipe(ctx context.Context, startX int64, startY int64, endX int64, endY int64, duration int64, params mobilerun.DeviceActionSwipeParams, opts ...option.RequestOption) error {
	return o.c.api.Devices.Actions.Swipe(ctx, o.id, params, bindOptions(opts, binding{"startX", "body", startX}, binding{"startY", "body", startY}, binding{"endX", "body", endX}, binding{"endY", "body", endY}, binding{"duration", "body", duration})...)
}

// Presses a key using the platform key code.
//
// Sends a single Android key event to the device, identified by its key code.
func (o *Device) Key(ctx context.Context, key int64, params mobilerun.DeviceKeyboardKeyParams, opts ...option.RequestOption) error {
	return o.c.api.Devices.Keyboard.Key(ctx, o.id, params, bindOptions(opts, binding{"key", "body", key})...)
}

// Launches the app identified by the path package name, optionally starting a
// specific activity given in the request body. Protected packages cannot be
// started.
//
// The object layer sets deviceId; values for it in params are replaced.
func (o *Device) OpenApp(ctx context.Context, packageName string, params mobilerun.DeviceAppStartParams, opts ...option.RequestOption) error {
	params.DeviceID = o.id
	return o.c.api.Devices.Apps.Start(ctx, packageName, params, opts...)
}

// Force-stops the app identified by the path package name. When clearData is set
// in the request body, the app's data is also cleared. Protected packages cannot
// be stopped.
//
// The object layer sets deviceId; values for it in params are replaced.
func (o *Device) StopApp(ctx context.Context, packageName string, params mobilerun.DeviceAppStopParams, opts ...option.RequestOption) error {
	params.DeviceID = o.id
	return o.c.api.Devices.Apps.Stop(ctx, packageName, params, opts...)
}

// Presses the Home button.
//
// Performs a global system action on the device, such as navigating back or going
// to the home screen, identified by an action code.
//
// The object layer sets action; values for it in params are replaced.
func (o *Device) Home(ctx context.Context, params mobilerun.DeviceActionGlobalParams, opts ...option.RequestOption) error {
	return o.c.api.Devices.Actions.Global(ctx, o.id, params, bindOptions(opts, binding{"action", "body", 2})...)
}

// Presses the Back button.
//
// Performs a global system action on the device, such as navigating back or going
// to the home screen, identified by an action code.
//
// The object layer sets action; values for it in params are replaced.
func (o *Device) Back(ctx context.Context, params mobilerun.DeviceActionGlobalParams, opts ...option.RequestOption) error {
	return o.c.api.Devices.Actions.Global(ctx, o.id, params, bindOptions(opts, binding{"action", "body", 1})...)
}

// Returns the current accessibility UI state as a structured tree of on-screen
// elements.
//
// Returns the current accessibility UI state of the device as a structured tree of
// on-screen elements. An optional filter query reduces the result to interactive
// elements.
func (o *Device) Ui(ctx context.Context, params mobilerun.DeviceStateUiParams, opts ...option.RequestOption) (*mobilerun.DeviceStateUiResponse, error) {
	return o.c.api.Devices.State.Ui(ctx, o.id, params, opts...)
}

// Returns the device wall-clock time as an RFC 3339 timestamp.
//
// Returns the device's current wall-clock time as an RFC 3339 timestamp.
func (o *Device) Time(ctx context.Context, params mobilerun.DeviceStateTimeParams, opts ...option.RequestOption) (*string, error) {
	return o.c.api.Devices.State.Time(ctx, o.id, params, opts...)
}

// DeviceActions groups the actions methods of a Device.
type DeviceActions struct {
	o *Device
}

func newDeviceActions(o *Device) *DeviceActions {
	n := &DeviceActions{o: o}

	return n
}

func (n *DeviceActions) String() string { return objectString("DeviceActions", nil) }

// Performs a global system action on the device, such as navigating back or going
// to the home screen, identified by an action code.
func (n *DeviceActions) Global(ctx context.Context, params mobilerun.DeviceActionGlobalParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Actions.Global(ctx, o.id, params, opts...)
}

// Returns whether the accessibility overlay is currently visible on the device.
func (n *DeviceActions) OverlayVisible(ctx context.Context, params mobilerun.DeviceActionOverlayVisibleParams, opts ...option.RequestOption) (*mobilerun.DeviceActionOverlayVisibleResponse, error) {
	o := n.o
	return o.c.api.Devices.Actions.OverlayVisible(ctx, o.id, params, opts...)
}

// Shows or hides the accessibility overlay on the device based on the visibility
// flag in the request body.
func (n *DeviceActions) SetOverlayVisible(ctx context.Context, params mobilerun.DeviceActionSetOverlayVisibleParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Actions.SetOverlayVisible(ctx, o.id, params, opts...)
}

// DeviceApps groups the apps methods of a Device.
type DeviceApps struct {
	o *Device
}

func newDeviceApps(o *Device) *DeviceApps {
	n := &DeviceApps{o: o}

	return n
}

func (n *DeviceApps) String() string { return objectString("DeviceApps", nil) }

// Returns detailed information about apps installed on the device, including
// package name and label. System and protected apps are excluded unless the
// corresponding query parameters are set.
func (n *DeviceApps) List(ctx context.Context, params mobilerun.DeviceAppListParams, opts ...option.RequestOption) (*[]mobilerun.DeviceAppListResponse, error) {
	o := n.o
	return o.c.api.Devices.Apps.List(ctx, o.id, params, opts...)
}

// Uninstalls the app identified by the path package name from the device.
// Protected packages cannot be deleted.
//
// The object layer sets deviceId; values for it in params are replaced.
func (n *DeviceApps) Delete(ctx context.Context, packageName string, params mobilerun.DeviceAppDeleteParams, opts ...option.RequestOption) error {
	o := n.o
	params.DeviceID = o.id
	return o.c.api.Devices.Apps.Delete(ctx, packageName, params, opts...)
}

// Grants an Android runtime permission to the package named in the path. The
// permission is given by its short name (e.g. POST_NOTIFICATIONS).
//
// The object layer sets deviceId; values for it in params are replaced.
func (n *DeviceApps) GrantPermission(ctx context.Context, permission mobilerun.DeviceAppGrantPermissionParamsPermission, params mobilerun.DeviceAppGrantPermissionParams, opts ...option.RequestOption) error {
	o := n.o
	params.DeviceID = o.id
	return o.c.api.Devices.Apps.GrantPermission(ctx, permission, params, opts...)
}

// Requests an app install on the device. The request body must supply exactly one
// of an Android packageName or an iOS bundleId. For packageName installs, optional
// country and versionCode select an exact regional uploaded Android version, and
// protected packages are rejected. background (default false) selects the response
// contract: false installs inline and returns the outcome directly (200 on
// success, an error status on failure); true accepts the request and runs the
// download + install in the background, returning 202 immediately — poll
// list-app-installs for the backend's view of that attempt's status. Refuses with
// 409 once 2 other installs are already running on the device, in either mode; a
// repeat request for an app that already has an install running is also refused
// with 409 rather than superseding it — retry once that attempt reaches a terminal
// state. Synchronous failures report APP_DOWNLOAD_FAILED (502) when the app file
// cannot be downloaded or APP_INSTALL_TIMEOUT (504) when the install times out.
func (n *DeviceApps) Install(ctx context.Context, params mobilerun.DeviceAppInstallParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Apps.Install(ctx, o.id, params, opts...)
}

// Reports the backend's durable view of background app-install attempts on this
// device — status reflects the install ATTEMPT, not device ground truth; list-apps
// remains authoritative for what is actually installed. Terminal and projected
// timeout records are kept ~15 minutes. Not gated on device readiness, so it also
// answers while the device is offline or crashed.
func (n *DeviceApps) ListInstalls(ctx context.Context, params mobilerun.DeviceAppListInstallsParams, opts ...option.RequestOption) (*mobilerun.DeviceAppListInstallsResponse, error) {
	o := n.o
	return o.c.api.Devices.Apps.ListInstalls(ctx, o.id, params, opts...)
}

// Revokes an Android runtime permission from the package named in the path. The
// permission is given by its short name (e.g. POST_NOTIFICATIONS).
//
// The object layer sets deviceId; values for it in params are replaced.
func (n *DeviceApps) RevokePermission(ctx context.Context, permission mobilerun.DeviceAppRevokePermissionParamsPermission, params mobilerun.DeviceAppRevokePermissionParams, opts ...option.RequestOption) error {
	o := n.o
	params.DeviceID = o.id
	return o.c.api.Devices.Apps.RevokePermission(ctx, permission, params, opts...)
}

// DeviceEsim groups the esim methods of a Device.
type DeviceEsim struct {
	o   *Device
	Apn *DeviceEsimApn
}

func newDeviceEsim(o *Device) *DeviceEsim {
	n := &DeviceEsim{o: o}
	n.Apn = newDeviceEsimApn(n.o)
	return n
}

func (n *DeviceEsim) String() string { return objectString("DeviceEsim", nil) }

// Returns the eSIM subscriptions currently provisioned on the device.
func (n *DeviceEsim) List(ctx context.Context, params mobilerun.DeviceEsimListParams, opts ...option.RequestOption) (*[]mobilerun.DeviceEsimListResponse, error) {
	o := n.o
	return o.c.api.Devices.Esim.List(ctx, o.id, params, opts...)
}

// Download profile and/or enable subscription.
func (n *DeviceEsim) Activate(ctx context.Context, params mobilerun.DeviceEsimActivateParams, opts ...option.RequestOption) (*mobilerun.DeviceEsimActivateResponse, error) {
	o := n.o
	return o.c.api.Devices.Esim.Activate(ctx, o.id, params, opts...)
}

// Enables the eSIM subscription identified by the subId in the request body so it
// becomes the active subscription.
func (n *DeviceEsim) Enable(ctx context.Context, params mobilerun.DeviceEsimEnableParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Esim.Enable(ctx, o.id, params, opts...)
}

// Deletes the eSIM subscription identified by the subId query parameter from the
// device.
func (n *DeviceEsim) Remove(ctx context.Context, params mobilerun.DeviceEsimRemoveParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Esim.Remove(ctx, o.id, params, opts...)
}

// Enables or disables data roaming for the device's eSIM based on the enabled flag
// in the request body.
func (n *DeviceEsim) SetRoaming(ctx context.Context, params mobilerun.DeviceEsimSetRoamingParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Esim.SetRoaming(ctx, o.id, params, opts...)
}

// Returns the connectivity status of the device's eSIM subscriptions.
func (n *DeviceEsim) Status(ctx context.Context, params mobilerun.DeviceEsimStatusParams, opts ...option.RequestOption) (*[]mobilerun.DeviceEsimStatusResponse, error) {
	o := n.o
	return o.c.api.Devices.Esim.Status(ctx, o.id, params, opts...)
}

// DeviceEsimApn groups the apn methods of a Device.
type DeviceEsimApn struct {
	o *Device
}

func newDeviceEsimApn(o *Device) *DeviceEsimApn {
	n := &DeviceEsimApn{o: o}

	return n
}

func (n *DeviceEsimApn) String() string { return objectString("DeviceEsimApn", nil) }

// Returns the access point names (APNs) configured for the device's active eSIM
// subscriptions.
func (n *DeviceEsimApn) List(ctx context.Context, params mobilerun.DeviceEsimApnListParams, opts ...option.RequestOption) (*[]mobilerun.DeviceEsimApnListResponse, error) {
	o := n.o
	return o.c.api.Devices.Esim.Apn.List(ctx, o.id, params, opts...)
}

// Marks an existing APN, identified by apnId, as the preferred APN for the given
// eSIM subscription in the request body.
func (n *DeviceEsimApn) Select(ctx context.Context, params mobilerun.DeviceEsimApnSelectParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Esim.Apn.Select(ctx, o.id, params, opts...)
}

// Creates an access point name (APN) from the request body and applies it to the
// given eSIM subscription. Type, protocol, and roaming protocol default to common
// values when omitted.
func (n *DeviceEsimApn) Set(ctx context.Context, params mobilerun.DeviceEsimApnSetParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Esim.Apn.Set(ctx, o.id, params, opts...)
}

// DeviceFiles groups the files methods of a Device.
type DeviceFiles struct {
	o *Device
}

func newDeviceFiles(o *Device) *DeviceFiles {
	n := &DeviceFiles{o: o}

	return n
}

func (n *DeviceFiles) String() string { return objectString("DeviceFiles", nil) }

// Lists the files at the directory path given in the path query parameter,
// returning each entry's metadata along with the path and total count.
func (n *DeviceFiles) List(ctx context.Context, params mobilerun.DeviceFileListParams, opts ...option.RequestOption) (*mobilerun.DeviceFileListResponse, error) {
	o := n.o
	return o.c.api.Devices.Files.List(ctx, o.id, params, opts...)
}

// Deletes the file at the path given in the path query parameter from the device.
func (n *DeviceFiles) Delete(ctx context.Context, params mobilerun.DeviceFileDeleteParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Files.Delete(ctx, o.id, params, opts...)
}

// Pulls the file at the given path query parameter from the device and returns its
// raw bytes as an octet-stream.
func (n *DeviceFiles) Download(ctx context.Context, params mobilerun.DeviceFileDownloadParams, opts ...option.RequestOption) (*string, error) {
	o := n.o
	return o.c.api.Devices.Files.Download(ctx, o.id, params, opts...)
}

// Uploads a file to the device via multipart form data, writing it into the
// directory given by the path query parameter using the uploaded file's name.
func (n *DeviceFiles) Upload(ctx context.Context, params mobilerun.DeviceFileUploadParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Files.Upload(ctx, o.id, params, opts...)
}

// DeviceKeyboard groups the keyboard methods of a Device.
type DeviceKeyboard struct {
	o *Device
}

func newDeviceKeyboard(o *Device) *DeviceKeyboard {
	n := &DeviceKeyboard{o: o}

	return n
}

func (n *DeviceKeyboard) String() string { return objectString("DeviceKeyboard", nil) }

// Clears the contents of the currently focused text input field.
func (n *DeviceKeyboard) Clear(ctx context.Context, params mobilerun.DeviceKeyboardClearParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Keyboard.Clear(ctx, o.id, params, opts...)
}

// DeviceLocation groups the location methods of a Device.
type DeviceLocation struct {
	o *Device
}

func newDeviceLocation(o *Device) *DeviceLocation {
	n := &DeviceLocation{o: o}

	return n
}

func (n *DeviceLocation) String() string { return objectString("DeviceLocation", nil) }

// Returns the device's current simulated GPS location as latitude and longitude.
// Devices without geo support return an unsupported-feature error.
func (n *DeviceLocation) Get(ctx context.Context, params mobilerun.DeviceLocationGetParams, opts ...option.RequestOption) (*shared.Location, error) {
	o := n.o
	return o.c.api.Devices.Location.Get(ctx, o.id, params, opts...)
}

// Clears any simulated GPS location and restores the device's default location
// behavior. Devices without geo support return an unsupported-feature error.
func (n *DeviceLocation) Reset(ctx context.Context, params mobilerun.DeviceLocationResetParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Location.Reset(ctx, o.id, params, opts...)
}

// Sets the device's simulated GPS location to the latitude and longitude in the
// request body. Devices without geo support return an unsupported-feature error.
func (n *DeviceLocation) Set(ctx context.Context, params mobilerun.DeviceLocationSetParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Location.Set(ctx, o.id, params, opts...)
}

// DevicePackages groups the packages methods of a Device.
type DevicePackages struct {
	o *Device
}

func newDevicePackages(o *Device) *DevicePackages {
	n := &DevicePackages{o: o}

	return n
}

func (n *DevicePackages) String() string { return objectString("DevicePackages", nil) }

// Returns the package names of apps installed on the device. System and protected
// packages are excluded unless the corresponding query parameters are set.
func (n *DevicePackages) List(ctx context.Context, params mobilerun.DevicePackageListParams, opts ...option.RequestOption) (*[]string, error) {
	o := n.o
	return o.c.api.Devices.Packages.List(ctx, o.id, params, opts...)
}

// DeviceProfile groups the profile methods of a Device.
type DeviceProfile struct {
	o *Device
}

func newDeviceProfile(o *Device) *DeviceProfile {
	n := &DeviceProfile{o: o}

	return n
}

func (n *DeviceProfile) String() string { return objectString("DeviceProfile", nil) }

// Apply a profile to a device
func (n *DeviceProfile) Update(ctx context.Context, params mobilerun.DeviceProfileUpdateParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Profile.Update(ctx, o.id, params, opts...)
}

// DeviceProxy groups the proxy methods of a Device.
type DeviceProxy struct {
	o *Device
}

func newDeviceProxy(o *Device) *DeviceProxy {
	n := &DeviceProxy{o: o}

	return n
}

func (n *DeviceProxy) String() string { return objectString("DeviceProxy", nil) }

// Routes the device's traffic through a SOCKS5 proxy supplied in the request body,
// replacing any existing connection. A smartIp option can be used to select an IP
// automatically; the legacy flat host/port/user/password fields remain supported.
func (n *DeviceProxy) Connect(ctx context.Context, params mobilerun.DeviceProxyConnectParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Proxy.Connect(ctx, o.id, params, opts...)
}

// Disconnects the device's active proxy connection and clears its stored proxy
// state. Returns successfully if no proxy is connected.
func (n *DeviceProxy) Disconnect(ctx context.Context, params mobilerun.DeviceProxyDisconnectParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Proxy.Disconnect(ctx, o.id, params, opts...)
}

// Returns the device's current proxy connection state, including whether a proxy
// is connected and its protocol and name.
func (n *DeviceProxy) Status(ctx context.Context, params mobilerun.DeviceProxyStatusParams, opts ...option.RequestOption) (*mobilerun.DeviceProxyStatusResponse, error) {
	o := n.o
	return o.c.api.Devices.Proxy.Status(ctx, o.id, params, opts...)
}

// DeviceTasks groups the tasks methods of a Device.
type DeviceTasks struct {
	o *Device
}

func newDeviceTasks(o *Device) *DeviceTasks {
	n := &DeviceTasks{o: o}

	return n
}

func (n *DeviceTasks) String() string { return objectString("DeviceTasks", nil) }

// Returns a paginated list of tasks that have run on the device, along with
// pagination metadata.
//
// It returns the first page of Task handles; ListAutoPaging iterates every page.
func (n *DeviceTasks) List(ctx context.Context, params mobilerun.DeviceTaskListParams, opts ...option.RequestOption) (*ObjectPage[*Task, *pagination.Page[mobilerun.DeviceTaskListResponse]], error) {
	o := n.o
	var wrap func(page *pagination.Page[mobilerun.DeviceTaskListResponse]) *ObjectPage[*Task, *pagination.Page[mobilerun.DeviceTaskListResponse]]
	wrap = func(page *pagination.Page[mobilerun.DeviceTaskListResponse]) *ObjectPage[*Task, *pagination.Page[mobilerun.DeviceTaskListResponse]] {
		result := &ObjectPage[*Task, *pagination.Page[mobilerun.DeviceTaskListResponse]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newTask(o.c, objectID(info, "taskId")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*Task, *pagination.Page[mobilerun.DeviceTaskListResponse]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := o.c.api.Devices.Tasks.List(ctx, o.id, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the Task handles of every page.
func (n *DeviceTasks) ListAutoPaging(ctx context.Context, params mobilerun.DeviceTaskListParams, opts ...option.RequestOption) *ObjectAutoPager[*Task] {
	page, err := n.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// DeviceClipboard groups the clipboard methods of a Device.
type DeviceClipboard struct {
	o *Device
}

func newDeviceClipboard(o *Device) *DeviceClipboard {
	n := &DeviceClipboard{o: o}

	return n
}

func (n *DeviceClipboard) String() string { return objectString("DeviceClipboard", nil) }

// Returns the current text content of the device's clipboard. Devices without
// clipboard support return an unsupported-feature error.
func (n *DeviceClipboard) Get(ctx context.Context, params mobilerun.DeviceClipboardGetParams, opts ...option.RequestOption) (*mobilerun.DeviceClipboardGetResponse, error) {
	o := n.o
	return o.c.api.Devices.Clipboard.Get(ctx, o.id, params, opts...)
}

// Replaces the device's clipboard content with the text in the request body; an
// empty text clears the clipboard. Devices without clipboard support return an
// unsupported-feature error.
func (n *DeviceClipboard) Set(ctx context.Context, params mobilerun.DeviceClipboardSetParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Clipboard.Set(ctx, o.id, params, opts...)
}

// DeviceTimezone groups the timezone methods of a Device.
type DeviceTimezone struct {
	o *Device
}

func newDeviceTimezone(o *Device) *DeviceTimezone {
	n := &DeviceTimezone{o: o}

	return n
}

func (n *DeviceTimezone) String() string { return objectString("DeviceTimezone", nil) }

// Returns the device's current timezone identifier. Devices that do not support
// timezone control return an unsupported-feature error.
func (n *DeviceTimezone) Get(ctx context.Context, params mobilerun.DeviceTimezoneGetParams, opts ...option.RequestOption) (*mobilerun.DeviceTimezoneGetResponse, error) {
	o := n.o
	return o.c.api.Devices.Timezone.Get(ctx, o.id, params, opts...)
}

// Sets the device timezone to the identifier in the request body. Devices that do
// not support timezone control return an unsupported-feature error.
func (n *DeviceTimezone) Set(ctx context.Context, params mobilerun.DeviceTimezoneSetParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Timezone.Set(ctx, o.id, params, opts...)
}

// DeviceLanguage groups the language methods of a Device.
type DeviceLanguage struct {
	o *Device
}

func newDeviceLanguage(o *Device) *DeviceLanguage {
	n := &DeviceLanguage{o: o}

	return n
}

func (n *DeviceLanguage) String() string { return objectString("DeviceLanguage", nil) }

// Returns the device's current language/locale as a BCP-47 locale string.
func (n *DeviceLanguage) Get(ctx context.Context, params mobilerun.DeviceLanguageGetParams, opts ...option.RequestOption) (*mobilerun.DeviceLanguageGetResponse, error) {
	o := n.o
	return o.c.api.Devices.Language.Get(ctx, o.id, params, opts...)
}

// Sets the device language/locale to the BCP-47 locale in the request body. An
// optional restart flag applies the change immediately by restarting the Android
// system framework (open apps are closed) instead of waiting for the next reboot;
// some devices ignore it.
func (n *DeviceLanguage) Set(ctx context.Context, params mobilerun.DeviceLanguageSetParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Language.Set(ctx, o.id, params, opts...)
}

// DeviceDeepLink groups the deep_link methods of a Device.
type DeviceDeepLink struct {
	o *Device
}

func newDeviceDeepLink(o *Device) *DeviceDeepLink {
	n := &DeviceDeepLink{o: o}

	return n
}

func (n *DeviceDeepLink) String() string { return objectString("DeviceDeepLink", nil) }

// Opens a deep link on the device. On Android the link is dispatched as an intent
// — packageName optionally pins it to a specific app and action overrides the
// default android.intent.action.VIEW. On iOS the URL is opened directly and the
// optional fields must be omitted. Protected packages are rejected.
func (n *DeviceDeepLink) ExecuteDeepLink(ctx context.Context, params mobilerun.DeviceDeepLinkExecuteDeepLinkParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.DeepLink.ExecuteDeepLink(ctx, o.id, params, opts...)
}

// DeviceBrowser groups the browser methods of a Device.
type DeviceBrowser struct {
	o *Device
}

func newDeviceBrowser(o *Device) *DeviceBrowser {
	n := &DeviceBrowser{o: o}

	return n
}

func (n *DeviceBrowser) String() string { return objectString("DeviceBrowser", nil) }

// Evaluates a JavaScript expression exactly once in the device's single, confirmed
// foreground Chrome tab via the Chrome DevTools Protocol and returns its
// JSON-serialized result. A responsive browser state without exactly one visible
// page returns DEVICE_NO_BROWSER_TARGET (400). Transient discovery, page-dial, or
// visibility-probe transport failures return CDP_TRANSPORT_UNAVAILABLE (503) with
// Retry-After and an executionOutcome; the service may start Chrome once and retry
// selection before Runtime.evaluate, but dial or probe recovery is skipped for a
// backgrounded Chrome, and a user script is never retried after its write or read.
// Devices without browser support return an unsupported-feature error.
func (n *DeviceBrowser) ExecuteScript(ctx context.Context, params mobilerun.DeviceBrowserExecuteScriptParams, opts ...option.RequestOption) (*mobilerun.DeviceBrowserExecuteScriptResponse, error) {
	o := n.o
	return o.c.api.Devices.Browser.ExecuteScript(ctx, o.id, params, opts...)
}

// DeviceKiosk groups the kiosk methods of a Device.
type DeviceKiosk struct {
	o *Device
}

func newDeviceKiosk(o *Device) *DeviceKiosk {
	n := &DeviceKiosk{o: o}

	return n
}

func (n *DeviceKiosk) String() string { return objectString("DeviceKiosk", nil) }

// Disables Android lock-task (kiosk) mode on the device, releasing it from the
// locked app.
func (n *DeviceKiosk) Disable(ctx context.Context, params mobilerun.DeviceKioskDisableParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Kiosk.Disable(ctx, o.id, params, opts...)
}

// Locks the device to the package named in the request body using Android
// lock-task (kiosk) mode, preventing the user from leaving the app.
func (n *DeviceKiosk) Enable(ctx context.Context, params mobilerun.DeviceKioskEnableParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Devices.Kiosk.Enable(ctx, o.id, params, opts...)
}

// DeviceNumber groups the number methods of a Device.
type DeviceNumber struct {
	o *Device
}

func newDeviceNumber(o *Device) *DeviceNumber {
	n := &DeviceNumber{o: o}

	return n
}

func (n *DeviceNumber) String() string { return objectString("DeviceNumber", nil) }

// Returns the cloud number bound to the device and the SIM line the device
// presents (carrier, number).
func (n *DeviceNumber) Get(ctx context.Context, params mobilerun.DeviceNumberGetParams, opts ...option.RequestOption) (*mobilerun.DeviceNumberGetResponse, error) {
	o := n.o
	return o.c.api.Devices.Number.Get(ctx, o.id, params, opts...)
}

// Binds one of your cloud numbers to the device. The device's SIM shows the number
// and every SMS it receives is delivered into the device's messages app. Available
// on devices whose capabilities report telephony. A number can be bound to one
// device, and a device holds one number; unbind first to move it.
func (n *DeviceNumber) Bind(ctx context.Context, params mobilerun.DeviceNumberBindParams, opts ...option.RequestOption) (*mobilerun.DeviceNumberBindResponse, error) {
	o := n.o
	return o.c.api.Devices.Number.Bind(ctx, o.id, params, opts...)
}

// Releases the number bound to the device. The device returns to its own default
// line; SMS to the number stop reaching the device.
func (n *DeviceNumber) Unbind(ctx context.Context, params mobilerun.DeviceNumberUnbindParams, opts ...option.RequestOption) (*mobilerun.DeviceNumberUnbindResponse, error) {
	o := n.o
	return o.c.api.Devices.Number.Unbind(ctx, o.id, params, opts...)
}

// SavedScreenshot is a handle to one SavedScreenshot.
//
// Field accessors read the last response (Info) without making a request.
type SavedScreenshot struct {
	c        *clientRef
	deviceID string
	id       string
	state    objectState
}

func newSavedScreenshot(c *clientRef, deviceID string, id string) *SavedScreenshot {
	o := &SavedScreenshot{c: c, deviceID: deviceID, id: id}

	return o
}

func (o *SavedScreenshot) withInfo(info any, root string) *SavedScreenshot {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *SavedScreenshot) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *SavedScreenshot) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *SavedScreenshot) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *SavedScreenshot) String() string {
	if o.Info() == nil {
		return objectString("SavedScreenshot", &o.id)
	}
	return objectString("SavedScreenshot", &o.id, "capturedAt", o.CapturedAt())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *SavedScreenshot) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// CapturedAt reads capturedAt. It reads only the last New/List snapshot without
// making a request; FromID handles have no data.
func (o *SavedScreenshot) CapturedAt() string { return resultString(o.state.field("capturedAt")) }

// DeviceID returns the bound Device ID.
func (o *SavedScreenshot) DeviceID() string { return o.deviceID }

// Delete a device screenshot
//
// The object layer sets deviceId; values for it in params are replaced.
func (o *SavedScreenshot) Delete(ctx context.Context, opts ...option.RequestOption) error {
	return o.c.api.Devices.Screenshots.Delete(ctx, o.id, mobilerun.DeviceScreenshotDeleteParams{DeviceID: o.deviceID}, opts...)
}

// Redirects to a presigned PNG URL valid for 15 minutes.
//
// Image returns the raw response bytes.
//
// The object layer sets deviceId; values for it in params are replaced.
func (o *SavedScreenshot) Image(ctx context.Context, opts ...option.RequestOption) ([]byte, error) {
	var body []byte
	var response *http.Response
	err := o.c.api.Devices.Screenshots.Image(ctx, o.id, mobilerun.DeviceScreenshotImageParams{DeviceID: o.deviceID}, append(opts, option.WithResponseBodyInto(&body), option.WithResponseInto(&response))...)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Recording is a handle to one Recording.
//
// Field accessors read the last response (Info) without making a request.
type Recording struct {
	c        *clientRef
	deviceID string
	id       string
	state    objectState
}

func newRecording(c *clientRef, deviceID string, id string) *Recording {
	o := &Recording{c: c, deviceID: deviceID, id: id}

	return o
}

func (o *Recording) withInfo(info any, root string) *Recording {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *Recording) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *Recording) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *Recording) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *Recording) String() string {
	if o.Info() == nil {
		return objectString("Recording", &o.id)
	}
	return objectString("Recording", &o.id, "name", o.Name(), "status", o.Status())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *Recording) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	_, err := o.c.api.Devices.Recordings.Stop(ctx, o.id, mobilerun.DeviceRecordingStopParams{DeviceID: o.deviceID}, opts...)
	return err
}

// Refresh fetches the object and updates Info.
func (o *Recording) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.DeviceRecordingStatusResponse, error) {
	res, err := o.c.api.Devices.Recordings.Status(ctx, o.id, mobilerun.DeviceRecordingStatusParams{DeviceID: o.deviceID}, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "")
	return res, nil
}

// Name reads name. It reads the last response without making a request.
func (o *Recording) Name() string { return resultString(o.state.field("name")) }

// Status reads status. It reads the last response without making a request.
func (o *Recording) Status() string { return resultString(o.state.field("status")) }

// Wait for the declared success condition; failure raises with the last response.
//
// Wait polls every 1000 ms for up to 300000 ms (WithPollInterval,
// WithWaitTimeout); failure states return a WaitFailureError and the deadline a
// WaitTimeoutError.
func (o *Recording) Wait(ctx context.Context, opts ...option.RequestOption) (*mobilerun.DeviceRecordingStatusResponse, error) {
	controls, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	poll := func(ctx context.Context, pollOpts []option.RequestOption) (*mobilerun.DeviceRecordingStatusResponse, bool, error) {
		requestOpts := append(append([]option.RequestOption{}, opts...), pollOpts...)
		res, err := o.c.api.Devices.Recordings.Status(ctx, o.id, mobilerun.DeviceRecordingStatusParams{DeviceID: o.deviceID}, requestOpts...)
		if err != nil {
			return nil, false, err
		}
		o.state.setInfo(res, "")
		done, err := enumState(res, "status", []string{"completed"}, []string{"failed", "expired"}, []string{"recording", "uploading"}, false)
		return res, done, err
	}
	return waitFor(ctx, durationOr(controls.waitTimeout, 300000*time.Millisecond), durationOr(controls.pollInterval, 1000*time.Millisecond), 30000*time.Millisecond, poll, retryPoll)
}

// DeviceID returns the bound Device ID.
func (o *Recording) DeviceID() string { return o.deviceID }

// Delete a device recording
//
// The object layer sets deviceId; values for it in params are replaced.
func (o *Recording) Delete(ctx context.Context, opts ...option.RequestOption) error {
	return o.c.api.Devices.Recordings.Delete(ctx, o.id, mobilerun.DeviceRecordingDeleteParams{DeviceID: o.deviceID}, opts...)
}

// Screenshots saved into the recording with
// `GET /devices/{deviceId}/screenshot?save=true`, oldest first, each with a
// presigned download URL.
//
// The object layer sets deviceId; values for it in params are replaced.
func (o *Recording) Screenshots(ctx context.Context, opts ...option.RequestOption) (*[]mobilerun.DeviceRecordingScreenshotsResponse, error) {
	return o.c.api.Devices.Recordings.Screenshots(ctx, o.id, mobilerun.DeviceRecordingScreenshotsParams{DeviceID: o.deviceID}, opts...)
}

// Stop a device recording
//
// The object layer sets deviceId; values for it in params are replaced.
func (o *Recording) Stop(ctx context.Context, opts ...option.RequestOption) (*mobilerun.DeviceRecordingStopResponse, error) {
	return o.c.api.Devices.Recordings.Stop(ctx, o.id, mobilerun.DeviceRecordingStopParams{DeviceID: o.deviceID}, opts...)
}

// Get a device recording video
//
// Video returns the raw response bytes.
//
// The object layer sets deviceId; values for it in params are replaced.
func (o *Recording) Video(ctx context.Context, opts ...option.RequestOption) ([]byte, error) {
	var body []byte
	var response *http.Response
	err := o.c.api.Devices.Recordings.Video(ctx, o.id, mobilerun.DeviceRecordingVideoParams{DeviceID: o.deviceID}, append(opts, option.WithResponseBodyInto(&body), option.WithResponseInto(&response))...)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Get a device recording trajectory
//
// Trajectory returns the raw response bytes.
//
// The object layer sets deviceId; values for it in params are replaced.
func (o *Recording) Trajectory(ctx context.Context, opts ...option.RequestOption) ([]byte, error) {
	var body []byte
	var response *http.Response
	err := o.c.api.Devices.Recordings.Trajectory(ctx, o.id, mobilerun.DeviceRecordingTrajectoryParams{DeviceID: o.deviceID}, append(opts, option.WithResponseBodyInto(&body), option.WithResponseInto(&response))...)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// TrafficSession is a handle to one TrafficSession.
//
// Field accessors read the last response (Info) without making a request.
type TrafficSession struct {
	c        *clientRef
	deviceID string
	id       string
	state    objectState
}

func newTrafficSession(c *clientRef, deviceID string, id string) *TrafficSession {
	o := &TrafficSession{c: c, deviceID: deviceID, id: id}

	return o
}

func (o *TrafficSession) withInfo(info any, root string) *TrafficSession {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *TrafficSession) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *TrafficSession) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *TrafficSession) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *TrafficSession) String() string {
	if o.Info() == nil {
		return objectString("TrafficSession", &o.id)
	}
	return objectString("TrafficSession", &o.id, "state", o.State())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *TrafficSession) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *TrafficSession) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.DeviceTrafficSessionGetResponse, error) {
	res, err := o.c.api.Devices.TrafficSessions.Get(ctx, o.id, mobilerun.DeviceTrafficSessionGetParams{DeviceID: o.deviceID}, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "")
	return res, nil
}

// State reads state. It reads the last response without making a request.
func (o *TrafficSession) State() string { return resultString(o.state.field("state")) }

// DeviceID returns the bound Device ID.
func (o *TrafficSession) DeviceID() string { return o.deviceID }

// Stop device traffic inspection
//
// The object layer sets deviceId; values for it in params are replaced.
func (o *TrafficSession) Delete(ctx context.Context, opts ...option.RequestOption) (*mobilerun.DeviceTrafficSessionDeleteResponse, error) {
	return o.c.api.Devices.TrafficSessions.Delete(ctx, o.id, mobilerun.DeviceTrafficSessionDeleteParams{DeviceID: o.deviceID}, opts...)
}

// A stored proxy configuration.
//
// Field accessors read the last response (Info) without making a request.
type ProxyConfiguration struct {
	c *clientRef

	id    string
	state objectState
}

func newProxyConfiguration(c *clientRef, id string) *ProxyConfiguration {
	o := &ProxyConfiguration{c: c, id: id}

	return o
}

func (o *ProxyConfiguration) withInfo(info any, root string) *ProxyConfiguration {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *ProxyConfiguration) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *ProxyConfiguration) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *ProxyConfiguration) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *ProxyConfiguration) String() string {
	if o.Info() == nil {
		return objectString("ProxyConfiguration", &o.id)
	}
	return objectString("ProxyConfiguration", &o.id, "name", o.Name(), "protocol", o.Protocol())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *ProxyConfiguration) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *ProxyConfiguration) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ProxyGetResponse, error) {
	res, err := o.c.api.Proxies.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// Name reads name. It reads the last response without making a request.
func (o *ProxyConfiguration) Name() string { return resultString(o.state.field("name")) }

// Protocol reads protocol. It reads the last response without making a request.
func (o *ProxyConfiguration) Protocol() string { return resultString(o.state.field("protocol")) }

// Replaces the proxy config identified by `proxyId` with the provided body. As
// with creation, the body is a `protocol`-discriminated union of socks5 or
// wireguard fields. Returns the updated config.
func (o *ProxyConfiguration) Update(ctx context.Context, params mobilerun.ProxyUpdateParams, opts ...option.RequestOption) (*mobilerun.ProxyUpdateResponse, error) {
	return o.c.api.Proxies.Update(ctx, o.id, params, opts...)
}

// Permanently deletes the proxy config identified by `proxyId` and returns the
// deleted config. Returns not found if no matching config exists.
func (o *ProxyConfiguration) Delete(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ProxyDeleteResponse, error) {
	return o.c.api.Proxies.Delete(ctx, o.id, opts...)
}

// ComputeFunction is a handle to one ComputeFunction.
//
// Field accessors read the last response (Info) without making a request.
type ComputeFunction struct {
	c *clientRef

	id       string
	state    objectState
	Versions *ComputeFunctionVersions
	Runs     *ComputeFunctionRuns
}

func newComputeFunction(c *clientRef, id string) *ComputeFunction {
	o := &ComputeFunction{c: c, id: id}
	o.Versions = newComputeFunctionVersions(o)
	o.Runs = newComputeFunctionRuns(o)
	return o
}

func (o *ComputeFunction) withInfo(info any, root string) *ComputeFunction {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *ComputeFunction) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *ComputeFunction) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *ComputeFunction) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *ComputeFunction) String() string {
	if o.Info() == nil {
		return objectString("ComputeFunction", &o.id)
	}
	return objectString("ComputeFunction", &o.id, "name", o.Name())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *ComputeFunction) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *ComputeFunction) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ComputeFunctionGetResponse, error) {
	res, err := o.c.api.Compute.Functions.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// Name reads name. It reads the last response without making a request.
func (o *ComputeFunction) Name() string { return resultString(o.state.field("name")) }

// Rename a function or update its description
func (o *ComputeFunction) Update(ctx context.Context, params mobilerun.ComputeFunctionUpdateParams, opts ...option.RequestOption) (*mobilerun.ComputeFunctionUpdateResponse, error) {
	return o.c.api.Compute.Functions.Update(ctx, o.id, params, opts...)
}

// Delete a function
func (o *ComputeFunction) Delete(ctx context.Context, opts ...option.RequestOption) error {
	return o.c.api.Compute.Functions.Delete(ctx, o.id, opts...)
}

// Read verified function source
func (o *ComputeFunction) Source(ctx context.Context, params mobilerun.ComputeFunctionSourceParams, opts ...option.RequestOption) (*mobilerun.ComputeFunctionSourceResponse, error) {
	return o.c.api.Compute.Functions.Source(ctx, o.id, params, opts...)
}

// ComputeFunctionVersions groups the versions methods of a ComputeFunction.
type ComputeFunctionVersions struct {
	o *ComputeFunction
}

func newComputeFunctionVersions(o *ComputeFunction) *ComputeFunctionVersions {
	n := &ComputeFunctionVersions{o: o}

	return n
}

func (n *ComputeFunctionVersions) String() string {
	return objectString("ComputeFunctionVersions", nil)
}

// Upload a new immutable version
func (n *ComputeFunctionVersions) New(ctx context.Context, params mobilerun.ComputeFunctionVersionNewParams, opts ...option.RequestOption) (*mobilerun.ComputeFunctionVersionNewResponse, error) {
	o := n.o
	return o.c.api.Compute.Functions.Versions.New(ctx, o.id, params, opts...)
}

// Get an exact version
//
// The object layer sets functionId; values for it in params are replaced.
func (n *ComputeFunctionVersions) Get(ctx context.Context, version int64, opts ...option.RequestOption) (*mobilerun.ComputeFunctionVersionGetResponse, error) {
	o := n.o
	return o.c.api.Compute.Functions.Versions.Get(ctx, version, mobilerun.ComputeFunctionVersionGetParams{FunctionID: o.id}, opts...)
}

// List active versions, newest first
func (n *ComputeFunctionVersions) List(ctx context.Context, params mobilerun.ComputeFunctionVersionListParams, opts ...option.RequestOption) (*mobilerun.ComputeFunctionVersionListResponse, error) {
	o := n.o
	return o.c.api.Compute.Functions.Versions.List(ctx, o.id, params, opts...)
}

// Delete a version except the last active version
//
// The object layer sets functionId; values for it in params are replaced.
func (n *ComputeFunctionVersions) Delete(ctx context.Context, version int64, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Compute.Functions.Versions.Delete(ctx, version, mobilerun.ComputeFunctionVersionDeleteParams{FunctionID: o.id}, opts...)
}

// ComputeFunctionRuns groups the runs methods of a ComputeFunction.
type ComputeFunctionRuns struct {
	o *ComputeFunction
}

func newComputeFunctionRuns(o *ComputeFunction) *ComputeFunctionRuns {
	n := &ComputeFunctionRuns{o: o}

	return n
}

func (n *ComputeFunctionRuns) String() string { return objectString("ComputeFunctionRuns", nil) }

// List the runs of a function, newest first
//
// It returns the first page of Run handles; ListAutoPaging iterates every page.
func (n *ComputeFunctionRuns) List(ctx context.Context, params mobilerun.ComputeFunctionRunListParams, opts ...option.RequestOption) (*ObjectPage[*Run, *mobilerun.ComputeFunctionRunListResponse], error) {
	o := n.o
	var load func(page int64) (*ObjectPage[*Run, *mobilerun.ComputeFunctionRunListResponse], error)
	load = func(page int64) (*ObjectPage[*Run, *mobilerun.ComputeFunctionRunListResponse], error) {
		pageOpts := opts
		if page > 0 {
			pageOpts = bindOptions(opts, binding{"page", "query", page})
		}
		res, err := o.c.api.Compute.Functions.Runs.List(ctx, o.id, params, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*Run, *mobilerun.ComputeFunctionRunListResponse]{Response: res}
		for _, info := range objectRows(res, "data.items") {
			result.Items = append(result.Items, newRun(o.c, objectID(info, "id")).withInfo(info, ""))
		}
		next, ok, err := nextPageNumber(res, "data.pagination.page", "data.pagination.hasNext", "data.pagination.pages", page)
		if err != nil {
			return nil, err
		}
		if ok {
			result.next = func() (*ObjectPage[*Run, *mobilerun.ComputeFunctionRunListResponse], error) { return load(next) }
		}
		return result, nil
	}
	return load(0)
}

// ListAutoPaging iterates the Run handles of every page.
func (n *ComputeFunctionRuns) ListAutoPaging(ctx context.Context, params mobilerun.ComputeFunctionRunListParams, opts ...option.RequestOption) *ObjectAutoPager[*Run] {
	page, err := n.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// Returns 200 with the finished run when it ends within `waitS`, else 202 with the
// run.
//
// New returns a Run handle; the handle does not own cleanup.
func (n *ComputeFunctionRuns) New(ctx context.Context, params mobilerun.ComputeFunctionRunNewParams, opts ...option.RequestOption) (*Run, error) {
	o := n.o
	res, err := o.c.api.Compute.Functions.Runs.New(ctx, o.id, params, opts...)
	if err != nil {
		return nil, err
	}
	return newRun(o.c, objectID(res, "data.id")).withInfo(res, "data"), nil
}

// Run is a handle to one Run.
//
// Field accessors read the last response (Info) without making a request.
type Run struct {
	c *clientRef

	id    string
	state objectState
}

func newRun(c *clientRef, id string) *Run {
	o := &Run{c: c, id: id}

	return o
}

func (o *Run) withInfo(info any, root string) *Run {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *Run) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *Run) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *Run) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *Run) String() string {
	if o.Info() == nil {
		return objectString("Run", &o.id)
	}
	return objectString("Run", &o.id, "status", o.Status(), "outcome", o.Outcome())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *Run) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *Run) Refresh(ctx context.Context, params mobilerun.ComputeRunGetParams, opts ...option.RequestOption) (*mobilerun.ComputeRunGetResponse, error) {
	res, err := o.c.api.Compute.Runs.Get(ctx, o.id, params, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// Status reads status. It reads the last response without making a request.
func (o *Run) Status() string { return resultString(o.state.field("status")) }

// Outcome reads outcome. It reads the last response without making a request.
func (o *Run) Outcome() string { return resultString(o.state.field("outcome")) }

// Wait for the declared success condition; failure raises with the last response.
//
// Wait polls every 1000 ms for up to 600000 ms (WithPollInterval,
// WithWaitTimeout); failure states return a WaitFailureError and the deadline a
// WaitTimeoutError.
func (o *Run) Wait(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ComputeRunGetResponse, error) {
	controls, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	poll := func(ctx context.Context, pollOpts []option.RequestOption) (*mobilerun.ComputeRunGetResponse, bool, error) {
		requestOpts := append(append([]option.RequestOption{}, opts...), pollOpts...)
		res, err := o.c.api.Compute.Runs.Get(ctx, o.id, mobilerun.ComputeRunGetParams{}, requestOpts...)
		if err != nil {
			return nil, false, err
		}
		o.state.setInfo(res, "data")
		done, err := enumState(res, "data.status", []string{"done"}, []string{"failed"}, []string{"running"}, false)
		return res, done, err
	}
	return waitFor(ctx, durationOr(controls.waitTimeout, 600000*time.Millisecond), durationOr(controls.pollInterval, 1000*time.Millisecond), 30000*time.Millisecond, poll, retryPoll)
}

// Wait for the declared success condition; failure raises with the last response.
//
// WaitLogs polls every 1000 ms for up to 60000 ms (WithPollInterval,
// WithWaitTimeout); failure states return a WaitFailureError and the deadline a
// WaitTimeoutError.
func (o *Run) WaitLogs(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ComputeRunGetResponse, error) {
	controls, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	poll := func(ctx context.Context, pollOpts []option.RequestOption) (*mobilerun.ComputeRunGetResponse, bool, error) {
		requestOpts := append(append([]option.RequestOption{}, opts...), pollOpts...)
		res, err := o.c.api.Compute.Runs.Get(ctx, o.id, mobilerun.ComputeRunGetParams{}, requestOpts...)
		if err != nil {
			return nil, false, err
		}
		o.state.setInfo(res, "data")
		done, err := enumState(res, "data.logs", []string{"available"}, []string{"none", "gone"}, []string{"pending"}, false)
		return res, done, err
	}
	return waitFor(ctx, durationOr(controls.waitTimeout, 60000*time.Millisecond), durationOr(controls.pollInterval, 1000*time.Millisecond), 30000*time.Millisecond, poll, retryPoll)
}

// Cancel a run
func (o *Run) Cancel(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ComputeRunCancelResponse, error) {
	return o.c.api.Compute.Runs.Cancel(ctx, o.id, opts...)
}

// Get a short-lived download link for the full log of a run
func (o *Run) Logs(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ComputeRunLogsResponse, error) {
	return o.c.api.Compute.Runs.Logs(ctx, o.id, opts...)
}

// Sandbox is a handle to one Sandbox.
//
// Field accessors read the last response (Info) without making a request.
type Sandbox struct {
	c *clientRef

	id       string
	state    objectState
	Commands *SandboxCommands
	Files    *SandboxFiles
}

func newSandbox(c *clientRef, id string) *Sandbox {
	o := &Sandbox{c: c, id: id}
	o.Commands = newSandboxCommands(o)
	o.Files = newSandboxFiles(o)
	return o
}

func (o *Sandbox) withInfo(info any, root string) *Sandbox {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *Sandbox) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *Sandbox) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *Sandbox) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *Sandbox) String() string {
	if o.Info() == nil {
		return objectString("Sandbox", &o.id)
	}
	return objectString("Sandbox", &o.id, "name", o.Name(), "status", o.Status())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *Sandbox) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	_, err := o.c.api.Compute.Sandboxes.Delete(ctx, o.id, opts...)
	return err
}

// Refresh fetches the object and updates Info.
func (o *Sandbox) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ComputeSandboxGetResponse, error) {
	res, err := o.c.api.Compute.Sandboxes.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// Name reads name. It reads the last response without making a request.
func (o *Sandbox) Name() string { return resultString(o.state.field("name")) }

// Status reads status. It reads the last response without making a request.
func (o *Sandbox) Status() string { return resultString(o.state.field("status")) }

// Stops the sandbox; a 202 response means shutdown is still in progress; once
// stopped, the sandbox stays listed with status stopped.
//
// Stop a sandbox
func (o *Sandbox) Delete(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ComputeSandboxDeleteResponse, error) {
	return o.c.api.Compute.Sandboxes.Delete(ctx, o.id, opts...)
}

// Execute a shell command
func (o *Sandbox) Exec(ctx context.Context, params mobilerun.ComputeSandboxExecParams, opts ...option.RequestOption) (*mobilerun.ComputeSandboxExecResponse, error) {
	return o.c.api.Compute.Sandboxes.Exec(ctx, o.id, params, opts...)
}

// SandboxCommands groups the commands methods of a Sandbox.
type SandboxCommands struct {
	o *Sandbox
}

func newSandboxCommands(o *Sandbox) *SandboxCommands {
	n := &SandboxCommands{o: o}

	return n
}

func (n *SandboxCommands) String() string { return objectString("SandboxCommands", nil) }

// Get a command result
//
// The object layer sets sandboxId; values for it in params are replaced.
func (n *SandboxCommands) Get(ctx context.Context, commandID string, params mobilerun.ComputeSandboxCommandGetParams, opts ...option.RequestOption) (*mobilerun.ComputeSandboxCommandGetResponse, error) {
	o := n.o
	params.SandboxID = o.id
	return o.c.api.Compute.Sandboxes.Commands.Get(ctx, commandID, params, opts...)
}

// List the newest commands
func (n *SandboxCommands) List(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ComputeSandboxCommandListResponse, error) {
	o := n.o
	return o.c.api.Compute.Sandboxes.Commands.List(ctx, o.id, opts...)
}

// SandboxFiles groups the files methods of a Sandbox.
type SandboxFiles struct {
	o *Sandbox
}

func newSandboxFiles(o *Sandbox) *SandboxFiles {
	n := &SandboxFiles{o: o}

	return n
}

func (n *SandboxFiles) String() string { return objectString("SandboxFiles", nil) }

// Upload a file
func (n *SandboxFiles) Upload(ctx context.Context, bodyValue io.Reader, params mobilerun.ComputeSandboxFileUploadParams, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Compute.Sandboxes.Files.Upload(ctx, o.id, bodyValue, params, opts...)
}

// Download a file
//
// Download returns the raw response bytes.
func (n *SandboxFiles) Download(ctx context.Context, params mobilerun.ComputeSandboxFileGetParams, opts ...option.RequestOption) ([]byte, error) {
	o := n.o
	var body []byte
	var response *http.Response
	_, err := o.c.api.Compute.Sandboxes.Files.Get(ctx, o.id, params, append(opts, option.WithResponseBodyInto(&body), option.WithResponseInto(&response))...)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// ConnectProxy is a handle to one ConnectProxy.
//
// Field accessors read the last response (Info) without making a request.
type ConnectProxy struct {
	c *clientRef

	id    string
	state objectState
}

func newConnectProxy(c *clientRef, id string) *ConnectProxy {
	o := &ConnectProxy{c: c, id: id}

	return o
}

func (o *ConnectProxy) withInfo(info any, root string) *ConnectProxy {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *ConnectProxy) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *ConnectProxy) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *ConnectProxy) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *ConnectProxy) String() string {
	if o.Info() == nil {
		return objectString("ConnectProxy", &o.id)
	}
	return objectString("ConnectProxy", &o.id, "name", o.Name(), "status", o.Status())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *ConnectProxy) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *ConnectProxy) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ConnectProxyGetResponse, error) {
	res, err := o.c.api.Connect.Proxies.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "")
	return res, nil
}

// Name reads name. It reads the last response without making a request.
func (o *ConnectProxy) Name() string { return resultString(o.state.field("name")) }

// Status reads status. It reads the last response without making a request.
func (o *ConnectProxy) Status() string { return resultString(o.state.field("status")) }

// Renames the proxy identified by the path ID.
func (o *ConnectProxy) Update(ctx context.Context, params mobilerun.ConnectProxyUpdateParams, opts ...option.RequestOption) (*mobilerun.ConnectProxyUpdateResponse, error) {
	return o.c.api.Connect.Proxies.Update(ctx, o.id, params, opts...)
}

// Deletes the proxy identified by the path ID and releases its provisioning.
// Returns 404 if no such proxy exists for the caller.
func (o *ConnectProxy) Cancel(ctx context.Context, opts ...option.RequestOption) error {
	return o.c.api.Connect.Proxies.Cancel(ctx, o.id, opts...)
}

// Refreshes a pending dedicated proxy's checkout on explicit payment intent.
// Already-paid proxies (provisioning, active, or cancelling) return their current
// status without a payment URL. An expired checkout requires a new purchase.
func (o *ConnectProxy) Checkout(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ConnectProxyCheckoutResponse, error) {
	return o.c.api.Connect.Proxies.Checkout(ctx, o.id, opts...)
}

// Checks right now that the proxy's upstream accepts a connection through its
// current session. If the upstream refuses the session and the proxy is a rotating
// (sticky-session) proxy, the proxy is moved onto a fresh session that has been
// checked to work; its credentials and country stay the same, its exit IP changes.
// Intended for callers that are about to depend on the proxy (e.g. while
// provisioning a phone). After `rotated`, the gateway uses the new session within
// about 30 seconds. A `refused` or `unreachable` result is not an error: the
// response describes the proxy's state.
func (o *ConnectProxy) EnsureHealthy(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ConnectProxyEnsureHealthyResponse, error) {
	return o.c.api.Connect.Proxies.EnsureHealthy(ctx, o.id, opts...)
}

// Returns the connection history recorded for this proxy, one item per connection
// (aggregated across the connection's lifetime). Supports filtering on every
// property plus ordering and pagination. Returns 503 when connection history is
// temporarily unavailable.
func (o *ConnectProxy) ListConnections(ctx context.Context, params mobilerun.ConnectProxyListConnectionsParams, opts ...option.RequestOption) (*pagination.Page[mobilerun.Connection], error) {
	return o.c.api.Connect.Proxies.ListConnections(ctx, o.id, params, append(append([]option.RequestOption{}, opts...), callerContext(ctx))...)
}

// Returns the connection history recorded for this proxy, one item per connection
// (aggregated across the connection's lifetime). Supports filtering on every
// property plus ordering and pagination. Returns 503 when connection history is
// temporarily unavailable.
func (o *ConnectProxy) ListConnectionsAutoPaging(ctx context.Context, params mobilerun.ConnectProxyListConnectionsParams, opts ...option.RequestOption) *pagination.PageAutoPager[mobilerun.Connection] {
	return o.c.api.Connect.Proxies.ListConnectionsAutoPaging(ctx, o.id, params, append(append([]option.RequestOption{}, opts...), callerContext(ctx))...)
}

// Returns the most recent cached network-latency measurement for the proxy,
// sampled periodically by connecting through the proxy to a fixed target.
// `latency` is null when no measurement is available yet (e.g. the proxy is not
// active, or it has not been sampled since coming online).
func (o *ConnectProxy) Ping(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ConnectProxyPingResponse, error) {
	return o.c.api.Connect.Proxies.Ping(ctx, o.id, opts...)
}

// ConnectUser is a handle to one ConnectUser.
//
// Field accessors read the last response (Info) without making a request.
type ConnectUser struct {
	c *clientRef

	id    string
	state objectState
}

func newConnectUser(c *clientRef, id string) *ConnectUser {
	o := &ConnectUser{c: c, id: id}

	return o
}

func (o *ConnectUser) withInfo(info any, root string) *ConnectUser {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *ConnectUser) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *ConnectUser) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *ConnectUser) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *ConnectUser) String() string {
	if o.Info() == nil {
		return objectString("ConnectUser", &o.id)
	}
	return objectString("ConnectUser", &o.id, "username", o.Username())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *ConnectUser) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *ConnectUser) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.ConnectUserGetResponse, error) {
	res, err := o.c.api.Connect.Users.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "")
	return res, nil
}

// Username reads username. It reads the last response without making a request.
func (o *ConnectUser) Username() string { return resultString(o.state.field("username")) }

// Rebind the user to a different proxy (or detach it by passing null).
func (o *ConnectUser) Update(ctx context.Context, params mobilerun.ConnectUserUpdateParams, opts ...option.RequestOption) (*mobilerun.ConnectUserUpdateResponse, error) {
	return o.c.api.Connect.Users.Update(ctx, o.id, params, opts...)
}

// Deletes the SOCKS5 user identified by the path ID, revoking its credentials and
// any proxy binding. Returns 404 if no such user exists for the caller.
func (o *ConnectUser) Delete(ctx context.Context, opts ...option.RequestOption) error {
	return o.c.api.Connect.Users.Delete(ctx, o.id, opts...)
}

// Returns the connection history recorded for this user, one item per connection
// (aggregated across the connection's lifetime). Supports filtering on every
// property plus ordering and pagination. Returns 503 when connection history is
// temporarily unavailable.
func (o *ConnectUser) ListConnections(ctx context.Context, params mobilerun.ConnectUserListConnectionsParams, opts ...option.RequestOption) (*pagination.Page[mobilerun.Connection], error) {
	return o.c.api.Connect.Users.ListConnections(ctx, o.id, params, append(append([]option.RequestOption{}, opts...), callerContext(ctx))...)
}

// Returns the connection history recorded for this user, one item per connection
// (aggregated across the connection's lifetime). Supports filtering on every
// property plus ordering and pagination. Returns 503 when connection history is
// temporarily unavailable.
func (o *ConnectUser) ListConnectionsAutoPaging(ctx context.Context, params mobilerun.ConnectUserListConnectionsParams, opts ...option.RequestOption) *pagination.PageAutoPager[mobilerun.Connection] {
	return o.c.api.Connect.Users.ListConnectionsAutoPaging(ctx, o.id, params, append(append([]option.RequestOption{}, opts...), callerContext(ctx))...)
}

// Task is a handle to one Task.
//
// Field accessors read the last response (Info) without making a request.
type Task struct {
	c *clientRef

	id          string
	state       objectState
	Screenshots *TaskScreenshots
	UiStates    *TaskUiStates
	Approvals   *TaskApprovals
}

func newTask(c *clientRef, id string) *Task {
	o := &Task{c: c, id: id}
	o.Screenshots = newTaskScreenshots(o)
	o.UiStates = newTaskUiStates(o)
	o.Approvals = newTaskApprovals(o)
	return o
}

func (o *Task) withInfo(info any, root string) *Task {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *Task) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *Task) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *Task) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *Task) String() string {
	if o.Info() == nil {
		return objectString("Task", &o.id)
	}
	return objectString("Task", &o.id, "status", o.Status())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *Task) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *Task) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.TaskResponse, error) {
	res, err := o.c.api.Tasks.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "task")
	return res, nil
}

// Status reads status. It is not available on handles from `device.tasks.list()`
// until Refresh.
func (o *Task) Status() string { return resultString(o.state.field("status")) }

// Wait for the declared success condition; failure raises with the last response.
//
// Wait polls every 1000 ms for up to 600000 ms (WithPollInterval,
// WithWaitTimeout); failure states return a WaitFailureError and the deadline a
// WaitTimeoutError.
func (o *Task) Wait(ctx context.Context, opts ...option.RequestOption) (*mobilerun.TaskGetStatusResponse, error) {
	controls, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	poll := func(ctx context.Context, pollOpts []option.RequestOption) (*mobilerun.TaskGetStatusResponse, bool, error) {
		requestOpts := append(append([]option.RequestOption{}, opts...), pollOpts...)
		res, err := o.c.api.Tasks.GetStatus(ctx, o.id, requestOpts...)
		if err != nil {
			return nil, false, err
		}
		o.state.setInfo(res, "")
		done, err := enumState(res, "status", []string{"completed"}, []string{"failed", "cancelled"}, []string{"prepared", "queued", "created", "running", "cancelling"}, false)
		return res, done, err
	}
	return waitFor(ctx, durationOr(controls.waitTimeout, 600000*time.Millisecond), durationOr(controls.pollInterval, 1000*time.Millisecond), 30000*time.Millisecond, poll, retryPoll)
}

// WaitOutput runs Wait and decodes the output at output into out (a pointer, as
// for json.Unmarshal). A missing or undecodable output returns an
// OutputValidationError.
func (o *Task) WaitOutput(ctx context.Context, out any, opts ...option.RequestOption) (*mobilerun.TaskGetStatusResponse, error) {
	res, err := o.Wait(ctx, opts...)
	if err != nil {
		return nil, err
	}
	if err := decodeOutput(o.id, res, "output", out); err != nil {
		return res, err
	}
	return res, nil
}

// Attach to a running task and receive its events as an SSE stream.
func (o *Task) Attach(ctx context.Context, opts ...option.RequestOption) error {
	return o.c.api.Tasks.Attach(ctx, o.id, opts...)
}

// Get the status of a task.
func (o *Task) GetStatus(ctx context.Context, opts ...option.RequestOption) (*mobilerun.TaskGetStatusResponse, error) {
	return o.c.api.Tasks.GetStatus(ctx, o.id, opts...)
}

// Get the trajectory of a task.
func (o *Task) GetTrajectory(ctx context.Context, opts ...option.RequestOption) (*mobilerun.TaskGetTrajectoryResponse, error) {
	return o.c.api.Tasks.GetTrajectory(ctx, o.id, opts...)
}

// Send a message to a running agent task. The message ID is delivered via SSE
// (UserMessageEvent with action=queued).
func (o *Task) SendMessage(ctx context.Context, params mobilerun.TaskSendMessageParams, opts ...option.RequestOption) (*mobilerun.TaskSendMessageResponse, error) {
	return o.c.api.Tasks.SendMessage(ctx, o.id, params, opts...)
}

// Cancel a running task. Returns an error if the task is already in a terminal
// state.
func (o *Task) Stop(ctx context.Context, opts ...option.RequestOption) (*mobilerun.TaskStopResponse, error) {
	return o.c.api.Tasks.Stop(ctx, o.id, opts...)
}

// TaskScreenshots groups the screenshots methods of a Task.
type TaskScreenshots struct {
	o *Task
}

func newTaskScreenshots(o *Task) *TaskScreenshots {
	n := &TaskScreenshots{o: o}

	return n
}

func (n *TaskScreenshots) String() string { return objectString("TaskScreenshots", nil) }

// Get a specific screenshot by index.
//
// The object layer sets task_id; values for it in params are replaced.
func (n *TaskScreenshots) Get(ctx context.Context, index int64, opts ...option.RequestOption) (*mobilerun.TaskScreenshotGetResponse, error) {
	o := n.o
	return o.c.api.Tasks.Screenshots.Get(ctx, index, mobilerun.TaskScreenshotGetParams{TaskID: o.id}, opts...)
}

// List all screenshot URLs for a task.
func (n *TaskScreenshots) List(ctx context.Context, opts ...option.RequestOption) (*mobilerun.TaskScreenshotListResponse, error) {
	o := n.o
	return o.c.api.Tasks.Screenshots.List(ctx, o.id, opts...)
}

// TaskUiStates groups the ui_states methods of a Task.
type TaskUiStates struct {
	o *Task
}

func newTaskUiStates(o *Task) *TaskUiStates {
	n := &TaskUiStates{o: o}

	return n
}

func (n *TaskUiStates) String() string { return objectString("TaskUiStates", nil) }

// Get a specific UI state by index.
//
// The object layer sets task_id; values for it in params are replaced.
func (n *TaskUiStates) Get(ctx context.Context, index int64, opts ...option.RequestOption) (*mobilerun.TaskUiStateGetResponse, error) {
	o := n.o
	return o.c.api.Tasks.UiStates.Get(ctx, index, mobilerun.TaskUiStateGetParams{TaskID: o.id}, opts...)
}

// List all UI state URLs for a task.
func (n *TaskUiStates) List(ctx context.Context, opts ...option.RequestOption) (*mobilerun.TaskUiStateListResponse, error) {
	o := n.o
	return o.c.api.Tasks.UiStates.List(ctx, o.id, opts...)
}

// TaskApprovals groups the approvals methods of a Task.
type TaskApprovals struct {
	o *Task
}

func newTaskApprovals(o *Task) *TaskApprovals {
	n := &TaskApprovals{o: o}

	return n
}

func (n *TaskApprovals) String() string { return objectString("TaskApprovals", nil) }

// Approve or deny the approval a running task is waiting for (see
// `pendingApproval` on the task). Deny ends the task as failed.
//
// The object layer sets task_id; values for it in params are replaced.
func (n *TaskApprovals) Decide(ctx context.Context, approvalID string, params mobilerun.TaskApprovalDecideParams, opts ...option.RequestOption) (*mobilerun.TaskApprovalDecideResponse, error) {
	o := n.o
	params.TaskID = o.id
	return o.c.api.Tasks.Approvals.Decide(ctx, approvalID, params, opts...)
}

// Trigger is a handle to one Trigger.
//
// Field accessors read the last response (Info) without making a request.
type Trigger struct {
	c *clientRef

	id    string
	state objectState
}

func newTrigger(c *clientRef, id string) *Trigger {
	o := &Trigger{c: c, id: id}

	return o
}

func (o *Trigger) withInfo(info any, root string) *Trigger {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *Trigger) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *Trigger) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *Trigger) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *Trigger) String() string {
	if o.Info() == nil {
		return objectString("Trigger", &o.id)
	}
	return objectString("Trigger", &o.id, "name", o.Name(), "activation", o.Activation())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *Trigger) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *Trigger) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WorkflowTriggerGetResponse, error) {
	res, err := o.c.api.Workflows.Triggers.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// Name reads name. It reads the last response without making a request.
func (o *Trigger) Name() string { return resultString(o.state.field("name")) }

// Activation reads activation. It reads the last response without making a
// request.
func (o *Trigger) Activation() string { return resultString(o.state.field("activation")) }

// Partially update a trigger; all fields are optional. When `activation` is
// changed, the type-specific field rules are re-validated. Schedule triggers
// accept stored `payload` values (at most 1 MiB); null clears them. The resulting
// payload must match customPayloadSchema (422 on mismatch, including schema
// changes that invalidate stored values). Absent or null values are validated as
// an empty object when a schema is configured. Changes apply from the next run.
// Returns 404 if the trigger does not exist.
func (o *Trigger) Update(ctx context.Context, params mobilerun.WorkflowTriggerUpdateParams, opts ...option.RequestOption) (*mobilerun.WorkflowTriggerUpdateResponse, error) {
	return o.c.api.Workflows.Triggers.Update(ctx, o.id, params, opts...)
}

// Delete a trigger by its ID. Returns 404 if no trigger matches.
func (o *Trigger) Delete(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WorkflowTriggerDeleteResponse, error) {
	return o.c.api.Workflows.Triggers.Delete(ctx, o.id, opts...)
}

// Invoke a custom trigger directly with an arbitrary JSON payload.
//
// Fan-out: a trigger may be referenced by multiple flows (workflows). Firing it
// enqueues one execution per enabled, non-deleted flow attached to this trigger,
// each receiving the same payload. The `enqueuedCount` in the response reports how
// many were enqueued (0 if no flows are attached, or if all matching flows are
// gated by a cooldown).
//
// Payload validation:
//
//   - If the trigger has a `customPayloadSchema`, the payload is validated against
//     it (JSON Schema).
//   - If no schema is configured, the payload only needs to be a JSON object — any
//     keys and values are accepted.
//
// Only triggers with `activation = "custom"` can be fired through this endpoint;
// event and schedule triggers return 409.
//
// Supports an optional `Idempotency-Key` header (1-255 printable ASCII
// characters), which maps to the derived `invocationId` used for fan-out
// deduplication. The payload is bound to the key on the first accepted fire,
// before fan-out. A repeat with the same key and an identical payload returns 202
// with `Idempotent-Replayed: true` (`deduplicated: true` when flows were skipped;
// `enqueuedCount` counts only executions enqueued by that call); a changed payload
// under the same key returns 422 `idempotency_key_reused`.
func (o *Trigger) Fire(ctx context.Context, params mobilerun.WorkflowTriggerFireParams, opts ...option.RequestOption) (*mobilerun.WorkflowTriggerFireResponse, error) {
	return o.c.api.Workflows.Triggers.Fire(ctx, o.id, params, opts...)
}

// CatalogEntry is a handle to one CatalogEntry.
//
// Field accessors read the last response (Info) without making a request.
type CatalogEntry struct {
	c *clientRef

	id    string
	state objectState
}

func newCatalogEntry(c *clientRef, id string) *CatalogEntry {
	o := &CatalogEntry{c: c, id: id}

	return o
}

func (o *CatalogEntry) withInfo(info any, root string) *CatalogEntry {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *CatalogEntry) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *CatalogEntry) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *CatalogEntry) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *CatalogEntry) String() string {
	if o.Info() == nil {
		return objectString("CatalogEntry", &o.id)
	}
	return objectString("CatalogEntry", &o.id, "name", o.Name(), "service", o.Service())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *CatalogEntry) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *CatalogEntry) Refresh(ctx context.Context, params mobilerun.WorkflowActionCatalogGetParams, opts ...option.RequestOption) (*mobilerun.WorkflowActionCatalogGetResponse, error) {
	res, err := o.c.api.Workflows.ActionCatalog.Get(ctx, o.id, params, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// Name reads name. It reads the last response without making a request.
func (o *CatalogEntry) Name() string { return resultString(o.state.field("name")) }

// Service reads service. It reads the last response without making a request.
func (o *CatalogEntry) Service() string { return resultString(o.state.field("service")) }

// WorkflowAction is a handle to one WorkflowAction.
//
// Field accessors read the last response (Info) without making a request.
type WorkflowAction struct {
	c *clientRef

	id    string
	state objectState
}

func newWorkflowAction(c *clientRef, id string) *WorkflowAction {
	o := &WorkflowAction{c: c, id: id}

	return o
}

func (o *WorkflowAction) withInfo(info any, root string) *WorkflowAction {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *WorkflowAction) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *WorkflowAction) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *WorkflowAction) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *WorkflowAction) String() string {
	if o.Info() == nil {
		return objectString("WorkflowAction", &o.id)
	}
	return objectString("WorkflowAction", &o.id, "name", o.Name(), "service", o.Service())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *WorkflowAction) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *WorkflowAction) Refresh(ctx context.Context, params mobilerun.WorkflowActionGetParams, opts ...option.RequestOption) (*mobilerun.WorkflowActionGetResponse, error) {
	res, err := o.c.api.Workflows.Actions.Get(ctx, o.id, params, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// Name reads name. It reads the last response without making a request.
func (o *WorkflowAction) Name() string { return resultString(o.state.field("name")) }

// Service reads service. It reads the last response without making a request.
func (o *WorkflowAction) Service() string { return resultString(o.state.field("service")) }

// Partially update an action's name, description, or params; all fields are
// optional. Returns 404 if the action does not exist.
func (o *WorkflowAction) Update(ctx context.Context, params mobilerun.WorkflowActionUpdateParams, opts ...option.RequestOption) (*mobilerun.WorkflowActionUpdateResponse, error) {
	return o.c.api.Workflows.Actions.Update(ctx, o.id, params, opts...)
}

// Delete an action by its ID. Returns 404 if no action matches.
func (o *WorkflowAction) Delete(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WorkflowActionDeleteResponse, error) {
	return o.c.api.Workflows.Actions.Delete(ctx, o.id, opts...)
}

// Flow is a handle to one Flow.
//
// Field accessors read the last response (Info) without making a request.
type Flow struct {
	c *clientRef

	id            string
	state         objectState
	ScheduledRuns *ScheduledRunManager
	Actions       *FlowActions
	State         *FlowState
}

func newFlow(c *clientRef, id string) *Flow {
	o := &Flow{c: c, id: id}
	o.ScheduledRuns = newScheduledRunManager(o.c, o.id)
	o.Actions = newFlowActions(o)
	o.State = newFlowState(o)
	return o
}

func (o *Flow) withInfo(info any, root string) *Flow {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *Flow) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *Flow) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *Flow) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *Flow) String() string {
	if o.Info() == nil {
		return objectString("Flow", &o.id)
	}
	return objectString("Flow", &o.id, "name", o.Name())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *Flow) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *Flow) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WorkflowFlowGetResponse, error) {
	res, err := o.c.api.Workflows.Flows.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// Name reads name. It reads the last response without making a request.
func (o *Flow) Name() string { return resultString(o.state.field("name")) }

// Partially update a flow's settings — name, trigger binding, enabled state,
// target devices, cooldown, or notifications; all fields are optional. The owner
// can enable a routine directly with `enabled: true` or
// `lifecycleStatus: "enabled"`, without verification; at least one step is
// required. For the verified path used by agent deploys, use POST
// /flows/{flowId}/verify then POST /flows/{flowId}/activate instead. Actions are
// managed through the flow-actions endpoints, not here. Returns 404 if the flow
// does not exist.
func (o *Flow) Update(ctx context.Context, params mobilerun.WorkflowFlowUpdateParams, opts ...option.RequestOption) (*mobilerun.WorkflowFlowUpdateResponse, error) {
	return o.c.api.Workflows.Flows.Update(ctx, o.id, params, opts...)
}

// Terminally archive a flow by its ID. Archived flows cannot be restored and are
// hidden from customer reads. Repeating the request is idempotent for the owner.
// Files saved by its runs are deleted with it.
func (o *Flow) Delete(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WorkflowFlowDeleteResponse, error) {
	return o.c.api.Workflows.Flows.Delete(ctx, o.id, opts...)
}

// The verified activation path used by agent deploys: first run POST
// /flows/{flowId}/verify, then activate the disabled flow only when the supplied
// verification execution succeeded and the flow graph has not changed since
// verification. Repeating activation is idempotent for an already-enabled flow.
// Alternatively, the owner can enable a routine with at least one step directly
// through PATCH /flows/{flowId} using `enabled: true` or
// `lifecycleStatus: "enabled"`, without verification.
func (o *Flow) Activate(ctx context.Context, params mobilerun.WorkflowFlowActivateParams, opts ...option.RequestOption) (*mobilerun.WorkflowFlowActivateResponse, error) {
	return o.c.api.Workflows.Flows.Activate(ctx, o.id, params, opts...)
}

// Return the recording/delivery readiness for a flow (recording mode, whether
// delivery can include the recording, whether a files.upload action exists, the
// first direct OneDrive/Google Drive upload step if any) plus, per destination,
// whether the flow owner has an active connection. Returns 404 if the flow does
// not exist.
func (o *Flow) DeliveryOptions(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WorkflowFlowDeliveryOptionsResponse, error) {
	return o.c.api.Workflows.Flows.DeliveryOptions(ctx, o.id, opts...)
}

// Simulate this flow firing without storing events, enqueuing jobs, or consuming
// cooldown/rate-limit slots.
//
// Works for every trigger activation type:
//
//   - `event`: validates the payload against the event catalog schema and evaluates
//     the trigger conditions.
//   - `custom`: validates the payload against the custom payload schema (conditions
//     do not apply).
//   - `schedule`: validates the stored trigger payload and its step references, and
//     reports the next fire time; request payload overrides are ignored.
//
// The response reports `wouldFire` — whether the flow would actually run right now
// — alongside the gates that decide it (enabled, device attached, blocked,
// cooldown). `rateLimited` is informational and is not folded into `wouldFire`.
func (o *Flow) DryRun(ctx context.Context, params mobilerun.WorkflowFlowDryRunParams, opts ...option.RequestOption) (*mobilerun.WorkflowFlowDryRunResponse, error) {
	return o.c.api.Workflows.Flows.DryRun(ctx, o.id, params, opts...)
}

// List self-healing repair episodes
func (o *Flow) ListRepairs(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WorkflowFlowListRepairsResponse, error) {
	return o.c.api.Workflows.Flows.ListRepairs(ctx, o.id, opts...)
}

// Body: {document}. Preserves enabled state. Omitted settings preserves all stored
// settings; present settings replaces them, with absent fields defaulted. Agent
// routines require a new version deployed through the agent. Omitted description,
// devices and trigger options clear their values. Shared triggers return
// trigger_shared and agent-managed routines return routine_agent_managed (409);
// shared action definitions are never modified. To enable a disabled routine with
// at least one step, the owner can PATCH /flows/{flowId} with enabled: true or
// lifecycleStatus: "enabled", without verification. For the verified path used by
// agent deploys, call POST /flows/{flowId}/verify then POST
// /flows/{flowId}/activate; activation requires a successful verification of the
// current graph.
func (o *Flow) ReplaceDocument(ctx context.Context, params mobilerun.WorkflowFlowReplaceDocumentParams, opts ...option.RequestOption) (*mobilerun.WorkflowFlowReplaceDocumentResponse, error) {
	return o.c.api.Workflows.Flows.ReplaceDocument(ctx, o.id, params, opts...)
}

// Returns the canonical routine document, including delivery, recording, health,
// cooldown, owner wake and scheduled inputs. Unsupported stored behavior returns
// document_unsupported with JSON Pointer paths.
func (o *Flow) GetDocument(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WorkflowFlowGetDocumentResponse, error) {
	return o.c.api.Workflows.Flows.GetDocument(ctx, o.id, opts...)
}

// Get a self-healing repair episode
//
// The object layer sets flowId; values for it in params are replaced.
func (o *Flow) GetRepair(ctx context.Context, episodeID string, opts ...option.RequestOption) (*mobilerun.WorkflowFlowGetRepairResponse, error) {
	return o.c.api.Workflows.Flows.GetRepair(ctx, episodeID, mobilerun.WorkflowFlowGetRepairParams{FlowID: o.id}, opts...)
}

// Immediately enqueue one live execution on every device bound to the flow,
// regardless of whether its trigger is event-based, custom, or scheduled. A
// routine without steps is refused with 409 `flow_has_no_steps`.
//
// Supports an optional `Idempotency-Key` header (1-255 printable ASCII
// characters). The run identity is derived from the key, so a retry never starts a
// second run. Replays with the same key and an identical body return the original
// status and the original `executionIds`; a changed body under the same key
// returns 422 `idempotency_key_reused`.
func (o *Flow) Run(ctx context.Context, params mobilerun.WorkflowFlowRunParams, opts ...option.RequestOption) (*mobilerun.WorkflowFlowRunResponse, error) {
	return o.c.api.Workflows.Flows.Run(ctx, o.id, params, opts...)
}

// Clear a flow's blocked status after fixing the underlying issue. Idempotent —
// safe to call on already-healthy flows.
func (o *Flow) Unblock(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WorkflowFlowUnblockResponse, error) {
	return o.c.api.Workflows.Flows.Unblock(ctx, o.id, opts...)
}

// Create a copy of an existing flow, including its actions and settings. The
// optional body can override the new flow's `name` and target `deviceIds`. Returns
// 404 if the source flow does not exist.
//
// Supports an optional `Idempotency-Key` header (1-255 printable ASCII
// characters). Replays with the same key and an identical body return the original
// 201; a changed body under the same key returns 422 `idempotency_key_reused`.
//
// Clone returns a Flow handle; the handle does not own cleanup.
func (o *Flow) Clone(ctx context.Context, params mobilerun.WorkflowFlowCloneParams, opts ...option.RequestOption) (*Flow, error) {
	res, err := o.c.api.Workflows.Flows.Clone(ctx, o.id, params, opts...)
	if err != nil {
		return nil, err
	}
	return newFlow(o.c, objectID(res, "data.id")).withInfo(res, "data"), nil
}

// Performs a real verification run on a device.
//
// Run a bound, deduplicated **verification** of this flow on exactly one device as
// a real run.
//
// Unlike a normal trigger firing, verification:
//
//   - runs a flow that is `enabled` OR `disabled` (an archived flow returns 404);
//   - refuses routines without steps with 409 `flow_has_no_steps`;
//   - targets one device (`deviceId`, or the single bound device);
//   - is idempotent per `invocationId`: a repeat returns the existing run
//     (`deduplicated: true`, HTTP 200) instead of enqueuing another (HTTP 202);
//   - never mutates flow health, `lastTriggeredAt`, or the repair loop.
//
// Per-activation payload gate:
//
//   - `custom`: validates the payload against the custom payload schema (422 on
//     failure).
//   - `event`: validates the payload against the event catalog schema and evaluates
//     the trigger conditions; a failing gate returns 409 `conditions_not_met` with
//     the dry-run report.
//   - `schedule`: uses the stored trigger payload as event.payload; request payload
//     overrides are ignored.
//
// Verify returns a Execution handle; the handle does not own cleanup.
func (o *Flow) Verify(ctx context.Context, params mobilerun.WorkflowFlowVerifyParams, opts ...option.RequestOption) (*Execution, error) {
	res, err := o.c.api.Workflows.Flows.Verify(ctx, o.id, params, opts...)
	if err != nil {
		return nil, err
	}
	return newExecution(o.c, objectID(res, "data.executionId")).withInfo(res, "data"), nil
}

// FlowActions groups the actions methods of a Flow.
type FlowActions struct {
	o *Flow
}

func newFlowActions(o *Flow) *FlowActions {
	n := &FlowActions{o: o}

	return n
}

func (n *FlowActions) String() string { return objectString("FlowActions", nil) }

// Return the ordered list of actions attached to a flow, including any nested
// child actions. Returns 404 if the flow does not exist.
func (n *FlowActions) List(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WorkflowFlowActionListResponse, error) {
	o := n.o
	return o.c.api.Workflows.Flows.Actions.List(ctx, o.id, opts...)
}

// Append a single action to a flow at the given `position`, optionally nesting it
// under a `parentFlowActionId` or supplying its own `children`. Supports a
// `nameOverride`, param `overrides`, and `continueOnError`. Returns 404 if the
// flow does not exist.
//
// Supports an optional `Idempotency-Key` header (1-255 printable ASCII
// characters). Replays with the same key and an identical body return the original
// 201; a changed body under the same key returns 422 `idempotency_key_reused`.
func (n *FlowActions) Add(ctx context.Context, params mobilerun.WorkflowFlowActionAddParams, opts ...option.RequestOption) (*mobilerun.WorkflowFlowActionAddResponse, error) {
	o := n.o
	return o.c.api.Workflows.Flows.Actions.Add(ctx, o.id, params, opts...)
}

// Remove a single action from a flow by its `flowActionId`. Returns 404 if the
// flow or flow action does not exist.
//
// The object layer sets flowId; values for it in params are replaced.
func (n *FlowActions) Remove(ctx context.Context, flowActionID string, opts ...option.RequestOption) (*mobilerun.WorkflowFlowActionRemoveResponse, error) {
	o := n.o
	return o.c.api.Workflows.Flows.Actions.Remove(ctx, flowActionID, mobilerun.WorkflowFlowActionRemoveParams{FlowID: o.id}, opts...)
}

// Replace a flow's entire action list with the supplied set (at least one
// required). Each action references an `actionId` and a unique `position`, and may
// include nested `children`, a `nameOverride`, param `overrides`, and a
// `continueOnError` flag. Returns 404 if the flow does not exist.
func (n *FlowActions) Replace(ctx context.Context, params mobilerun.WorkflowFlowActionReplaceParams, opts ...option.RequestOption) (*mobilerun.WorkflowFlowActionReplaceResponse, error) {
	o := n.o
	return o.c.api.Workflows.Flows.Actions.Replace(ctx, o.id, params, opts...)
}

// FlowState groups the state methods of a Flow.
type FlowState struct {
	o *Flow
}

func newFlowState(o *Flow) *FlowState {
	n := &FlowState{o: o}

	return n
}

func (n *FlowState) String() string { return objectString("FlowState", nil) }

// List the keys a routine has saved with `state.save`, with the metadata of each
// key's newest version.
func (n *FlowState) List(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WorkflowFlowStateListResponse, error) {
	o := n.o
	return o.c.api.Workflows.Flows.State.List(ctx, o.id, opts...)
}

// Delete every version of one state key. The next `state.load` of the key reports
// `found: false`. Deleting an unknown key succeeds.
//
// The object layer sets flowId; values for it in params are replaced.
func (n *FlowState) Delete(ctx context.Context, key string, opts ...option.RequestOption) (*mobilerun.WorkflowFlowStateDeleteResponse, error) {
	o := n.o
	return o.c.api.Workflows.Flows.State.Delete(ctx, key, mobilerun.WorkflowFlowStateDeleteParams{FlowID: o.id}, opts...)
}

// Return the retained versions of one state key, newest first, including each
// saved value. An unknown key returns an empty list.
//
// The object layer sets flowId; values for it in params are replaced.
func (n *FlowState) ListVersions(ctx context.Context, key string, opts ...option.RequestOption) (*mobilerun.WorkflowFlowStateListVersionsResponse, error) {
	o := n.o
	return o.c.api.Workflows.Flows.State.ListVersions(ctx, key, mobilerun.WorkflowFlowStateListVersionsParams{FlowID: o.id}, opts...)
}

// ScheduledRun is a handle to one ScheduledRun.
//
// Field accessors read the last response (Info) without making a request.
type ScheduledRun struct {
	c      *clientRef
	flowID string
	id     string
	state  objectState
}

func newScheduledRun(c *clientRef, flowID string, id string) *ScheduledRun {
	o := &ScheduledRun{c: c, flowID: flowID, id: id}

	return o
}

func (o *ScheduledRun) withInfo(info any, root string) *ScheduledRun {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *ScheduledRun) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *ScheduledRun) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *ScheduledRun) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *ScheduledRun) String() string {
	if o.Info() == nil {
		return objectString("ScheduledRun", &o.id)
	}
	return objectString("ScheduledRun", &o.id, "status", o.Status(), "runAt", o.RunAt())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *ScheduledRun) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Status reads status. It reads only the last New/List snapshot without making a
// request; FromID handles have no data.
func (o *ScheduledRun) Status() string { return resultString(o.state.field("status")) }

// RunAt reads runAt. It reads only the last New/List snapshot without making a
// request; FromID handles have no data.
func (o *ScheduledRun) RunAt() string { return resultString(o.state.field("runAt")) }

// FlowID returns the bound Flow ID.
func (o *ScheduledRun) FlowID() string { return o.flowID }

// Cancel a planned flow run
//
// The object layer sets flowId; values for it in params are replaced.
func (o *ScheduledRun) Cancel(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WorkflowFlowScheduledRunCancelResponse, error) {
	return o.c.api.Workflows.Flows.ScheduledRuns.Cancel(ctx, o.id, mobilerun.WorkflowFlowScheduledRunCancelParams{FlowID: o.flowID}, opts...)
}

// Execution is a handle to one Execution.
//
// Field accessors read the last response (Info) without making a request.
type Execution struct {
	c *clientRef

	id          string
	state       objectState
	Screenshots *ExecutionScreenshots
}

func newExecution(c *clientRef, id string) *Execution {
	o := &Execution{c: c, id: id}
	o.Screenshots = newExecutionScreenshots(o)
	return o
}

func (o *Execution) withInfo(info any, root string) *Execution {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *Execution) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *Execution) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *Execution) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *Execution) String() string {
	if o.Info() == nil {
		return objectString("Execution", &o.id)
	}
	return objectString("Execution", &o.id, "status", o.Status())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *Execution) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *Execution) Refresh(ctx context.Context, params mobilerun.WorkflowExecutionGetParams, opts ...option.RequestOption) (*mobilerun.WorkflowExecutionGetResponse, error) {
	res, err := o.c.api.Workflows.Executions.Get(ctx, o.id, params, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// Status reads status. It reads the last response without making a request.
func (o *Execution) Status() string { return resultString(o.state.field("status")) }

// Stops the execution between steps and marks it cancelled. Idempotent-ish:
// already-terminal executions return 409.
func (o *Execution) Abort(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WorkflowExecutionAbortResponse, error) {
	return o.c.api.Workflows.Executions.Abort(ctx, o.id, opts...)
}

// ExecutionScreenshots groups the screenshots methods of a Execution.
type ExecutionScreenshots struct {
	o *Execution
}

func newExecutionScreenshots(o *Execution) *ExecutionScreenshots {
	n := &ExecutionScreenshots{o: o}

	return n
}

func (n *ExecutionScreenshots) String() string { return objectString("ExecutionScreenshots", nil) }

// Redirects (302) to a freshly signed, time-limited image URL for a screenshot
// captured during this execution. Owner-gated: 404 if the execution or the
// screenshot does not belong to the caller.
//
// The object layer sets executionId; values for it in params are replaced.
func (n *ExecutionScreenshots) Get(ctx context.Context, screenshotID string, opts ...option.RequestOption) error {
	o := n.o
	return o.c.api.Workflows.Executions.Screenshots.Get(ctx, screenshotID, mobilerun.WorkflowExecutionScreenshotGetParams{ExecutionID: o.id}, opts...)
}

// info is the last create/get snapshot and may contain the create-time secret.
// Rotation revokes it and returns the new secret without updating info.
//
// Field accessors read the last response (Info) without making a request.
type Webhook struct {
	c *clientRef

	id         string
	state      objectState
	Deliveries *WebhookDeliveryManager
}

func newWebhook(c *clientRef, id string) *Webhook {
	o := &Webhook{c: c, id: id}
	o.Deliveries = newWebhookDeliveryManager(o.c, o.id)
	return o
}

func (o *Webhook) withInfo(info any, root string) *Webhook {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *Webhook) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *Webhook) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *Webhook) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *Webhook) String() string {
	if o.Info() == nil {
		return objectString("Webhook", &o.id)
	}
	return objectString("Webhook", &o.id, "url", o.URL())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *Webhook) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *Webhook) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WebhookGetResponse, error) {
	res, err := o.c.api.Webhooks.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// URL reads url. It reads the last response without making a request.
func (o *Webhook) URL() string { return resultString(o.state.field("url")) }

// Updates a webhook subscription. Any combination of the subscribed event types,
// state (ACTIVE or DISABLED), and description may be changed, and at least one
// field must be supplied. Setting state to ACTIVE re-enables a subscription that
// was auto-blocked after sustained delivery failures. The URL or integration
// target cannot be changed; delete and recreate the webhook instead.
func (o *Webhook) Update(ctx context.Context, params mobilerun.WebhookUpdateParams, opts ...option.RequestOption) (*mobilerun.WebhookUpdateResponse, error) {
	return o.c.api.Webhooks.Update(ctx, o.id, params, opts...)
}

// Deletes a webhook subscription so it stops receiving deliveries. Returns 204 No
// Content on success.
func (o *Webhook) Delete(ctx context.Context, opts ...option.RequestOption) error {
	return o.c.api.Webhooks.Delete(ctx, o.id, opts...)
}

// Rotates and revokes the old secret. The last create/get snapshot in info is
// preserved; any create-time secret in it is now revoked. Returns the new secret.
//
// Generates a new signing secret for the webhook subscription and returns it once
// in the response. The previous secret is replaced immediately, so any signature
// verification on your endpoint must be updated to use the new value. Only `http`
// webhooks have a signing secret; for `integration` webhooks this returns 400.
func (o *Webhook) RotateSecret(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WebhookRotateSecretResponse, error) {
	return o.c.api.Webhooks.RotateSecret(ctx, o.id, opts...)
}

// Sends a single test payload to the webhook subscription URL (or a test message
// to its integration target) to verify connectivity. The response reports whether
// the attempt succeeded along with the returned HTTP status code or error, if any.
func (o *Webhook) Test(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WebhookTestDeliveryResponse, error) {
	return o.c.api.Webhooks.TestDelivery(ctx, o.id, opts...)
}

// WebhookDelivery is a handle to one WebhookDelivery.
//
// Field accessors read the last response (Info) without making a request.
type WebhookDelivery struct {
	c         *clientRef
	webhookID string
	id        string
	state     objectState
}

func newWebhookDelivery(c *clientRef, webhookID string, id string) *WebhookDelivery {
	o := &WebhookDelivery{c: c, webhookID: webhookID, id: id}

	return o
}

func (o *WebhookDelivery) withInfo(info any, root string) *WebhookDelivery {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *WebhookDelivery) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *WebhookDelivery) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *WebhookDelivery) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *WebhookDelivery) String() string {
	if o.Info() == nil {
		return objectString("WebhookDelivery", &o.id)
	}
	return objectString("WebhookDelivery", &o.id, "status", o.Status())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *WebhookDelivery) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *WebhookDelivery) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.WebhookDeliveryGetAttemptsResponse, error) {
	res, err := o.c.api.Webhooks.Deliveries.GetAttempts(ctx, o.id, mobilerun.WebhookDeliveryGetAttemptsParams{ID: o.webhookID}, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// Status reads status. It reads the last response without making a request.
func (o *WebhookDelivery) Status() string { return resultString(o.state.field("status")) }

// WebhookID returns the bound Webhook ID.
func (o *WebhookDelivery) WebhookID() string { return o.webhookID }

// Mailbox is a handle to one Mailbox.
//
// Field accessors read the last response (Info) without making a request.
type Mailbox struct {
	c *clientRef

	id       string
	state    objectState
	Messages *MailboxMessageManager
}

func newMailbox(c *clientRef, id string) *Mailbox {
	o := &Mailbox{c: c, id: id}
	o.Messages = newMailboxMessageManager(o.c, o.id)
	return o
}

func (o *Mailbox) withInfo(info any, root string) *Mailbox {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *Mailbox) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *Mailbox) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *Mailbox) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *Mailbox) String() string {
	if o.Info() == nil {
		return objectString("Mailbox", &o.id)
	}
	return objectString("Mailbox", &o.id, "address", o.Address(), "status", o.Status())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *Mailbox) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *Mailbox) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.MailboxGetResponse, error) {
	res, err := o.c.api.Mailboxes.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// Address reads address. It reads the last response without making a request.
func (o *Mailbox) Address() string { return resultString(o.state.field("address")) }

// Status reads status. It reads the last response without making a request.
func (o *Mailbox) Status() string { return resultString(o.state.field("status")) }

// Only codes received after this wait started are considered unless after is
// passed. The comparison uses the client clock. Returns the OTP data including
// code, received time, sender, message ID and confidence; leaves info unchanged.
//
// WaitForOtp polls every 1000 ms for up to 120000 ms (WithPollInterval,
// WithWaitTimeout); failure states return a WaitFailureError and the deadline a
// WaitTimeoutError.
func (o *Mailbox) WaitForOtp(ctx context.Context, params mobilerun.MailboxOtpParams, opts ...option.RequestOption) (mobilerun.MailboxOtpResponseData, error) {
	controls, opts, err := splitOptions(opts)
	if err != nil {
		var zero mobilerun.MailboxOtpResponseData
		return zero, err
	}
	waitStart := time.Now()
	var defaults []binding
	if !params.After.Valid() {
		defaults = append(defaults, binding{"after", "query", waitStart})
	}
	poll := func(ctx context.Context, pollOpts []option.RequestOption) (*mobilerun.MailboxOtpResponse, bool, error) {
		requestOpts := append(append([]option.RequestOption{}, opts...), pollOpts...)
		res, err := o.c.api.Mailboxes.Otp(ctx, o.id, params, bindOptions(requestOpts, defaults...)...)
		if err != nil {
			return nil, false, err
		}
		return res, conditionMet(lookup(res, "data.code"), "non_empty"), nil
	}
	res, err := waitFor(ctx, durationOr(controls.waitTimeout, 120000*time.Millisecond), durationOr(controls.pollInterval, 1000*time.Millisecond), 30000*time.Millisecond, poll, retryPoll)
	if err != nil {
		var zero mobilerun.MailboxOtpResponseData
		return zero, err
	}
	return res.Data, nil
}

// Updates the label of a mailbox.
func (o *Mailbox) Update(ctx context.Context, params mobilerun.MailboxUpdateParams, opts ...option.RequestOption) (*mobilerun.MailboxUpdateResponse, error) {
	return o.c.api.Mailboxes.Update(ctx, o.id, params, opts...)
}

// Cancels the mailbox rental or schedules cancellation for an active paid mailbox.
// Existing addresses and messages are retained.
//
// Cancels a pending mailbox or schedules an active paid mailbox for cancellation.
// Existing addresses and messages are retained. Repeating the request is safe. An
// external inbox (Gmail) cannot be cancelled here; disconnect the link instead.
func (o *Mailbox) Delete(ctx context.Context, opts ...option.RequestOption) (*mobilerun.MailboxDeleteResponse, error) {
	return o.c.api.Mailboxes.Delete(ctx, o.id, opts...)
}

// Removes only this mailbox link. The Gmail connection used by your agents stays
// in place.
func (o *Mailbox) Disconnect(ctx context.Context, opts ...option.RequestOption) (*mobilerun.MailboxDisconnectResponse, error) {
	return o.c.api.Mailboxes.Disconnect(ctx, o.id, opts...)
}

// Returns the most likely recent OTP for the mailbox. Returns 204 when no matching
// code is available.
func (o *Mailbox) Otp(ctx context.Context, params mobilerun.MailboxOtpParams, opts ...option.RequestOption) (*mobilerun.MailboxOtpResponse, error) {
	return o.c.api.Mailboxes.Otp(ctx, o.id, params, opts...)
}

// Restarts a paid mailbox rental. Uses included capacity when available unless
// paid service is requested.
//
// Restarts an archived mailbox with the same address. Uses included capacity when
// available unless paid service is requested.
func (o *Mailbox) Restart(ctx context.Context, params mobilerun.MailboxRestartParams, opts ...option.RequestOption) (*mobilerun.MailboxRestartResponse, error) {
	return o.c.api.Mailboxes.Restart(ctx, o.id, params, opts...)
}

// Withdraws a scheduled cancellation. Only available while cancellation is
// pending.
func (o *Mailbox) Uncancel(ctx context.Context, opts ...option.RequestOption) (*mobilerun.MailboxUncancelResponse, error) {
	return o.c.api.Mailboxes.Uncancel(ctx, o.id, opts...)
}

// MailboxConnection is a handle to one MailboxConnection.
//
// Field accessors read the last response (Info) without making a request.
type MailboxConnection struct {
	c *clientRef

	id    string
	state objectState
}

func newMailboxConnection(c *clientRef, id string) *MailboxConnection {
	o := &MailboxConnection{c: c, id: id}

	return o
}

func (o *MailboxConnection) withInfo(info any, root string) *MailboxConnection {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *MailboxConnection) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *MailboxConnection) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *MailboxConnection) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *MailboxConnection) String() string {
	if o.Info() == nil {
		return objectString("MailboxConnection", &o.id)
	}
	return objectString("MailboxConnection", &o.id, "address", o.Address(), "status", o.Status())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *MailboxConnection) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *MailboxConnection) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.MailboxConnectionGetResponse, error) {
	res, err := o.c.api.Mailboxes.Connections.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "")
	return res, nil
}

// Address reads address. It reads the last response without making a request.
func (o *MailboxConnection) Address() string { return resultString(o.state.field("address")) }

// Status reads status. It reads the last response without making a request.
func (o *MailboxConnection) Status() string { return resultString(o.state.field("status")) }

// MailboxMessage is a handle to one MailboxMessage.
//
// Field accessors read the last response (Info) without making a request.
type MailboxMessage struct {
	c         *clientRef
	mailboxID string
	id        string
	state     objectState
}

func newMailboxMessage(c *clientRef, mailboxID string, id string) *MailboxMessage {
	o := &MailboxMessage{c: c, mailboxID: mailboxID, id: id}

	return o
}

func (o *MailboxMessage) withInfo(info any, root string) *MailboxMessage {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *MailboxMessage) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *MailboxMessage) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *MailboxMessage) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *MailboxMessage) String() string {
	if o.Info() == nil {
		return objectString("MailboxMessage", &o.id)
	}
	return objectString("MailboxMessage", &o.id, "subject", o.Subject())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *MailboxMessage) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *MailboxMessage) Refresh(ctx context.Context, params mobilerun.MailboxMessageGetParams, opts ...option.RequestOption) (*mobilerun.MailboxMessageGetResponse, error) {
	params.MailboxID = o.mailboxID
	res, err := o.c.api.Mailboxes.Messages.Get(ctx, o.id, params, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// Subject reads subject. It reads the last response without making a request.
func (o *MailboxMessage) Subject() string { return resultString(o.state.field("subject")) }

// MailboxID returns the bound Mailbox ID.
func (o *MailboxMessage) MailboxID() string { return o.mailboxID }

// RoutineFile is a handle to one RoutineFile.
//
// Field accessors read the last response (Info) without making a request.
type RoutineFile struct {
	c *clientRef

	id    string
	state objectState
}

func newRoutineFile(c *clientRef, id string) *RoutineFile {
	o := &RoutineFile{c: c, id: id}

	return o
}

func (o *RoutineFile) withInfo(info any, root string) *RoutineFile {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *RoutineFile) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *RoutineFile) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *RoutineFile) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *RoutineFile) String() string {
	if o.Info() == nil {
		return objectString("RoutineFile", &o.id)
	}
	return objectString("RoutineFile", &o.id, "filename", o.Filename(), "folder", o.Folder(), "kind", o.Kind())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *RoutineFile) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Filename reads filename. It reads only the last New/List snapshot without making
// a request; FromID handles have no data.
func (o *RoutineFile) Filename() string { return resultString(o.state.field("filename")) }

// Folder reads folder. It reads only the last New/List snapshot without making a
// request; FromID handles have no data.
func (o *RoutineFile) Folder() string { return resultString(o.state.field("folder")) }

// Kind reads kind. It reads only the last New/List snapshot without making a
// request; FromID handles have no data.
func (o *RoutineFile) Kind() string { return resultString(o.state.field("kind")) }

// Permanently deletes this saved routine file.
//
// Deletes one saved file permanently and frees its storage (the file disappears at
// once; the stored object is removed right after). Idempotent: an already deleted
// file returns 204 too.
func (o *RoutineFile) Delete(ctx context.Context, opts ...option.RequestOption) error {
	return o.c.api.RoutineFiles.Delete(ctx, o.id, opts...)
}

// Redirects (302) to a time-limited (at most 5 minutes) download URL for the file.
// 404 when the file does not exist or belongs to someone else.
//
// Download returns the raw response bytes.
func (o *RoutineFile) Download(ctx context.Context, opts ...option.RequestOption) ([]byte, error) {
	var body []byte
	var response *http.Response
	err := o.c.api.RoutineFiles.Download(ctx, o.id, append(opts, option.WithResponseBodyInto(&body), option.WithResponseInto(&response))...)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// AgentFile is a handle to one AgentFile.
//
// Field accessors read the last response (Info) without making a request.
type AgentFile struct {
	c *clientRef

	id    string
	state objectState
}

func newAgentFile(c *clientRef, id string) *AgentFile {
	o := &AgentFile{c: c, id: id}

	return o
}

func (o *AgentFile) withInfo(info any, root string) *AgentFile {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *AgentFile) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *AgentFile) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *AgentFile) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *AgentFile) String() string {
	if o.Info() == nil {
		return objectString("AgentFile", &o.id)
	}
	return objectString("AgentFile", &o.id, "filename", o.Filename(), "displayName", o.DisplayName())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *AgentFile) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Filename reads filename. It reads only the last New/List snapshot without making
// a request; FromID handles have no data.
func (o *AgentFile) Filename() string { return resultString(o.state.field("filename")) }

// DisplayName reads displayName. It reads only the last New/List snapshot without
// making a request; FromID handles have no data.
func (o *AgentFile) DisplayName() string { return resultString(o.state.field("displayName")) }

// Partial update of `displayName` and/or `enabled`. Only files with `zone=skills`
// are mutable; other zones return 422 `unsupported_zone`.
func (o *AgentFile) Update(ctx context.Context, params mobilerun.FileUpdateParams, opts ...option.RequestOption) (*mobilerun.FileUpdateResponse, error) {
	return o.c.api.Files.Update(ctx, o.id, params, opts...)
}

// Hard-delete a file.
func (o *AgentFile) Delete(ctx context.Context, opts ...option.RequestOption) (*mobilerun.FileDeleteResponse, error) {
	return o.c.api.Files.Delete(ctx, o.id, opts...)
}

// Soft-cancels an in-flight upload before confirm, transitioning it to expired and
// deleting the underlying object. Only acts on `pending` uploads — refuses to
// touch `ready` files, so confirmed files are never removed. Idempotent:
// `{ cancelled: false }` if the file exists but is no longer pending.
func (o *AgentFile) CancelPending(ctx context.Context, opts ...option.RequestOption) (*mobilerun.FileCancelPendingResponse, error) {
	return o.c.api.Files.CancelPending(ctx, o.id, opts...)
}

// Confirm a file upload after validating the stored object with a server-side HEAD
// request.
func (o *AgentFile) Confirm(ctx context.Context, opts ...option.RequestOption) (*mobilerun.FileConfirmResponse, error) {
	return o.c.api.Files.Confirm(ctx, o.id, opts...)
}

// Redirects to a short-lived presigned download URL. Missing, non-ready, and files
// owned by another tenant all return the same 404 response.
//
// Download returns the raw response bytes.
func (o *AgentFile) Download(ctx context.Context, opts ...option.RequestOption) ([]byte, error) {
	var body []byte
	var response *http.Response
	err := o.c.api.Files.Download(ctx, o.id, append(opts, option.WithResponseBodyInto(&body), option.WithResponseInto(&response))...)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Agent is a handle to one Agent.
//
// Field accessors read the last response (Info) without making a request.
type Agent struct {
	c *clientRef

	id            string
	state         objectState
	Files         *AgentFiles
	Conversations *AgentConversations
	Requests      *AgentRequests
	Routines      *AgentRoutines
	Profile       *AgentProfile
}

func newAgent(c *clientRef, id string) *Agent {
	o := &Agent{c: c, id: id}
	o.Files = newAgentFiles(o)
	o.Conversations = newAgentConversations(o)
	o.Requests = newAgentRequests(o)
	o.Routines = newAgentRoutines(o)
	o.Profile = newAgentProfile(o)
	return o
}

func (o *Agent) withInfo(info any, root string) *Agent {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *Agent) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *Agent) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *Agent) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *Agent) String() string {
	if o.Info() == nil {
		return objectString("Agent", &o.id)
	}
	return objectString("Agent", &o.id, "displayName", o.DisplayName(), "status", o.Status())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *Agent) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// DisplayName reads profile.displayName. It reads only the last New/List snapshot
// without making a request; FromID handles have no data.
func (o *Agent) DisplayName() string { return resultString(o.state.field("profile.displayName")) }

// Status reads status. It reads only the last New/List snapshot without making a
// request; FromID handles have no data.
func (o *Agent) Status() string { return resultString(o.state.field("status")) }

// AgentFiles groups the files methods of a Agent.
type AgentFiles struct {
	o *Agent
}

func newAgentFiles(o *Agent) *AgentFiles {
	n := &AgentFiles{o: o}

	return n
}

func (n *AgentFiles) String() string { return objectString("AgentFiles", nil) }

// List the ready files of one agent, optionally filtered by zone.
//
// It returns the first page of AgentFile handles; ListAutoPaging iterates every
// page.
func (n *AgentFiles) List(ctx context.Context, params mobilerun.AgentFileListParams, opts ...option.RequestOption) (*ObjectPage[*AgentFile, *mobilerun.AgentFileListResponse], error) {
	o := n.o
	load := func() (*ObjectPage[*AgentFile, *mobilerun.AgentFileListResponse], error) {
		pageOpts := opts
		res, err := o.c.api.Agents.Files.List(ctx, o.id, params, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*AgentFile, *mobilerun.AgentFileListResponse]{Response: res}
		for _, info := range objectRows(res, "files") {
			result.Items = append(result.Items, newAgentFile(o.c, objectID(info, "fileId")).withInfo(info, ""))
		}
		return result, nil
	}
	return load()
}

// ListAutoPaging iterates the AgentFile handles of every page.
func (n *AgentFiles) ListAutoPaging(ctx context.Context, params mobilerun.AgentFileListParams, opts ...option.RequestOption) *ObjectAutoPager[*AgentFile] {
	page, err := n.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// AgentConversations groups the conversations methods of a Agent.
type AgentConversations struct {
	o *Agent
}

func newAgentConversations(o *Agent) *AgentConversations {
	n := &AgentConversations{o: o}

	return n
}

func (n *AgentConversations) String() string { return objectString("AgentConversations", nil) }

// Create a named chat session bound to one agent. Idempotent via the
// `Idempotency-Key` header, like the unscoped create.
func (n *AgentConversations) New(ctx context.Context, params mobilerun.AgentConversationNewParams, opts ...option.RequestOption) (*mobilerun.AgentConversationNewResponse, error) {
	o := n.o
	return o.c.api.Agents.Conversations.New(ctx, o.id, params, opts...)
}

// List the user's named chat sessions for one agent. Supports the same `kind`,
// `workflowId` and `mine` filters as the unscoped list.
func (n *AgentConversations) List(ctx context.Context, params mobilerun.AgentConversationListParams, opts ...option.RequestOption) (*mobilerun.AgentConversationListResponse, error) {
	o := n.o
	return o.c.api.Agents.Conversations.List(ctx, o.id, params, opts...)
}

// AgentRequests groups the requests methods of a Agent.
type AgentRequests struct {
	o *Agent
}

func newAgentRequests(o *Agent) *AgentRequests {
	n := &AgentRequests{o: o}

	return n
}

func (n *AgentRequests) String() string { return objectString("AgentRequests", nil) }

// Owner-scoped. Returns newest requests where the Agent is source or target. Use
// nextCursor as before to read older pages. Invalid cursors return 400
// invalid_cursor.
//
// It returns the first page of AgentRequest handles; ListForAgentAutoPaging
// iterates every page.
func (n *AgentRequests) ListForAgent(ctx context.Context, params mobilerun.AgentRequestListForAgentParams, opts ...option.RequestOption) (*ObjectPage[*AgentRequest, *mobilerun.AgentRequestListForAgentResponse], error) {
	o := n.o
	load := func() (*ObjectPage[*AgentRequest, *mobilerun.AgentRequestListForAgentResponse], error) {
		pageOpts := opts
		res, err := o.c.api.Agents.Requests.ListForAgent(ctx, o.id, params, pageOpts...)
		if err != nil {
			return nil, err
		}
		result := &ObjectPage[*AgentRequest, *mobilerun.AgentRequestListForAgentResponse]{Response: res}
		for _, info := range objectRows(res, "requests") {
			result.Items = append(result.Items, newAgentRequest(o.c, objectID(info, "id")).withInfo(info, ""))
		}
		return result, nil
	}
	return load()
}

// ListForAgentAutoPaging iterates the AgentRequest handles of every page.
func (n *AgentRequests) ListForAgentAutoPaging(ctx context.Context, params mobilerun.AgentRequestListForAgentParams, opts ...option.RequestOption) *ObjectAutoPager[*AgentRequest] {
	page, err := n.ListForAgent(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}

// AgentRoutines groups the routines methods of a Agent.
type AgentRoutines struct {
	o *Agent
}

func newAgentRoutines(o *Agent) *AgentRoutines {
	n := &AgentRoutines{o: o}

	return n
}

func (n *AgentRoutines) String() string { return objectString("AgentRoutines", nil) }

// Forwards the scheduler CreateFlowBody unchanged and assigns the Agent selected
// in the path. Requires Idempotency-Key; reuse it after a transient failure.
func (n *AgentRoutines) New(ctx context.Context, params mobilerun.AgentRoutineNewParams, opts ...option.RequestOption) (*mobilerun.AgentRoutineNewResponse, error) {
	o := n.o
	return o.c.api.Agents.Routines.New(ctx, o.id, params, opts...)
}

// The source must belong to the authenticated owner and be assigned to the same
// Agent or unassigned. The target is the Agent in the path. Requires
// Idempotency-Key.
//
// The object layer sets agentId; values for it in params are replaced.
func (n *AgentRoutines) Clone(ctx context.Context, routineID string, params mobilerun.AgentRoutineCloneParams, opts ...option.RequestOption) (*mobilerun.AgentRoutineCloneResponse, error) {
	o := n.o
	params.AgentID = o.id
	return o.c.api.Agents.Routines.Clone(ctx, routineID, params, opts...)
}

// AgentProfile groups the profile methods of a Agent.
type AgentProfile struct {
	o *Agent
}

func newAgentProfile(o *Agent) *AgentProfile {
	n := &AgentProfile{o: o}

	return n
}

func (n *AgentProfile) String() string { return objectString("AgentProfile", nil) }

// Owner-scoped partial update of `displayName`, `instructions`, `modelProfile` and
// the avatar. A null modelProfile uses the environment default; an unknown id
// returns unknown_model_profile with details.validProfiles. Returns the updated
// Agent; an Agent of another owner or an unknown id is a 404.
//
// Update returns a Agent handle; the handle does not own cleanup.
func (n *AgentProfile) Update(ctx context.Context, params mobilerun.AgentProfileUpdateParams, opts ...option.RequestOption) (*Agent, error) {
	o := n.o
	res, err := o.c.api.Agents.Profile.Update(ctx, o.id, params, opts...)
	if err != nil {
		return nil, err
	}
	return newAgent(o.c, objectID(res, "id")).withInfo(res, ""), nil
}

// AgentRequest is a handle to one AgentRequest.
//
// Field accessors read the last response (Info) without making a request.
type AgentRequest struct {
	c *clientRef

	id    string
	state objectState
}

func newAgentRequest(c *clientRef, id string) *AgentRequest {
	o := &AgentRequest{c: c, id: id}

	return o
}

func (o *AgentRequest) withInfo(info any, root string) *AgentRequest {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *AgentRequest) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *AgentRequest) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *AgentRequest) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *AgentRequest) String() string {
	if o.Info() == nil {
		return objectString("AgentRequest", &o.id)
	}
	return objectString("AgentRequest", &o.id, "status", o.Status())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *AgentRequest) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *AgentRequest) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.AgentRequestGetResponse, error) {
	res, err := o.c.api.Agents.Requests.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "")
	return res, nil
}

// Status reads status. It reads the last response without making a request.
func (o *AgentRequest) Status() string { return resultString(o.state.field("status")) }

// NumberMessage is a handle to one NumberMessage.
//
// Field accessors read the last response (Info) without making a request.
type NumberMessage struct {
	c *clientRef

	id    string
	state objectState
}

func newNumberMessage(c *clientRef, id string) *NumberMessage {
	o := &NumberMessage{c: c, id: id}

	return o
}

func (o *NumberMessage) withInfo(info any, root string) *NumberMessage {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *NumberMessage) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *NumberMessage) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *NumberMessage) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *NumberMessage) String() string {
	if o.Info() == nil {
		return objectString("NumberMessage", &o.id)
	}
	return objectString("NumberMessage", &o.id, "status", o.Status(), "direction", o.Direction())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *NumberMessage) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *NumberMessage) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.MessageGetResponse, error) {
	res, err := o.c.api.Messages.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// Status reads status. It reads the last response without making a request.
func (o *NumberMessage) Status() string { return resultString(o.state.field("status")) }

// Direction reads direction. It reads the last response without making a request.
func (o *NumberMessage) Direction() string { return resultString(o.state.field("direction")) }

// PhoneNumber is a handle to one PhoneNumber.
//
// Field accessors read the last response (Info) without making a request.
type PhoneNumber struct {
	c *clientRef

	id       string
	state    objectState
	Messages *PhoneNumberMessages
}

func newPhoneNumber(c *clientRef, id string) *PhoneNumber {
	o := &PhoneNumber{c: c, id: id}
	o.Messages = newPhoneNumberMessages(o)
	return o
}

func (o *PhoneNumber) withInfo(info any, root string) *PhoneNumber {
	o.state.setInfo(info, root)
	return o
}

// ID returns the object ID.
func (o *PhoneNumber) ID() string { return o.id }

// Info returns the last response this handle saw, or nil for an unfetched handle.
// Its dynamic type is the response type of the call that produced it.
func (o *PhoneNumber) Info() any {
	info, _ := o.state.snapshot()
	return info
}

// InfoJSON returns the raw JSON of Info.
func (o *PhoneNumber) InfoJSON() string { return rawJSON(o.Info()) }

// String shows the class, ID and declared fields, never Info or the client.
func (o *PhoneNumber) String() string {
	if o.Info() == nil {
		return objectString("PhoneNumber", &o.id)
	}
	return objectString("PhoneNumber", &o.id, "phoneNumber", o.PhoneNumber(), "state", o.State())
}

// Close disposes the object when this handle created it. Handles from Get, List,
// FromID or other methods never dispose; a second Close is a no-op. opts apply to
// the dispose request.
func (o *PhoneNumber) Close(ctx context.Context, opts ...option.RequestOption) error {
	if !o.state.takeOwnership() {
		return nil
	}
	return nil
}

// Refresh fetches the object and updates Info.
func (o *PhoneNumber) Refresh(ctx context.Context, opts ...option.RequestOption) (*mobilerun.NumberGetResponse, error) {
	res, err := o.c.api.Numbers.Get(ctx, o.id, opts...)
	if err != nil {
		return nil, err
	}
	o.state.setInfo(res, "data")
	return res, nil
}

// PhoneNumber reads phoneNumber. It is not available on handles from
// `numbers.create()` until Refresh.
func (o *PhoneNumber) PhoneNumber() string { return resultString(o.state.field("phoneNumber")) }

// State reads state. It reads the last response without making a request.
func (o *PhoneNumber) State() string { return resultString(o.state.field("state")) }

// Wait for the number to become active. awaiting_payment means checkout is not
// completed; the last response carried by WaitFailureError includes the checkout
// URL in data.checkoutUrl (data.checkout_url in Python).
//
// WaitUntilActive polls every 1000 ms for up to 300000 ms (WithPollInterval,
// WithWaitTimeout); failure states return a WaitFailureError and the deadline a
// WaitTimeoutError.
func (o *PhoneNumber) WaitUntilActive(ctx context.Context, opts ...option.RequestOption) (*mobilerun.NumberGetResponse, error) {
	controls, opts, err := splitOptions(opts)
	if err != nil {
		return nil, err
	}
	poll := func(ctx context.Context, pollOpts []option.RequestOption) (*mobilerun.NumberGetResponse, bool, error) {
		requestOpts := append(append([]option.RequestOption{}, opts...), pollOpts...)
		res, err := o.c.api.Numbers.Get(ctx, o.id, requestOpts...)
		if err != nil {
			return nil, false, err
		}
		o.state.setInfo(res, "data")
		done, err := enumState(res, "data.state", []string{"active", "cancel_scheduled"}, []string{"expired", "failed", "awaiting_payment"}, []string{"provisioning"}, false)
		return res, done, err
	}
	return waitFor(ctx, durationOr(controls.waitTimeout, 300000*time.Millisecond), durationOr(controls.pollInterval, 1000*time.Millisecond), 30000*time.Millisecond, poll, retryPoll)
}

// Updates the display label. Omitting `label` leaves it unchanged; null or an
// empty string clears it.
func (o *PhoneNumber) Update(ctx context.Context, params mobilerun.NumberUpdateParams, opts ...option.RequestOption) (*mobilerun.NumberUpdateResponse, error) {
	return o.c.api.Numbers.Update(ctx, o.id, params, opts...)
}

// Cancels the phone number rental.
//
// Cancels a pending purchase or schedules cancellation of an active paid phone
// number. Repeating a scheduled cancellation is safe. Returns 409 when
// cancellation is not available.
func (o *PhoneNumber) Delete(ctx context.Context, opts ...option.RequestOption) (*mobilerun.NumberDeleteResponse, error) {
	return o.c.api.Numbers.Delete(ctx, o.id, opts...)
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
//
// Send returns a NumberMessage handle; the handle does not own cleanup.
func (o *PhoneNumber) Send(ctx context.Context, params mobilerun.NumberMessageSendParams, opts ...option.RequestOption) (*NumberMessage, error) {
	res, err := o.c.api.Numbers.Messages.Send(ctx, o.id, params, opts...)
	if err != nil {
		return nil, err
	}
	return newNumberMessage(o.c, objectID(res, "data.id")).withInfo(res, "data"), nil
}

// PhoneNumberMessages groups the messages methods of a PhoneNumber.
type PhoneNumberMessages struct {
	o *PhoneNumber
}

func newPhoneNumberMessages(o *PhoneNumber) *PhoneNumberMessages {
	n := &PhoneNumberMessages{o: o}

	return n
}

func (n *PhoneNumberMessages) String() string { return objectString("PhoneNumberMessages", nil) }

// Returns SMS messages on the number, inbound and outbound, scoped to the
// authenticated user. Newest first. Messages stay with the number across device
// switches.
//
// It returns the first page of NumberMessage handles; ListAutoPaging iterates
// every page.
func (n *PhoneNumberMessages) List(ctx context.Context, params mobilerun.NumberMessageListParams, opts ...option.RequestOption) (*ObjectPage[*NumberMessage, *pagination.Page[mobilerun.NumberMessageListResponse]], error) {
	o := n.o
	var wrap func(page *pagination.Page[mobilerun.NumberMessageListResponse]) *ObjectPage[*NumberMessage, *pagination.Page[mobilerun.NumberMessageListResponse]]
	wrap = func(page *pagination.Page[mobilerun.NumberMessageListResponse]) *ObjectPage[*NumberMessage, *pagination.Page[mobilerun.NumberMessageListResponse]] {
		result := &ObjectPage[*NumberMessage, *pagination.Page[mobilerun.NumberMessageListResponse]]{Response: page}
		for i := range page.Items {
			info := &page.Items[i]
			result.Items = append(result.Items, newNumberMessage(o.c, objectID(info, "id")).withInfo(info, ""))
		}
		result.next = func() (*ObjectPage[*NumberMessage, *pagination.Page[mobilerun.NumberMessageListResponse]], error) {
			next, err := page.GetNextPage()
			if err != nil || next == nil {
				return nil, err
			}
			return wrap(next), nil
		}
		return result
	}
	pageOpts := append(append([]option.RequestOption{}, opts...), callerContext(ctx))
	page, err := o.c.api.Numbers.Messages.List(ctx, o.id, params, pageOpts...)
	if err != nil {
		return nil, err
	}
	return wrap(page), nil
}

// ListAutoPaging iterates the NumberMessage handles of every page.
func (n *PhoneNumberMessages) ListAutoPaging(ctx context.Context, params mobilerun.NumberMessageListParams, opts ...option.RequestOption) *ObjectAutoPager[*NumberMessage] {
	page, err := n.List(ctx, params, opts...)
	return newObjectAutoPager(page, err)
}
