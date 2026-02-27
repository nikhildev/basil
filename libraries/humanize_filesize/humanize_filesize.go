package humanize_filesize

import (
	"errors"
	"fmt"
)

var ErrNegativeSize = errors.New("size cannot be negative")

// GetHumanizedFilesize takes size_in_bytes as an int32 pointer and returns the size in megabytes.
// Returns an error if the pointer is nil or the value is negative.
func GetHumanizedFilesize(size_in_bytes *int32) (string, error) {
	if size_in_bytes == nil {
		return "", errors.New("size_in_bytes must not be nil")
	}
	if *size_in_bytes < 0 {
		return "", ErrNegativeSize
	}
	size_in_megabytes := float64(*size_in_bytes) / (1024 * 1024)
	return fmt.Sprintf("%.4f MB", size_in_megabytes), nil
}
