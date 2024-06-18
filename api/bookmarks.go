package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ReynirPY/library-managment-system/internal/handlers"
	"github.com/ReynirPY/library-managment-system/internal/models"
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
