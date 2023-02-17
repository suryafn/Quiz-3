package controllers

import (
	"Quiz-3/database"
	"Quiz-3/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
import "Quiz-3/respository"

func GetAllBooks(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := respository.GetAllBooks(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var book structs.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		panic(err)
	}

	err = respository.InsertBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert Book",
	})
}

func UpdateBook(c *gin.Context) {
	var book structs.Book

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.ID = int64(id)

	err = respository.UpdateBook(database.DbConnection, book)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Book",
	})
}

func DeleteBook(c *gin.Context) {
	var book structs.Book
	id, err := strconv.Atoi(c.Param("id"))

	book.ID = int64(id)

	err = respository.DeleteBook(database.DbConnection, book)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"return": "Success Delete Book",
	})
}
