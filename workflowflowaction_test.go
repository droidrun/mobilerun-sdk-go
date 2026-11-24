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

func TestWorkflowFlowActionList(t *testing.T) {
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
	_, err := client.Workflows.Flows.Actions.List(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowFlowActionAddWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.Actions.Add(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.WorkflowFlowActionAddParams{
			ActionID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Position: 0,
			Children: []mobilerun.WorkflowFlowActionAddParamsChild{{
				ActionID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				Position:        0,
				ContinueOnError: mobilerun.Bool(true),
				Key:             mobilerun.String("key"),
				NameOverride:    mobilerun.String("x"),
				OverridesProperty: mobilerun.WorkflowFlowActionAddParamsChildOverrides{
					Params: map[string]any{
						"foo": "bar",
					},
				},
				RecordingEnabled: mobilerun.Bool(true),
			}},
			ContinueOnError: mobilerun.Bool(true),
			Key:             mobilerun.String("key"),
			NameOverride:    mobilerun.String("x"),
			OverridesProperty: mobilerun.WorkflowFlowActionAddParamsOverrides{
				Params: map[string]any{
					"foo": "bar",
				},
			},
			ParentFlowActionID: mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			RecordingEnabled:   mobilerun.Bool(true),
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

func TestWorkflowFlowActionRemove(t *testing.T) {
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
	_, err := client.Workflows.Flows.Actions.Remove(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.WorkflowFlowActionRemoveParams{
			FlowID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestWorkflowFlowActionReplace(t *testing.T) {
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
	_, err := client.Workflows.Flows.Actions.Replace(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.WorkflowFlowActionReplaceParams{
			Actions: []mobilerun.WorkflowFlowActionReplaceParamsAction{{
				ActionID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				Position: 0,
				Children: []mobilerun.WorkflowFlowActionReplaceParamsActionChild{{
					ActionID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					Position:        0,
					ContinueOnError: mobilerun.Bool(true),
					Key:             mobilerun.String("key"),
					NameOverride:    mobilerun.String("x"),
					OverridesProperty: mobilerun.WorkflowFlowActionReplaceParamsActionChildOverrides{
						Params: map[string]any{
							"foo": "bar",
						},
					},
					RecordingEnabled: mobilerun.Bool(true),
				}},
				ContinueOnError: mobilerun.Bool(true),
				Key:             mobilerun.String("key"),
				NameOverride:    mobilerun.String("x"),
				OverridesProperty: mobilerun.WorkflowFlowActionReplaceParamsActionOverrides{
					Params: map[string]any{
						"foo": "bar",
					},
				},
				RecordingEnabled: mobilerun.Bool(true),
			}},
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
