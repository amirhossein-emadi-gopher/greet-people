package main

import (
	"fmt"

	"github.com/amirhossein-emadi-gopher/greet-people/greetings"
)

func main() {
	message := greetings.Hello("Amirhossein")
	fmt.Println(message)
}
