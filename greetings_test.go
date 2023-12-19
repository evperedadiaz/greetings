package greetings

import (
	"regexp"
	"testing"
)

func TestHelloNameOK(t *testing.T) {
	name := "Pepe"
	nameRegExp := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Pepe")
	if !nameRegExp.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Pepe") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, nameRegExp)
	}
}

func TestHelloEmptyName(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere un string vacío, error`, msg, err)
	}
}
