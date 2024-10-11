package date

import (
	"testing"
)

func BenchmarkNewDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewDate(15, 8, 2024)
	}
}

func BenchmarkAddDays(b *testing.B) {
	date, _ := NewDate(15, 8, 2024)
	for i := 0; i < b.N; i++ {
		date.AddDays(1)
	}
}

func BenchmarkSubtractDays(b *testing.B) {
	date, _ := NewDate(15, 8, 2024)
	for i := 0; i < b.N; i++ {
		date.SubtractDays(1)
	}
}

func BenchmarkCompareDates(b *testing.B) {
	date1, _ := NewDate(15, 8, 2024)
	date2, _ := NewDate(16, 8, 2024)
	for i := 0; i < b.N; i++ {
		_ = date1.Compare(*date2)
	}
}

func BenchmarkDateFormat(b *testing.B) {
	date, _ := NewDate(15, 8, 2024)
	for i := 0; i < b.N; i++ {
		_ = date.Format("{dd}/{mm}/{yyyy}")
	}
}
