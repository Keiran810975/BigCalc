package utils

import (
	"time"
)

func GetCurrentTime() string {
	//获取当前时间
	now := time.Now()
	//将时间格式化为字符串
	formattedTime := now.Format("2006-01-02 15:04:05")
	return formattedTime
}