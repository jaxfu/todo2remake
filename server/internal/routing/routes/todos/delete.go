package todos

import (
	"fmt"
	"furrj/todo_2_remake/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type deletePayload struct {
	ID uint `json:"id"`
}

// DeleteTodo deletes a todo by id
// TODO: accept username and delete specific user's todo
func DeleteTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// init payload object to allow binding
		var payload deletePayload
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
		fmt.Printf("Delete ID: %d\n", payload.ID)

		// remove todo from slice by ID
		// TODO: remove from db
		TestTodos = filter(TestTodos, payload.ID)

		// set response to valid and send
		response.Valid = true
		ctx.JSON(http.StatusOK, response)
	}
}

// filter iterates through todos and removes one with matching ID
func filter(todos []types.Todo, id uint) []types.Todo {
	var filtered []types.Todo

	for _, todo := range todos {
		if todo.ID != id {
			filtered = append(filtered, todo)
		}
	}

	return filtered
}
