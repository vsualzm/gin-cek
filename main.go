package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	//1
	router := gin.Default()

	//2
	router.GET("/", rootHandler)

	router.GET("/hello", helloHandler)

	router.GET("/books/:id/:title", booksHandler)

	router.GET("/query", queryHandler)

	router.POST("/books", postBooksHandler)

	//3
	router.Run()

}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "ilham maulana",
		"bio":  "Seorang programmer",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "helohelo",
		"sub":  "bau kemem",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

type BookInput struct {
	Title    string `json:"title" binding:"required"` // validasi untuk gin
	Price    int    `json:"price" binding:"required"` // validasi untuk gin
	SubTitle string `json:"sub_title"`
}

func postBooksHandler(c *gin.Context) {
	// title dan price

	// menangkap dari struct

	var BookInput BookInput

	err := c.ShouldBindJSON(&BookInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":     BookInput.Title,
		"price":     BookInput.Price,
		"sub_title": BookInput.SubTitle,
	})

}
