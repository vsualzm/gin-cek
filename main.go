package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vsualzm/gin-cek/book"
	"github.com/vsualzm/gin-cek/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// connection database
	dsn := "root:@tcp(127.0.0.1:3306)/books_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	// migrate
	db.AutoMigrate(&book.Book{})
	fmt.Println("database Connection succeed")

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	//1
	router := gin.Default()
	//2
	v1 := router.Group("/v1")
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBooksByID)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	//3
	router.Run()

	//main
	//handler
	//service
	//repository
	//db
	//mysql

}
