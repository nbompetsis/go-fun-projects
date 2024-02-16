package greetings

import (
	"regexp"
	"testing"
)

func TestHelloFunction(t *testing.T) {
	name := "Nikolas"
	msg, err := HelloFunction(name)
	want := regexp.MustCompile(name)
	if !want.MatchString("Nikolas") || err != nil {
		t.Fatalf("Hello(\"Nikolas\") = %q, %v, want match for %#q, nil", msg, err, want)
	}
}

func TestHelloFunctionEmpty(t *testing.T) {
	msg, err := HelloFunction("")
	if msg != "" || err == nil {
		t.Fatalf("Hello(\"\") = %q, %v, want \"\", error", msg, err)
	}
}
