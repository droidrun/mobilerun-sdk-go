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

func TestMailboxMessageGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Mailboxes.Messages.Get(
		context.TODO(),
		"x",
		mobilerun.MailboxMessageGetParams{
			MailboxID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			IncludeHTML: mobilerun.MailboxMessageGetParamsIncludeHTMLTrue,
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

func TestMailboxMessageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Mailboxes.Messages.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.MailboxMessageListParams{
			After:  mobilerun.Time(time.Now()),
			Cursor: mobilerun.String("cursor"),
			HasOtp: mobilerun.MailboxMessageListParamsHasOtpTrue,
			Limit:  mobilerun.Int(1),
			Sender: mobilerun.String("sender"),
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
