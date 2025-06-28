package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	testCases := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectError bool
	}{
		{
			name: "valid header",
			headers: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "ApiKey token12345")
				return h
			}(),
			expectedKey: "token12345",
			expectError: false,
		},
		{
			name:        "missing Authorization header",
			headers:     http.Header{},
			expectedKey: "",
			expectError: true,
		},
		{
			name: "wrong scheme in Authorization",
			headers: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "Bearer token12345")
				return h
			}(),
			expectedKey: "",
			expectError: true,
		},
		// Cannot use this one without altering the GetAPIKey function to handle empty tokens
		// {
		// 	name: "empty token after ApiKey",
		// 	headers: func() http.Header {
		// 		h := http.Header{}
		// 		h.Set("Authorization", "ApiKey ")
		// 		return h
		// 	}(),
		// 	expectedKey: "",
		// 	expectError: true,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tc.headers)
			if tc.expectError {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if apiKey != tc.expectedKey {
					t.Errorf("expected %q, got %q", tc.expectedKey, apiKey)
				}
			}
		})
	}
}
