package main

import (
	"log/slog"
	"math/rand"
	"os"

	"github.com/nikhildev/basil/libraries/humanize_filesize"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	v := rand.Int31n(1000000)
	humanized := humanize_filesize.GetHumanizedFilesize(&v)

	logger.Info("generated random filesize",
		"bytes", v,
		"humanized", humanized,
	)
}
