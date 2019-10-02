package fizzbuzz_test

import (
	"testing"

	"github.com/kdnakt/go-snippets/testing-codes/fizzbuzz"
)

func TestGetMsg(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{
			num:  15,
			want: "FizzBuzz",
		},
		{
			num:  5,
			want: "Buzz",
		},
		{
			num:  3,
			want: "Fizz",
		},
		{
			num:  4,
			want: "4",
		},
	}
	for _, tt := range tests {
		if got := fizzbuzz.GetMsg(tt.num); got != tt.want {
			t.Errorf("GetMsg(%d) = %s, want %s", tt.num, got, tt.want)
		}
	}
}
