package todos

import (
	"fmt"
	"furrj/todo_2_remake/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateTodo accepts a todo object and updates the todo
// TODO: accept username and update specific user's todo
func UpdateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// init payload to allow binding
		var payload types.Todo
		// init response object as invalid
		response := types.ResponseValid{
			Valid: false,
		}

		// bind request body to payload
		if err := ctx.BindJSON(&payload); err != nil {
			fmt.Printf("Error binding payload: %+v\n", err)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}
		fmt.Printf("Login Payload: %+v\n", payload)

		// update todo in slice
		// TODO: update in db
		updateTodo(TestTodos, payload)

		// set response to valid and send
		response.Valid = true
		ctx.JSON(http.StatusOK, response)
	}
}

// updateTodo updates todo with same ID in slice
func updateTodo(todos []types.Todo, todo types.Todo) {
	for i := range todos {
		if todos[i].ID == todo.ID {
			todos[i] = todo
		}
	}
}
