package date

// isLeapYear Function to determine if a year is a leap year
func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
