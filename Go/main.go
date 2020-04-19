package main

import (
	"net/http"

	"github.com/pluralsight/webservice/Go/controllers"
)

func main() {
	// build with go build .
	// within folder similar to maven in java
	// ToDo look over panics
	// panic("panic message")
	// https://medium.com/swlh/simple-guide-to-panic-handling-and-recovery-in-golang-72d6181ae3e8
	// https://blog.golang.org/defer-panic-and-recover
	// https://medium.com/rungo/defer-panic-and-recover-in-go-689dfa7f8802

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
