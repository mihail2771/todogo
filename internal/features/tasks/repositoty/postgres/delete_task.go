package task_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/mihail2771/todogo/internal/core/errors"
)

func (r *TasksRepository) DeleteTask(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `
	DELETE FROM todoapp.tasks
	WHERE id = $1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec query: %w,", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("task with id='%d': %w", id, core_errors.ErrNotFound)
	}

	return nil

}
