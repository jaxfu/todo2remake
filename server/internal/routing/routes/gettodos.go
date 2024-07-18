package routes

import (
	"furrj/todo_2_remake/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

var testTodos = []types.Todo{
	{Title: "test", Content: "test"},
	{Title: "test2", Content: "test2"},
	{Title: "test3", Content: "test3"},
}

func GetTodos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, testTodos)
	}
}
