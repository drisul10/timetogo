package main

import (
	"fmt"
	"andrisul.dev/greetings"
)

func main() {
	message := greetings.Hello("Andri")
	fmt.Println(message)
}