package main

import "log"

// This is the code used for solve the problem 1 from project Euler https://projecteuler.net/problem=1
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

func main() {
	res := sumMultiples([]int{3, 5}, 10)
	log.Println(res)
}

// sumMultiples will return the sum of all the numbers multiples
// of the given slice of multiples and below the given value
func sumMultiples(multiples []int, value int) int {
	sum := 0
	count := 1
	for count < value {
		for _, multiple := range multiples {
			if res := count % multiple; res == 0 {
				sum += count
				break
			}
		}
		count++
	}

	return sum
}
