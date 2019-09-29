package sample_test

import (
	"testing"

	"github.com/kdnakt/go-snippets/testing-codes/sample"
)

func TestInStatusList(t *testing.T) {
	var x string
	var want bool

	x = "deleted"
	want = true
	if got := sample.InStatusList(x); got != want {
		t.Errorf("InStatusList(%s) = %t, but want %t", x, got, want)
	}
}
