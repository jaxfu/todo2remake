package types

type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestRegister struct {
	Username       string `json:"username"`
	FirstPassword  string `json:"first_password"`
	SecondPassword string `json:"second_password"`
}
