package cache

import (
	"fmt"
	"strconv"
)

const (
	// DailyRankKey rank of the day
	DailyRankKey = "rank:daily"
)

// VideoViewKey key of count of video  has been visisted.
// view:video:1 -> 100
// view:video:2 -> 150
func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}
