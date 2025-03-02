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
	usersTable = "users"
	expTable   = "statistic"
)

type UsersRepo struct {
	storage *postgres.Postgres
}

func New(storage *postgres.Postgres) *UsersRepo {
	return &UsersRepo{storage: storage}
}

func (u *UsersRepo) CreateUser(ctx context.Context, userData *dto.UserDataHash) (*models.User, error) {
	const op = "UserRepo.CreateUser"

	tx, err := u.storage.Pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s - u.Pool.Begin: %w", op, mapping.MapErrors(err))
	}
	defer func() { _ = tx.Rollback(ctx) }()

	sql, args, err := u.storage.Builder.
		Insert(usersTable).
		Columns("username", "pass_hash").
		Values(userData.Username, userData.PassHash).
		Suffix("RETURNING id, username, created_at").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("%s - u.storage.Builder: %w", op, mapping.MapErrors(err))
	}

	var user models.User

	err = u.storage.Pool.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Username, &user.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("%s - u.storage.Pool.QueryRow: %w", op, mapping.MapErrors(err))
	}

	sql, args, err = u.storage.Builder.
		Insert(expTable).
		Columns("exp_value", "user_id").
		Values(0, user.ID).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("%s - u.storage.Builder: %w", op, mapping.MapErrors(err))
	}

	_, err = u.storage.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return nil, fmt.Errorf("%s - u.storage.Pool.Exec: %w", op, mapping.MapErrors(err))
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s - tx.Commit: %w", op, mapping.MapErrors(err))
	}

	return &user, nil
}

func (u *UsersRepo) GetUserID(ctx context.Context, userData *dto.UserDataHash) (*dto.Identify, error) {
	sql, args, _ := u.storage.Builder.
		Select("id").
		From(usersTable).
		Where("username = ? and pass_hash = ?", userData.Username, userData.PassHash).
		ToSql()

	var userIdentify dto.Identify

	err := u.storage.Pool.QueryRow(ctx, sql, args...).Scan(&userIdentify.ID)

	if err != nil {
		return nil, fmt.Errorf("UserRepo.GetUser - ur.storage.Pool.QueryRow: %w", mapping.MapErrors(err))
	}

	return &userIdentify, nil
}
