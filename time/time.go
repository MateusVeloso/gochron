package time

import (
	"fmt"
	"sync"
	"time"
)

// Time represents a moment in time with an associated time zone.
type Time struct {
	time     *time.Time     // The time value
	location *time.Location // The associated time zone
}

var (
	// locationCache stores time zones based on their UTC offset.
	locationCache = make(map[int]*time.Location)
	// cacheMutex protects access to locationCache for concurrent use.
	cacheMutex sync.RWMutex
)

// NewFromUTC returns a Time instance with the specified NewFromUTC offset.
// If no offset is provided, it defaults to NewFromUTC+0.
func NewFromUTC(offsets ...int) (Time, error) {
	offset := 0
	if len(offsets) > 0 {
		offset = offsets[0]
		// Validate the offset
		if offset < -12 || offset > 14 {
			return Time{}, fmt.Errorf("invalid offset: %d. Offset must be between -12 and +14", offset)
		}
	}

	loc := getCachedLocation(offset)
	now := time.Now().In(loc)
	return Time{
		time:     &now,
		location: loc,
	}, nil
}

// getCachedLocation returns a *time.Location from the cache or creates a new one if it doesn't exist.
func getCachedLocation(offset int) *time.Location {
	cacheMutex.RLock()
	loc, ok := locationCache[offset]
	cacheMutex.RUnlock()
	if ok {
		return loc
	}

	// Not in cache, create new location
	locName := fmt.Sprintf("NewFromUTC%+d", offset)
	loc = time.FixedZone(locName, offset*3600)

	// Store in cache
	cacheMutex.Lock()
	locationCache[offset] = loc
	cacheMutex.Unlock()

	return loc
}

// Time returns the time.Time instance of the Time struct.
func (t *Time) Time() time.Time {
	return *t.time
}

// Location returns the time.Location instance of the Time struct.
func (t *Time) Location() *time.Location {
	return t.location
}

func (t *Time) String() string {
	return t.time.String()
}
