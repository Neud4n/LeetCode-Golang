package main

import "fmt"

/*

Given an integer x, return true if x is a palindrome, and false otherwise.

Constraints:
* -2³¹ <= x <= 2³¹ - 1

*/

func IsPalindrome(x int) bool {

	var sing string
	if x > 0 {
		sing = "positive"
	} else {
		sing = "negative"
	}

	var reversedDigit int
	xBkp := x
	for x > 0 {
		lastDigit := x % 10
		reversedDigit = (reversedDigit * 10) + lastDigit

		x /= 10
	}

	if sing == "negative" {
		reversedDigit = reversedDigit * -1
	}

	if xBkp == reversedDigit {
		return true
	}
	return false
}

func main() {
	/*
		INPUTS:
			x = 121 - expected: true
			x = -121 - expected: false
			x = 10 - expexted: false
	*/

	fmt.Println(IsPalindrome(121))
	fmt.Println(IsPalindrome(-121))
	fmt.Println(IsPalindrome(10))
}
