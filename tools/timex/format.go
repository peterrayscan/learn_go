package timex

import "time"

var format = "2006-01-02 15:04:05.000000000"

func GetTimeByString(timeStr string) (time.Time, error) {
	return time.Parse(format, timeStr)
}

func GetStringByTime(time time.Time) string {
	return time.Format(format)
}

func GetNowString() string {
	return GetStringByTime(time.Now())
}
