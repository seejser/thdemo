package utils

import "time"

// NowPtr 返回当前时间指针
func NowPtr() *time.Time {
	t := time.Now()
	return &t
}
