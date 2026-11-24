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

func TestWorkflowTriggerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Triggers.New(context.TODO(), mobilerun.WorkflowTriggerNewParams{
		Activation: mobilerun.WorkflowTriggerNewParamsActivationEvent,
		Name:       "x",
		Conditions: mobilerun.WorkflowTriggerNewParamsConditions{
			All: []any{map[string]any{}},
			Any: []any{map[string]any{}},
		},
		CustomPayloadSchema: map[string]any{
			"foo": "bar",
		},
		Description: mobilerun.String("description"),
		EventType:   mobilerun.String("eventType"),
		Payload: map[string]any{
			"foo": "bar",
		},
		ScheduleRule: mobilerun.WorkflowTriggerNewParamsScheduleRule{
			Type:       "once",
			DateTime:   mobilerun.String("dateTime"),
			Expression: mobilerun.String("expression"),
			Jitter: mobilerun.WorkflowTriggerNewParamsScheduleRuleJitter{
				AfterMinutes:  mobilerun.Int(0),
				BeforeMinutes: mobilerun.Int(0),
			},
			Rrule: mobilerun.String("rrule"),
		},
		Timezone: mobilerun.String("timezone"),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowTriggerGet(t *testing.T) {
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
	_, err := client.Workflows.Triggers.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowTriggerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Triggers.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.WorkflowTriggerUpdateParams{
			Activation: mobilerun.WorkflowTriggerUpdateParamsActivationEvent,
			Conditions: mobilerun.WorkflowTriggerUpdateParamsConditions{
				All: []any{map[string]any{}},
				Any: []any{map[string]any{}},
			},
			CustomPayloadSchema: map[string]any{
				"foo": "bar",
			},
			Description: mobilerun.String("description"),
			EventType:   mobilerun.String("eventType"),
			Name:        mobilerun.String("x"),
			Payload: map[string]any{
				"foo": "bar",
			},
			ScheduleRule: mobilerun.WorkflowTriggerUpdateParamsScheduleRule{
				Type:       "once",
				DateTime:   mobilerun.String("dateTime"),
				Expression: mobilerun.String("expression"),
				Jitter: mobilerun.WorkflowTriggerUpdateParamsScheduleRuleJitter{
					AfterMinutes:  mobilerun.Int(0),
					BeforeMinutes: mobilerun.Int(0),
				},
				Rrule: mobilerun.String("rrule"),
			},
			Timezone: mobilerun.String("timezone"),
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

func TestWorkflowTriggerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Triggers.List(context.TODO(), mobilerun.WorkflowTriggerListParams{
		Activation:       mobilerun.WorkflowTriggerListParamsActivationEvent,
		EventType:        mobilerun.String("eventType"),
		OrderBy:          mobilerun.WorkflowTriggerListParamsOrderByName,
		OrderByDirection: mobilerun.WorkflowTriggerListParamsOrderByDirectionAsc,
		Page:             mobilerun.Int(1),
		PageSize:         mobilerun.Int(1),
		Search:           mobilerun.String("x"),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowTriggerDelete(t *testing.T) {
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
	_, err := client.Workflows.Triggers.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowTriggerFire(t *testing.T) {
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
	_, err := client.Workflows.Triggers.Fire(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.WorkflowTriggerFireParams{
			Payload: map[string]any{
				"foo": "bar",
			},
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
