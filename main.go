package main

import (
	"log"

	"github.com/Uttkarsh-Raj/Proxie/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routes.IncomingRoutes(router)
	log.Fatal(router.Run(":7000"))
}
