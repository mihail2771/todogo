package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/mihail2771/todogo/iternal/core/logger"
	core_http_request "github.com/mihail2771/todogo/iternal/core/transport/http/request"
	core_http_response "github.com/mihail2771/todogo/iternal/core/transport/http/response"
)

type DeleteTaskResponse TaskDTOResponse

func (h *TasksHTTPHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responceHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responceHandler.ErrorResponse(err, "failed to get task id path value")
		return
	}

	if err := h.tasksService.DeleteTask(ctx, *taskID); err != nil {
		responceHandler.ErrorResponse(err, "failed to delete task")
		return
	}

	responceHandler.NoContentResponse()
}
