package main

import "testing"

var Debug bool = false

func TestIsOne(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("Skip!!")
	}
	v := IsONe(i)

	if !v {
		t.Error("%v != %v ", i, 1)
	}
}
