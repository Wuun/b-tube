package serializer

import (
	"btube/model"
)

// Video is the serialzer of video.
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	URL       string `json:"url"`
	Avatar    string `json:"avatar"`
	View      uint64 `json:"view"`
	Author    string `json:"author"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo serialze the video.
func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		URL:       item.VideoURL(),
		Avatar:    item.AvatarURL(),
		View:      item.View(),
		Author:    item.Author,
		CreatedAt: item.CreatedAt.Unix(),
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
