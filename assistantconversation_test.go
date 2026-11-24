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

func TestAssistantConversationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Assistant.Conversations.New(context.TODO(), mobilerun.AssistantConversationNewParams{
		Title:       "x",
		Agent:       mobilerun.String("x"),
		Description: mobilerun.String("description"),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAssistantConversationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Assistant.Conversations.Update(
		context.TODO(),
		"550e8400-e29b-41d4-a716-446655440000",
		mobilerun.AssistantConversationUpdateParams{
			Description: mobilerun.String("description"),
			Pinned:      mobilerun.Bool(true),
			Status:      mobilerun.AssistantConversationUpdateParamsStatusActive,
			Title:       mobilerun.String("x"),
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

func TestAssistantConversationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Assistant.Conversations.List(context.TODO(), mobilerun.AssistantConversationListParams{
		Kind:       mobilerun.AssistantConversationListParamsKindChat,
		Mine:       mobilerun.AssistantConversationListParamsMineTrue,
		WorkflowID: mobilerun.String("x"),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAssistantConversationAbortWithOptionalParams(t *testing.T) {
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
	_, err := client.Assistant.Conversations.Abort(context.TODO(), mobilerun.AssistantConversationAbortParams{
		SessionID:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		ExpectedTurnID: mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAssistantConversationAnswerPermissionWithOptionalParams(t *testing.T) {
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
	_, err := client.Assistant.Conversations.AnswerPermission(context.TODO(), mobilerun.AssistantConversationAnswerPermissionParams{
		PermissionID: "x",
		Response:     mobilerun.AssistantConversationAnswerPermissionParamsResponseOnce,
		SessionID:    mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAssistantConversationAnswerQuestionWithOptionalParams(t *testing.T) {
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
	_, err := client.Assistant.Conversations.AnswerQuestion(context.TODO(), mobilerun.AssistantConversationAnswerQuestionParams{
		Answers: [][]mobilerun.AssistantConversationAnswerQuestionParamsAnswerUnion{{{
			OfAssistantConversationAnswerQuestionsAnswerLabel: &mobilerun.AssistantConversationAnswerQuestionParamsAnswerLabel{
				Label: "x",
			},
		}}},
		QuestionID: "x",
		SessionID:  mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAssistantConversationHistoryWithOptionalParams(t *testing.T) {
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
	_, err := client.Assistant.Conversations.History(
		context.TODO(),
		"550e8400-e29b-41d4-a716-446655440000",
		mobilerun.AssistantConversationHistoryParams{
			Before: mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Limit:  mobilerun.Int(1),
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

func TestAssistantConversationRejectQuestionWithOptionalParams(t *testing.T) {
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
	_, err := client.Assistant.Conversations.RejectQuestion(context.TODO(), mobilerun.AssistantConversationRejectQuestionParams{
		QuestionID: "x",
		SessionID:  mobilerun.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAssistantConversationSendWithOptionalParams(t *testing.T) {
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
	_, err := client.Assistant.Conversations.Send(context.TODO(), mobilerun.AssistantConversationSendParams{
		Message:   "x",
		SessionID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		Agent:     mobilerun.String("agent"),
		Model:     mobilerun.String("x"),
	})
	if err != nil {
		var apierr *mobilerun.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
