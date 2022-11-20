package news

import (
	"time"
	"tracking-server/application"
	"tracking-server/shared"
	"tracking-server/shared/dto"
)

type (
	ViewService interface {
		CreateNews(data dto.CreateNewsDto) (dto.CreateNewsResponse, error)
	}
	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) CreateNews(data dto.CreateNewsDto) (dto.CreateNewsResponse, error) {
	var (
		news     *dto.News
		response dto.CreateNewsResponse
	)

	news = &dto.News{
		Title:     data.Title,
		Detail:    data.Detail,
		CreatedAt: time.Now(),
	}

	err := v.application.NewsService.Create(news)
	if err != nil {
		v.shared.Logger.Errorf("error when inserting news to database, err: %s", err.Error())
		return response, err
	}

	response = news.ToCreateNewsResponse()

	return response, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}
