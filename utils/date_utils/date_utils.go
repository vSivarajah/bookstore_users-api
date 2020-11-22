package date_utils

import "time"

const (
	apiDateLayout = "Jan 2, 2006 at 3:04pm (MST)"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
