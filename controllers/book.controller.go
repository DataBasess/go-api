package controllers

import (
	"fmt"
	"go-api/models"
	"go-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	fmt.Println("GetBooks")
	result, err := services.GetBooks()
	if err != nil {
		c.JSON(http.StatusExpectationFailed, err)
	}
	c.JSON(http.StatusOK, result)
}

func GetBook(c *gin.Context) {
	fmt.Println("GetBook")
	ID := c.Param("id")
	fmt.Println("ID : ", ID)
	result, err := services.GetBook(ID)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, err)
	}
	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	book := models.Book{}
	validator := c.ShouldBindJSON(&book)
	if validator != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "BadRequest"})
	} else {
		fmt.Println("InsertBook")
		fmt.Printf("%+v\n", book)
		result, err := services.InsertBook(book)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, err)
		}
		c.JSON(http.StatusOK, result)
	}
}

func UpdateBook(c *gin.Context) {
	book := models.Book{}
	validator := c.ShouldBindJSON(&book)
	ID := c.Params.ByName("id")
	if validator != nil || ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "BadRequest"})
	} else {
		fmt.Println("UpdateBook")
		fmt.Printf("%+v\n", book)
		result, err := services.UpdateBook(ID, book)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, err)
		}
		c.JSON(http.StatusOK, result)
	}
}

func DeleteBook(c *gin.Context) {
	fmt.Println("DeleteBook")
	ID := c.Params.ByName("id")
	result, err := services.DeleteBook(ID)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, err)
	}
	c.JSON(http.StatusOK, result)
}
