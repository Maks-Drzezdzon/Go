package main_test

import (
	"testing"
)

// every test has to start with Test
// add _test to file name for blackbox test
// pass in a pointer for *testing
// whitebox tests are in the same package
// ToDo look up go test flag docs
// use 'go test' to run tests
// go test -cover to see test coverage
// or go test -coverprofile cover.out
// to inspect you need to use
// go test cover -html or -func cover.out
// -covermode count for more metrics
func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 4
	if got != expected {
		t.Errorf("TestAddition Failed Got: %v expected: %v", got, expected)
	}
}
