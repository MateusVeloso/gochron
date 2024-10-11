package time

import "fmt"
import "strings"

// Format Method to format the output of Time as a string based on a given format
func (t Time) Format(format string) string {
	format = strings.Replace(format, "{hh}", fmt.Sprintf("%02d", t.hour), -1)
	format = strings.Replace(format, "{min}", fmt.Sprintf("%02d", t.minute), -1)
	format = strings.Replace(format, "{ss}", fmt.Sprintf("%02d", t.second), -1)
	return format
}
