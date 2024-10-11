package datetime

import (
	"testing"
)

func BenchmarkNewDateTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewDateTime(15, 8, 2024, 14, 30, 45)
	}
}

func BenchmarkAddDuration(b *testing.B) {
	datetime, _ := NewDateTime(15, 8, 2024, 14, 30, 45)
	for i := 0; i < b.N; i++ {
		datetime.AddDuration(1, 3600) // Add 1 day and 3600 seconds (1 hour)
	}
}

func BenchmarkSubtractDuration(b *testing.B) {
	datetime, _ := NewDateTime(15, 8, 2024, 14, 30, 45)
	for i := 0; i < b.N; i++ {
		datetime.SubtractDuration(1, 3600) // Subtract 1 day and 3600 seconds (1 hour)
	}
}

func BenchmarkCompareDateTimes(b *testing.B) {
	datetime1, _ := NewDateTime(15, 8, 2024, 14, 30, 45)
	datetime2, _ := NewDateTime(16, 8, 2024, 14, 30, 45)
	for i := 0; i < b.N; i++ {
		_ = datetime1.Compare(*datetime2)
	}
}

func BenchmarkDateTimeFormat(b *testing.B) {
	datetime, _ := NewDateTime(15, 8, 2024, 14, 30, 45)
	for i := 0; i < b.N; i++ {
		_ = datetime.Format("{dd}/{mm}/{yyyy} {hh}:{min}:{ss}")
	}
}
