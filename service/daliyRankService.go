package service

import (
	"btube/conf"
	"btube/model"
	"btube/serializer"
	"strconv"
)

//DailyRankService is use to return  daily rank.
type DailyRankService struct{}

//Rank return daily rank.
func (srv *DailyRankService)Rank() *serializer.Response {
	var videos []model.Video
	vids, err := conf.RedisConnect.ZRevRange("b-tube::todyview", 0, 9).Result()
	if err != nil {
		return &serializer.Response{
			StatusCode: 4001,
			Msg:        "error when try to ge rank of the day.",
		}
	}
	for _, v := range vids {
		var vid model.Video
		id, _ := strconv.Atoi(v)
		conf.MySQLConnect.Where("id = ?", id).Find(&vid)
		videos = append(videos, vid)
	}
	sVideos := serializer.BuildVideos(videos)
	return &serializer.Response{
		StatusCode: 0,
		Data:       sVideos,
	}
}
