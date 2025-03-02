package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"TicTacToe/internal/domain/models"
	httperr "TicTacToe/internal/infrastructure/controller/http/errors"
	"TicTacToe/internal/interfaces/convert"
	"TicTacToe/internal/interfaces/dto"
)

type EditInfo interface {
	CreateUser(ctx context.Context, userData *dto.UserData) (*models.User, error)
	GenerateToken(ctx context.Context, userData *dto.UserData) (*dto.AuthToken, error)
	ParseToken(token string) (string, error)
}

type AuthsRoutes struct {
	userInfo EditInfo
}

func newAuthRoutes(handler *gin.RouterGroup, userInfo EditInfo) {
	r := &AuthsRoutes{userInfo}

	h := handler.Group("/auth")
	{
		h.POST("/sign-up", r.signUp)
		h.POST("/sign-in", r.signIn)
	}
}

func (u *AuthsRoutes) signUp(c *gin.Context) {
	var input dto.SignUpParams

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{})
		return
	}

	userData := convert.SignUpParamsToUserData(&input)

	user, err := u.userInfo.CreateUser(c.Request.Context(), userData)

	if err != nil {
		c.JSON(httperr.MapErrors(err), map[string]interface{}{})
		return
	}

	c.JSON(http.StatusOK, dto.SignUpResp{
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

	auth, err := u.userInfo.GenerateToken(c.Request.Context(), userData)

	if err != nil {
		c.JSON(httperr.MapErrors(err), map[string]interface{}{})
		return
	}

	c.JSON(http.StatusOK, dto.SignInResp{
		Token: auth.Token,
	})
}
