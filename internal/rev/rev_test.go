package rev_test

import (
	"reflect"
	"testing"

	"example.com/revcopy/internal/rev"
)

func TestReverseCopy(t *testing.T) {
	in := []int{1, 2, 3}
	out := rev.Reverse(in)
	if !reflect.DeepEqual(out, []int{3, 2, 1}) {
		t.Fatalf("out=%v", out)
	}
	if !reflect.DeepEqual(in, []int{1, 2, 3}) {
		t.Fatalf("input mutated: %v", in)
	}
}
