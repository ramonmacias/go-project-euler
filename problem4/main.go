package main

import (
	"log"
	"strconv"
)

// This is the code used for solve the problem 4 from project Euler https://projecteuler.net/problem=4
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

func main() {
	upper, lower := 999, 0

	i := upper
	max := 0
	for i >= lower {
		j := i
		for j >= lower {
			product := i * j
			if product < max {
				break
			}

			if isPalindrome(product) && product > max {
				max = product
			}
			j--
		}
		i--
	}
	log.Println(max)
}

// isPalindrome checks if the given number is a
// palindrome or not, by returning true or false
func isPalindrome(number int) bool {
	str := strconv.Itoa(number)
	start := 0
	end := len(str) - 1

	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}
