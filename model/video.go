package model

import (
	"btube/conf"
	"os"
	"strconv"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Video model of video
type Video struct {
	gorm.Model
	Title    string
	Info     string
	Video    string
	Avatar   string
	AuthorID uint
	View     uint64 //view is the view of the day before,and total is the view num in redis plus this one.
}

// AvatarURL avatar of video.
func (video *Video) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(video.Avatar, oss.HTTPGet, 600)
	return signedGetURL
}

// VideoURL url of video.
func (video *Video) VideoURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(video.Video, oss.HTTPGet, 600)
	return signedGetURL
}

//AddView add 1 to today's view in redis.
func (video *Video) AddView() {
	id := strconv.Itoa(int(video.ID))
	if _, err := conf.RedisConnect.ZScore("b-tube::todyview", id).Result(); err != nil {
		conf.RedisConnect.ZAdd("b-tube::todyview", redis.Z{Score: 0, Member: id})
	}
	conf.RedisConnect.ZIncrBy("b-tube::todyview", 1, id)
}
