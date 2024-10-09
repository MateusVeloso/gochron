package time

import (
	"fmt"
	"sync"
	"testing"
)

// Test NewFromUTC function without offset (defaults to NewFromUTC+0)
func TestUTC_NoOffset(t *testing.T) {
	tm, err := NewFromUTC()
	if err != nil {
		t.Fatalf("Error calling NewFromUTC(): %v", err)
	}
	if tm.Location().String() != "NewFromUTC+0" {
		t.Errorf("Expected location 'NewFromUTC+0', but got '%s'", tm.Location().String())
	}
}

// Test NewFromUTC function with a valid offset
func TestUTC_WithValidOffset(t *testing.T) {
	offset := -3
	tm, err := NewFromUTC(offset)
	if err != nil {
		t.Fatalf("Error calling NewFromUTC(%d): %v", offset, err)
	}
	expectedLoc := "NewFromUTC-3"
	if tm.Location().String() != expectedLoc {
		t.Errorf("Expected location '%s', but got '%s'", expectedLoc, tm.Location().String())
	}
}

// Test NewFromUTC function with an invalid offset (below minimum)
func TestUTC_WithInvalidOffsetLow(t *testing.T) {
	offset := -13
	_, err := NewFromUTC(offset)
	if err == nil {
		t.Fatalf("Expected error when calling NewFromUTC(%d), but got none", offset)
	}
}

// Test NewFromUTC function with an invalid offset (above maximum)
func TestUTC_WithInvalidOffsetHigh(t *testing.T) {
	offset := 15
	_, err := NewFromUTC(offset)
	if err == nil {
		t.Fatalf("Expected error when calling NewFromUTC(%d), but got none", offset)
	}
}

// Concurrent test for NewFromUTC function with multiple offsets
func TestUTC_Concurrent(t *testing.T) {
	offsets := []int{-12, -6, 0, 6, 14}
	var wg sync.WaitGroup
	for _, offset := range offsets {
		wg.Add(1)
		go func(offset int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				tm, err := NewFromUTC(offset)
				if err != nil {
					t.Errorf("Error calling NewFromUTC(%d): %v", offset, err)
					return
				}
				expectedLoc := fmt.Sprintf("NewFromUTC%+d", offset)
				if tm.Location().String() != expectedLoc {
					t.Errorf("Expected location '%s', but got '%s'", expectedLoc, tm.Location().String())
				}
			}
		}(offset)
	}
	wg.Wait()
}

// Test to verify if the cache is working correctly
func TestUTC_Cache(t *testing.T) {
	offset := -3
	loc1 := getCachedLocation(offset)
	loc2 := getCachedLocation(offset)
	if loc1 != loc2 {
		t.Errorf("Cache failed: locations should be the same")
	}
}
