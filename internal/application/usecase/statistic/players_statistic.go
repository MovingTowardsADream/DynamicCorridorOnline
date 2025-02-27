package statistic

import (
	"context"

	usecaseerr "TicTacToe/internal/application/usecase/errors"
	"TicTacToe/internal/interfaces/dto"
	"TicTacToe/pkg/logger"
)

type Statist interface {
	GetExpByUserId(ctx context.Context, identify *dto.Identify) (*dto.UserExp, error)
	EditExpByUserId(ctx context.Context, updater *dto.UpdateExp) error
	DeltaExpByUserId(ctx context.Context, adder *dto.AddExp) error
	GetLeaderBoard(ctx context.Context, limitBoard *dto.LimitsBoard) (*dto.LeaderBoard, error)
}

type StatisticsInfo struct {
	statistData Statist
	log         logger.Logger
}

func New(log logger.Logger, statistData Statist) *StatisticsInfo {
	return &StatisticsInfo{
		statistData: statistData,
		log:         log,
	}
}

func (s *StatisticsInfo) GetExpByUserId(ctx context.Context, identify *dto.Identify) (*dto.UserExp, error) {
	const op = "StatisticsInfo - GetExpByUserId"

	userExp, err := s.statistData.GetExpByUserId(ctx, identify)

	if err != nil {
		if ok, err := usecaseerr.MapErrors(err); ok {
			return nil, err
		}

		s.log.Error(op, s.log.Err(err))

		return nil, err
	}

	return userExp, nil
}

func (s *StatisticsInfo) EditExpByUserId(ctx context.Context, updater *dto.UpdateExp) error {
	const op = "StatisticsInfo - EditExpByUserId"

	err := s.statistData.EditExpByUserId(ctx, updater)

	if err != nil {
		if ok, err := usecaseerr.MapErrors(err); ok {
			return err
		}

		s.log.Error(op, s.log.Err(err))

		return err
	}

	return nil
}

func (s *StatisticsInfo) DeltaExpByUserId(ctx context.Context, adder *dto.AddExp) error {
	const op = "StatisticsInfo - DeltaExpByUserId"

	err := s.statistData.DeltaExpByUserId(ctx, adder)

	if err != nil {
		if ok, err := usecaseerr.MapErrors(err); ok {
			return err
		}

		s.log.Error(op, s.log.Err(err))

		return err
	}

	return nil
}

func (s *StatisticsInfo) GetLeaderBoard(ctx context.Context, limitBoard *dto.LimitsBoard) (*dto.LeaderBoard, error) {
	const op = "StatisticsInfo - GetLeaderBoard"

	leaderBoard, err := s.statistData.GetLeaderBoard(ctx, limitBoard)

	if err != nil {
		if ok, err := usecaseerr.MapErrors(err); ok {
			return nil, err
		}

		s.log.Error(op, s.log.Err(err))

		return nil, err
	}

	return leaderBoard, nil
}
