package main

import (
	"go-api/controllers"
	"go-api/tests"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// Unit Test
	tests.UnitTest()

	// Init Router
	router := gin.Default()

	// group: bookRoute
	bookRoute := router.Group("/api/book")
	{
		bookRoute.GET("/", controllers.GetBooks)
		bookRoute.GET("/:id", controllers.GetBook)
		bookRoute.POST("/", controllers.InsertBook)
		bookRoute.PUT("/:id", controllers.UpdateBook)
		bookRoute.DELETE("/:id", controllers.DeleteBook)
	}

	// Listen and serve on 0.0.0.0:8000
	log.Println("Listen and serve http://127.0.0.1:8000")
	router.Run(":8000")

}
