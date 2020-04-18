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
	// TestHello()
	// for types use %T and %v for value
	age := 24
	fmt.Printf("%v is now a %T and was an %T", age, float32(age), age)

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// go also has pointers
	// pointers are declared by the data type *string *int
	// they have to me init with what type they will be so string
	var firstName *string = new(string)
	// now this can be assigned a val
	*firstName = "Maks"
	// you need  to ref the pointer again ofc with * to get the data not the address
	fmt.Println(*firstName)

	lastName := "D"
	fmt.Println(lastName)

}

func hello(data string) string {
	return data
}
