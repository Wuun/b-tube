package cache

import (
	"fmt"
	"strconv"
)

const (
	// DailyRankKey rank of the day
	DailyRankKey = "rank:daily"
)

// VideoViewKey key of count of video  has been visited.

func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}
