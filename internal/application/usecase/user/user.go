package auth

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"

	"TicTacToe/internal/domain/models"
	"TicTacToe/internal/interfaces/dto"
	"TicTacToe/pkg/hasher"
	"TicTacToe/pkg/logger"

	usecaseerr "TicTacToe/internal/application/usecase/errors"
)

type Users interface {
	CreateUser(ctx context.Context, userData *dto.UserDataHash) (*models.User, error)
	GetUserID(ctx context.Context, identify *dto.UserDataHash) (*dto.Identify, error)
}

type UsersInfo struct {
	usersData  Users
	log        logger.Logger
	hash       hasher.PasswordHash
	tokenTLL   time.Duration
	signingKey string
}

func New(log logger.Logger, hash hasher.PasswordHash, usersData Users, tokenTLL time.Duration, signingKey string) *UsersInfo {
	return &UsersInfo{
		usersData:  usersData,
		hash:       hash,
		log:        log,
		tokenTLL:   tokenTLL,
		signingKey: signingKey,
	}
}

func (e *UsersInfo) CreateUser(ctx context.Context, userData *dto.UserData) (*models.User, error) {
	const op = "UsersInfo - CreateUser"

	userDataHash := &dto.UserDataHash{
		Username: userData.Username,
		PassHash: e.hash.Hash(userData.Password),
	}

	user, err := e.usersData.CreateUser(ctx, userDataHash)

	if err != nil {
		if ok, err := usecaseerr.MapErrors(err); ok {
			return nil, err
		}

		e.log.Error(op, e.log.Err(err))

		return nil, err
	}

	return user, nil
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

func (e *UsersInfo) GenerateToken(ctx context.Context, userData *dto.UserData) (*dto.AuthToken, error) {
	const op = "UsersInfo - GenerateToken"

	userDataHash := &dto.UserDataHash{
		Username: userData.Username,
		PassHash: e.hash.Hash(userData.Password),
	}

	identify, err := e.usersData.GetUserID(ctx, userDataHash)

	if err != nil {
		if ok, err := usecaseerr.MapErrors(err); ok {
			return nil, err
		}

		e.log.Error(op, e.log.Err(err))

		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(e.tokenTLL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		identify.ID,
	})

	tokenWithSignKey, err := token.SignedString([]byte(e.signingKey))

	if err != nil {
		e.log.Error(op, e.log.Err(err))

		return nil, usecaseerr.ErrAddSignKey
	}

	return &dto.AuthToken{Token: tokenWithSignKey}, nil
}
