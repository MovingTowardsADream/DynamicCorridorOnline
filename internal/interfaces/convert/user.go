package convert

import (
	"TicTacToe/internal/interfaces/dto"
)

func SignUpParamsToUserData(input *dto.SignUpParams) *dto.UserData {
	return &dto.UserData{
		Username: input.Username,
		Password: input.Password,
	}
}

func SignInParamsToUserData(input *dto.SignInParams) *dto.UserData {
	return &dto.UserData{
		Username: input.Username,
		Password: input.Password,
	}
}
