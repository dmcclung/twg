package math

import (
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum([]int{10, -2, 3})
	if sum != 11 {
		t.Errorf("Wanted 11 but received %d", sum)
	}
}

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	if sum != 3 {
		t.Errorf("Wanted 3 but received %d", sum)
	}
}
