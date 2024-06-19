package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ReynirPY/library-managment-system/internal/handlers"
	"github.com/ReynirPY/library-managment-system/internal/models"
	"github.com/gorilla/mux"
)

func AddBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)

	var bookmark models.Bookmark
	err := json.NewDecoder(r.Body).Decode(&bookmark)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookmark.UserID = user.ID

	err = handlers.AddBookmark(bookmark)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "failed to insert bookmark", http.StatusInternalServerError)
		return
	}
}

func DeleteBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "wrong id data", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	err = handlers.DeleteBookmark(user.ID, id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "failed to delete bookmark", http.StatusInternalServerError)
		return
	}

}
