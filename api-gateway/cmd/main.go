package main

import (
	"carrental/api-gateway/internal/client"
	"carrental/api-gateway/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	clients := client.NewClients()
	h := handlers.NewHandler(clients)

	api := r.Group("/api")
	{
		api.GET("/users/:id", h.GetUser)
		api.GET("/cars/:id", h.GetCar)
		api.GET("/rentals/:id", h.GetRental)
	}

	r.Run(":8080")
}
