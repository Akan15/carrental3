package main

import (
	"time"

	"github.com/Akan15/carrental3/api-gateway/internal/client"
	"github.com/Akan15/carrental3/api-gateway/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	clients := client.InitClients()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handlers.RegisterUserRoutes(r, clients)
	handlers.RegisterCarRoutes(r, clients)
	handlers.RegisterRentalRoutes(r, clients)
	r.Run(":8080")
}
