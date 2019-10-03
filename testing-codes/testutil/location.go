package testutil

import (
	"testing"
	"time"
)

// GetJstLocation returns Asia/Tokyo Location
func GetJstLocation(t *testing.T) *time.Location {
	t.Helper()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal("Failed to load JST time Location")
	}
	return jst
}
