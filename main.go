package main

import (
	"Quiz-3/controllers"
	"Quiz-3/database"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed to load file environment")
	} else {
		fmt.Println("success to load file environment")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	DB, err := sql.Open("postgres", psqlInfo)

	err = DB.Ping()

	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	router := gin.Default()

	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"admin":  "password",
		"editor": "secret",
	}))

	// bangun-datar
	router.GET("/bangun-datar/segitiga-sama-sisi", controllers.SegitigaSamaSisi)
	router.GET("/bangun-datar/persegi", controllers.Persegi)
	router.GET("/bangun-datar/persegi-panjang", controllers.PersegiPanjang)
	router.GET("/bangun-datar/lingkaran", controllers.Lingkaran)

	// category
	router.GET("/categories", controllers.GetAllCategories)
	authorized.POST("/categories", controllers.InsertCategory)
	authorized.PUT("/categories/:id", controllers.UpdateCategory)
	authorized.DELETE("/categories/:id", controllers.DeleteCategory)
	router.GET("/categories/:id/books", controllers.GetAllBookByCategoryId)

	// book
	router.GET("/books", controllers.GetAllBooks)
	authorized.POST("/books", controllers.InsertBook)
	authorized.PUT("/books/:id", controllers.UpdateBook)
	authorized.DELETE("/books/:id", controllers.DeleteBook)

	router.Run("localhost:8080")

}
