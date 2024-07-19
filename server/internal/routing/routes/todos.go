package routes

import (
	"fmt"
	"furrj/todo_2_remake/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

var testTodos = []types.Todo{
	{ID: 1, Title: "test", Content: "test"},
	{ID: 2, Title: "test2", Content: "test2"},
	{ID: 3, Title: "test3", Content: "test3"},
}

func GetTodos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, testTodos)
	}
}

type ResponseValid struct {
	Valid bool `json:"valid"`
}

func AddTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload types.Todo
		response := ResponseValid{
			Valid: false,
		}

		if err := ctx.BindJSON(&payload); err != nil {
			fmt.Printf("Error binding payload: %+v\n", err)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}
		fmt.Printf("Login Payload: %+v\n", payload)

		payload.ID = uint(len(testTodos) + 1)
		testTodos = append(testTodos, payload)
		response.Valid = true

		ctx.JSON(http.StatusOK, response)
	}
}

func UpdateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, testTodos)
	}
}

type deletePayload struct {
	ID uint `json:"id"`
}

func DeleteTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload deletePayload
		response := ResponseValid{
			Valid: false,
		}

		if err := ctx.BindJSON(&payload); err != nil {
			fmt.Printf("Error binding payload: %+v\n", err)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}
		fmt.Printf("Delete ID: %d\n", payload.ID)

		testTodos = filter(testTodos, payload.ID)
		response.Valid = true

		ctx.JSON(http.StatusOK, response)
	}
}

func filter(todos []types.Todo, id uint) []types.Todo {
	var filtered []types.Todo

	for _, todo := range todos {
		if todo.ID != id {
			filtered = append(filtered, todo)
		}
	}

	return filtered
}
