package validator

import (
	"net/http/httptest"
	"testing"
)

func TestValidateToken(t *testing.T) {
	tests := []struct {
		name        string
		headerValue string
		expectedErr string
	}{
		{
			name:        "Valid token",
			headerValue: "Bearer validtoken",
			expectedErr: "",
		},
		{
			name:        "Missing authorization header",
			headerValue: "",
			expectedErr: "authorization header is missing",
		},
		{
			name:        "Invalid authorization header format",
			headerValue: "Bearer",
			expectedErr: "invalid Authorization header format",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			req.Header.Set("Authorization", test.headerValue)

			err := ValidateToken(req, "validtoken")

			if test.expectedErr == "" {
				if err != nil {
					t.Errorf("Expected no error, got: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("Expected error '%s', got nil", test.expectedErr)
				} else if err.Error() != test.expectedErr {
					t.Errorf("Expected error '%s', got '%s'", test.expectedErr, err.Error())
				}
			}
		})
	}
}
