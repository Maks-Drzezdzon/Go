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

const (
	baseID = iota
	id     = iota
)
const (
	newIota = iota
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
	// iota will inc every time you ref it again, it resets between constant blocks
	// so newIota will reset back to 0 and not 2 because its in a new const block
	fmt.Println(baseID, id, newIota)

	array := [3]int{1, 2, 3}
	fmt.Println(array)

	// slices are based ontop of arrays
	// any changes made to the arrya or slices will reflect in the other
	// think of slices like references / pointers to arrays
	slice := array[:]
	fmt.Println(slice)
	// slices are not fixed size unlike arrays
	sliceTwo := []int{1, 2, 3}
	fmt.Println(sliceTwo)
	sliceTwo = append(sliceTwo, 4)
	fmt.Println(sliceTwo)
	// slices have the same func as python with slice[start:stop]
	fmt.Println(sliceTwo[:2])

}

func hello(data string) string {
	return data
}
