package main

import (
	"net/http"

	"github.com/pluralsight/webservice/Go/controllers"
)

func main() {
	// build with go build .
	// within folder similar to maven in java
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
