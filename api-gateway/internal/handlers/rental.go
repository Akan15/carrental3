package handlers

import (
	"context"
	"net/http"

	"github.com/Akan15/carrental3/api-gateway/internal/client"
	pb "github.com/Akan15/carrental3/rental-service/proto"
	"github.com/gin-gonic/gin"
)

func RegisterRentalRoutes(r *gin.Engine, c *client.Clients) {
	// Получить активные аренды
	r.GET("/rentals/active", func(ctx *gin.Context) {
		resp, err := c.Rental.GetActiveRentals(ctx, &pb.Empty{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, resp.Rentals)
	})

	r.POST("/rentals", func(ctx *gin.Context) {
		var req struct {
			UserId string `json:"userId"`
			CarId  string `json:"carId"`
			Type   string `json:"type"`
		}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		rental, err := c.Rental.CreateRental(context.Background(), &pb.CreateRentalRequest{
			UserId: req.UserId,
			CarId:  req.CarId,
			Type:   req.Type,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, rental)
	})

}
