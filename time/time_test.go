package time

import (
	"testing"
)

func TestNewTime(t *testing.T) {
	time := NewTime(14, 30, 45)
	if time.GetHour() != 14 || time.GetMinute() != 30 || time.GetSecond() != 45 {
		t.Errorf("Time is incorrect: got %02d:%02d:%02d, want 14:30:45", time.GetHour(), time.GetMinute(), time.GetSecond())
	}
}

func TestAddSeconds(t *testing.T) {
	time := NewTime(14, 59, 50)
	time.AddSeconds(15)
	if time.GetHour() != 15 || time.GetMinute() != 0 || time.GetSecond() != 5 {
		t.Errorf("Time after adding seconds is incorrect: got %02d:%02d:%02d, want 15:00:05", time.GetHour(), time.GetMinute(), time.GetSecond())
	}

	time = NewTime(23, 59, 50)
	time.AddSeconds(20)
	if time.GetHour() != 0 || time.GetMinute() != 0 || time.GetSecond() != 10 {
		t.Errorf("Time after adding seconds is incorrect: got %02d:%02d:%02d, want 00:00:10", time.GetHour(), time.GetMinute(), time.GetSecond())
	}
}

func TestSubtractSeconds(t *testing.T) {
	time := NewTime(0, 0, 10)
	time.SubtractSeconds(15)
	if time.GetHour() != 23 || time.GetMinute() != 59 || time.GetSecond() != 55 {
		t.Errorf("Time after subtracting seconds is incorrect: got %02d:%02d:%02d, want 23:59:55", time.GetHour(), time.GetMinute(), time.GetSecond())
	}

	time = NewTime(1, 0, 0)
	time.SubtractSeconds(3600)
	if time.GetHour() != 0 || time.GetMinute() != 0 || time.GetSecond() != 0 {
		t.Errorf("Time after subtracting seconds is incorrect: got %02d:%02d:%02d, want 00:00:00", time.GetHour(), time.GetMinute(), time.GetSecond())
	}
}

func TestCompareTimes(t *testing.T) {
	time1 := NewTime(14, 30, 45)
	time2 := NewTime(15, 0, 0)

	if time1.Compare(*time2) != -1 {
		t.Error("Expected time1 to be earlier than time2")
	}

	if time2.Compare(*time1) != 1 {
		t.Error("Expected time2 to be later than time1")
	}

	time3 := NewTime(14, 30, 45)
	if time1.Compare(*time3) != 0 {
		t.Error("Expected time1 to be equal to time3")
	}
}

func TestTimeFormat(t *testing.T) {
	time := NewTime(5, 7, 9)
	formatted := time.Format("{hh}:{min}:{ss}")
	if formatted != "05:07:09" {
		t.Errorf("Formatted time is incorrect: got %s, want 05:07:09", formatted)
	}
}
