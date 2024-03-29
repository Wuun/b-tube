package service

import (
	"btube/conf"
	"btube/model"
	"btube/serializer"
)

//UploadVideoService is used to recoding video information to mysql.
type UploadVideoService struct {
	Title    string `form:"title"`
	Info     string `form:"info"`
	Video    string `form:"video"`
	Avatar   string `form:"avatar"`
	AuthorID uint   `form:"author_id"`
}

//Upload is ued to recording video.
func (upSrv *UploadVideoService) Upload() *serializer.Response {
	video := &model.Video{
		Title:    upSrv.Title,
		Info:     upSrv.Info,
		Video:    upSrv.Video,
		Avatar:   upSrv.Avatar,
		AuthorID: upSrv.AuthorID,
	}

	if err := conf.MySQLConnect.Create(&video).Error; err != nil {
		return &serializer.Response{
			StatusCode: 40002,
			Msg:        "error when try to recording video information",
		}
	}
	return &serializer.Response{
		StatusCode: 0,
		Msg:        "record successfully.",
	}
}
