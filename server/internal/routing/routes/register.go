package routes

import (
	"fmt"
	"furrj/todo_2_remake/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload types.RequestRegister
		response := types.ResponseValid{Valid: false}

		// bind request body to payload
		if err := ctx.BindJSON(&payload); err != nil {
			fmt.Printf("Error binding payload: %+v\n", err)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}
		fmt.Printf("Register Payload: %+v\n", payload)

		ctx.JSON(http.StatusOK, types.ResponseValid{Valid: true})
	}
}
