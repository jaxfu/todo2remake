package todos

import (
	"fmt"
	"furrj/todo_2_remake/internal/routing/routes"
	"furrj/todo_2_remake/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateTodo() gin.HandlerFunc {
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

		updateTodo(TestTodos, payload)
		response.Valid = true
		ctx.JSON(http.StatusOK, response)
	}
}

func updateTodo(todos []types.Todo, todo types.Todo) {
	for i := range todos {
		if todos[i].ID == todo.ID {
			todos[i] = todo
		}
	}
}
