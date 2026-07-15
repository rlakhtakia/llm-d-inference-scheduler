package handlers

import (
	"errors"
	"testing"
)

func TestIsInfrastructureError(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "no pods available in datastore",
			err:      errors.New("inference error: Internal - no pods available in datastore"),
			expected: true,
		},
		{
			name:     "failed to find endpoint candidates",
			err:      errors.New("inference error: ServiceUnavailable - failed to find endpoint candidates for serving the request"),
			expected: true,
		},
		{
			name:     "datastore not synced",
			err:      errors.New("datastore not synced yet"),
			expected: true,
		},
		{
			name:     "not serving",
			err:      errors.New("EPP is not serving requests"),
			expected: true,
		},
		{
			name:     "client bad request payload",
			err:      errors.New("invalid request body JSON format"),
			expected: false,
		},
		{
			name:     "model not found policy error",
			err:      errors.New("model 'llama3' not authorized for user"),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isInfrastructureError(tc.err)
			if got != tc.expected {
				t.Errorf("isInfrastructureError(%v) = %v, want %v", tc.err, got, tc.expected)
			}
		})
	}
}
