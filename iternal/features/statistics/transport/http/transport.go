package statistics_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/mihail2771/todogo/iternal/core/domain"
	core_http_server "github.com/mihail2771/todogo/iternal/core/transport/http/server"
)

type StatisticsHTTPHandler struct {
	statisticSevice StatisticsService
}

type StatisticsService interface {
	GetStatistics(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) (domain.Staistics, error)
}

func NewStatisticHTTPHandler(
	statisticSevice StatisticsService,
) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		statisticSevice: statisticSevice,
	}
}

func (h *StatisticsHTTPHandler) Routers() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}
