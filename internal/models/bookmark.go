package models

import "time"

type Bookmark struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	BookID    int       `json:"book_id" db:"book_id"`
	CreatedAt time.Time `json:"created-at" db:"created_at"`
}
