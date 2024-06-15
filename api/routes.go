package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "helloworld")
	})
	//book actions routes
	r.HandleFunc("/books", GetBooksHandler).Methods("GET")
	r.HandleFunc("/books", PostBookHandler).Methods("POST")
	r.HandleFunc("/books/{id}", GetBookHandler).Methods("GET")
	r.HandleFunc("/books/{id}", DeleteBookHandler).Methods("DELETE")
	r.HandleFunc("/books/{id}", PutBookHandler).Methods("PUT")

	//auth routes
	r.HandleFunc("/registration", RegistrationHandler).Methods("POST")

	return r
}
