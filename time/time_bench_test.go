package time

import (
	"testing"
)

// Benchmark for NewFromUTC function without offset
func BenchmarkUTC_NoOffset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewFromUTC()
	}
}

// Benchmark for NewFromUTC function with offset
func BenchmarkUTC_WithOffset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewFromUTC(-3)
	}
}
