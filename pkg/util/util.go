package util

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomNumeric(size int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if size <= 0 {
		panic("{ size : " + strconv.Itoa(size) + " } must be more than 0 ")
	}
	value := ""
	for index := 0; index < size; index++ {
		value += strconv.Itoa(r.Intn(10))
	}

	return value
}

func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}

// adaptiveTime converts a timestamp (could be sec or ms) to time.Time
func AdaptiveTime(ts int64) time.Time {
	if ts <= 0 {
		return time.Time{}
	}
	// If ts > 1e12, it's likely milliseconds (e.g., 2026-03-03 is ~1.7e9 sec, but 1.7e12 ms)
	if ts > 1000000000000 {
		return time.UnixMilli(ts)
	}
	return time.Unix(ts, 0)
}
