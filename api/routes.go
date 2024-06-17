package api

import (
	"fmt"
	"net/http"

	"github.com/ReynirPY/library-managment-system/internal/auth"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	// creating router with gorilla mux
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "helloworld")
	})

	//=====book actions routes=====

	//public book routes
	r.HandleFunc("/books", GetBooksHandler).Methods("GET")
	r.HandleFunc("/books/{id}", GetBookHandler).Methods("GET")

	//admin only routes
	books := r.PathPrefix("/books").Subrouter()
	books.Use(auth.JWTMiddlewareAdmin)
	books.HandleFunc("", PostBookHandler).Methods("POST")
	books.HandleFunc("/{id}", DeleteBookHandler).Methods("DELETE")
	books.HandleFunc("/{id}", PutBookHandler).Methods("PUT")

	//auth user routes
	//r.HandleFunc("/user/registration", RegistrationHandler).Methods("POST")
	//r.HandleFunc("/user/login", LoginHandler).Methods("POST")

	//auth admin routes
	r.HandleFunc("/admins/sign-up/", AdminRegistrationHandler).Methods("POST")
	r.HandleFunc("/admins/sign-in/", AdminLoginHandler).Methods("POST")
	return r
}
