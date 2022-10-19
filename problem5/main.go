package main

import "log"

// This is the code used for solve the problem 5 from project Euler https://projecteuler.net/problem=5
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

func main() {
	number := 1
	for true {
		if itIsEvenDisible(20, number) {
			log.Println("found", number)
			return
		}
		number++
	}
}

func itIsEvenDisible(maxDivider int, number int) bool {
	divider := 1
	for divider <= maxDivider {
		if number%divider != 0 {
			return false
		}
		divider++
	}
	return true
}
