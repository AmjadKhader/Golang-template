package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	httpPkg "main/pkg/http"
	model "main/pkg/model"
)

var (
	mockBooksData = []byte(`{"data":[{"id":1,"Year":1974,"Title":"Carrie","handle":"carrie","Publisher":"Doubleday","ISBN":"978-0-385-08695-0","Pages":199,"Notes":[""],"created_at":"2023-11-13T23:48:47.848Z","villains":[{"name":"Tina Blake","url":"https://stephen-king-api.onrender.com/api/villain/4"},{"name":"Cindi","url":"https://stephen-king-api.onrender.com/api/villain/14"},{"name":"Myra Crewes","url":"https://stephen-king-api.onrender.com/api/villain/16"},{"name":"Billy deLois","url":"https://stephen-king-api.onrender.com/api/villain/25"},{"name":"Kenny Garson","url":"https://stephen-king-api.onrender.com/api/villain/38"},{"name":"Mary Lila Grace","url":"https://stephen-king-api.onrender.com/api/villain/44"},{"name":"Christine Hargensen","url":"https://stephen-king-api.onrender.com/api/villain/49"},{"name":"Vic Mooney","url":"https://stephen-king-api.onrender.com/api/villain/75"},{"name":"The Mortimer Snerds","url":"https://stephen-king-api.onrender.com/api/villain/77"},{"name":"Billy Nolan","url":"https://stephen-king-api.onrender.com/api/villain/80"},{"name":"Elenor Richmond","url":"https://stephen-king-api.onrender.com/api/villain/90"},{"name":"Rachel Spies","url":"https://stephen-king-api.onrender.com/api/villain/94"},{"name":"Jackie Talbot","url":"https://stephen-king-api.onrender.com/api/villain/99"},{"name":"Donna and Mary Lila Grace Thibodeau","url":"https://stephen-king-api.onrender.com/api/villain/102"},{"name":"Jessica Upshaw","url":"https://stephen-king-api.onrender.com/api/villain/108"},{"name":"Norma Watson","url":"https://stephen-king-api.onrender.com/api/villain/109"},{"name":"Margaret White","url":"https://stephen-king-api.onrender.com/api/villain/113"}]},{"id":2,"Year":1975,"Title":"Salem's Lot","handle":"salem-s-lot","Publisher":"Doubleday","ISBN":"978-0-385-00751-1","Pages":439,"Notes":["Nominee, World Fantasy Award, 1976[2]"],"created_at":"2023-11-13T23:48:48.098Z","villains":[{"name":"Kurt Barlow","url":"https://stephen-king-api.onrender.com/api/villain/2"},{"name":"Richard Straker","url":"https://stephen-king-api.onrender.com/api/villain/98"}]},{"id":3,"Year":1977,"Title":"The Shining","handle":"the-shining","Publisher":"Doubleday","ISBN":"978-0-385-12167-5","Pages":447,"Notes":["Runner-up (4th place), Locus Award for Best Fantasy Novel, 1978[2]"],"created_at":"2023-11-13T23:48:48.219Z","villains":[{"name":"Horace M. Derwent","url":"https://stephen-king-api.onrender.com/api/villain/26"},{"name":"Delbert Grady","url":"https://stephen-king-api.onrender.com/api/villain/45"},{"name":"Jack Torrance","url":"https://stephen-king-api.onrender.com/api/villain/106"}]}]}`)
	mockBookData  = []byte(`{"data":{"id":1,"Year":1974,"Title":"Carrie","handle":"carrie","Publisher":"Doubleday","ISBN":"978-0-385-08695-0","Pages":199,"Notes":[""],"created_at":"2023-11-13T23:48:47.848Z","villains":[{"name":"Tina Blake","url":"https://stephen-king-api.onrender.com/api/villain/4"},{"name":"Cindi","url":"https://stephen-king-api.onrender.com/api/villain/14"},{"name":"Myra Crewes","url":"https://stephen-king-api.onrender.com/api/villain/16"},{"name":"Billy deLois","url":"https://stephen-king-api.onrender.com/api/villain/25"},{"name":"Kenny Garson","url":"https://stephen-king-api.onrender.com/api/villain/38"},{"name":"Mary Lila Grace","url":"https://stephen-king-api.onrender.com/api/villain/44"},{"name":"Christine Hargensen","url":"https://stephen-king-api.onrender.com/api/villain/49"},{"name":"Vic Mooney","url":"https://stephen-king-api.onrender.com/api/villain/75"},{"name":"The Mortimer Snerds","url":"https://stephen-king-api.onrender.com/api/villain/77"},{"name":"Billy Nolan","url":"https://stephen-king-api.onrender.com/api/villain/80"},{"name":"Elenor Richmond","url":"https://stephen-king-api.onrender.com/api/villain/90"},{"name":"Rachel Spies","url":"https://stephen-king-api.onrender.com/api/villain/94"},{"name":"Jackie Talbot","url":"https://stephen-king-api.onrender.com/api/villain/99"},{"name":"Donna and Mary Lila Grace Thibodeau","url":"https://stephen-king-api.onrender.com/api/villain/102"},{"name":"Jessica Upshaw","url":"https://stephen-king-api.onrender.com/api/villain/108"},{"name":"Norma Watson","url":"https://stephen-king-api.onrender.com/api/villain/109"},{"name":"Margaret White","url":"https://stephen-king-api.onrender.com/api/villain/113"}]}}`)
)

func TestPublic_GetAllBooks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/public/book", httpPkg.Public_GetAllBooks)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/public/book", nil)
	assert.Nil(t, err)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := &httpPkg.BookResponse{Books: []model.PublicBook{}}
	err = json.Unmarshal(mockBooksData, expected)
	assert.Nil(t, err)

	assert.ObjectsAreEqual(len(expected.Books), 3)
}

func TestPublic_GetBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/public/book/1", httpPkg.Public_GetBook)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/public/book/1", nil)
	assert.Nil(t, err)

	router.ServeHTTP(w, req)

	expected := &httpPkg.SingleBookResponse{Book: model.PublicBook{}}
	err = json.Unmarshal(mockBookData, expected)
	assert.Nil(t, err)

	assert.ObjectsAreEqual(expected.Book.ID, 1)
	assert.ObjectsAreEqual(expected.Book.Title, "Carrie")
}
