package routes

import (
	"catalogolivros/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		book := main.Group("books")
		{
			book.POST("/create", controllers.CreateBook)
			book.GET("/read/:id", controllers.ReadBook)
			book.GET("/read_all", controllers.ReadAllBooks)
			book.PUT("/update", controllers.UpdateBook)
			book.DELETE("/delete/:id", controllers.DeleteBook)
		}
	}

	return router
}
