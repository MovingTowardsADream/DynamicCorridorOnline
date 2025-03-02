package dto

import (
	"time"
)

type SuccessResponse struct {
	Description string `json:"description"`
}

type SignUpResp struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type SignInResp struct {
	Token string `json:"token"`
}
