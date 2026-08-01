package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/mihail2771/todogo/iternal/core/errors"
)

func GetIntPathValue(r *http.Request, key string) (*int, error) {
	PathValue := r.PathValue(key)
	if PathValue == "" {
		return nil, fmt.Errorf(
			"PathValue %s by key=%s not a valid integer: %w",
			PathValue,
			key,
			core_errors.ErrInvalidArgument,
		)
	}
	val, err := strconv.Atoi(PathValue)
	if err != nil {
		return nil, fmt.Errorf(
			"PathValue %s by key=%s not a valid integer: %v: %w",
			PathValue,
			key,
			err,
			core_errors.ErrInvalidArgument,
		)
	}
	return &val, nil
}
