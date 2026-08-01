package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	core_errors "github.com/mihail2771/todogo/iternal/core/errors"
)

func GetIntQueryParams(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}
	val, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf(
			"param %s by key=%s not a valid integer: %%v: %w",
			param,
			key,
			core_errors.ErrInvalidArgument,
		)
	}
	return &val, nil
}

func GetDateQueryParams(r *http.Request, key string) (*time.Time, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}
	val, err := time.Parse("2006-01-02", param)
	if err != nil {
		return nil, fmt.Errorf(
			"param %s by key=%s not a valid date: %%v: %w",
			param,
			key,
			core_errors.ErrInvalidArgument,
		)
	}
	return &val, nil
}
