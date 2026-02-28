package main

import (
	"fmt"
	"math/rand"

	"github.com/nikhildev/basil/libraries/humanize_filesize"
)

func main() {
	v := rand.Int63n(1000000)
	fmt.Printf("%d bytes = %s\n", v, humanize_filesize.HumanizeFilesize(v))
}
