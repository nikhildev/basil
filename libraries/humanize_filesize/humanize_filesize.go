package humanize_filesize

import (
	"errors"
	"fmt"
)

var ErrNegativeSize = errors.New("size cannot be negative")

// HumanizeFilesize takes a size in bytes and returns a human-readable string
// with an appropriate unit (B, KB, MB, GB, TB).
func HumanizeFilesize(sizeInBytes int64) string {
	if sizeInBytes < 0 {
		return "0 B"
	}

	switch {
	case sizeInBytes >= 1<<40:
		return fmt.Sprintf("%.2f TB", float64(sizeInBytes)/float64(1<<40))
	case sizeInBytes >= 1<<30:
		return fmt.Sprintf("%.2f GB", float64(sizeInBytes)/float64(1<<30))
	case sizeInBytes >= 1<<20:
		return fmt.Sprintf("%.2f MB", float64(sizeInBytes)/float64(1<<20))
	case sizeInBytes >= 1<<10:
		return fmt.Sprintf("%.2f KB", float64(sizeInBytes)/float64(1<<10))
	default:
		return fmt.Sprintf("%d B", sizeInBytes)
	}
}
