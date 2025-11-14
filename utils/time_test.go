package utils

import (
	"testing"
	"time"
)

func TestTimeAgo(t *testing.T) {
	now := time.Now().Unix()

	tests := []struct {
		timestamp int64
		expected  string
	}{
		{now - 1, "1 second ago"},
		{now - 30, "30 seconds ago"},
		{now - 90, "1 minute ago"},
		{now - 150, "2 minutes ago"},
		{now - 3_600, "1 hour ago"},
		{now - 86_400, "1 day ago"},
		{now - 172_800, "2 days ago"},
	}

	for _, test := range tests {
		result := TimeAgo(test.timestamp)
		if result != test.expected {
			t.Errorf("TimeAgo(%d) = %s; want %s", test.timestamp, result, test.expected)
		}
	}
}
