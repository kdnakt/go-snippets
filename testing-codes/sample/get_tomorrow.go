package sample

import "time"

// GetTomorrow returns tomorrow as time.Time
func GetTomorrow(tm time.Time) time.Time {
	return tm.AddDate(0, 0, 1)
}
