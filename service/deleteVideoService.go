package service

import (
	"btube/conf"
	"btube/model"
	"btube/serializer"
)

//DeleteVideoService delete video.
type DeleteVideoService struct{}

//Delete delete video.
func (srv *DeleteVideoService) Delete(id string) *serializer.Response {
	if err := conf.MySQLConnect.Where("id = ?", id).Delete(&model.Video{}).Error; err != nil {
		return &serializer.Response{
			StatusCode: 40001,
			Msg:        "Delete Failed.",
		}
	}

	return &serializer.Response{
		StatusCode: 0,
		Msg:        "Delete Successfully.",
	}
}
