package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	httperr "TicTacToe/internal/infrastructure/controller/http/errors"
	"TicTacToe/internal/interfaces/dto"
	"TicTacToe/internal/interfaces/middleware"
)

const userIDCtx = "UserID"

type StatistInfo interface {
	GetExpByUserId(ctx context.Context, identify *dto.Identify) (*dto.UserExp, error)
	EditExpByUserId(ctx context.Context, updater *dto.UpdateExp) error
	DeltaExpByUserId(ctx context.Context, adder *dto.AddExp) error
	GetLeaderBoard(ctx context.Context, limitBoard *dto.LimitsBoard) (*dto.LeaderBoard, error)
}

type StatistRoutes struct {
	statistInfo StatistInfo
}

func newStatistRoutes(handler *gin.RouterGroup, statistInfo StatistInfo) {
	r := &StatistRoutes{statistInfo: statistInfo}

	statistic := handler.Group("/statistic")
	{
		players := statistic.Group("/players")
		{
			players.GET("", r.GetLeaderBoard)

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
	id, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{})
		return
	}

	userExp, err := r.statistInfo.GetExpByUserId(c.Request.Context(), &dto.Identify{ID: id})

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
	id, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{})
		return
	}

	err = r.statistInfo.EditExpByUserId(c.Request.Context(), &dto.UpdateExp{ID: id})

	if err != nil {
		c.JSON(httperr.MapErrors(err), map[string]interface{}{})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Description: "success",
	})
}

func (r *StatistRoutes) DeltaExperience(c *gin.Context) {
	id, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{})
		return
	}

	var exp dto.AddExpReq

	if err := c.BindJSON(&exp); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{})
		return
	}

	err = r.statistInfo.DeltaExpByUserId(c.Request.Context(), &dto.AddExp{ID: id, AddExpValue: exp.ExpValue})

	if err != nil {
		c.JSON(httperr.MapErrors(err), map[string]interface{}{})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Description: "success",
	})
}

const limitParam = "limit"

func (r *StatistRoutes) GetLeaderBoard(c *gin.Context) {
	limit, err := strconv.ParseUint(c.Query(limitParam), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{})
		return
	}

	userExp, err := r.statistInfo.GetLeaderBoard(c.Request.Context(), &dto.LimitsBoard{Limit: limit})

	if err != nil {
		c.JSON(httperr.MapErrors(err), map[string]interface{}{})
		return
	}

	type respond struct {
		Leaders []dto.UserExp
	}

	c.JSON(http.StatusOK, respond{
		Leaders: userExp.Leaders,
	})
}
