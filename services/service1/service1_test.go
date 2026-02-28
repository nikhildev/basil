package main

import (
	"strings"
	"testing"
)

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name        string
		input       int32
		wantPrefix  string
		wantContain string
	}{
		{
			name:        "zero bytes",
			input:       0,
			wantPrefix:  "0 bytes = ",
			wantContain: "MB",
		},
		{
			name:        "2048 bytes",
			input:       2048,
			wantPrefix:  "2048 bytes = ",
			wantContain: "MB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatSize(tt.input)
			if !strings.HasPrefix(result, tt.wantPrefix) {
				t.Errorf("formatSize(%d) = %q, want prefix %q", tt.input, result, tt.wantPrefix)
			}
			if !strings.Contains(result, tt.wantContain) {
				t.Errorf("formatSize(%d) = %q, want to contain %q", tt.input, result, tt.wantContain)
			}
		})
	}
}
