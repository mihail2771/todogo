package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/mihail2771/todogo/internal/core/logger"
	core_http_request "github.com/mihail2771/todogo/internal/core/transport/http/request"
	core_http_response "github.com/mihail2771/todogo/internal/core/transport/http/response"
)

type GetTasksResponse []TaskDTOResponse

// GetTasks 	godoc
// @Summary 	Список задач
// @Description Список задач с пагинацией
// @Tags 		tasks
// @Produce 	json
// @Param 		user_id query int false "Фильтрация по ID автора задачи"
// @Param 		limit query int false "Размер страницы с задачами"
// @Param 		offset query int false "Смещение страницы с задачами"
// @Success 	200 {object} GetTasksResponse "Список задач"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/tasks [get]
func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, limit, offset, err := getUserIdLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get query params",
		)
		return
	}

	tasksDomains, err := h.tasksService.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get tasks")
		return
	}

	response := GetTasksResponse(usersDTOFromDomains(tasksDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIdLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
		userIdQueryParamKey = "user_id"
	)

	userID, err := core_http_request.GetIntQueryParams(r, userIdQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get userID: %w", err)
	}

	limit, err := core_http_request.GetIntQueryParams(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get limit: %w", err)
	}
	offset, err := core_http_request.GetIntQueryParams(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get offset: %w", err)
	}
	return userID, limit, offset, nil
}
