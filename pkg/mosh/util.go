package mosh

import (
	internals "../../internal/mosh"

	"time"
)

func Timestamp(t time.Time) uint64 {
	return uint64(t.UnixNano() / int64(time.Millisecond/time.Nanosecond))
}

func GetVersion() string {
	return internals.BUILD_VERSION
}
