package task_postgres_repository

import (
	core_postgres_pool "github.com/mihail2771/todogo/iternal/core/repository/postgres/pool"
)

type TasksRepository struct {
	pool core_postgres_pool.Pool
}

func NewTasksRepository(
	pool core_postgres_pool.Pool,
) *TasksRepository {
	return &TasksRepository{
		pool: pool,
	}
}
