package http

import (
	"net/http"

	model "main/pkg/model"
	bookService "main/pkg/service"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(context *gin.Context) {
	books, err := bookService.GetAllBooks()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, books)
}

func GetBook(context *gin.Context) {
	bookId := context.Param("id")
	book, err := bookService.GetBook(bookId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, book)
}

func AddBook(context *gin.Context) {
	var book model.Book
	if err := context.BindJSON(&book); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := bookService.AddBook(book)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, output)
}

func UpdateBook(context *gin.Context) {
	bookId := context.Param("id")

	var book model.Book
	if err := context.BindJSON(&book); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := bookService.UpdateBook(bookId, book)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, output)
}

func DeleteBook(context *gin.Context) {
	bookId := context.Param("id")
	err := bookService.DeleteBook(bookId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, bookId)
}
