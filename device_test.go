// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/testutil"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/shared"
)

func TestDeviceNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Devices.New(context.TODO(), mobilerun.DeviceNewParams{
		DeviceSpec: shared.DeviceSpecParam{
			AndroidVersion: mobilerun.Int(0),
			Apps:           []string{"string"},
			Carrier: shared.DeviceCarrierParam{
				GsmOperatorAlpha:         "GsmOperatorAlpha",
				GsmOperatorNumeric:       0,
				GsmSimOperatorAlpha:      "GsmSimOperatorAlpha",
				GsmSimOperatorISOCountry: "GsmSimOperatorIsoCountry",
				GsmSimOperatorNumeric:    0,
				PersistSysTimezone:       "PersistSysTimezone",
			},
			Country: mobilerun.String("country"),
			Files:   []string{"string"},
			Identifiers: shared.DeviceIdentifiersParam{
				BootloaderSerialNumber: "BootloaderSerialNumber",
				IdentifierAndroidID:    "IdentifierAndroidID",
				IdentifierAppSetID:     "IdentifierAppSetID",
				IdentifierBluetoothMac: "IdentifierBluetoothMAC",
				IdentifierGaid:         "IdentifierGAID",
				IdentifierGsfid:        "IdentifierGSFID",
				IdentifierIccid:        "IdentifierICCID",
				IdentifierImei:         "IdentifierIMEI",
				IdentifierImsi:         "IdentifierIMSI",
				IdentifierMediaDrmid:   "IdentifierMediaDRMID",
				IdentifierMeid:         "IdentifierMEID",
				IdentifierPhoneNumber:  "IdentifierPhoneNumber",
				IdentifierSerial:       "IdentifierSerial",
				IdentifierWifiMac:      "IdentifierWifiMAC",
				SerialNumber:           "SerialNumber",
			},
			Locale: mobilerun.String("locale"),
			Location: shared.LocationParam{
				Latitude:  0,
				Longitude: 0,
			},
			Name: mobilerun.String("name"),
			Proxy: shared.DeviceSpecProxyParam{
				Connect: shared.DeviceSpecProxyConnectParam{
					ID:      mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Country: mobilerun.String("SE"),
				},
				Name:    mobilerun.String("name"),
				SmartIP: mobilerun.Bool(true),
				Socks5: shared.DeviceSpecProxySocks5Param{
					Host:     "host",
					Password: "password",
					Port:     0,
					User:     "user",
				},
			},
			Timezone: mobilerun.String("timezone"),
		},
		Billing:    mobilerun.DeviceNewParamsBillingAuto,
		Country:    mobilerun.String("country"),
		DeviceType: mobilerun.DeviceNewParamsDeviceTypeAndroidCloudPhone,
		ProfileID:  mobilerun.String("profileId"),
		Template:   mobilerun.String("onboarding"),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Devices.Get(context.TODO(), "deviceId")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Devices.List(context.TODO(), mobilerun.DeviceListParams{
		Country:          mobilerun.String("country"),
		CreatedAfter:     mobilerun.Time(time.Now()),
		CreatedBefore:    mobilerun.Time(time.Now()),
		CreatedBy:        mobilerun.String("createdBy"),
		Mine:             mobilerun.Bool(true),
		Name:             mobilerun.String("name"),
		OrderBy:          mobilerun.DeviceListParamsOrderByID,
		OrderByDirection: mobilerun.DeviceListParamsOrderByDirectionAsc,
		Page:             mobilerun.Int(0),
		PageSize:         mobilerun.Int(0),
		Platform:         mobilerun.DeviceListParamsPlatformAndroid,
		ProviderID:       mobilerun.String("providerId"),
		Search:           mobilerun.String("search"),
		State:            []string{"creating"},
		Type:             mobilerun.DeviceListParamsTypeAndroidCloudPhone,
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceCount(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Devices.Count(context.TODO())
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceFingerprintWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Devices.Fingerprint(
		context.TODO(),
		"deviceId",
		mobilerun.DeviceFingerprintParams{
			XDeviceDisplayID: mobilerun.Int(0),
		},
	)
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceReboot(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Devices.Reboot(context.TODO(), "deviceId")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceResetWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Devices.Reset(
		context.TODO(),
		"deviceId",
		mobilerun.DeviceResetParams{
			Mode:     mobilerun.DeviceResetParamsModeFull,
			Preserve: []string{"proxy"},
		},
	)
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceResume(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Devices.Resume(context.TODO(), "deviceId")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceGetCapabilities(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Devices.GetCapabilities(context.TODO(), "deviceId")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceSetName(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Devices.SetName(
		context.TODO(),
		"deviceId",
		mobilerun.DeviceSetNameParams{
			Name: "x",
		},
	)
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceStop(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Devices.Stop(context.TODO(), "deviceId")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceSummaryWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Devices.Summary(context.TODO(), mobilerun.DeviceSummaryParams{
		State: []string{"creating"},
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceTerminateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Devices.Terminate(
		context.TODO(),
		"deviceId",
		mobilerun.DeviceTerminateParams{
			PreviousDeviceID: mobilerun.String("previousDeviceId"),
			TerminateAt:      mobilerun.Time(time.Now()),
		},
	)
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceUpdateIdleTimeout(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Devices.UpdateIdleTimeout(
		context.TODO(),
		"deviceId",
		mobilerun.DeviceUpdateIdleTimeoutParams{
			Minutes: mobilerun.Int(0),
		},
	)
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceWaitReady(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mobilerun.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Devices.WaitReady(context.TODO(), "deviceId")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
