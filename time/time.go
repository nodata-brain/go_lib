package time

import (
	"time"
)

func SubTime(t int) bool {
	start := time.Now().Unix()

	var end int64
	for {
		end = time.Now().Unix()
		sub := end - start

		if int64(t) == sub {
			break
		}
		if int64(t) < sub {
			return false
		}
	}
	return true
}
