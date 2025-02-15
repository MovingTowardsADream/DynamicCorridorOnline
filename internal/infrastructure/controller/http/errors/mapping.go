package errors

import (
	"errors"
	"net/http"

	usecaseerr "TicTacToe/internal/application/usecase/errors"
)

func MapErrors(err error) int {
	switch {
	case errors.Is(err, usecaseerr.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, usecaseerr.ErrTimeout):
		return http.StatusRequestTimeout
	case errors.Is(err, usecaseerr.ErrAlreadyExists):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
