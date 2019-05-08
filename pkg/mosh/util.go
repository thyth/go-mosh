package mosh

import "time"

func Timestamp(t time.Time) uint64 {
	return uint64(t.UnixNano() / int64(time.Millisecond/time.Nanosecond))
}
