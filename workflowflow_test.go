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

func TestWorkflowFlowNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.New(context.TODO(), mobilerun.WorkflowFlowNewParams{
		Name:      "x",
		TriggerID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		Actions: []mobilerun.WorkflowFlowNewParamsAction{{
			ActionID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Position: 0,
			Children: []mobilerun.WorkflowFlowNewParamsActionChild{{
				ActionID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				Position:        0,
				ContinueOnError: mobilerun.Bool(true),
				Key:             mobilerun.String("key"),
				NameOverride:    mobilerun.String("x"),
				OverridesProperty: mobilerun.WorkflowFlowNewParamsActionChildOverrides{
					Params: map[string]any{
						"foo": "bar",
					},
				},
				RecordingEnabled: mobilerun.Bool(true),
			}},
			ContinueOnError: mobilerun.Bool(true),
			Key:             mobilerun.String("key"),
			NameOverride:    mobilerun.String("x"),
			OverridesProperty: mobilerun.WorkflowFlowNewParamsActionOverrides{
				Params: map[string]any{
					"foo": "bar",
				},
			},
			RecordingEnabled: mobilerun.Bool(true),
		}},
		CooldownScope:   mobilerun.WorkflowFlowNewParamsCooldownScopeFlow,
		CooldownSeconds: mobilerun.Int(0),
		Delivery: mobilerun.WorkflowFlowNewParamsDelivery{
			Destination: "mobilerun",
			Folder:      mobilerun.String("x"),
			Recording: mobilerun.WorkflowFlowNewParamsDeliveryRecording{
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
		OwnerWake:               mobilerun.WorkflowFlowNewParamsOwnerWakeFinal,
		RecordingEnabled:        mobilerun.Bool(true),
		RecordingPolicy: mobilerun.WorkflowFlowNewParamsRecordingPolicy{
			Mode: "off",
		},
		SaveSandboxLogs:        mobilerun.Bool(true),
		SelfHealingEnabled:     mobilerun.Bool(true),
		SelfHealingMaxAttempts: mobilerun.Int(1),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowFlowGet(t *testing.T) {
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
	_, err := client.Workflows.Flows.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowFlowUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.WorkflowFlowUpdateParams{
			CooldownScope:   mobilerun.WorkflowFlowUpdateParamsCooldownScopeFlow,
			CooldownSeconds: mobilerun.Int(0),
			Delivery: mobilerun.WorkflowFlowUpdateParamsDelivery{
				Destination: "mobilerun",
				Folder:      mobilerun.String("x"),
				Recording: mobilerun.WorkflowFlowUpdateParamsDeliveryRecording{
					Filename: "x",
				},
				Screenshots: map[string]any{},
			},
			Description:             mobilerun.String("description"),
			DeviceIDs:               []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			Enabled:                 mobilerun.Bool(true),
			HealthMonitoringEnabled: mobilerun.Bool(true),
			LifecycleStatus:         mobilerun.WorkflowFlowUpdateParamsLifecycleStatusEnabled,
			Name:                    mobilerun.String("x"),
			NotifyOnFailure:         mobilerun.Bool(true),
			NotifyOnSuccess:         mobilerun.Bool(true),
			NotifyWebhookID:         mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			OwnerWake:               mobilerun.WorkflowFlowUpdateParamsOwnerWakeFinal,
			RecordingEnabled:        mobilerun.Bool(true),
			RecordingPolicy: mobilerun.WorkflowFlowUpdateParamsRecordingPolicy{
				Mode: "off",
			},
			SaveSandboxLogs:        mobilerun.Bool(true),
			SelfHealingEnabled:     mobilerun.Bool(true),
			SelfHealingMaxAttempts: mobilerun.Int(1),
			TriggerID:              mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestWorkflowFlowListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.List(context.TODO(), mobilerun.WorkflowFlowListParams{
		AgentID:          mobilerun.WorkflowFlowListParamsAgentIDUnassigned,
		CreatedBy:        mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Enabled:          mobilerun.WorkflowFlowListParamsEnabledTrue,
		Mine:             mobilerun.WorkflowFlowListParamsMineTrue,
		OrderBy:          mobilerun.WorkflowFlowListParamsOrderByName,
		OrderByDirection: mobilerun.WorkflowFlowListParamsOrderByDirectionAsc,
		Page:             mobilerun.Int(1),
		PageSize:         mobilerun.Int(1),
		Search:           mobilerun.String("x"),
		Status:           []string{"healthy"},
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

func TestWorkflowFlowDelete(t *testing.T) {
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
	_, err := client.Workflows.Flows.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowFlowActivate(t *testing.T) {
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
	_, err := client.Workflows.Flows.Activate(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.WorkflowFlowActivateParams{
			VerificationExecutionID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestWorkflowFlowCapacityWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.Capacity(context.TODO(), mobilerun.WorkflowFlowCapacityParams{
		AgentID: mobilerun.WorkflowFlowCapacityParamsAgentIDUnassigned,
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowFlowCloneWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.Clone(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.WorkflowFlowCloneParams{
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

func TestWorkflowFlowNewFromDocumentWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.NewFromDocument(context.TODO(), mobilerun.WorkflowFlowNewFromDocumentParams{
		Document: map[string]any{},
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowFlowDeliveryOptions(t *testing.T) {
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
	_, err := client.Workflows.Flows.DeliveryOptions(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowFlowDocumentSchemaWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.DocumentSchema(context.TODO(), mobilerun.WorkflowFlowDocumentSchemaParams{
		Profile: mobilerun.WorkflowFlowDocumentSchemaParamsProfileAgentRoutine,
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowFlowDryRunWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.DryRun(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.WorkflowFlowDryRunParams{
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

func TestWorkflowFlowListRepairs(t *testing.T) {
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
	_, err := client.Workflows.Flows.ListRepairs(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowFlowReplaceDocumentWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.ReplaceDocument(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.WorkflowFlowReplaceDocumentParams{
			Document: map[string]any{},
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

func TestWorkflowFlowGetDocument(t *testing.T) {
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
	_, err := client.Workflows.Flows.GetDocument(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowFlowGetRepair(t *testing.T) {
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
	_, err := client.Workflows.Flows.GetRepair(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.WorkflowFlowGetRepairParams{
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

func TestWorkflowFlowRunWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.Run(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.WorkflowFlowRunParams{
			CaptureScreenshots: mobilerun.Bool(true),
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

func TestWorkflowFlowTemplateContextWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.TemplateContext(context.TODO(), mobilerun.WorkflowFlowTemplateContextParams{
		TemplateResolutionVersion: mobilerun.Int(2),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowFlowUnblock(t *testing.T) {
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
	_, err := client.Workflows.Flows.Unblock(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowFlowValidateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.Validate(context.TODO(), mobilerun.WorkflowFlowValidateParams{
		OfObject: &mobilerun.WorkflowFlowValidateParamsBodyObject{
			Actions: []mobilerun.WorkflowFlowValidateParamsBodyObjectActionUnion{{
				OfWorkflowFlowValidatesBodyObjectActionFlowActionInput: &mobilerun.WorkflowFlowValidateParamsBodyObjectActionFlowActionInput{
					ActionID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					Position: 0,
					Children: []mobilerun.WorkflowFlowValidateParamsBodyObjectActionFlowActionInputChild{{
						ActionID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
						Position:        0,
						ContinueOnError: mobilerun.Bool(true),
						Key:             mobilerun.String("key"),
						NameOverride:    mobilerun.String("x"),
						OverridesProperty: mobilerun.WorkflowFlowValidateParamsBodyObjectActionFlowActionInputChildOverrides{
							Params: map[string]any{
								"foo": "bar",
							},
						},
						RecordingEnabled: mobilerun.Bool(true),
					}},
					ContinueOnError: mobilerun.Bool(true),
					Key:             mobilerun.String("key"),
					NameOverride:    mobilerun.String("x"),
					OverridesProperty: mobilerun.WorkflowFlowValidateParamsBodyObjectActionFlowActionInputOverrides{
						Params: map[string]any{
							"foo": "bar",
						},
					},
					RecordingEnabled: mobilerun.Bool(true),
				},
			}},
			Name:            "x",
			CooldownScope:   "flow",
			CooldownSeconds: mobilerun.Int(0),
			Delivery: mobilerun.WorkflowFlowValidateParamsBodyObjectDelivery{
				Destination: "mobilerun",
				Folder:      mobilerun.String("x"),
				Recording: mobilerun.WorkflowFlowValidateParamsBodyObjectDeliveryRecording{
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
			OwnerWake:               "final",
			RecordingEnabled:        mobilerun.Bool(true),
			RecordingPolicy: mobilerun.WorkflowFlowValidateParamsBodyObjectRecordingPolicy{
				Mode: "off",
			},
			SaveSandboxLogs:        mobilerun.Bool(true),
			SelfHealingEnabled:     mobilerun.Bool(true),
			SelfHealingMaxAttempts: mobilerun.Int(1),
			TriggerID:              mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowFlowValidateDocumentWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.ValidateDocument(context.TODO(), mobilerun.WorkflowFlowValidateDocumentParams{
		Document: map[string]any{},
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowFlowVerifyWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Flows.Verify(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mobilerun.WorkflowFlowVerifyParams{
			InvocationID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			DeviceID:     mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
