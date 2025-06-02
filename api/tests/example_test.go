package tests

import (
	"testing"
)

func Add(number1, number2 int) int {
	return number1 + number2
}

func TestExample(t *testing.T) {
	want := 3
	got := Add(1, 2)

	if want != got {
		t.Errorf("want %d but go %d", want, got)
	}
}
