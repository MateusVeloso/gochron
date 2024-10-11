package datetime

import (
	"fmt"
	"strings"
)

// Format Method to format the output of DateTime as a string based on a given format
func (dt DateTime) Format(format string) string {
	format = strings.Replace(format, "{dd}", fmt.Sprintf("%02d", dt.Date.GetDay()), -1)
	format = strings.Replace(format, "{mm}", fmt.Sprintf("%02d", dt.Date.GetMonth()), -1)
	format = strings.Replace(format, "{yyyy}", fmt.Sprintf("%04d", dt.Date.GetYear()), -1)
	format = strings.Replace(format, "{hh}", fmt.Sprintf("%02d", dt.Time.GetHour()), -1)
	format = strings.Replace(format, "{min}", fmt.Sprintf("%02d", dt.Time.GetMinute()), -1)
	format = strings.Replace(format, "{ss}", fmt.Sprintf("%02d", dt.Time.GetSecond()), -1)
	return format
}
