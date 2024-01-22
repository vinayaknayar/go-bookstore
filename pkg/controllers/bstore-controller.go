package controllers

import (
	"fmt"
	"strconv"

	"github.com/vinayaknayar/go-bookstore/pkg/models"
	"github.com/vinayaknayar/go-bookstore/pkg/utils"
	"gofr.dev/pkg/gofr"
)

var NewBook models.Book

func GetBook(ctx *gofr.Context) (interface{}, error) {
	books := models.GetAllBooks()
	return books, nil
}

func GetBookById(ctx *gofr.Context) (interface{}, error) {
	bookId := ctx.PathParam("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("invalid book ID: %v", err)
	}

	bookDetails, db := models.GetBookById(ID)
	if db.Error != nil {
		return nil, fmt.Errorf("error fetching book details: %v", db.Error)
	}

	return bookDetails, nil
}

func CreateBook(ctx *gofr.Context) (interface{}, error) {
	createBook := &models.Book{}
	utils.ParseBody(ctx.Request(), createBook)
	b, err := createBook.CreateBook()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func DeleteBook(ctx *gofr.Context) (interface{}, error) {
	bookId := ctx.PathParam("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("invalid book ID: %v", err)
	}
	book := models.DeleteBook(ID)

	return book, nil
}

func UpdateBook(ctx *gofr.Context) (interface{}, error) {
	var updateBook = &models.Book{}
	utils.ParseBody(ctx.Request(), updateBook)

	bookID := ctx.PathParam("bookId")
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("invalid book ID: %v", err)
	}

	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	return bookDetails, nil
}