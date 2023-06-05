package main_test

import (
	"testing"

	. "github.com/willianpc/playing-with-go"
)

func TestSimpleMath(t *testing.T) {

	tests := []struct {
		Name     string
		Expected int
		Actual   int
	}{
		{"Sum should be 6", 6, SimpleMath(2, 4)},
		{"Sum should be -2", -2, SimpleMath(2, -4)},
		{"Sum should be 10", 10, SimpleMath(6, 4)},
		{"Sum should be 200", 200, SimpleMath(2, 198)},
	}

	for _, test := range tests {
		if test.Actual != test.Expected {
			t.Errorf("expected %d but got %d", test.Expected, test.Actual)
		}
	}

}
