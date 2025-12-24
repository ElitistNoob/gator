package timeutils

import (
	"fmt"
	"time"
)

func ParseTime(s string) (time.Time, error) {
	layouts := []string{
		time.RFC1123Z,
		time.RFC3339,
		time.DateOnly,
		time.TimeOnly,
		time.DateTime,
		"2006-01-02",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unknown time format %q", s)
}
