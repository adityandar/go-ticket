package router

import (
	"fmt"
	"go-ticket/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()
	router.POST("/register", controller.RegisterAsOrganizer)
	router.POST("/login", controller.LoginAsOrganizer)
	router.PATCH("/organizer/:id", controller.UpdateOrganizer)
	router.DELETE("/organizer/:id", controller.DeleteOrganizer)

	port := 8100
	log.Println("Server running in port:", port)
	router.Run(fmt.Sprintf(":%d", port))
}
