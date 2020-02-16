package main

import (
	"e-commenerce/configuration"
	"e-commenerce/controller"
	"e-commenerce/db"
	"e-commenerce/validation"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := configuration.Readconfigurationfile(); err != nil {

		log.Fatal(err)
	}
	if err := db.Initializedatabase(); err != nil {
		log.Fatal(err)
	}

	if err := validation.Createcustomvalidations(); err != nil {
		log.Fatal(err)
	}

	gin.SetMode(configuration.Configurationins.Development.Mode)
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")

	engine.GET("/", controller.Showhomepage)
	engine.POST("/signup", controller.Signup)
	engine.GET("/signin", controller.Loginpage)
	engine.POST("/signin", controller.Login)

	log.Fatal(engine.Run(configuration.Configurationins.Development.Port))
}
