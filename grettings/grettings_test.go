package grettings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Jonh Nino"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Jonh Nino")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Jonh Nino") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
