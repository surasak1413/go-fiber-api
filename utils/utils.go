package utils

import "time"

func GetCurrentEpochTime() int64 {
	return time.Now().Unix()
}
