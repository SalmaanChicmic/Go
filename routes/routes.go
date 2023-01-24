package routes

import (
	"github.com/SalmaanChicmic/Golang/config"
	"github.com/SalmaanChicmic/Golang/controller"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.New()

	router.SetTrustedProxies([]string{"112.196.113.2"})

	router.GET("/", controller.Test)
	// books
	router.GET("/books/:id", controller.GetBook)
	router.GET("/books", controller.GetBooks)
	router.POST("/books", controller.SaveBooks)

	router.Run(config.PORT)
}
