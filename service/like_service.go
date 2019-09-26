package service

import (
	"btube/conf"
	"btube/serializer"
	"os"
	"strings"
)

//LikeService is use to operate thumbsup ad thumbsdouwn.
type LikeService struct {
	UserID  string
	VideoID string
	Code    int
}

//Operation is use to deal with user's operation for like or dislike
//a video.
func (thumb *LikeService) Operation() *serializer.Response {
	key := buildKey(thumb.UserID, thumb.VideoID)
	if err := conf.RedisConnect.HSet(os.Getenv("REDIS_LIKE_KEY"), key, thumb.Code).Err(); err != nil {
		return &serializer.Response{
			StatusCode: 4002,
			Msg:        "Operation falid.",
		}
	}
	return &serializer.Response{
		StatusCode: 0,
		Msg:        "Operation successful.",
	}
}

func buildKey(userID string, videoID string) string {
	builder := &strings.Builder{}
	builder.WriteString(userID)
	builder.WriteString("::")
	builder.WriteString(videoID)
	return builder.String()
}
