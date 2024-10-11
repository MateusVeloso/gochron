package datetime

import (
	"github.com/mateusveloso/gochron/date"
	"github.com/mateusveloso/gochron/time"
)

// DateTime structure to combine Date and Time into a single structure.
// This structure represents a specific date and time, enabling comprehensive date-time operations.
type DateTime struct {
	Date date.Date
	Time time.Time
}

// NewDateTime Function to create a new instance of DateTime
// Combines Date and Time into a single structure to represent a complete date and time.
func NewDateTime(day, month, year, hour, minute, second int) (*DateTime, error) {
	date, err := date.NewDate(day, month, year)
	if err != nil {
		return nil, err
	}

	time := time.NewTime(hour, minute, second)

	return &DateTime{Date: *date, Time: *time}, nil
}

// AddDuration Method to add a duration to the DateTime structure
// Adds a specified number of days and seconds to the DateTime.
func (dt *DateTime) AddDuration(days int, seconds int) {
	dt.Date.AddDays(days)
	dt.Time.AddSeconds(seconds)
}

// SubtractDuration Method to subtract a duration from the DateTime structure
// Subtracts a specified number of days and seconds from the DateTime.
func (dt *DateTime) SubtractDuration(days int, seconds int) {
	dt.Date.SubtractDays(days)
	dt.Time.SubtractSeconds(seconds)
}

// Compare Method to compare two DateTime structures
// Returns 1 if the current DateTime is greater, -1 if smaller, and 0 if equal.
func (dt DateTime) Compare(other DateTime) int {
	dateComparison := dt.Date.Compare(other.Date)
	if dateComparison != 0 {
		return dateComparison
	}
	return dt.Time.Compare(other.Time)
}
