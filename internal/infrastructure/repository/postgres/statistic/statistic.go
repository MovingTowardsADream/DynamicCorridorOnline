package statistic

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	repoerr "TicTacToe/internal/infrastructure/repository/errors"
	"TicTacToe/internal/infrastructure/repository/postgres"
	"TicTacToe/internal/infrastructure/repository/postgres/mapping"
	"TicTacToe/internal/interfaces/dto"
)

const (
	usersTable     = "users"
	statisticTable = "statistic"
)

type StatisticsRepo struct {
	storage *postgres.Postgres
}

func New(storage *postgres.Postgres) *StatisticsRepo {
	return &StatisticsRepo{storage: storage}
}

func (r *StatisticsRepo) GetExpByUserId(ctx context.Context, identify *dto.Identify) (*dto.UserExp, error) {
	sql, args, _ := r.storage.Builder.
		Select("users.id", "users.username", "statistic.exp_value").
		From(usersTable).
		InnerJoin("statistic on users.id = statistic.user_id").
		Where("users.id = ?", identify.ID).
		ToSql()

	var userExp dto.UserExp

	err := r.storage.Pool.QueryRow(ctx, sql, args).Scan(
		&userExp.ID,
		&userExp.Username,
		&userExp.ExpValue,
	)

	if err != nil {
		return nil, fmt.Errorf("StatisticsRepo.GetExpByUserId - r.storage.Pool.QueryRow: %w", mapping.MapErrors(err))
	}

	return &userExp, nil
}

func (r *StatisticsRepo) EditExpByUserId(ctx context.Context, updater *dto.UpdateExp) error {
	sql, args, _ := r.storage.Builder.
		Update(statisticTable).
		Set("exp_value", updater.ExpValue).
		Where("user_id = ?", updater.ID).
		ToSql()

	row, err := r.storage.Pool.Exec(ctx, sql, args)

	if err != nil {
		return fmt.Errorf("StatisticsRepo.EditExpByUserId - r.storage.Pool.QueryRow: %w", mapping.MapErrors(err))
	}

	if row.RowsAffected() == 0 {
		return fmt.Errorf("StatisticsRepo.EditExpByUserId - r.storage.Pool.QueryRow: %w", repoerr.ErrNotFound)
	}

	return nil
}

func (r *StatisticsRepo) DeltaExpByUserId(ctx context.Context, adder *dto.AddExp) error {
	const op = "StatisticsRepo.DeltaExpByUserId"

	sql, args, _ := r.storage.Builder.
		Update(statisticTable).
		Set("exp_value", squirrel.Expr("LEAST(GREATEST(0, exp_value + ?), 2147483647)", adder.AddExpValue)).
		Where("user_id = ?", adder.ID).
		ToSql()

	row, err := r.storage.Pool.Exec(ctx, sql, args)

	if err != nil {
		return fmt.Errorf("%s - r.storage.Pool.QueryRow: %w", op, mapping.MapErrors(err))
	}

	if row.RowsAffected() == 0 {
		return fmt.Errorf("%s - r.storage.Pool.QueryRow: %w", op, repoerr.ErrNotFound)
	}

	return nil
}

func (r *StatisticsRepo) GetLeaderBoard(ctx context.Context, limitBoard *dto.LimitsBoard) (*dto.LeaderBoard, error) {
	sql, args, _ := r.storage.Builder.
		Select("users.id, users.username, statistic.exp_value").
		From(usersTable).
		InnerJoin("statistic on users.id = statistic.user_id").
		OrderBy("statistic.exp_value desc").
		Limit(limitBoard.Limit).
		ToSql()

	rows, err := r.storage.Pool.Query(ctx, sql, args)

	if err != nil {
		return nil, fmt.Errorf("StatisticsRepo.EditExpByUserId - r.storage.Pool.QueryRow: %w", mapping.MapErrors(err))
	}

	var leaderBoard dto.LeaderBoard

	for rows.Next() {
		var userExp dto.UserExp
		if err := rows.Scan(&userExp.ID, &userExp.Username, &userExp.ExpValue); err != nil {
			return nil, fmt.Errorf("StatisticsRepo.EditExpByUserId - rows.Next: %w", mapping.MapErrors(err))
		}
		leaderBoard.Leaders = append(leaderBoard.Leaders, userExp)
	}

	return &leaderBoard, nil
}
