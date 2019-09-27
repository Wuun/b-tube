package service

import (
	"btube/conf"
	"btube/model"
	"btube/serializer"
)

//VideoDetailService use to show video detail.
type VideoDetailService struct{}

//Show show detail.
func (srv *VideoDetailService) Show(id string) *serializer.Response {
	var (
		videoModel model.Video
		video      serializer.Video
	)
	if err := conf.MySQLConnect.Where("id = ?", id).Find(&videoModel).Error; err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "error w hen try to get video.",
		}
	}
	video = serializer.BuildVideo(videoModel)
	//add 1 ciew to redis for today's view.
	videoModel.AddView()
	return &serializer.Response{
		StatusCode: 0,
		Data:       video,
	}
}
