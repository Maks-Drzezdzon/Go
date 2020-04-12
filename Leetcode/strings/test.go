package strings

import (
	"fmt"
	"testing"
)

func main() {
	// https://blog.alexellis.io/golang-writing-unit-tests/
	fmt.Println(DefangIPaddr("1.1.1.1"))
}

func defangIPaddrTest(t *testing.T) {
	got := DefangIPaddr("1.1.1.1")
	var want string = "1[.]1[.]1[.]1"

	if got != want {
		t.Errorf("defangIPaddrTest failed, got: %v, want: %v", got, want)
	} else {
		fmt.Println("defangIPaddrTest Test Passed")
	}

}
