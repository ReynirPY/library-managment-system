package handlers

import (
	"fmt"

	"github.com/ReynirPY/library-managment-system/config"
	"github.com/ReynirPY/library-managment-system/internal/models"
)

func FetchBooks() ([]*models.Book, error) {
	var books []*models.Book
	err := config.DB.Select(&books, `SELECT id, title, author, "year", isbn
	FROM book;`)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch books %w", err)
	}

	return books, nil
}

func FetchBook(id int) (*models.Book, error) {
	var book models.Book
	fetchBookStr := `SELECT id, title, author, "year", isbn
	FROM book WHERE id=$1;`
	row := config.DB.QueryRow(fetchBookStr, id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.Isbn)
	if err != nil {
		return nil, fmt.Errorf("failed to take row %w", err)
	}
	return &book, nil
}

func DeleteBook(id int) error {
	deleteBookStr := `DELETE FROM public.book
	WHERE id=$1;`
	result, err := config.DB.Exec(deleteBookStr, id)
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

func InsertBook(book models.Book) error {
	insertBookStr := `INSERT INTO book
	(title, author, "year", isbn)
	VALUES(:title, :author, :year,:isbn);`
	_, err := config.DB.NamedExec(insertBookStr, &book)

	if err != nil {
		return fmt.Errorf("failed to incert book %w", err)
	}

	return nil
}

func UpdateBook(id int, book models.Book) error {
	updateBookStr := `UPDATE book SET title=$1, author=$2, year=$3, isbn=$4 WHERE id= $5`
	_, err := config.DB.Exec(updateBookStr, book.Title, book.Author, book.Year, book.Isbn, id)
	if err != nil {
		return fmt.Errorf("failed to update book %w", err)
	}

	return nil
}
