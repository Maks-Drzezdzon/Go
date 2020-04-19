package controllers

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/pluralsight/webservice/Go/models"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			println("get all data")
			// uc.getAll(w, r)
		case http.MethodPost:
			println("print all data")
			// uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		id, err := strconv.Atoi(matches[1])

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		switch r.Method {
		case http.MethodGet:
			uc.getAll(id, w)
		case http.MethodPost:
			uc.post(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (uc *userController) getAll(id int, w http.ResponseWriter) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc *userController) post(id int, w http.ResponseWriter, r *http.Request) {
	// post
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	// del
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	// get
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
