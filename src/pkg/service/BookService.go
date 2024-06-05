package service

import (
	"database/sql"
	config "main/pkg/config"
	model "main/pkg/model"
)

func GetAllBooks() ([]model.Book, error) {
	rows, err := config.DB.Query("SELECT id, title, auther, total_pages FROM books")
	if err != nil {
		return nil, err
	}

	defer rows.Close() // for resource cleanup
	var books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Auther, &book.NumberOfPages)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func GetBook(id string) (*model.Book, error) {
	var book model.Book

	output := config.DB.QueryRow("SELECT id, title, auther, total_pages FROM books where id = "+id).Scan(&book.ID, &book.Title, &book.Auther, &book.NumberOfPages)
	if output == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	} else if output != nil {
		return nil, sql.ErrConnDone
	}

	return &book, nil
}

func AddBook(input model.Book) (*model.Book, error) {
	var book model.Book

	output := config.DB.QueryRow("INSERT INTO books (title, auther, total_pages) VALUES ($1, $2, $3) RETURNING id, title, auther, total_pages", input.Title, input.Auther, input.NumberOfPages).Scan(&book.ID, &book.Title, &book.Auther, &book.NumberOfPages)
	if output != nil {
		return nil, sql.ErrConnDone
	}

	return &book, nil
}

func UpdateBook(bookId string, input model.Book) (*model.Book, error) {

	result, err := config.DB.Exec("UPDATE books SET title=$1, auther=$2, total_pages=$3 WHERE id=$4", input.Title, input.Auther, input.NumberOfPages, bookId)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}
	input.ID = bookId

	return &input, nil
}

func DeleteBook(bookId string) error {
	result, err := config.DB.Exec("DELETE FROM books WHERE id=$1", bookId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
