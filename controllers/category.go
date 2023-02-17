package controllers

import (
	"Quiz-3/database"
	"Quiz-3/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
import "Quiz-3/respository"

func GetAllCategories(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := respository.GetAllCategories(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var category structs.Category

	err := c.ShouldBindJSON(&category)

	if err != nil {
		panic(err)
	}

	err = respository.InsertCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert Category",
	})
}

func UpdateCategory(c *gin.Context) {
	var category structs.Category

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.ID = int64(id)

	err = respository.UpdateCategory(database.DbConnection, category)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Category",
	})
}

func DeleteCategory(c *gin.Context) {
	var category structs.Category
	id, err := strconv.Atoi(c.Param("id"))

	category.ID = int64(id)

	err = respository.DeleteCategory(database.DbConnection, category)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"return": "Success Delete Category",
	})
}

func GetAllBookByCategoryId(c *gin.Context) {
	var (
		result gin.H
	)

	var category structs.Category
	id, err := strconv.Atoi(c.Param("id"))

	category.ID = int64(id)

	err, books := respository.GetAllBookByCategoryId(database.DbConnection, category)

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
