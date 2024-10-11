package time

// Time structure to represent hours, minutes, and seconds.
// This structure represents a specific time of day with hour, minute, and second fields.
// Fields are private to ensure consistent data manipulation through defined methods.
type Time struct {
	hour   int
	minute int
	second int
}

// NewTime Function to create a new instance of Time
// Initializes a new Time instance with given hour, minute, and second values.
func NewTime(hour, minute, second int) *Time {
	return &Time{hour: hour, minute: minute, second: second}
}

// AddSeconds Method to add seconds to the Time structure
// Manages overflow from seconds to minutes, and minutes to hours.
func (t *Time) AddSeconds(seconds int) {
	t.second += seconds

	t.minute += t.second / 60
	t.second = t.second % 60

	t.hour += t.minute / 60
	t.minute = t.minute % 60

	t.hour = t.hour % 24
}

// SubtractSeconds Method to subtract seconds from the Time structure
// Handles underflow and adjusts minutes and hours accordingly.
func (t *Time) SubtractSeconds(seconds int) {
	t.second -= seconds
	for t.second < 0 {
		t.second += 60
		t.minute--
	}
	for t.minute < 0 {
		t.minute += 60
		t.hour--
	}
	for t.hour < 0 {
		t.hour += 24
	}
}

// Compare Method to compare two Time structures
// Returns 1 if the current time is greater, -1 if smaller, and 0 if equal.
func (t Time) Compare(other Time) int {
	if t.hour != other.hour {
		if t.hour > other.hour {
			return 1
		}
		return -1
	}
	if t.minute != other.minute {
		if t.minute > other.minute {
			return 1
		}
		return -1
	}
	if t.second != other.second {
		if t.second > other.second {
			return 1
		}
		return -1
	}
	return 0
}

// GetHour Method to get the hour of the Time structure
func (t Time) GetHour() int {
	return t.hour
}

// GetMinute Method to get the minute of the Time structure
func (t Time) GetMinute() int {
	return t.minute
}

// GetSecond Method to get the second of the Time structure
func (t Time) GetSecond() int {
	return t.second
}
