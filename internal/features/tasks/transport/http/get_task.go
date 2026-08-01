package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/mihail2771/todogo/internal/core/logger"
	core_http_request "github.com/mihail2771/todogo/internal/core/transport/http/request"
	core_http_response "github.com/mihail2771/todogo/internal/core/transport/http/response"
)

type GetTaskResponse TaskDTOResponse

// GetTask	 	godoc
// @Summary 	Получение задачи
// @Description Получение задачи
// @Tags 		tasks
// @Accept 		json
// @Param 		id path int true "ID задачи"
// @Success 	200 {object} GetTaskResponse "Задача успешно найдена"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	404 {object} core_http_response.ErrorResponse "User not found"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/tasks/{id} [get]
func (h *TasksHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHangler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHangler.ErrorResponse(
			err,
			"failed to get task id path value",
		)
		return
	}

	task, err := h.tasksService.GetTask(ctx, *taskID)
	if err != nil {
		responseHangler.ErrorResponse(
			err,
			"failed to get task",
		)
		return
	}

	response := GetTaskResponse(TaskDTOResponse(task))

	responseHangler.JSONResponse(response, http.StatusOK)
}
