package mosh

import "time"

func Timestamp(t time.Time) uint64 {
	return uint64(t.UnixNano() / int64(time.Millisecond/time.Nanosecond))
}

// TODO find a way to read the VERSION from the associated build of Mosh and make it available as a constant here
