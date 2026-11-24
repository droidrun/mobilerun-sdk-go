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

func TestAgentRoutineNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Routines.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.AgentRoutineNewParams{
			Actions: []mobilerun.AgentRoutineNewParamsAction{{
				ActionID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				Position: 0,
			}},
			Name:            "x",
			TriggerID:       "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			CooldownScope:   mobilerun.AgentRoutineNewParamsCooldownScopeFlow,
			CooldownSeconds: mobilerun.Int(0),
			Delivery: mobilerun.AgentRoutineNewParamsDelivery{
				Destination: "one_drive",
				Folder:      mobilerun.String("x"),
				Recording: mobilerun.AgentRoutineNewParamsDeliveryRecording{
					Filename: "x",
				},
				Screenshots: map[string]any{},
			},
			Description:             mobilerun.String("description"),
			DeviceIDs:               []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			Enabled:                 mobilerun.Bool(true),
			HealthMonitoringEnabled: mobilerun.Bool(true),
			NotifyOnFailure:         mobilerun.Bool(true),
			NotifyOnSuccess:         mobilerun.Bool(true),
			NotifyWebhookID:         mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			RecordingEnabled:        mobilerun.Bool(true),
			RecordingPolicy: mobilerun.AgentRoutineNewParamsRecordingPolicy{
				Mode: "off",
			},
			SelfHealingEnabled:     mobilerun.Bool(true),
			SelfHealingMaxAttempts: mobilerun.Int(1),
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

func TestAgentRoutineCloneWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Routines.Clone(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.AgentRoutineCloneParams{
			AgentID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			DeviceIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			Name:      mobilerun.String("x"),
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
