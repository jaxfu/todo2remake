// Package dbHandler handles communication with the Postgres database
package dbHandler

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DBHandler object contains a pointer to pgx connection, one per server
type DBHandler struct {
	Conn *pgxpool.Pool
}

// InitDBHandler is the constructor for DBHandler, takes a database connection
// string and returns a pointer to instantiated DBHandler.
func InitDBHandler(connectionString string) (*DBHandler, error) {
	var newDBHandler DBHandler

	db, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return &newDBHandler, err
	}

	newDBHandler.Conn = db
	return &newDBHandler, nil
}
