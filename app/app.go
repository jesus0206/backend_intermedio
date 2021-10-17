package app

import (
	"fmt"
	"log"
	"os"

	"yofio/intermedio/controllers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// StartApp function runner server
func StartApp() {
	router := RouterInitial()
	log.Fatal(router.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}

// RouterInitial exported
func RouterInitial() *gin.Engine {

	if os.Getenv("DEBUG") == "TRUE" {
		gin.SetMode(gin.DebugMode)
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}
	// router.Use(cors.Default())

	controller := controllers.NewController()
	router.POST("/credit-assignment", controller.CreditAssignment)
	return router
}
