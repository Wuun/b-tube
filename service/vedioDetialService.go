package service

import (
	"btube/conf"
	"btube/model"
	"btube/serializer"
)

type VideoDetailService struct{}

func (srv *VideoDetailService) Show(id string) *serializer.Response {
	var (
		videoModel model.Video
	)
	if err := conf.MySQLConnect.Where("id = ?", id).Find(&videoModel).Error; err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "error when try to get video.",
		}
	}
	video := serializer.Video{
		ID:        videoModel.ID,
		Title:     videoModel.Title,
		Info:      videoModel.Info,
		URL:       videoModel.VideoURL(),
		Avatar:    videoModel.AvatarURL(),
		Author:    videoModel.AuthorID,
		CreatedAt: videoModel.CreatedAt.Unix(),
		TotalView: videoModel.TotalView,
	}

	return &serializer.Response{
		StatusCode: 0,
		Data:       video,
	}
}
