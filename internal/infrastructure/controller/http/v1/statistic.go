package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	httperr "TicTacToe/internal/infrastructure/controller/http/errors"
	"TicTacToe/internal/interfaces/dto"
)

type StatistInfo interface {
	GetExpByUserId(ctx context.Context, identify *dto.Identify) (*dto.UserExp, error)
	EditExpByUserId(ctx context.Context, updater *dto.UpdateExp) error
	DeltaExpByUserId(ctx context.Context, adder *dto.AddExp) error
	GetLeaderBoard(ctx context.Context, limitBoard *dto.LimitsBoard) (*dto.LeaderBoard, error)
}

type StatistRoutes struct {
	statistInfo StatistInfo
}

func NewStatistRoutes(handler *gin.RouterGroup, statistInfo StatistInfo) {
	r := &StatistRoutes{statistInfo: statistInfo}

	statistic := handler.Group("/statistic")
	{
		players := statistic.Group("/players")
		{
			players.GET("/leaderboard", r.GetLeaderBoard)

			experience := players.Group("/experience")
			{
				experience.GET("", r.GetExperience)
				experience.PUT("", r.EditExperience)
				experience.PATCH("", r.DeltaExperience)
			}
		}
	}
}

func (r *StatistRoutes) GetExperience(c *gin.Context) {
	identify := &dto.Identify{}

	userExp, err := r.statistInfo.GetExpByUserId(c, identify)

	if err != nil {
		c.JSON(httperr.MapErrors(err), map[string]interface{}{})
		return
	}

	type respond struct {
		ID       string
		Username string
		ExpValue int
	}

	c.JSON(http.StatusOK, respond{
		ID:       userExp.ID,
		Username: userExp.Username,
		ExpValue: userExp.ExpValue,
	})
}

func (r *StatistRoutes) EditExperience(c *gin.Context) {
}

func (r *StatistRoutes) DeltaExperience(c *gin.Context) {

}

func (r *StatistRoutes) GetLeaderBoard(c *gin.Context) {

}
