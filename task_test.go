// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/droidrun/mobilerun-sdk-go/v6"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/testutil"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
)

func TestTaskGet(t *testing.T) {
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
	_, err := client.Tasks.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskListWithOptionalParams(t *testing.T) {
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
	_, err := client.Tasks.List(context.TODO(), mobilerun.TaskListParams{
		CreatedBy:        mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Mine:             mobilerun.Bool(true),
		OrderBy:          mobilerun.TaskListParamsOrderByID,
		OrderByDirection: mobilerun.TaskListParamsOrderByDirectionAsc,
		Page:             mobilerun.Int(1),
		PageSize:         mobilerun.Int(1),
		Query:            mobilerun.String("query"),
		Source:           mobilerun.TaskListParamsSourceAPI,
		Status:           mobilerun.TaskListParamsStatusPrepared,
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskAttach(t *testing.T) {
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
	err := client.Tasks.Attach(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskGetStatus(t *testing.T) {
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
	_, err := client.Tasks.GetStatus(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskGetTrajectory(t *testing.T) {
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
	_, err := client.Tasks.GetTrajectory(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskRunWithOptionalParams(t *testing.T) {
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
	_, err := client.Tasks.Run(context.TODO(), mobilerun.TaskRunParams{
		DeviceID:          "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		Task:              "x",
		Accessibility:     mobilerun.Bool(true),
		Apps:              []string{"string"},
		ContinueOnFailure: mobilerun.Bool(true),
		Credentials: []mobilerun.TaskRunParamsCredential{{
			CredentialNames: []string{"string"},
			PackageName:     "packageName",
		}},
		DisplayID:        mobilerun.Int(0),
		ExecutionTimeout: mobilerun.Int(1),
		Files:            []string{"string"},
		Humanize:         mobilerun.Bool(true),
		Internal:         mobilerun.Bool(true),
		LlmModel:         mobilerun.String("llmModel"),
		MaxSteps:         mobilerun.Int(0),
		OutputSchema: map[string]any{
			"foo": "bar",
		},
		Reasoning:        mobilerun.Bool(true),
		RecordingEnabled: mobilerun.Bool(true),
		Stealth:          mobilerun.Bool(true),
		SubagentModel:    mobilerun.String("subagentModel"),
		SystemPrompt:     mobilerun.String("systemPrompt"),
		Temperature:      mobilerun.Float(0),
		Vision:           mobilerun.Bool(true),
		VpnCountry:       mobilerun.TaskRunParamsVpnCountryUs,
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskRunStreamedWithOptionalParams(t *testing.T) {
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
	_, err := client.Tasks.RunStreamed(context.TODO(), mobilerun.TaskRunStreamedParams{
		DeviceID:          "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		Task:              "x",
		Accessibility:     mobilerun.Bool(true),
		Apps:              []string{"string"},
		ContinueOnFailure: mobilerun.Bool(true),
		Credentials: []mobilerun.TaskRunStreamedParamsCredential{{
			CredentialNames: []string{"string"},
			PackageName:     "packageName",
		}},
		DisplayID:        mobilerun.Int(0),
		ExecutionTimeout: mobilerun.Int(1),
		Files:            []string{"string"},
		Humanize:         mobilerun.Bool(true),
		Internal:         mobilerun.Bool(true),
		LlmModel:         mobilerun.String("llmModel"),
		MaxSteps:         mobilerun.Int(0),
		OutputSchema: map[string]any{
			"foo": "bar",
		},
		Reasoning:        mobilerun.Bool(true),
		RecordingEnabled: mobilerun.Bool(true),
		Stealth:          mobilerun.Bool(true),
		SubagentModel:    mobilerun.String("subagentModel"),
		SystemPrompt:     mobilerun.String("systemPrompt"),
		Temperature:      mobilerun.Float(0),
		Vision:           mobilerun.Bool(true),
		VpnCountry:       mobilerun.TaskRunStreamedParamsVpnCountryUs,
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskSendMessage(t *testing.T) {
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
	_, err := client.Tasks.SendMessage(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.TaskSendMessageParams{
			Message: "x",
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

func TestTaskStop(t *testing.T) {
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
	_, err := client.Tasks.Stop(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
