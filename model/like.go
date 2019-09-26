package model

import "github.com/jinzhu/gorm"

//Like is the like nums add dislike nums of all video
type Like struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	VideoID uint `json:"video_id"`
	Type    int  `json:"type"`
}
