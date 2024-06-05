package Routes

import (
	handlers "main/pkg/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()

	router.GET("/api/book", handlers.GetAllBooks)
	router.GET("/api/book/:id", handlers.GetBook)
	router.POST("/api/book", handlers.AddBook)
	router.PUT("/api/book/:id", handlers.UpdateBook)
	router.DELETE("/api/book/:id", handlers.DeleteBook)

	router.GET("/api/public/book", handlers.Public_GetAllBooks)
	router.GET("/api/public/book/:id", handlers.Public_GetBook)

	return router
}
