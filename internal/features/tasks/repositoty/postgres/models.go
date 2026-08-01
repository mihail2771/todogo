package task_postgres_repository

import (
	"time"

	"github.com/mihail2771/todogo/internal/core/domain"
)

type TaskModel struct {
	ID      int
	Version int

	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time

	AuthorUserID int
}

func TaskDomainsFromModels(taskModels []TaskModel) []domain.Task {
	taskDomains := make([]domain.Task, len(taskModels))
	for i, taskModels := range taskModels {
		taskDomains[i] = domain.NewTask(
			taskModels.ID,
			taskModels.Version,
			taskModels.Title,
			taskModels.Description,
			taskModels.Completed,
			taskModels.CreatedAt,
			taskModels.CompletedAt,
			taskModels.AuthorUserID,
		)
	}
	return taskDomains

}
