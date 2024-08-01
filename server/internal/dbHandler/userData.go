package dbHandler

import (
	"context"
	"furrj/todo_2_remake/internal/types"
)

// Gets
const QGetUserIDByUsername = `
	SELECT user_id
	FROM users
	WHERE username=$1
`

func (dbHandler *DBHandler) GetUserIDByUsername(username string) (types.UserID, error) {
	var id types.UserID
	err := dbHandler.Conn.QueryRow(context.Background(), QGetUserIDByUsername, username).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

const QGetUserDataByUserID = `
	SELECT user_id, username, password
	FROM users
	WHERE user_id=$1
`

func (dbHandler *DBHandler) GetUserDataByUserID(UserID types.UserID) (types.UserData, error) {
	var UserData types.UserData
	err := dbHandler.Conn.QueryRow(context.Background(), QGetUserDataByUserID, UserID).Scan(
		&UserData.UserID,
		&UserData.Username,
		&UserData.Password,
	)
	if err != nil {
		return UserData, err
	}
	return UserData, nil
}

const QGetUserDataByUsername = `
	SELECT user_id, username, password
	FROM users
	WHERE username=$1
`

func (dbHandler *DBHandler) GetUserDataByUsername(username string) (types.UserData, error) {
	var UserData types.UserData
	err := dbHandler.Conn.QueryRow(context.Background(), QGetUserDataByUsername, username).Scan(
		&UserData.UserID,
		&UserData.Username,
		&UserData.Password,
	)
	if err != nil {
		return UserData, err
	}
	return UserData, nil
}

// Inserts

const EInsertUser = `
	INSERT INTO users (username, password)
	VALUES ($1, $2)
	RETURNING user_id
`

func (dbHandler *DBHandler) InsertUser(username, password string) (types.UserID, error) {
	var userID types.UserID
	err := dbHandler.Conn.QueryRow(context.Background(), EInsertUser, username, password).Scan(&userID)
	if err != nil {
		return userID, err
	}
	return userID, nil
}
