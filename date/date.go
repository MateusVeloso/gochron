package date

import "errors"

// Date structure to represent basic dates with maximum efficiency.
// This structure represents a specific date with day, month, and year fields.
// The fields are private to ensure data integrity and consistent validation through constructor methods.
// We work directly with integers for better performance.
type Date struct {
	day   int
	month int
	year  int
}

// Precomputed days in each month, non-leap year.
var daysInMonth = [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// NewDate Function to create a new instance of Date
// Ensures that the created date is valid by performing validation checks.
func NewDate(day, month, year int) (*Date, error) {
	date := Date{day: day, month: month, year: year}
	if !date.IsValid() {
		return nil, errors.New("invalid date")
	}
	return &date, nil
}

// AddDays Method to add days to the Date structure
// Handles overflow of days between months and years.
func (d *Date) AddDays(days int) {
	d.day += days
	for {
		days := daysInMonth[d.month-1]
		if d.month == 2 && isLeapYear(d.year) {
			days = 29
		}

		if d.day <= days {
			break
		}
		d.day -= days
		d.month++
		if d.month > 12 {
			d.month = 1
			d.year++
		}
	}
}

// SubtractDays Method to subtract days from the Date structure
// Handles underflow of days between months and years.
func (d *Date) SubtractDays(days int) {
	d.day -= days
	for d.day <= 0 {
		d.month--
		if d.month < 1 {
			d.month = 12
			d.year--
		}
		days := daysInMonth[d.month-1]
		if d.month == 2 && isLeapYear(d.year) {
			days = 29
		}
		d.day += days
	}
}

// Compare Method to compare two Date structures
// Returns 1 if the current date is greater, -1 if smaller, and 0 if equal.
func (d Date) Compare(other Date) int {
	if d.year != other.year {
		if d.year > other.year {
			return 1
		}
		return -1
	}
	if d.month != other.month {
		if d.month > other.month {
			return 1
		}
		return -1
	}
	if d.day != other.day {
		if d.day > other.day {
			return 1
		}
		return -1
	}
	return 0
}

// GetDay Method to get the day of the Date structure
func (d Date) GetDay() int {
	return d.day
}

// GetMonth Method to get the month of the Date structure
func (d Date) GetMonth() int {
	return d.month
}

// GetYear Method to get the year of the Date structure
func (d Date) GetYear() int {
	return d.year
}
