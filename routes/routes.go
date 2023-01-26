package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SalmaanChicmic/Go/config"
	"github.com/SalmaanChicmic/Go/controller"
	"github.com/SalmaanChicmic/Go/middlewares"
)

func Router() {
	// test route
	http.Handle("/", middlewares.LogReq(http.HandlerFunc(controller.Test)))
	// books
	http.Handle("/book", middlewares.LogReq(http.HandlerFunc(controller.GetBook)))
	http.Handle("/books", middlewares.LogReq(http.HandlerFunc(controller.HandleBooks)))

	fmt.Printf("Listening to port: %v \n", config.PORT)

	log.Fatal(http.ListenAndServe(config.PORT, nil))
}
