package main

import (
	"fmt"
	"log"

	"github.com/amirhossein-emadi-gopher/greet-people/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Amirhossein", "Ali", "Amir", "Amin", "Armin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
