package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ReynirPY/library-managment-system/internal/handlers"
	"github.com/ReynirPY/library-managment-system/internal/models"
	"github.com/gorilla/mux"
)

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := handlers.FetchBooks()
	if err != nil {
		http.Error(w, "failed to fetch books", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	response, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "failed to convert fetched books to json", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "wrong id data", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	book, err := handlers.FetchBook(id)
	if err != nil {
		http.Error(w, "failed to fetch book", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	response, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "failed to convert fetched book to json", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "wrong id data", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	err = handlers.DeleteBook(id)
	if err != nil {
		if err.Error() == fmt.Sprintf("no book found with id %d", id) {
			http.Error(w, "not such book found", http.StatusNotFound)
			log.Println(err.Error())
			return
		}
		http.Error(w, "failed to delete", http.StatusInternalServerError)
		log.Println(err.Error())
	}

}

func PostBookHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "wrong input", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	log.Println(book)

	err = handlers.InsertBook(book)
	if err != nil {
		http.Error(w, "failed to add book", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	log.Println("book was inserted")

}

func PutBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "wrong id data", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	var book models.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "wrong input", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	log.Println(book)

	err = handlers.UpdateBook(id, book)
	if err != nil {
		http.Error(w, "failed to update book", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	log.Println("book was updated")

}
