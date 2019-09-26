package sample_test

import (
	"testing"

	"github.com/kdnakt/go-snippets/testing-codes/sample"
)

func TestSayHello(t *testing.T) {
	want := "hello"
	if got := sample.SayHello(); got != want {
		t.Errorf("SayHello() = %s, want %s", got, want)
	}
}
