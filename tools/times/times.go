package times

import (
	"fmt"
	"time"
)

// HourAppoint 获取今日指定小时时间戳
func HourAppoint(hour int) (timestamp int64) {
	now := time.Now()
	timestamp = time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, now.Location()).Unix()
	return
}

func Countdown() (countdown int64) {
	now := time.Now()
	timestamp := time.Date(now.Year(), now.Month(), now.Day(), 24, 0, 0, 0, now.Location()).Unix()
	fmt.Println(timestamp)
	return (timestamp - now.Unix())
}
