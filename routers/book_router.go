package routers

import (
	"belajar-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)

	router.PUT("/books/:book_ID", controllers.UpdateBook)

	router.GET("/books/:book_ID", controllers.GetBook)

	router.DELETE("/books/:book_ID", controllers.DeleteBook)


	return router
}
