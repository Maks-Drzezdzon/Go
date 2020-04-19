package main

import (
	"net/http"

	"github.com/pluralsight/webservice/Go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
