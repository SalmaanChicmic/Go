package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"net/http"

	"github.com/SalmaanChicmic/Go/constants"
	"github.com/SalmaanChicmic/Go/models"
	"github.com/SalmaanChicmic/Go/services"
	"github.com/gin-gonic/gin"
)

/*
 */
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func GetBook(c *gin.Context) {
	var books []models.Book

	bookId, _ := strconv.Atoi(c.Param("id"))

	file, err := ioutil.ReadFile("books.json")
	if err != nil {
		// if file does not exist
		c.JSON(http.StatusOK, gin.H{
			"message": "No books found",
		})

		return
	}

	json.Unmarshal(file, &books)

	if len(books) <= bookId {

		c.JSON(http.StatusOK, gin.H{
			"message": constants.BOOK_NOT_FOUND,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": constants.SUCCESS,
		"data":    books[bookId],
	})
}

func GetBooks(c *gin.Context) {
	var books []models.Book

	file, err := ioutil.ReadFile("books.json")
	if err != nil {
		// if file does not exist
		c.JSON(http.StatusOK, gin.H{
			"message": constants.BOOK_NOT_FOUND,
		})
	} else {
		json.Unmarshal(file, &books)

		c.JSON(http.StatusOK, gin.H{
			"message": constants.SUCCESS,
			"data":    books,
		})
	}
}

func SaveBooks(c *gin.Context) {
	books, err := services.GetExistingBooks()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": constants.BOOK_NOT_FOUND,
		})

		return
	}

	var book models.Book

	err = c.ShouldBindJSON(&book)
	if err != nil {
		fmt.Println(constants.ERROR_OCCURED, err)
		panic(constants.ERROR_OCCURED)
	}

	book.Id = len(books)

	books = append(books, book)

	bookData, _ := json.MarshalIndent(books, "", " ")

	_ = ioutil.WriteFile("books.json", bookData, 0644)

	c.JSON(http.StatusOK, gin.H{
		"message": "book saved",
	})
}
