package fizzbuzz

import "strconv"

// GetMsg returns fizzbuzz message
func GetMsg(num int) string {
	var res string
	switch {
	case num%15 == 0:
		res = "FizzBuzz"
	case num%3 == 0:
		res = "Fizz"
	case num%5 == 0:
		res = "Buzz"
	default:
		res = strconv.Itoa(num)
	}
	return res
}
