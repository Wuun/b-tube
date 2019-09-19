package model

import (
	"btube/cache"
	"btube/conf"
	"os"
	"strconv"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
)

// Video model of video
type Video struct {
	gorm.Model
	Title     string
	Info      string
	Video     string
	Avatar    string
	AuthorID  uint
	TotalView uint64
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

// View count of video has been visited.
func (video *Video) View() uint64 {
	countStr, _ := conf.RedisConnection.Get(cache.VideoViewKey(video.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView plus one when video is visited.
func (video *Video) AddView() {
	conf.RedisConnection.Incr(cache.VideoViewKey(video.ID))
	conf.RedisConnection.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(video.ID)))
}
