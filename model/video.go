package model

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"os"
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
