package date_utils

import "time"

const (
	apiDateLayout = "Jan 2, 2006 at 3:04pm (MST)"
	apiDbLayout   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
