package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/nikhildev/basil/libraries/humanize_filesize"
)

func main() {
	v := rand.Int31n(1000000)
	humanized, err := humanize_filesize.GetHumanizedFilesize(&v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d bytes = %s\n", v, humanized)
}
