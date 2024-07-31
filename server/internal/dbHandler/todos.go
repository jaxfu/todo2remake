package dbHandler

import (
	"context"
	"furrj/todo_2_remake/internal/types"
)

// GETS

const QGetTodosByUserID = `
	SELECT todo_id, title, content
	FROM todos
	WHERE user_id=$1
`

func (dbHandler *DBHandler) GetTodosByUserID(userID types.UserID) ([]types.Todo, error) {
	var todos []types.Todo

	rows, err := dbHandler.Conn.Query(context.Background(), QGetTodosByUserID, userID)
	if err != nil {
		return todos, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo types.Todo

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Content)
		if err != nil {
			return todos, err
		}

		todos = append(todos, todo)
	}

	if rows.Err() != nil {
		return todos, rows.Err()
	}

	return todos, nil
}
