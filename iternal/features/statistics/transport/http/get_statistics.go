package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mihail2771/todogo/iternal/core/domain"
	core_logger "github.com/mihail2771/todogo/iternal/core/logger"
	core_http_request "github.com/mihail2771/todogo/iternal/core/transport/http/request"
	core_http_response "github.com/mihail2771/todogo/iternal/core/transport/http/response"
)

type GetStatisticsResponse struct {
	TaskCreated                int      `json:"task_created" example:"3"`
	TasksCompleted             int      `json:"task_completed" example:"10"`
	TasksCompletedRate         *float64 `json:"task_completed_rate" example:"20"`
	TasksAverageCompletionTime *string  `json:"task_average_completion_time" example:"1m30s"`
}

// GetStatistics 	godoc
// @Summary 		Получение статистики
// @Description 	Получение статистики с опциональной фильтрацией по user_id и времени создания
// @Tags 			statistics
// @Produce 		json
// @Param 			user_id query int false "Фильтр по автору задачи"
// @Param 			from query string false "Начало времени статистики"
// @Param 			to query string false "Окончание времени статистики"
// @Success 		200 {object} GetStatisticsResponse "Успешное получение статистики"
// @Failure 		400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 		500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 			/statistics [get]
func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, from, to, err := getUserIdFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID/from/to query params",
		)
		return
	}

	statistics, err := h.statisticsService.GetStatistics(ctx, userID, from, to)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get GetStatistics",
		)
		return
	}

	response := GetStatisticsResponse(toDTOFromDomain(statistics))

	responseHandler.JSONResponse(response, http.StatusOK)

}

func toDTOFromDomain(statistics domain.Statistics) GetStatisticsResponse {
	var avgTime *string
	if statistics.TasksAverageCompletionTime != nil {
		duration := statistics.TasksAverageCompletionTime.String()
		avgTime = &duration
	}
	return GetStatisticsResponse{
		TaskCreated:                statistics.TaskCreated,
		TasksCompleted:             statistics.TasksCompleted,
		TasksCompletedRate:         statistics.TasksCompletedRate,
		TasksAverageCompletionTime: avgTime,
	}
}

func getUserIdFromToQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		fromQueryParamKey   = "from"
		toQueryParamKey     = "to"
		userIdQueryParamKey = "user_id"
	)

	userID, err := core_http_request.GetIntQueryParams(r, userIdQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get userID: %w", err)
	}

	from, err := core_http_request.GetDateQueryParams(r, fromQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get from: %w", err)
	}

	to, err := core_http_request.GetDateQueryParams(r, toQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get to: %w", err)
	}
	return userID, from, to, nil
}
