package handlers

import (
	"fmt"
	"strconv"

	"github.com/ReynirPY/library-managment-system/config"
	"github.com/ReynirPY/library-managment-system/internal/models"
)

func FetchBooks(author string, year int, isbn string) ([]*models.Book, error) {
	query := `SELECT id, title, author, "year", isbn
	FROM book WHERE 1=1`
	var books []*models.Book
	args := []interface{}{}
	argIndex := 1

	if author != "" {
		query += " AND author=$" + strconv.Itoa(argIndex)
		args = append(args, author)
		argIndex++
	}

	if year != 0 {
		query += " AND year=$" + strconv.Itoa(argIndex)
		args = append(args, year)
		argIndex++
	}

	if isbn != "" {
		query += " AND isbn=$" + strconv.Itoa(argIndex)
		args = append(args, isbn)
		argIndex++
	}

	err := config.DB.Select(&books, query, args...)

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
