package handlers

import (
	"net/http"

	"github.com/Akan15/carrental3/api-gateway/internal/client"
	userPb "github.com/Akan15/carrental3/user-service/proto"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, c *client.Clients) {
	r.POST("/register", func(ctx *gin.Context) {
		var req struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := c.User.Register(ctx, &userPb.RegisterRequest{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": resp.Message})
	})

	r.POST("/login", func(ctx *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := c.User.Login(ctx, &userPb.LoginRequest{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"token": resp.Token})
	})

	r.POST("/verify", func(ctx *gin.Context) {
		var req struct {
			Token string `json:"token"`
		}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := c.User.Verify(ctx, &userPb.VerifyRequest{
			Token: req.Token,
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"user_id": resp.UserId})
	})

	r.PUT("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var req struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := c.User.UpdateUser(ctx, &userPb.UpdateUserRequest{
			Id:    id,
			Name:  req.Name,
			Email: req.Email,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, resp)
	})

	r.POST("/resend-verification", func(ctx *gin.Context) {
		var req struct {
			Email string `json:"email" binding:"required,email"`
		}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
			return
		}

		resp, err := c.User.ResendVerification(ctx, &userPb.ResendVerificationRequest{
			Email: req.Email,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": resp.Message})
	})

	r.POST("/change-password", func(ctx *gin.Context) {
		var req struct {
			Email       string `json:"email"`
			OldPassword string `json:"old_password"`
			NewPassword string `json:"new_password"`
		}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		resp, err := c.User.ChangePassword(ctx, &userPb.ChangePasswordRequest{
			Email:       req.Email,
			OldPassword: req.OldPassword,
			NewPassword: req.NewPassword,
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": resp.Message})
	})

}
