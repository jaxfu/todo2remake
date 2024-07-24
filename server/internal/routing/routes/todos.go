package routes

import (
	"fmt"
	"furrj/todo_2_remake/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

var testTodos = []types.Todo{
	{ID: 5, Title: "test5", Content: "test5"},
	{ID: 2, Title: "test2", Content: "test2"},
	{ID: 3, Title: "test3", Content: "test3"},
	{ID: 1, Title: "test", Content: "test"},
	{ID: 4, Title: "test4", Content: "test4"},
}

func GetTodos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		testTodos = quickSortTodosByID(testTodos)
		fmt.Printf("%+v\n", testTodos)
		ctx.JSON(http.StatusOK, testTodos)
	}
}

func quickSortTodosByID(todos []types.Todo) []types.Todo {
	if todos == nil || len(todos) <= 1 {
		return todos
	}

	quicksort(todos, 0, len(todos)-1)
	return todos
}

func quicksort(todos []types.Todo, low, high int) {
	if low < high {
		p := partition(todos, low, high)
		quicksort(todos, low, p-1)
		quicksort(todos, p+1, high)
	}
}

func partition(todos []types.Todo, low, high int) int {
	pivot := todos[high].ID
	i := low - 1

	for j := low; j < high; j++ {
		if todos[j].ID < pivot {
			i++
			todos[i], todos[j] = todos[j], todos[i]
		}
	}

	todos[i+1], todos[high] = todos[high], todos[i+1]
	return i + 1
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

		updateTodo(testTodos, payload)
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
