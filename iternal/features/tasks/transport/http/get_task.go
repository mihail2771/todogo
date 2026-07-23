package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/mihail2771/todogo/iternal/core/logger"
	core_http_request "github.com/mihail2771/todogo/iternal/core/transport/http/request"
	core_http_response "github.com/mihail2771/todogo/iternal/core/transport/http/response"
)

type GetTaskResponse TaskDTOResponse

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
