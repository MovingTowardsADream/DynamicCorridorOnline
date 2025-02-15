package user

import (
	"context"
	"fmt"

	"TicTacToe/internal/domain/models"
	"TicTacToe/internal/infrastructure/repository/postgres"
	"TicTacToe/internal/infrastructure/repository/postgres/mapping"
	"TicTacToe/internal/interfaces/dto"
)

const (
	usersTable = "user"
)

type UsersRepo struct {
	storage *postgres.Postgres
}

func New(storage *postgres.Postgres) *UsersRepo {
	return &UsersRepo{storage: storage}
}

func (u *UsersRepo) CreateUser(ctx context.Context, userData *dto.UserDataHash) (*models.User, error) {
	sql, args, _ := u.storage.Builder.
		Insert(usersTable).
		Columns("username", "pass_hash").
		Values(userData.Username, userData.PassHash).
		Suffix("RETURNING id, username, created_at").
		ToSql()

	var user models.User

	err := u.storage.Pool.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Username, &user.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("UserRepo.CreateUser - ur.storage.Pool.QueryRow: %w", mapping.MapErrors(err))
	}

	return &user, nil
}

func (u *UsersRepo) GetUserID(ctx context.Context, userData *dto.UserDataHash) (*dto.Identify, error) {
	sql, args, _ := u.storage.Builder.
		Select("username", "pass_hash").
		From(usersTable).
		Where("username = ? and pass_hash = ?", userData.Username, userData.PassHash).
		Suffix("RETURNING id").
		ToSql()

	var userIdentify dto.Identify

	err := u.storage.Pool.QueryRow(ctx, sql, args).Scan(&userIdentify.ID)

	if err != nil {
		return nil, fmt.Errorf("UserRepo.GetUser - ur.storage.Pool.QueryRow: %w", mapping.MapErrors(err))
	}

	return &userIdentify, nil
}
