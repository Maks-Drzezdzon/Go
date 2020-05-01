package strings

import (
	"fmt"
	"testing"
)

func main() {
	// https://blog.alexellis.io/golang-writing-unit-tests/
	fmt.Println(DefangIPaddr("1.1.1.1"))
	fmt.Println(NumJewelsInStones("aA", "aAAbbbb"))
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

func NumJewelsInStonesTest(t *testing.T) {
	got := NumJewelsInStones("aA", "aAAbbbb")
	var want int = 3

	if got != want {
		t.Errorf("NumJewelsInStonesTest failed, got: %v, want: %v", got, want)
	} else {
		fmt.Println("NumJewelsInStonesTest Test Passed")
	}

}
