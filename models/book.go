package models

type Book struct {
	Id     int
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}
