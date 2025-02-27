package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	userIDCtx           = "userID"
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

	header, ok := c.Get(authorizationHeader)
	if header == "" || !ok {
		return "", false
	}

	headerStr, ok := header.(string)

	if !ok {
		return "", false
	}

	if len(headerStr) > len(prefix) && strings.EqualFold(headerStr[:len(prefix)], prefix) {
		return headerStr[len(prefix):], true
	}

	return "", false
}
