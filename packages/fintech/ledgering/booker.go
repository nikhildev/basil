package main

import (
	lib "basil/packages/fintech/invoicing"
	"fmt"
)

func main() {
	i := lib.GetNewInvoice()
	fmt.Println("Generate new Invoice with ID:", i.ID)
}
