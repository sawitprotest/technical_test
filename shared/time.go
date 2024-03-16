package shared

import "time"

const (
	FormatDateTime  = `2006-01-02 15:04:05`
	TimeAsiaJakarta = `Asia/Jakarta`
)

func UTC7(t time.Time) time.Time {
	location, err := time.LoadLocation(TimeAsiaJakarta)
	if err != nil {
		return time.Now()
	}
	return t.In(location)
}
