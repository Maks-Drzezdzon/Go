package main_test

import (
	"testing"
)

func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 4
	if got != expected {
		t.Errorf("TestAddition Failed Got: %v expected: %v", got, expected)
	}
}
