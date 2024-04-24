package time_limit

import (
	"os"
	"time"
)

func GetTimeLimit() time.Duration {
	var time_limit time.Duration

	time_limit, err := time.ParseDuration(os.Getenv("TIMEOUT_FIBONACCI"))
	if err != nil {
		time_limit = 10 * time.Second
	}

	return time_limit
}
