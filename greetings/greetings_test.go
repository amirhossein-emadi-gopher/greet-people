package greetings

import (
	"fmt"
	"slices"
	"testing"
)

var formats = []string{
	"Hi, %s. Welcome!",
	"Great to see you, %s!",
	"Hail, %s! Well met!",
	"Hello, %s! Welcome to my App.",
}

func TestHelloName(t *testing.T) {
	name := "Amirhossein"
	msg, _ := Hello(name)
	var wanted string
	if !slices.ContainsFunc(formats, func(f string) bool {
		wanted = fmt.Sprintf(f, name)
		return wanted == msg
	}) {
		t.Errorf(`Hello("%s") = "%s", want "%s"`, name, msg, wanted)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if err == nil || msg != "" {
		t.Errorf(`Hello("") = "%s", %v, want "", error`, msg, err)
	}
}
