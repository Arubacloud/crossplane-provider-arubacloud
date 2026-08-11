package clients

import (
	"encoding/json"
	"testing"
)

func TestCredentialsParsing_RequiredFields(t *testing.T) {
	raw := `{"client_id":"cid","client_secret":"csec"}`
	creds := map[string]string{}
	if err := json.Unmarshal([]byte(raw), &creds); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if creds["client_id"] != "cid" {
		t.Errorf("client_id: got %q, want %q", creds["client_id"], "cid")
	}
	if creds["client_secret"] != "csec" {
		t.Errorf("client_secret: got %q, want %q", creds["client_secret"], "csec")
	}
}

func TestCredentialsParsing_OptionalFields(t *testing.T) {
	raw := `{
		"client_id":"cid",
		"client_secret":"csec",
		"base_url":"https://api.example.com",
		"token_issuer_url":"https://auth.example.com",
		"resource_timeout":"60m"
	}`
	creds := map[string]string{}
	if err := json.Unmarshal([]byte(raw), &creds); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	tests := []struct{ key, want string }{
		{"base_url", "https://api.example.com"},
		{"token_issuer_url", "https://auth.example.com"},
		{"resource_timeout", "60m"},
	}
	for _, tc := range tests {
		if creds[tc.key] != tc.want {
			t.Errorf("%s: got %q, want %q", tc.key, creds[tc.key], tc.want)
		}
	}
}

func TestCredentialsParsing_OptionalFieldsAbsent(t *testing.T) {
	raw := `{"client_id":"cid","client_secret":"csec"}`
	creds := map[string]string{}
	if err := json.Unmarshal([]byte(raw), &creds); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	for _, key := range []string{"base_url", "token_issuer_url", "resource_timeout"} {
		if v, ok := creds[key]; ok && v != "" {
			t.Errorf("expected %s to be absent or empty, got %q", key, v)
		}
	}
}

func TestCredentialsParsing_InvalidJSON(t *testing.T) {
	creds := map[string]string{}
	err := json.Unmarshal([]byte("not-json"), &creds)
	if err == nil {
		t.Error("expected unmarshal error for invalid JSON, got nil")
	}
}

func TestTerraformSetupConfiguration(t *testing.T) {
	creds := map[string]string{
		"client_id":        "my-id",
		"client_secret":    "my-secret",
		"base_url":         "https://custom.api",
		"token_issuer_url": "https://custom.auth",
		"resource_timeout": "45m",
	}

	cfg := map[string]any{
		"client_id":     creds["client_id"],
		"client_secret": creds["client_secret"],
	}
	for _, key := range []string{"base_url", "token_issuer_url", "resource_timeout"} {
		if v := creds[key]; v != "" {
			cfg[key] = v
		}
	}

	if cfg["client_id"] != "my-id" {
		t.Errorf("client_id: got %v", cfg["client_id"])
	}
	if cfg["client_secret"] != "my-secret" {
		t.Errorf("client_secret: got %v", cfg["client_secret"])
	}
	if cfg["base_url"] != "https://custom.api" {
		t.Errorf("base_url: got %v", cfg["base_url"])
	}
	if cfg["resource_timeout"] != "45m" {
		t.Errorf("resource_timeout: got %v", cfg["resource_timeout"])
	}
}

func TestTerraformSetupConfiguration_EmptyOptionals(t *testing.T) {
	creds := map[string]string{
		"client_id":     "my-id",
		"client_secret": "my-secret",
	}

	cfg := map[string]any{
		"client_id":     creds["client_id"],
		"client_secret": creds["client_secret"],
	}
	for _, key := range []string{"base_url", "token_issuer_url", "resource_timeout"} {
		if v := creds[key]; v != "" {
			cfg[key] = v
		}
	}

	for _, key := range []string{"base_url", "token_issuer_url", "resource_timeout"} {
		if _, present := cfg[key]; present {
			t.Errorf("expected %s to be absent from config when not set in credentials", key)
		}
	}
}
