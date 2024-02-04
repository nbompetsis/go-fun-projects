package main

import (
	"fmt"
	"testing"
)

func TestJoinWithCommasNilSlice(t *testing.T) {
	var list []string
	result := JoinWithCommas(list)
	want := ""
	if result != want {
		t.Errorf(errorMessage(list, result, want))
	}
}

func TestJoinWithCommasEmptySlice(t *testing.T) {
	list := []string{}
	result := JoinWithCommas(list)
	want := ""
	if result != want {
		t.Errorf(errorMessage(list, result, want))
	}
}

func TestJoinWithCommasOneElement(t *testing.T) {
	list := []string{"apple"}
	result := JoinWithCommas(list)
	want := "apple"
	if result != want {
		t.Errorf(errorMessage(list, result, want))
	}
}

func TestJoinWithCommasTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	result := JoinWithCommas(list)
	want := "apple and orange"
	if result != want {
		t.Errorf(errorMessage(list, result, want))
	}
}

func TestJoinWithCommasThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	result := JoinWithCommas(list)
	want := "apple, orange, and pear"
	if result != want {
		t.Errorf(errorMessage(list, result, want))
	}
}

func errorMessage(list []string, result string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, result, want)
}
