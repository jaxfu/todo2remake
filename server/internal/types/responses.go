package types

type ResponseValid struct {
	Valid bool `json:"valid"`
}

type ResponseLoginRegister struct {
	UserID UserID `json:"user_id"`
}
