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

func TestWorkflowExecutionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Executions.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.WorkflowExecutionGetParams{
			Include: mobilerun.WorkflowExecutionGetParamsIncludeInput,
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

func TestWorkflowExecutionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Executions.List(context.TODO(), mobilerun.WorkflowExecutionListParams{
		FlowID:           mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		From:             mobilerun.Time(time.Now()),
		InvocationID:     mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Kind:             []string{"live"},
		OrderBy:          mobilerun.WorkflowExecutionListParamsOrderByStartedAt,
		OrderByDirection: mobilerun.WorkflowExecutionListParamsOrderByDirectionAsc,
		Page:             mobilerun.Int(1),
		PageSize:         mobilerun.Int(1),
		Search:           mobilerun.String("x"),
		Status:           mobilerun.WorkflowExecutionListParamsStatusPending,
		To:               mobilerun.Time(time.Now()),
		TriggerID:        mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowExecutionAbort(t *testing.T) {
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
	_, err := client.Workflows.Executions.Abort(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowExecutionGetMetricsWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Executions.GetMetrics(context.TODO(), mobilerun.WorkflowExecutionGetMetricsParams{
		FlowID:    mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		From:      mobilerun.Time(time.Now()),
		To:        mobilerun.Time(time.Now()),
		TriggerID: mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
