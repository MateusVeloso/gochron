package gochron

import (
	"github.com/mateusveloso/gochron/date"
	"github.com/mateusveloso/gochron/datetime"
	"github.com/mateusveloso/gochron/time"
)

// NewDate Function to create a new instance of Date
// Ensures that the created date is valid by performing validation checks.
func NewDate(day, month, year int) (*date.Date, error) {
	return date.NewDate(day, month, year)
}

// NewTime Function to create a new instance of Time
// Initializes a new Time instance with given hour, minute, and second values.
func NewTime(hour, minute, second int) *time.Time {
	return time.NewTime(hour, minute, second)
}

// NewDateTime Function to create a new instance of DateTime
// Combines Date and Time into a single structure to represent a complete date and time.
func NewDateTime(day, month, year, hour, minute, second int) (*datetime.DateTime, error) {
	return datetime.NewDateTime(day, month, year, hour, minute, second)
}
