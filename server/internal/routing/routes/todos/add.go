package todos

import (
	"fmt"
	"furrj/todo_2_remake/internal/routing/routes"
	"furrj/todo_2_remake/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload types.Todo
		response := routes.ResponseValid{
			Valid: false,
		}

		if err := ctx.BindJSON(&payload); err != nil {
			fmt.Printf("Error binding payload: %+v\n", err)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}
		fmt.Printf("Login Payload: %+v\n", payload)

		payload.ID = uint(len(TestTodos) + 1)
		TestTodos = append(TestTodos, payload)
		response.Valid = true

		ctx.JSON(http.StatusOK, response)
	}
}
