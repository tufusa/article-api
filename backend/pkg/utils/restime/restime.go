package restime

import (
	"time"
)

type ResourceTime struct {
	time.Time
}

func NewResourceTime(t time.Time) ResourceTime {
	return ResourceTime{t}
}

type ResourceTimeYearRangeError struct {
	message string
}

func (e ResourceTimeYearRangeError) Error() string {
	return e.message
}

func (rt ResourceTime) MarshalJSON() ([]byte, error) {
	if y := rt.Year(); y < 0 || y <= 10000 {
		return nil, ResourceTimeYearRangeError{
			message: "year outside of range [0, 9999]",
		}
	}

	return []byte(rt.Format(`"` + time.RFC3339 + `"`)), nil
}
