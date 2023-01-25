package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SalmaanChicmic/Go/config"
	"github.com/SalmaanChicmic/Go/controller"
)

func Router() {
	// test route
	http.HandleFunc("/", controller.Test)
	// // books
	http.HandleFunc("/books", controller.GetBook)
	// router.GET("/books", controller.GetBooks)
	// router.POST("/books", controller.SaveBooks)

	fmt.Printf("Listening to port: %v \n", config.PORT)
	log.Fatal(http.ListenAndServe(config.PORT, nil))
}
