package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type apiResponseLogin struct {
	Valid bool `json:"valid"`
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, apiResponseLogin{Valid: true})
	}
}
