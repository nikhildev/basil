package humanize_filesize

import (
	"errors"
	"testing"
)

func TestGetHumanizedFilesize(t *testing.T) {
	tests := []struct {
		name          string
		size_in_bytes *int32
		expected      string
		wantErr       bool
	}{
		{
			name:          "nil bytes returns error",
			size_in_bytes: nil,
			wantErr:       true,
		},
		{
			name:          "negative bytes returns error",
			size_in_bytes: int32Ptr(-1),
			wantErr:       true,
		},
		{
			name:          "2048 bytes",
			size_in_bytes: int32Ptr(2048),
			expected:      "0.0020 MB",
		},
		{
			name:          "0 bytes",
			size_in_bytes: int32Ptr(0),
			expected:      "0.0000 MB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetHumanizedFilesize(tt.size_in_bytes)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestErrNegativeSize(t *testing.T) {
	v := int32Ptr(-100)
	_, err := GetHumanizedFilesize(v)
	if !errors.Is(err, ErrNegativeSize) {
		t.Errorf("expected ErrNegativeSize, got %v", err)
	}
}

func int32Ptr(n int32) *int32 {
	return &n
}
