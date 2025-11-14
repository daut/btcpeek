package utils

import (
	"fmt"
	"time"
)

func FormatIso8601(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(time.RFC3339)
}

func TimeAgo(timestamp int64) string {
	now := time.Now().Unix()

	diff := now - timestamp

	if diff < 60 {
		return pluralize(diff, "second")
	} else if diff < 3600 {
		return pluralize(diff/60, "minute")
	} else if diff < 86400 {
		return pluralize(diff/3600, "hour")
	} else {
		return pluralize(diff/86400, "day")
	}
}

func pluralize(n int64, unit string) string {
	if n == 1 {
		return fmt.Sprintf("%d %s ago", n, unit)
	}
	return fmt.Sprintf("%d %ss ago", n, unit)
}
