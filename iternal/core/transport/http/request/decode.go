package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	core_errors "github.com/mihail2771/todogo/iternal/core/errors"
)

var requestValidator = validator.New()

type validatable interface {
	Validate() error
}

func DecodeAndValidateRequest(r *http.Request, dest any) error {

	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		fmt.Println(r.Body)
		return fmt.Errorf(
			"decode json: %v:%w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	var (
		err error
	)

	v, ok := dest.(validatable)
	if ok {
		err = v.Validate()
	} else {
		err = requestValidator.Struct(dest)
	}

	if err != nil {
		return fmt.Errorf(
			"validate request: %v:%w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
