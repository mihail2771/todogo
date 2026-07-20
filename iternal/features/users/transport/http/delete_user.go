package user_transport_http

import (
	"net/http"

	core_logger "github.com/mihail2771/todogo/iternal/core/logger"
	core_http_request "github.com/mihail2771/todogo/iternal/core/transport/http/request"
	core_http_response "github.com/mihail2771/todogo/iternal/core/transport/http/response"
)

type DeleteUserResponse UserDTOResponse

func (h *UserHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	log.Debug("invoce DeleteUser handler")
	responceHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responceHandler.ErrorResponse(err, "failed to get user id path value")
		return
	}

	if err := h.userService.DeleteUser(ctx, *userID); err != nil {
		responceHandler.ErrorResponse(err, "failed to delete user")
		return
	}

	responceHandler.NoContentResponse()
}
