package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test case 1: Valid API Key
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey my-secret-key")
	
	key, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if key != "my-secret-key" {
		t.Errorf("expected my-secret-key, got %s", key)
	}

	// Test case 2: Missing header
	_, err = GetAPIKey(http.Header{})
	if err == nil {
		t.Error("expected an error due to missing header, but got nil")
	}
}
