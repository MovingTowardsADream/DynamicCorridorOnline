package auth

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"TicTacToe/internal/domain/models"
	"TicTacToe/internal/interfaces/convert"
	"TicTacToe/internal/interfaces/dto"

	httperr "TicTacToe/internal/infrastructure/controller/http/errors"
)

type EditInfo interface {
	CreateUser(ctx context.Context, userData *dto.UserData) (*models.User, error)
	GenerateToken(ctx context.Context, userData *dto.UserData) (*dto.AuthToken, error)
}

type AuthsRoutes struct {
	userInfo EditInfo
	log      *slog.Logger
}

func NewAuthRoutes(log *slog.Logger, handler *gin.RouterGroup, userInfo EditInfo) {
	r := &AuthsRoutes{userInfo, log}

	h := handler.Group("/auth")
	{
		h.POST("sign-up", r.signUp)
		h.POST("sign-in", r.signIn)
	}
}

func (u *AuthsRoutes) signUp(c *gin.Context) {
	var input dto.SignUpParams

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{})
		return
	}

	userData := convert.SignUpParamsToUserData(&input)

	user, err := u.userInfo.CreateUser(c, userData)

	if err != nil {
		c.JSON(httperr.MapErrors(err), map[string]interface{}{})
		return
	}

	type respond struct {
		ID        int64
		Username  string
		CreatedAt time.Time
	}

	c.JSON(http.StatusOK, respond{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	})
}

func (u *AuthsRoutes) signIn(c *gin.Context) {
	var input dto.SignInParams

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{})
		return
	}

	userData := convert.SignInParamsToUserData(&input)

	auth, err := u.userInfo.GenerateToken(c, userData)

	if err != nil {
		c.JSON(httperr.MapErrors(err), map[string]interface{}{})
		return
	}

	type respond struct {
		token string
	}

	c.JSON(http.StatusOK, respond{
		token: auth.Token,
	})
}
