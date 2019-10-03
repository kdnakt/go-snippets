package sample_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kdnakt/go-snippets/testing-codes/sample"
)

func TestGetTomorrow(t *testing.T) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal("Failedto locad JST time Location")
	}
	tm := time.Date(2019, time.October, 4, 0, 0, 0, 0, jst)
	want := tm.AddDate(0, 0, 1)
	got := sample.GetTomorrow(tm)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("GetTomorrow() differs, diff(-got +want) %s", diff)
	}
}
