package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/SalmaanChicmic/Go/models"
)

func GetExistingBooks() ([]models.Book, error) {

	var books []models.Book

	file, err := ioutil.ReadFile("./books.json")
	if err != nil {
		// if file does not exist
		fmt.Println("File does not exist", err)
	} else {
		json.Unmarshal(file, &books)
	}

	return books, err
}

func SaveBookData(books []models.Book) {
	bookData, _ := json.MarshalIndent(books, "", " ")

	_ = ioutil.WriteFile("books.json", bookData, 0644)
}
