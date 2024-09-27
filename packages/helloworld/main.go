package main

import (
	"fmt"

	greeter "basil/packages/libs"
)

func main() {
	fmt.Println(greeter.GetGreeting("Gophers"))
}
