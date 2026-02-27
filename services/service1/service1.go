package main

import (
	"fmt"
	"math/rand"

	"github.com/nikhildev/basil/libraries/humanize_filesize"
)

func formatSize(v int32) string {
	return fmt.Sprintf("%d bytes = %s", v, humanize_filesize.GetHumanizedFilesize(&v))
}

func main() {
	v := rand.Int31n(1000000)
	fmt.Print(formatSize(v))
}
