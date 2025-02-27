package statistic

import (
	"github.com/gin-gonic/gin"
)

type StatistInfo interface {
}

type StatistRoutes struct {
	statistInfo StatistInfo
}

func NewStatistRoutes(handler *gin.RouterGroup, statistInfo StatistInfo) {
	r := &StatistRoutes{statistInfo: statistInfo}

	h := handler.Group("/statistic")
	{
	}
}
