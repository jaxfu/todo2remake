package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type apiResponseRegister struct {
	Valid bool `json:"valid"`
}

func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, apiResponseRegister{Valid: true})
	}
}
