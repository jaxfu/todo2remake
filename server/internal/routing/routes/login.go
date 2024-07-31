package routes

import (
	"fmt"
	"furrj/todo_2_remake/internal/dbHandler"
	"furrj/todo_2_remake/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(db *dbHandler.DBHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload types.RequestLogin
		response := types.ResponseValid{Valid: false}

		// bind request body to payload
		if err := ctx.BindJSON(&payload); err != nil {
			fmt.Printf("Error binding payload: %+v\n", err)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}
		fmt.Printf("Login Payload: %+v\n", payload)

		// get user data from db
		userData, err := db.GetUserDataByUsername(payload.Username)
		if err != nil {
			fmt.Printf("error getting user data: %v\n", err)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}

		ctx.JSON(http.StatusOK, types.ResponseValid{Valid: true})
	}
}
