package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	userIDCtx           = "UserID"
	authorizationHeader = "Authorization"
)

type Authorization interface {
	ParseToken(token string) (string, error)
}

type AuthMiddleware struct {
	auth Authorization
}

func (h *AuthMiddleware) UserIdentity() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := bearerToken(c)
		if !ok {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{})
			return
		}

		userID, err := h.auth.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{})
			return
		}

		c.Set(userIDCtx, userID)

		c.Next()
	}
}

func bearerToken(c *gin.Context) (string, bool) {
	const prefix = "Bearer "

	header := c.GetHeader(authorizationHeader)

	if header == "" {
		return "", false
	}

	if len(header) > len(prefix) && strings.EqualFold(header[:len(prefix)], prefix) {
		return header[len(prefix):], true
	}

	return "", false
}

var ErrUserIDEmpty = errors.New("user id empty")

func GetUserID(c *gin.Context) (string, error) {
	id, ok := c.Get(userIDCtx)
	if !ok {
		return "", ErrUserIDEmpty
	}

	return id.(string), nil
}
