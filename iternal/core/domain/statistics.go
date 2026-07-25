package domain

import "time"

type Statistics struct {
	TaskCreated                int
	TasksCompleted             int
	TasksCompletedRate         *float64
	TasksAverageCompletionTime *time.Duration
}

func NewStatistics(
	taskCreated int,
	tasksCompleted int,
	tasksCompletedRate *float64,
	tasksAverageCompletionTime *time.Duration,
) Statistics {
	return Statistics{
		TaskCreated:                taskCreated,
		TasksCompleted:             tasksCompleted,
		TasksCompletedRate:         tasksCompletedRate,
		TasksAverageCompletionTime: tasksAverageCompletionTime,
	}
}
