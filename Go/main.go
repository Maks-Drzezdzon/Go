// Licence space

//  what is this package doing

package main

// when creating a package provide a solution to a problem
// focus on the consumer of the package - simple to use - minimize api - encapsulate changes/abstract output while changing internals
// maximize reusability reduce dependencies

// avoid panics and provide error handling
import (
	"net/http"
	// you can use vendored imports that import a library
	// much like this one below
	"github.com/pluralsight/webservice/Go/controllers"
	// however you can make a folder called vendor
	// with a folder structure relative to the import above
	// and paste all the required info there
	// and everything will work as it did before
	// meanwhile letting you edit or have multiple versions of the same library
)

func main() {
	// build with go build .
	// within folder similar to maven in java
	// ToDo look over panics
	// panic("panic message")
	// https://medium.com/swlh/simple-guide-to-panic-handling-and-recovery-in-golang-72d6181ae3e8
	// https://blog.golang.org/defer-panic-and-recover
	// https://medium.com/rungo/defer-panic-and-recover-in-go-689dfa7f8802

	// when you use init in go all import variables will be instantiated first
	// however you dont have controll over what order they will be init
	// then any init in your main will be instantiated after
	// only called during package initialization
	// cant be called
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
