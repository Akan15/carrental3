package handlers

import (
	"net/http"

	"github.com/Akan15/carrental3/api-gateway/internal/client"

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
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user ID required"})
		return
	}
	ctx := c.Request.Context()
	resp, err := h.Clients.UserClient.GetUser(ctx, &userpb.GetUserRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCar(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "car ID required"})
		return
	}
	ctx := c.Request.Context()
	resp, err := h.Clients.CarClient.GetCar(ctx, &carpb.CarIdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetRental(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rental ID required"})
		return
	}
	ctx := c.Request.Context()
	resp, err := h.Clients.RentalClient.GetRental(ctx, &rentalpb.GetRentalRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
