package main

import "testing"

func TestSumMutliples(t *testing.T) {
	res := sumMultiples([]int{3, 5}, 10)
	if res != 23 {
		t.Fatalf("expected 23 but got %d", res)
	}

	res = sumMultiples([]int{3, 5}, 1000)
	if res != 233168 {
		t.Fatalf("expected 233168 but got %d", res)
	}
}
