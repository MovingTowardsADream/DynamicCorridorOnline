package dto

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDataHash struct {
	Username string `json:"username"`
	PassHash string `json:"pass_hash"`
}

type Identify struct {
	ID string `json:"id"`
}

type AuthToken struct {
	Token string `json:"token"`
}
