package todos

import (
	"fmt"
	"furrj/todo_2_remake/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestAddTodo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// AddTodo accepts a requestAddTodo object then creates and stores a new todo
// TODO: accept username and add a todo to specific user
func AddTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// init payload object to allow binding
		var payload requestAddTodo
		// init newTodo object
		var newTodo types.Todo
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

		// creat new todo object
		newTodo.Title = payload.Title
		newTodo.Content = payload.Content
		newTodo.ID = uint(len(TestTodos) + 1)

		// append new todo to slice
		// TODO: store in db
		TestTodos = append(TestTodos, newTodo)

		// set response to valid and send
		response.Valid = true
		ctx.JSON(http.StatusOK, response)
	}
}
