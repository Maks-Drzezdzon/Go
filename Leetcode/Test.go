package main

import (
	"fmt"
	"testing"
)

var (
	// init new vars and shows they can be related but that is a design decision
	name     string = "maks"
	lastName string = "d"
	age      int    = 24
)

func main() {
	/* This is my first sample program. */
	Hello(name)
	//TestHello()
	// for types use %T
	fmt.Printf("%T", age)

}

func Hello(data string) string {
	return data
}

func TestHello(t *testing.T) {
	got := Hello("")
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
