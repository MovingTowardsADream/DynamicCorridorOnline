package dto

type SignUpParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
