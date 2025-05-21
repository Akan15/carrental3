package main

import (
	"github.com/Akan15/carrental3/api-gateway/internal/client"
	"github.com/Akan15/carrental3/api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	clients := client.InitClients()
	r := gin.Default()

	handlers.RegisterUserRoutes(r, clients)
	handlers.RegisterCarRoutes(r, clients)
	handlers.RegisterRentalRoutes(r, clients)
	r.Run(":8080")
}
