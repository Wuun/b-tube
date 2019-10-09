package cron

import (
	"btube/cache"
	"btube/conf"
	"btube/model"
	"strconv"
)

//DeleteTodayView delete today's view and plus it to total view.
func DeleteTodayView() error {
	s, err := cache.RedisConnect.ZRange("b-tube::todyview", 0, -1).Result()
	if err != nil {
		return err
	}
	for _, v := range s {
		score, err := cache.RedisConnect.ZScore("b-tube::todyview", v).Result()
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
		cache.RedisConnect.ZRem("b-tube::todyview", v)
	}
	return nil
}
