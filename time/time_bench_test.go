package time

import (
	"testing"
)

func BenchmarkNewTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewTime(14, 30, 45)
	}
}

func BenchmarkAddSeconds(b *testing.B) {
	time := NewTime(14, 30, 45)
	for i := 0; i < b.N; i++ {
		time.AddSeconds(1)
	}
}

func BenchmarkSubtractSeconds(b *testing.B) {
	time := NewTime(14, 30, 45)
	for i := 0; i < b.N; i++ {
		time.SubtractSeconds(1)
	}
}

func BenchmarkCompareTimes(b *testing.B) {
	time1 := NewTime(14, 30, 45)
	time2 := NewTime(15, 0, 0)
	for i := 0; i < b.N; i++ {
		_ = time1.Compare(*time2)
	}
}

func BenchmarkTimeFormat(b *testing.B) {
	time := NewTime(14, 30, 45)
	for i := 0; i < b.N; i++ {
		_ = time.Format("{hh}:{min}:{ss}")
	}
}
