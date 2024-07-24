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

func AddTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload requestAddTodo
		var newTodo types.Todo
		response := types.ResponseValid{
			Valid: false,
		}

		if err := ctx.BindJSON(&payload); err != nil {
			fmt.Printf("Error binding payload: %+v\n", err)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}
		fmt.Printf("Login Payload: %+v\n", payload)

		newTodo.Title = payload.Title
		newTodo.Content = payload.Content
		newTodo.ID = uint(len(TestTodos) + 1)
		TestTodos = append(TestTodos, newTodo)

		response.Valid = true
		ctx.JSON(http.StatusOK, response)
	}
}
