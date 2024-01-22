package routes 

import (
	"gofr.dev/pkg/gofr"
	"github.com/vinayaknayar/go-bookstore/pkg/controllers"
)

func RegisterBookStoreRoutes () {
	app := gofr.New()

	app.POST("/book", controllers.CreateBook)
	app.GET("/book", controllers.GetBook)
	app.GET("/book/{bookId}", controllers.GetBookById)
	app.PUT("/book/{bookId}", controllers.UpdateBook)
	app.DELETE("/book/{bookId}", controllers.DeleteBook)

	app.Start()
}