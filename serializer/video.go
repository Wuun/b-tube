package serializer

import (
	"btube/conf"
	"btube/model"
)

// Video is the serialzer of video.
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	URL       string `json:"url"`
	Avatar    string `json:"avatar"`
	TotalView uint64 `json:"view"`
	Author    uint   `json:"author_id"`
	CreatedAt int64  `json:"created_at"`
	Like      int    `json:"like"`
	Dislike   int    `json:"dislike"`
}

// BuildVideo serialze the video.
func BuildVideo(item model.Video) Video {
	var (
		likeCount    int
		dislikeCount int
	)
	conf.MySQLConnect.Model(&model.Like{}).Find("where video_id = ? and type = 0", item.ID).Count(&likeCount)
	conf.MySQLConnect.Model(&model.Like{}).Find("where video_id = ? and type = 1", item.ID).Count(&dislikeCount)
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		URL:       item.VideoURL(),
		Avatar:    item.AvatarURL(),
		Author:    item.AuthorID,
		CreatedAt: item.CreatedAt.Unix(),
		Like:      likeCount,
		Dislike:   dislikeCount,
	}
}

// BuildVideos serialize a set of video.
func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}
