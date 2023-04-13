package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	Title string      `json:"title" binding:"required"`
	Price interface{} `json:"price" binding:"required,number"`
}

func postBooksHandler(c *gin.Context) {
	// title dan price
	// menangkap dari struct

	var BookInput BookInput

	err := c.ShouldBindJSON(&BookInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s, condition:%s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": BookInput.Title,
		"price": BookInput.Price,
	})

}
