package tasks_transport_http

import (
	"time"

	"github.com/mihail2771/todogo/internal/core/domain"
)

type TaskDTOResponse struct {
	ID           int        `json:"id" example:"15"`
	Version      int        `json:"version" example:"3"`
	Title        string     `json:"title" example:"Название задачи"`
	Description  *string    `json:"description" example:"Описание задачи"`
	Completed    bool       `json:"completed" example:"false"`
	CreatedAt    time.Time  `json:"created_at" example:"2026-02-26T10:30:00Z"`
	CompletedAt  *time.Time `json:"completed_at" example:"null"`
	AuthorUserID int        `json:"author_user_id" example:"2"`
}

func TaskDTOFromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func usersDTOFromDomains(tasks []domain.Task) []TaskDTOResponse {
	userDTO := make([]TaskDTOResponse, len(tasks))
	for i, task := range tasks {
		userDTO[i] = TaskDTOResponse(task)
	}

	return userDTO

}
