package main

import (
	"fmt"
	"github.com/martoranam/hello_world/helloworlder"
)

type Printable struct {
	text string
}

func main() {
	var helloworld Printable

	helloworld.text = "Empty"
	fmt.Printf(helloworld.text)

	helloworlder.Update(&helloworld.text)
	fmt.Printf("\n%s", helloworld.text)
}
