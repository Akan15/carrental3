package handlers

import (
	"net/http"

	"github.com/Akan15/carrental3/api-gateway/internal/client"
	pb "github.com/Akan15/carrental3/rental-service/proto"
	"github.com/gin-gonic/gin"
)

func RegisterRentalRoutes(r *gin.Engine, c *client.Clients) {
	r.GET("/rentals/active", func(ctx *gin.Context) {
		resp, err := c.Rental.GetActiveRentals(ctx, &pb.Empty{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, resp.Rentals)
	})
}
