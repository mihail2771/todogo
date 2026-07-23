package tasks_service

import (
	"context"
	"fmt"

	"github.com/mihail2771/todogo/iternal/core/domain"
	core_errors "github.com/mihail2771/todogo/iternal/core/errors"
)

func (s *TasksService) GetTasks(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
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

	users, err := s.tasksRepository.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get tasks: %w", err)
	}

	return users, nil
}
