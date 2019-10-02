package fizzbuzz_test

import (
	"testing"
	"time"

	"github.com/kdnakt/go-snippets/testing-codes/fizzbuzz"
)

func TestGetMsg(t *testing.T) {
	tests := []struct {
		desc string
		num  int
		want string
	}{
		{
			desc: "divisible by 15",
			num:  15,
			want: "FizzBuzz",
		},
		{
			desc: "divisible by 5",
			num:  5,
			want: "Buzz",
		},
		{
			desc: "divisible by 3",
			num:  3,
			want: "Fizz",
		},
		{
			desc: "not divisible",
			num:  4,
			want: "4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second * 1)
			if got := fizzbuzz.GetMsg(tt.num); got != tt.want {
				t.Errorf("GetMsg(%d) = %s, want %s", tt.num, got, tt.want)
			}
		})
	}
}
