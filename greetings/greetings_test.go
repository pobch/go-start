package greetings

import (
	"fmt"
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	inputName := "Pob"
	want := regexp.MustCompile(fmt.Sprintf(`\b%v\b`, inputName))
	msg, err := Hello(inputName)
	if !(want.MatchString(msg)) || err != nil {
		t.Fatalf(`Hello("%v") = %q, %v ... want match for %#q, nil`, inputName, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v ... want "", error`, msg, err)
	}
}
