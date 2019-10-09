package service

import (
	"btube/conf"
	"btube/model"
	"btube/serializer"
)

//ListVideoService is use to  list video user want.
type ListVideoService struct {
	Limit int 	`json:"limit" form:"limit" binding:"required"`
	Start int	`json:"start" form:"start" binding:"required"`
}

//List get the videos.
func (srv *ListVideoService) List() *serializer.Response {
	var (
		videos       []model.Video
		videosResult []serializer.Video
	)
	err := conf.MySQLConnect.Limit(srv.Limit).Offset(srv.Start).Find(&videos).Error
	if err != nil {
		return &serializer.Response{
			StatusCode: 40001,
			Msg:        "params is wrong.",
		}
	}

	for _, v := range videos {
		videosResult = append(videosResult, serializer.BuildVideo(v))
	}
	return &serializer.Response{
		StatusCode: 0,
		Data:       videosResult,
	}
}
