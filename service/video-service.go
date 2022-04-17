package service

import "github.com/mariarobertap/api-vidroglass/models"

type VideoService interface {

	Save(models.Video) models.Video
	FindAll() []models.Video

}

type videoService struct {
	videos []models.Video
}


func NewVideoService() VideoService{
	return &videoService{}
}

func (service *videoService) Save(video models.Video) models.Video{
	service.videos = append(service.videos, video)

	return video
}

func (service *videoService) FindAll() []models.Video{

	return service.videos
}