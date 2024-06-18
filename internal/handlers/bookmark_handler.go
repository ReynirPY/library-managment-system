package handlers

import (
	"fmt"

	"github.com/ReynirPY/library-managment-system/config"
	"github.com/ReynirPY/library-managment-system/internal/models"
)

func AddBookmark(bookmark models.Bookmark) error {
	_, err := config.DB.NamedExec("INSERT INTO bookmarks (user_id, book_id) VALUES (:user_id, :book_id)", &bookmark)
	if err != nil {
		return fmt.Errorf("error during insert bookmark %w", err)
	}
	return nil
}
