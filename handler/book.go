package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vsualzm/gin-cek/book"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "ilham maulana",
		"bio":  "Seorang programmer",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "helohelo",
		"sub":  "bau kemem",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

func PostBooksHandler(c *gin.Context) {
	// title dan price
	// menangkap dari struct

	var BookInput book.BookInput

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
