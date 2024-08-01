package todos

import (
	"fmt"
	"furrj/todo_2_remake/internal/dbHandler"
	"furrj/todo_2_remake/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddTodo accepts a requestAddTodo object then creates and stores a new todo
func AddTodo(db *dbHandler.DBHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// init payload object to allow binding
		var payload types.RequestAddTodo
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
		fmt.Printf("addTodo payload: %+v\n", payload)

		// create new todo object
		newTodo.Title = payload.Title
		newTodo.Content = payload.Content

		// store in db
		if err := db.InsertTodoByUserID(payload.UserID, newTodo); err != nil {
			fmt.Printf("error inserting: %+v\n", err)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}

		// set response to valid and send
		response.Valid = true
		ctx.JSON(http.StatusOK, response)
	}
}
