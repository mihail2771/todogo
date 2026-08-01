package users_service

import (
	"context"
	"fmt"

	"github.com/mihail2771/todogo/internal/core/domain"
	core_errors "github.com/mihail2771/todogo/internal/core/errors"
)

func (s *UserService) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit cannot be negative %sw",
			core_errors.ErrInvalidArgument,
		)
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset cannot be negative %w",
			core_errors.ErrInvalidArgument,
		)
	}

	users, err := s.userRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}

	return users, nil
}
