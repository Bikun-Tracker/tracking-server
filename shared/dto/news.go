package dto

import "time"

type (
	News struct {
		ID        uint      `gorm:"primaryKey;autoIncrement"`
		Title     string    `gorm:"column:title"`
		Detail    string    `gorm:"column:detail"`
		CreatedAt time.Time `gorm:"column:createdAt"`
	}

	// CreateNewsDto CreateNewsDto
	CreateNewsDto struct {
		Title  string `json:"title" validate:"required"`
		Detail string `json:"detail" validate:"required"`
	}

	// CreateNewsResponse CreateNewsResponse
	CreateNewsResponse struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Detail    string `json:"detail"`
		CreatedAt string `json:"createdAt"`
	}
)

func (n *News) ToCreateNewsResponse() CreateNewsResponse {
	return CreateNewsResponse{
		ID:        n.ID,
		Title:     n.Title,
		Detail:    n.Detail,
		CreatedAt: n.CreatedAt.String(),
	}
}
