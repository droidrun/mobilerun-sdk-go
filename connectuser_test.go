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
)

func TestConnectUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Connect.Users.New(context.TODO(), mobilerun.ConnectUserNewParams{
		Password: mobilerun.String("x"),
		ProxyID:  mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Username: mobilerun.String("x"),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectUserGet(t *testing.T) {
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
	_, err := client.Connect.Users.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectUserUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Connect.Users.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.ConnectUserUpdateParams{
			ProxyID: mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestConnectUserListWithOptionalParams(t *testing.T) {
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
	_, err := client.Connect.Users.List(context.TODO(), mobilerun.ConnectUserListParams{
		Page:     mobilerun.Int(1),
		PageSize: mobilerun.Int(1),
		ProxyID:  mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectUserDelete(t *testing.T) {
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
	err := client.Connect.Users.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectUserListConnectionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Connect.Users.ListConnections(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.ConnectUserListConnectionsParams{
			CloseReason:   mobilerun.String("closeReason"),
			Country:       mobilerun.String("country"),
			DstHost:       mobilerun.String("dstHost"),
			DstPort:       mobilerun.Int(0),
			EndedAfter:    mobilerun.Time(time.Now()),
			EndedBefore:   mobilerun.Time(time.Now()),
			MaxBytesIn:    mobilerun.Int(0),
			MaxBytesOut:   mobilerun.Int(0),
			MaxDurationMs: mobilerun.Int(0),
			MaxTotalBytes: mobilerun.Int(0),
			MinBytesIn:    mobilerun.Int(0),
			MinBytesOut:   mobilerun.Int(0),
			MinDurationMs: mobilerun.Int(0),
			MinTotalBytes: mobilerun.Int(0),
			Order:         mobilerun.ConnectUserListConnectionsParamsOrderAsc,
			OrderBy:       mobilerun.ConnectUserListConnectionsParamsOrderByStartedAt,
			Page:          mobilerun.Int(1),
			PageSize:      mobilerun.Int(1),
			Protocol:      mobilerun.ConnectUserListConnectionsParamsProtocolTcp,
			Provider:      mobilerun.String("provider"),
			ProxyID:       mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			SessionID:     mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			StartedAfter:  mobilerun.Time(time.Now()),
			StartedBefore: mobilerun.Time(time.Now()),
			Status:        mobilerun.ConnectUserListConnectionsParamsStatusActive,
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
