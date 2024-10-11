package date

import "fmt"
import "strings"

// IsValid Method to validate a date with direct integer comparisons for efficiency.
func (d Date) IsValid() bool {
	if d.year < 0 || d.month < 1 || d.month > 12 || d.day < 1 {
		return false
	}

	days := daysInMonth[d.month-1]
	if d.month == 2 && isLeapYear(d.year) {
		days = 29
	}

	if d.day > days {
		return false
	}

	return true
}

// Format Method to format the output of Date as a string based on a given format
func (d Date) Format(format string) string {
	format = strings.Replace(format, "{dd}", fmt.Sprintf("%02d", d.day), -1)
	format = strings.Replace(format, "{mm}", fmt.Sprintf("%02d", d.month), -1)
	format = strings.Replace(format, "{yyyy}", fmt.Sprintf("%04d", d.year), -1)
	return format
}
