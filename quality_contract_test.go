package onemsgsdk_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	sdk "github.com/1msg/1msg-sdk-go"
)

// TestSendMessageWireContract exercises the published API against an isolated local server.
// It verifies path, token encoding, required body, optional omission and response decoding.
func TestSendMessageWireContract(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.Method != http.MethodPost || r.URL.Path != "/fixture-instance/sendMessage" {
			t.Errorf("unexpected operation: %s %s", r.Method, r.URL.Path)
		}
		if r.URL.Query().Get("token") != "fixture+token&value" {
			t.Errorf("token did not round-trip")
		}
		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}
		if r.PostForm.Get("body") != "Hello 👋" || r.PostForm.Get("chatId") != "12020721369@c.us" {
			t.Errorf("wrong request body: %v", r.PostForm)
		}
		if r.PostForm.Has("phone") || r.PostForm.Has("quotedMsgId") {
			t.Errorf("unset optional fields were serialized")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"sent":true,"id":"fixture-message"}`))
	}))
	defer server.Close()
	config := sdk.NewConfiguration()
	config.Servers = sdk.ServerConfigurations{{URL: server.URL + "/fixture-instance"}}
	client := sdk.NewAPIClient(config)
	result, response, err := client.MessagingAPI.SendMessage(context.Background()).Token("fixture+token&value").Body("Hello 👋").ChatId("12020721369@c.us").Execute()
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if !result.GetSent() || result.GetId() != "fixture-message" || requests != 1 {
		t.Fatalf("unexpected response: %#v; requests=%d", result, requests)
	}
	_, _, err = client.MessagingAPI.SendMessage(context.Background()).Body("missing token").Execute()
	if err == nil || requests != 1 {
		t.Fatal("missing token must fail before transport")
	}
}

// TestRequiredAndOptionalModelContract proves required fields reject missing data and optional fields stay absent.
func TestRequiredAndOptionalModelContract(t *testing.T) {
	var request sdk.SendMessageRequest
	if err := json.Unmarshal([]byte(`{}`), &request); err == nil {
		t.Fatal("missing required body accepted")
	}
	if err := json.Unmarshal([]byte(`{"body":"hello","chatId":"12020721369@c.us"}`), &request); err != nil {
		t.Fatal(err)
	}
	encoded, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	var payload map[string]interface{}
	if err = json.Unmarshal(encoded, &payload); err != nil {
		t.Fatal(err)
	}
	if payload["body"] != "hello" || payload["chatId"] != "12020721369@c.us" {
		t.Fatalf("field names changed: %s", encoded)
	}
	if _, exists := payload["phone"]; exists {
		t.Fatal("unset phone must be omitted")
	}
}
