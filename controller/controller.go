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

	"github.com/SalmaanChicmic/Go/constants"
	"github.com/SalmaanChicmic/Go/models"
	"github.com/SalmaanChicmic/Go/services"
	// "github.com/SalmaanChicmic/Go/services"
)

// response types
type response struct {
	Message string
}

type bookRes struct {
	Message string
	Book    models.Book
}

type booksRes struct {
	Message string
	Books   []models.Book
}

func Test(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Server is healthy, world!\n")
}

func GetBook(w http.ResponseWriter, req *http.Request) {
	var books []models.Book

	bookId, _ := strconv.Atoi(req.URL.Query().Get("id"))

	file, err := ioutil.ReadFile("books.json")
	if err != nil {
		// if file does not exist
		res := bookRes{Message: constants.FILE_NOT_FOUND}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)

		return
	}

	json.Unmarshal(file, &books)

	if len(books) <= bookId {
		res := bookRes{Message: constants.BOOK_NOT_FOUND}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)

		return
	}

	res := bookRes{Message: constants.SUCCESS, Book: books[bookId]}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func getBooks(w http.ResponseWriter, req *http.Request) {

	var books []models.Book

	file, err := ioutil.ReadFile("books.json")
	if err != nil {
		// if file does not exist
		res := response{Message: constants.FILE_NOT_FOUND}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)

		return

	} else {
		json.Unmarshal(file, &books)

		res := booksRes{Message: constants.SUCCESS, Books: books}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func saveBooks(w http.ResponseWriter, req *http.Request) {
	books, err := services.GetExistingBooks()

	if err != nil {
		// if file does not exist
		res := response{Message: constants.FILE_NOT_FOUND}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)

		return
	}

	var book models.Book

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(constants.ERROR_OCCURED, err)
		panic(constants.ERROR_OCCURED)
	}

	json.Unmarshal(body, &book)

	book.Id = len(books)

	books = append(books, book)

	bookData, _ := json.MarshalIndent(books, "", " ")

	_ = ioutil.WriteFile("books.json", bookData, 0644)

	res := response{Message: constants.SUCCESS}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func HandleBooks(w http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		getBooks(w, req)
	}

	if req.Method == "POST" {
		saveBooks(w, req)
	}

}
