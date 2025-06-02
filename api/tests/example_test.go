package tests

import (
	"testing"
)

func Add(number1, number2 int) int {
	return number1 + number2
}

func TestExample(t *testing.T) {
	wanted := 3
	got := Add(1, 2)

	if wanted != got {
		t.Errorf("want %d but go %d", wanted, got)
	}
}
