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

	// mengambil data buku
	// books, err := bookRepository.FindAll()
	// for _, book := range books {
	// 	fmt.Println("title:", book.Title)
	// }

	// mengambil per id
	// book, err := bookRepository.FindByID(2)
	// fmt.Println("title:", book.Title)

	// create book
	// book := book.Book{
	// 	Title:       "gg geming",
	// 	Description: "Good book",
	// 	Price:       2000,
	// 	Rating:      30,
	// 	Discount:    20,
	// }

	// bookRequest := book.BookRequest{
	// 	Title: "gundam",
	// 	Price: "123123",
	// }

	// bookService.Create(bookRequest)
	// bookRepository.Create(book)

	//1
	router := gin.Default()
	//2
	v1 := router.Group("/v1")
	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	v1.GET("/books/:id/:title", bookHandler.BooksHandler)
	v1.GET("/query", bookHandler.QueryHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)
	//3
	router.Run()

	//main
	//service
	//repository
	//db
	//mysql

}
