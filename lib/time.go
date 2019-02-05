package lib

import "time"

func ParseUnix(val int64) time.Time {
	return time.Unix(val, 0).UTC()
}
