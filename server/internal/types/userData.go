package types

type UserID uint

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserID   UserID `json:"user_id"`
}
