package statistics_service

import (
	"context"
	"time"

	"github.com/mihail2771/todogo/internal/core/domain"
)

type StatisticsService struct {
	repositoryService StatisticsRepository
}

type StatisticsRepository interface {
	GetTasks(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) ([]domain.Task, error)
}

func NewStatisticsService(
	repositoryService StatisticsRepository,
) *StatisticsService {
	return &StatisticsService{
		repositoryService: repositoryService,
	}
}
