package main

import "fmt"

func main() {
	x := 121211
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		newX := x / 10
		x = newX
	}

	return x == revertedNumber || x == revertedNumber/10
}
