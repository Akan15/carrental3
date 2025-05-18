package handlers

import (
	"context"
	"net/http"

	"carrental/api-gateway/internal/client"
	carpb "carrental/api-gateway/proto/car"
	rentalpb "carrental/api-gateway/proto/rental"
	userpb "carrental/api-gateway/proto/user"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Clients *client.Clients
}

func NewHandler(clients *client.Clients) *Handler {
	return &Handler{Clients: clients}
}

func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.Clients.UserClient.GetUser(context.Background(), &userpb.GetUserRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCar(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.Clients.CarClient.GetCar(context.Background(), &carpb.GetCarRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetRental(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.Clients.RentalClient.GetRental(context.Background(), &rentalpb.GetRentalRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
