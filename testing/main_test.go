package main

import "testing"

func TestAdd(t *testing.T) {
	// arrange
	left, right := 1, 3
	expect := 4
	// act
	got := Add(left, right)
	// assert
	if expect != got {
		t.Errorf("Failed to add %v and %v", left, right)
	}
}
