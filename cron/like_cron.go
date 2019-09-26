package cron

import (
	"btube/conf"
	"btube/model"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/robfig/cron"
)

//Like is the crontab of like operation.
func Like() {
	c := cron.New()
	c.AddFunc("0 /5 * * * *", likeOperation)
	c.Start()
}

func likeOperation() {
	hashSet, err := conf.RedisConnect.HGetAll(os.Getenv("REDIS_LIKE_KEY")).Result()
	if err != nil {
		log.Println(err)
	}
	for k, v := range hashSet {
		err = logLikeData(k, v)
		if err != nil {
			log.Println(err)
			continue
		}
		conf.RedisConnect.HDel(os.Getenv("REDIS_LIKE_KEY"), k)
	}
}

func logLikeData(s string, v string) error {
	var likeModel model.Like
	kvPair := strings.Split(s, "::")
	value, err := strconv.Atoi(v)
	if err != nil {
		return err
	}
	uid, err := strconv.Atoi(kvPair[0])
	likeModel.UserID = uint(uid)
	if err != nil {
		return err
	}
	vid, err := strconv.Atoi(kvPair[1])
	likeModel.VideoID = uint(vid)
	if err != nil {
		return err
	}
	err = conf.MySQLConnect.Where("user_id = ? and video_id = ?", uid, vid).Delete(likeModel).Error
	if err != nil {
		return err
	}
	switch value {
	case 0, 1:
		err = conf.MySQLConnect.Create(likeModel).Error
		if err != nil {
			return err
		}
		return nil
	case 2, 3:
		return nil
	default:
		return errors.New("not such type of like")
	}
}
