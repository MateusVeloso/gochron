package datetime

import (
	"testing"
)

func TestNewDateTime(t *testing.T) {
	datetime, err := NewDateTime(15, 8, 2024, 14, 30, 45)
	if err != nil {
		t.Errorf("Error creating valid datetime: %v", err)
	}

	if datetime.Date.GetDay() != 15 || datetime.Date.GetMonth() != 8 || datetime.Date.GetYear() != 2024 ||
		datetime.Time.GetHour() != 14 || datetime.Time.GetMinute() != 30 || datetime.Time.GetSecond() != 45 {
		t.Errorf("DateTime is incorrect: got %02d-%02d-%04d %02d:%02d:%02d, want 15-08-2024 14:30:45",
			datetime.Date.GetDay(), datetime.Date.GetMonth(), datetime.Date.GetYear(),
			datetime.Time.GetHour(), datetime.Time.GetMinute(), datetime.Time.GetSecond())
	}
}

func TestAddDuration(t *testing.T) {
	datetime, _ := NewDateTime(28, 2, 2020, 23, 59, 50) // Leap year
	datetime.AddDuration(1, 20)                         // Add 1 day and 20 seconds

	if datetime.Date.GetDay() != 29 || datetime.Date.GetMonth() != 2 || datetime.Date.GetYear() != 2020 ||
		datetime.Time.GetHour() != 0 || datetime.Time.GetMinute() != 0 || datetime.Time.GetSecond() != 10 {
		t.Errorf("DateTime after adding duration is incorrect: got %02d-%02d-%04d %02d:%02d:%02d, want 29-02-2020 00:00:10",
			datetime.Date.GetDay(), datetime.Date.GetMonth(), datetime.Date.GetYear(),
			datetime.Time.GetHour(), datetime.Time.GetMinute(), datetime.Time.GetSecond())
	}
}

func TestSubtractDuration(t *testing.T) {
	datetime, _ := NewDateTime(1, 3, 2020, 0, 0, 10) // Leap year
	datetime.SubtractDuration(1, 20)                 // Subtract 1 day and 20 seconds

	if datetime.Date.GetDay() != 29 || datetime.Date.GetMonth() != 2 || datetime.Date.GetYear() != 2020 ||
		datetime.Time.GetHour() != 23 || datetime.Time.GetMinute() != 59 || datetime.Time.GetSecond() != 50 {
		t.Errorf("DateTime after subtracting duration is incorrect: got %02d-%02d-%04d %02d:%02d:%02d, want 29-02-2020 23:59:50",
			datetime.Date.GetDay(), datetime.Date.GetMonth(), datetime.Date.GetYear(),
			datetime.Time.GetHour(), datetime.Time.GetMinute(), datetime.Time.GetSecond())
	}
}

func TestCompareDateTimes(t *testing.T) {
	datetime1, _ := NewDateTime(15, 8, 2024, 14, 30, 45)
	datetime2, _ := NewDateTime(15, 8, 2024, 15, 0, 0)

	if datetime1.Compare(*datetime2) != -1 {
		t.Error("Expected datetime1 to be earlier than datetime2")
	}

	if datetime2.Compare(*datetime1) != 1 {
		t.Error("Expected datetime2 to be later than datetime1")
	}

	datetime3, _ := NewDateTime(15, 8, 2024, 14, 30, 45)
	if datetime1.Compare(*datetime3) != 0 {
		t.Error("Expected datetime1 to be equal to datetime3")
	}
}

func TestDateTimeFormat(t *testing.T) {
	datetime, _ := NewDateTime(5, 7, 2024, 5, 7, 9)
	formatted := datetime.Format("{dd}/{mm}/{yyyy} {hh}:{min}:{ss}")
	if formatted != "05/07/2024 05:07:09" {
		t.Errorf("Formatted datetime is incorrect: got %s, want 05/07/2024 05:07:09", formatted)
	}
}
