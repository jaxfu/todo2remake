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

func GetTodos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		TestTodos = quickSortTodosByID(TestTodos)
		fmt.Printf("%+v\n", TestTodos)
		ctx.JSON(http.StatusOK, TestTodos)
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
