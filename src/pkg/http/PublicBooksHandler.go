package http

import (
	"encoding/json"
	"io/ioutil"

	model "main/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

const baseUrl string = "https://stephen-king-api.onrender.com/api/"

type BookResponse struct {
	Books []model.PublicBook `json:"data"` // Assuming the field name is "books"
}

type SingleBookResponse struct {
	Book model.PublicBook `json:"data"` // Assuming the field name is "books"
}

func Public_GetAllBooks(context *gin.Context) {
	resposneByte, err := callApi("books/")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	books, err := retrieveBooks(resposneByte)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, books)
}

func Public_GetBook(context *gin.Context) {
	bookId := context.Param("id")
	resposneByte, err := callApi("book/" + bookId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	books, err := retrieveBook(resposneByte)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, books)
}

func callApi(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, baseUrl+url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func retrieveBooks(bodyBytes []byte) ([]model.PublicBook, error) {
	var response BookResponse

	err := json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return response.Books, nil
}

func retrieveBook(bodyBytes []byte) (model.PublicBook, error) {
	var response SingleBookResponse

	err := json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return model.PublicBook{}, err
	}

	return response.Book, nil
}
