package date

import (
	"testing"
)

func TestNewDate(t *testing.T) {
	date, err := NewDate(15, 8, 2024)
	if err != nil {
		t.Errorf("Error creating valid date: %v", err)
	}

	if date.GetDay() != 15 || date.GetMonth() != 8 || date.GetYear() != 2024 {
		t.Errorf("Date is incorrect: got %02d-%02d-%04d, want 15-08-2024", date.GetDay(), date.GetMonth(), date.GetYear())
	}
}

func TestNewDateInvalid(t *testing.T) {
	_, err := NewDate(32, 13, 2024)
	if err == nil {
		t.Error("Expected error for invalid date, got nil")
	}
}

func TestAddDays(t *testing.T) {
	date, _ := NewDate(28, 2, 2020) // Leap year
	date.AddDays(1)
	if date.GetDay() != 29 || date.GetMonth() != 2 || date.GetYear() != 2020 {
		t.Errorf("Date after adding days is incorrect: got %02d-%02d-%04d, want 29-02-2020", date.GetDay(), date.GetMonth(), date.GetYear())
	}

	date.AddDays(1)
	if date.GetDay() != 1 || date.GetMonth() != 3 || date.GetYear() != 2020 {
		t.Errorf("Date after adding days is incorrect: got %02d-%02d-%04d, want 01-03-2020", date.GetDay(), date.GetMonth(), date.GetYear())
	}
}

func TestSubtractDays(t *testing.T) {
	date, _ := NewDate(1, 3, 2020) // Leap year
	date.SubtractDays(1)
	if date.GetDay() != 29 || date.GetMonth() != 2 || date.GetYear() != 2020 {
		t.Errorf("Date after subtracting days is incorrect: got %02d-%02d-%04d, want 29-02-2020", date.GetDay(), date.GetMonth(), date.GetYear())
	}
}

func TestCompareDates(t *testing.T) {
	date1, _ := NewDate(15, 8, 2024)
	date2, _ := NewDate(16, 8, 2024)

	if date1.Compare(*date2) != -1 {
		t.Error("Expected date1 to be earlier than date2")
	}

	if date2.Compare(*date1) != 1 {
		t.Error("Expected date2 to be later than date1")
	}

	date3, _ := NewDate(15, 8, 2024)
	if date1.Compare(*date3) != 0 {
		t.Error("Expected date1 to be equal to date3")
	}
}

func TestDateFormat(t *testing.T) {
	date, _ := NewDate(5, 7, 2024)
	formatted := date.Format("{dd}/{mm}/{yyyy}")
	if formatted != "05/07/2024" {
		t.Errorf("Formatted date is incorrect: got %s, want 05/07/2024", formatted)
	}
}
