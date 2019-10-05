package sample_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kdnakt/go-snippets/testing-codes/sample"
	"github.com/kdnakt/go-snippets/testing-codes/testutil"
)

func TestGetTomorrow(t *testing.T) {
	tm := time.Date(2019, time.October, 4, 0, 0, 0, 0, testutil.GetJstLocation(t))
	want := tm.AddDate(0, 0, 1)
	got := sample.GetTomorrow(tm)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("GetTomorrow() differs, diff(-got +want) %s", diff)
	}
}
