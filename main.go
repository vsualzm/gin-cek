package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vsualzm/gin-cek/handler"
)

func main() {

	//1
	router := gin.Default()

	//2

	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	//3
	router.Run()

}
