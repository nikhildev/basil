package main

import (
	"log/slog"
	"math/rand"
	"os"

	"github.com/nikhildev/basil/libraries/humanize_filesize"
)

func formatSize(v int32) string {
	return fmt.Sprintf("%d bytes = %s", v, humanize_filesize.GetHumanizedFilesize(&v))
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	v := rand.Int31n(1000000)
	humanized := humanize_filesize.GetHumanizedFilesize(&v)

	logger.Info("generated random filesize",
		"bytes", v,
		"humanized", humanized,
	)
}
