package auth_test

import (
	"net/http"
	"testing"
	"errors"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		want      string
		expectErr error
	}{
		{
			name:      "Valid API Key",
			headers:   http.Header{"Authorization": {"ApiKey 12345"}},
			want:      "12345",
			expectErr: nil,
		},
		{
			name:      "No Authorization Header",
			headers:   http.Header{},
			want:      "",
			expectErr: auth.ErrNoAuthHeaderIncluded,
		},
		{
			name:      "Malformed Authorization Header",
			headers:   http.Header{"Authorization": {"Bearer 12345"}},
			want:      "",
			expectErr: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := auth.GetAPIKey(tt.headers)
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}
			if err != nil && err.Error() != tt.expectErr.Error() {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.expectErr)
			}
		})
	}
}
