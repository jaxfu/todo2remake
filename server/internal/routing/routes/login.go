package routes

import (
	"fmt"
	"furrj/todo_2_remake/internal/dbHandler"
	"furrj/todo_2_remake/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Login(db *dbHandler.DBHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload types.RequestLogin
		response := types.ResponseLoginRegister{UserID: 0}

		// bind request body to payload
		if err := ctx.BindJSON(&payload); err != nil {
			fmt.Printf("Error binding payload: %+v\n", err)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}
		fmt.Printf("Register Payload: %+v\n", payload)

		// get user data from db
		userData, err := db.GetUserDataByUsername(payload.Username)
		if err != nil {
			if err == pgx.ErrNoRows {
				fmt.Printf("username %s not found\n", payload.Username)
				ctx.JSON(http.StatusUnauthorized, response)
				return
			}
			fmt.Printf("error getting user data: %v\n", err)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}

		// if password incorrect
		if payload.Password != userData.Password {
			fmt.Printf("password incorrect")
			ctx.JSON(http.StatusUnauthorized, response)
			return
		}

		response.UserID = userData.UserID
		ctx.JSON(http.StatusOK, response)
	}
}
