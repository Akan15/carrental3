package handlers

import (
	"net/http"

	"github.com/Akan15/carrental3/api-gateway/internal/client"
	carPb "github.com/Akan15/carrental3/car-service/proto"
	"github.com/gin-gonic/gin"
)

func RegisterCarRoutes(r *gin.Engine, c *client.Clients) {
	r.GET("/cars", func(ctx *gin.Context) {
		resp, err := c.Car.ListCars(ctx, &carPb.Empty{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, resp.Cars)
	})

	r.POST("/cars", func(ctx *gin.Context) {
		var req struct {
			Brand     string  `json:"brand"`
			Model     string  `json:"model"`
			City      string  `json:"city"`
			Status    string  `json:"status"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		car, err := c.Car.CreateCar(ctx, &carPb.CreateCarRequest{
			Brand:     req.Brand,
			Model:     req.Model,
			City:      req.City,
			Status:    req.Status,
			Latitude:  req.Latitude,
			Longitude: req.Longitude,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, car)
	})

	r.GET("/cars/available", func(ctx *gin.Context) {
		resp, err := c.Car.GetAvailableCars(ctx, &carPb.Empty{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, resp.Cars)
	})

	r.GET("/cars/search", func(ctx *gin.Context) {
		model := ctx.Query("model")
		resp, err := c.Car.FindByModel(ctx, &carPb.ModelRequest{Model: model})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, resp.Cars)
	})

}
