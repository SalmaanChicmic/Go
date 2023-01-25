package controller

import (
	// "encoding/json"
	// "fmt"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"

	"net/http"

	"github.com/SalmaanChicmic/Go/models"
	// "github.com/SalmaanChicmic/Go/services"
)

/*
 */
func Test(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Server is healthy, world!\n")
}

func GetBook(w http.ResponseWriter, req *http.Request) {
	var books []models.Book

	bookId, _ := strconv.Atoi(req.URL.Query().Get("id"))

	file, err := ioutil.ReadFile("books.json")
	if err != nil {
		// // if file does not exist
		c.JSON(http.StatusOK, gin.H{
			"message": "No books found",
		})
		
		res := json.Marshal({
			"message": "No books found",
		})

		w.Header().Set("Content-Type", "application/json")
        w.Write()


		return
	}

	json.Unmarshal(file, &books)

	if len(books) <= bookId {

		// c.JSON(http.StatusOK, gin.H{
		// 	"message": constants.BOOK_NOT_FOUND,
		// })

		io.WriteString(w, "Server is healthy, world!\n")

		return
	}

	io.WriteString(w, "Server is healthy, world!\n")

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": constants.SUCCESS,
	// 	"data":    books[bookId],
	// })
}

// func GetBooks(w http.ResponseWriter, req *http.Request) {
// 	var books []models.Book

// 	file, err := ioutil.ReadFile("books.json")
// 	if err != nil {
// 		// if file does not exist
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": constants.BOOK_NOT_FOUND,
// 		})
// 	} else {
// 		json.Unmarshal(file, &books)

// 		c.JSON(http.StatusOK, gin.H{
// 			"message": constants.SUCCESS,
// 			"data":    books,
// 		})
// 	}
// }

// func SaveBooks(w http.ResponseWriter, req *http.Request) {
// 	books, err := services.GetExistingBooks()
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": constants.BOOK_NOT_FOUND,
// 		})

// 		return
// 	}

// 	var book models.Book

// 	err = c.ShouldBindJSON(&book)
// 	if err != nil {
// 		fmt.Println(constants.ERROR_OCCURED, err)
// 		panic(constants.ERROR_OCCURED)
// 	}

// 	book.Id = len(books)

// 	books = append(books, book)

// 	bookData, _ := json.MarshalIndent(books, "", " ")

// 	_ = ioutil.WriteFile("books.json", bookData, 0644)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "book saved",
// 	})

// }
