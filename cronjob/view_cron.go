package cron

import (
	"btube/conf"
	"btube/model"
	"strconv"
)

//DeleteTodayView delete today's view and polus it to total view.
func DeleteTodayView() error {
	s, err := conf.RedisConnect.ZRange("b-tube::todyview", 0, -1).Result()
	if err != nil {
		return err
	}
	for _, v := range s {
		score, err := conf.RedisConnect.ZScore("b-tube::todyview", v).Result()
		if err != nil {
			return err
		}

		var video model.Video
		id, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		video.ID = uint(id)
		conf.MySQLConnect.First(&video)
		conf.MySQLConnect.Model(&video).Update("view", uint64(score)+video.View)
		conf.RedisConnect.ZRem("b-tube::todyview", v)
	}
	return nil
}
