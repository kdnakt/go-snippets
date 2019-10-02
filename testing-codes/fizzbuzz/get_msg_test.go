package fizzbuzz_test

import (
	"testing"

	"github.com/kdnakt/go-snippets/testing-codes/fizzbuzz"
)

func TestGetMsg(t *testing.T) {
	var num int
	var want string

	num = 15
	want = "FizzBuzz"
	if got := fizzbuzz.GetMsg(num); got != want {
		t.Errorf("GetMsg(%d) = %s, want %s", num, got, want)
	}

	num = 5
	want = "Buzz"
	if got := fizzbuzz.GetMsg(num); got != want {
		t.Errorf("GetMsg(%d) = %s, want %s", num, got, want)
	}

	num = 3
	want = "Fizz"
	if got := fizzbuzz.GetMsg(num); got != want {
		t.Errorf("GetMsg(%d) = %s, want %s", num, got, want)
	}

	num = 4
	want = "4"
	if got := fizzbuzz.GetMsg(num); got != want {
		t.Errorf("GetMsg(%d) = %s, want %s", num, got, want)
	}
}
