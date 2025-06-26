package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		header     http.Header
		wantKey    string
		wantErrMsg string
	}{
		{
			name:       "missing Authorization header",
			header:     http.Header{},
			wantErrMsg: "authorization header missing",
		},
		{
			name: "wrong prefix",
			header: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			wantErrMsg: "authorization header must start with 'ApiKey '",
		},
		{
			name: "empty key",
			header: http.Header{
				"Authorization": []string{"ApiKey "},
			},
			wantErrMsg: "token key is empty",
		},
		{
			name: "valid key",
			header: http.Header{
				"Authorization": []string{"ApiKey abc123"},
			},
			wantKey: "abc123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.header)

			if tt.wantErrMsg != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErrMsg) {
					t.Errorf("expected error '%s', got '%v'", tt.wantErrMsg, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if key != tt.wantKey {
				t.Errorf("expected key '%s', got '%s'", tt.wantKey, key)
			}
		})
	}
}
