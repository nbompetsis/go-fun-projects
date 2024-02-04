package arithmetic

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1 + 2 is not equal 3")
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(8, 4) != 4 {
		t.Error("8 - 4 is not equal 4")
	}
}
