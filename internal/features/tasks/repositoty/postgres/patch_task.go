package task_postgres_repository

import (
	"context"
	"fmt"

	"github.com/mihail2771/todogo/internal/core/domain"
	core_errors "github.com/mihail2771/todogo/internal/core/errors"
	core_postgres_pool "github.com/mihail2771/todogo/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) PatchTask(
	ctx context.Context,
	id int,
	task domain.Task,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `
	UPDATE todoapp.tasks
	SET
		title=$1,
		description=$2,
		completed=$3,
		completed_at=$4,
		author_user_id=$5,
		version=version+1
	WHERE id=$6 and version=$7
	RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id
	`
	row := r.pool.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.CompletedAt,
		task.AuthorUserID,
		id,
		task.Version,
	)

	var taskModel TaskModel

	if err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorUserID,
	); err != nil {
		if err == core_postgres_pool.ErrNoRows {
			return domain.Task{}, fmt.Errorf(
				"task with id='%d' concurrantly accssed: %w",
				id,
				core_errors.ErrConflict,
			)
		}

		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	taskDomain := domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorUserID,
	)

	return taskDomain, nil
}
