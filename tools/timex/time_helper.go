package timex

import "time"

const (
	defaultTimeFormat = "2006-01-02 15:04:05"
	allTimeFormat     = "2006-01-02 15:04:05.999999999"
)

type TimeHelper struct {
	err       error
	innerTime time.Time
	format    string
}

func NewTimeHelper(t ...time.Time) *TimeHelper {
	if len(t) > 0 {
		return &TimeHelper{
			innerTime: t[0],
			format:    defaultTimeFormat,
		}
	}
	return &TimeHelper{
		format: defaultTimeFormat,
	}
}

func (th *TimeHelper) SetTime(t time.Time) *TimeHelper {
	if th.err != nil {
		return th
	}
	th.innerTime = t
	return th
}

func (th *TimeHelper) SetTimeByString(timeStr string) *TimeHelper {
	if th.err != nil {
		return th
	}
	th.innerTime, th.err = time.Parse(th.format, timeStr)
	return th
}

func (th *TimeHelper) SetFormat(format string) *TimeHelper {
	if th.err != nil {
		return th
	}
	th.format = format
	return th
}

func (th *TimeHelper) SetAllFormat() *TimeHelper {
	if th.err != nil {
		return th
	}
	th.format = allTimeFormat
	return th
}

func (th *TimeHelper) GetTime() time.Time {
	if th.err != nil {
		return time.Time{}
	}
	return th.innerTime
}

func (th *TimeHelper) GetTimeString() string {
	if th.err != nil {
		return ""
	}
	return th.innerTime.Format(th.format)
}

func (th *TimeHelper) GetError() error {
	return th.err
}
