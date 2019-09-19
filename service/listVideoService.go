package service

import (
	"btube/conf"
	"btube/model"
	"btube/serializer"
)

//ListVideoService is use to  list video user want.
type ListVideoService struct {
	Start int
	limit int
}

//List get the videos.
func (srv *ListVideoService) List() *serializer.Response {
	var (
		videos       []model.Video
		videosResult []serializer.Video
	)
	if err := conf.MySQLConnect.Limit(srv.limit).Offset(srv.Start).Find(&videos).Error; err != nil {
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
