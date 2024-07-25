package todos

import (
	"fmt"
	"furrj/todo_2_remake/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

var TestTodos = []types.Todo{
	{ID: 5, Title: "test5", Content: "test5"},
	{ID: 2, Title: "test2", Content: "test2"},
	{ID: 3, Title: "test3", Content: "test3"},
	{ID: 1, Title: "test", Content: "test"},
	{ID: 4, Title: "test4", Content: "test4"},
}

// GetTodos quicksorts todos by id and returns all todos
// TODO: accept username and get specific user's todos
func GetTodos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// call quicksortTodosByID to sort todos
		TestTodos = quickSortTodosByID(TestTodos)
		fmt.Printf("%+v\n", TestTodos)

		// send response
		ctx.JSON(http.StatusOK, TestTodos)
	}
}

// quickSortTodosByID quick sorts slice of todos by id in place
func quickSortTodosByID(todos []types.Todo) []types.Todo {
	if todos == nil || len(todos) <= 1 {
		return todos
	}

	quicksort(todos, 0, len(todos)-1)
	return todos
}

// quicksort recursively quicksorts slice of todos in place
func quicksort(todos []types.Todo, low, high int) {
	if low < high {
		p := partition(todos, low, high)
		quicksort(todos, low, p-1)
		quicksort(todos, p+1, high)
	}
}

// partition partitions slice of todos in place
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
