package main

import (
	"fmt"
	"math/rand"

	"github.com/nikhildev/basil/libraries/humanize_filesize"
)

func main() {
	v := rand.Int31n(1000000)
	fmt.Printf(`%d bytes = %s`, v, humanize_filesize.GetHumanizedFilesize(&v))
}
