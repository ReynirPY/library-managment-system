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

func DeleteBookmark(userID, id int) error {
	query := "DELETE from bookmarks WHERE user_id=$1 AND id=$2"
	result, err := config.DB.Exec(query, userID, id)
	if err != nil {
		return fmt.Errorf("failed to delete %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no book found with id %d", id)
	}

	return nil
}
