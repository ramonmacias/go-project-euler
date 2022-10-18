package main

import "log"

// This is the code used for solve the problem 3 from project Euler https://projecteuler.net/problem=3
// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?
// It takes a lot of time, probably by applying some go routines we can speed up the process

func main() {
	res := getTheLargestPrimeFactor(600851475143)
	log.Println(res)
}

// getTheLargestPrimeFactor gets the biggest primer factor number
// of the given number as a limit, we try to find out the number by
// starting from the biggest number possible
func getTheLargestPrimeFactor(number int) int {
	aux := number
	for aux != 0 {
		if number%aux == 0 {
			if isPrimeFactor(aux) {
				return aux
			}
		}
		aux--
	}
	return 0
}

// isPrimeFactor check if the given number is
// a primer factor number or not
func isPrimeFactor(number int) bool {
	count := 2
	for count < number {
		if number%count == 0 {
			return false
		}
		count++
	}
	return true
}
