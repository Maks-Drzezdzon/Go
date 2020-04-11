package main

import (
	"fmt"
)

/*
unused code functions will case errors in go

# pay attention to methods too #
lowercase var is private
- name

uppercase var is public
- Name

blockcase var is block, not visable outside of block of code
- NAME

keep acronyms uppercase
- dbHTTP over dbHttp
- dbURL over dbUrl

*/

var (
	// init new vars and shows they can be related but that is a design decision
	name     string = "maks"
	lastName        = "d"
	age      int    = 20
	// cant use := in global variable scopes
)

func main() {
	/* This is my first sample program. */
	hello(name)
	//TestHello()
	// for types use %T and %v for value
	age := 24
	fmt.Printf("%v is now a %T and was an %T", age, float32(age), age)

}

func hello(data string) string {
	return data
}
