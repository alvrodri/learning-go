package main

import (
	"Twitter/main/routes"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func registerRoutes(router *gin.Engine) {
	router.GET("/user/:username", routes.GetUser)
	router.POST("/user/", routes.CreateUser)
}

func main() {
	err := mgm.SetDefaultConfig(nil, "twitter", options.Client().ApplyURI("mongodb://localhost:27017/app"))
	if err != nil {
		log.Fatal(err)
		return
	}

	router := gin.Default()

	registerRoutes(router)

	err = router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
