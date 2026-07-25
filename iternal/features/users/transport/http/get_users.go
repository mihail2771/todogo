package user_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/mihail2771/todogo/iternal/core/logger"
	core_http_request "github.com/mihail2771/todogo/iternal/core/transport/http/request"
	core_http_response "github.com/mihail2771/todogo/iternal/core/transport/http/response"
)

type GetUsersResponse []UserDTOResponse

// GetUsers 	godoc
// @Summary 	Список пользователей
// @Description Список пользователей с пагинацией
// @Tags 		users
// @Produce 	json
// @Param 		limit query int false "Размер страницы с пользователями"
// @Param 		offset query int false "Смещение страницы с пользователями"
// @Success 	200 {object} GetUsersResponse "Список пользователей"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/users [get]
func (h *UserHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	log.Debug("invoce GetUsers handler")
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit and offset")
		return
	}

	userDomains, err := h.userService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
		return
	}

	response := GetUsersResponse(usersDTOFromDomains(userDomains))

	responseHandler.JSONResponse(response, http.StatusOK)

}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)
	limit, err := core_http_request.GetIntQueryParams(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get limit: %w", err)
	}
	offset, err := core_http_request.GetIntQueryParams(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get offset: %w", err)
	}
	return limit, offset, nil
}
