package main

import "fmt"

type Printable struct {
	text string
}

func update(target *string) {
	*target += "World!"
}

func main() {
	var helloworld Printable

	helloworld.text = "Hello "
	fmt.Printf(helloworld.text)

	update(&helloworld.text)
	fmt.Printf("\n%s", helloworld.text)
}
