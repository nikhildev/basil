package humanize_filesize

import (
	"testing"
)

func TestHumanizeFilesize(t *testing.T) {
	tests := []struct {
		name        string
		sizeInBytes int64
		expected    string
	}{
		{
			name:        "zero bytes",
			sizeInBytes: 0,
			expected:    "0 B",
		},
		{
			name:        "negative bytes",
			sizeInBytes: -1,
			expected:    "0 B",
		},
		{
			name:        "500 bytes",
			sizeInBytes: 500,
			expected:    "500 B",
		},
		{
			name:        "2048 bytes",
			sizeInBytes: 2048,
			expected:    "2.00 KB",
		},
		{
			name:        "1 MB",
			sizeInBytes: 1048576,
			expected:    "1.00 MB",
		},
		{
			name:        "1.5 GB",
			sizeInBytes: 1610612736,
			expected:    "1.50 GB",
		},
		{
			name:        "2 TB",
			sizeInBytes: 2199023255552,
			expected:    "2.00 TB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HumanizeFilesize(tt.sizeInBytes)
			if result != tt.expected {
				t.Errorf("HumanizeFilesize(%d) = %s, want %s", tt.sizeInBytes, result, tt.expected)
			}
		})
	}
}
