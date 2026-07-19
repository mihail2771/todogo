package user_transport_http

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/mihail2771/todogo/iternal/core/domain"
	core_errors "github.com/mihail2771/todogo/iternal/core/errors"
	core_logger "github.com/mihail2771/todogo/iternal/core/logger"
	core_http_request "github.com/mihail2771/todogo/iternal/core/transport/http/request"
	core_http_response "github.com/mihail2771/todogo/iternal/core/transport/http/response"
	core_http_types "github.com/mihail2771/todogo/iternal/core/transport/http/types"
	core_http_utils "github.com/mihail2771/todogo/iternal/core/transport/http/utils"
)

type PatchUserRequst struct {
	FullName    core_http_types.Nullebel[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullebel[string] `json:"phone_number"`
}

func (r *PatchUserRequst) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("Fullname can't be NULL")
		}

		fullNameLen := len([]rune(*r.FullName.Value))
		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("Fullname length must be between 3 and 100")
		}
	}

	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			re := regexp.MustCompile(`^\+[0-9]{9,14}$`)
			if !re.MatchString(*r.PhoneNumber.Value) {
				return fmt.Errorf(
					"invalid phone number format %s:%w",
					*r.PhoneNumber.Value,
					core_errors.ErrInvalidArgument,
				)
			}
		}
	}

	return nil
}

type PatchUserResponse UserDTOResponse

func (h *UserHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	log.Debug("invoce PatchUser handler")
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)
		return
	}

	var request PatchUserRequst
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.userService.PatchUser(ctx, *userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)
		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JSONResponse(response, http.StatusOK)

}

func userPatchFromRequest(request PatchUserRequst) domain.UserPatch {
	return domain.UserPatch{
		FullName:    request.FullName.ToDomain(),
		PhoneNumber: request.PhoneNumber.ToDomain(),
	}
}
