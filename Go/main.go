package main

import (
	"errors"
	"fmt"

	"github.com/pluralsight/webservice/Go/models"
)

func main() error {
	u := models.User{ID: 2, FirstName: "m", LastName: "d"}
	fmt.Println(u)
	return errors.New("failed to run main")
}
