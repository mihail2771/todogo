package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/mihail2771/todogo/iternal/core/domain"
	core_errors "github.com/mihail2771/todogo/iternal/core/errors"
)

func (s *StatisticsService) GetStatistics(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) (domain.Staistics, error) {

	if from != nil && to != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Staistics{}, fmt.Errorf(
				"`to` must be after `from`: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	tasks, err := s.repositoryService.GetTasks(ctx, userID, from, to)
	if err != nil {
		return domain.Staistics{}, fmt.Errorf("get tasks from repository: %w", err)
	}

	statistics := calcStatistics(tasks)

	return statistics, nil

}

func calcStatistics(tasks []domain.Task) domain.Staistics {
	if len(tasks) == 0 {
		return domain.NewStaistics(0, 0, nil, nil)
	}

	tasksCreated := len(tasks)
	tasksCompleted := 0
	var totalCompletedDuration time.Duration
	for _, task := range tasks {
		if task.Completed {
			tasksCompleted++

		}

		completetionDuration := task.CompletetionDuration()
		if completetionDuration != nil {
			totalCompletedDuration += *completetionDuration
		}

	}

	tasksCompletedRate := float64(tasksCompleted) / float64(tasksCreated) * 100

	var tasksAverageCompletionTime *time.Duration
	if tasksCompleted > 0 {
		avg := totalCompletedDuration / time.Duration(tasksCompleted)
		tasksAverageCompletionTime = &avg
	}

	return domain.NewStaistics(
		tasksCreated,
		tasksCompleted,
		&tasksCompletedRate,
		tasksAverageCompletionTime,
	)
}
