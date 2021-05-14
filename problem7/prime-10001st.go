//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime number is 13.
//What is the 10001st prime number?

/*
	-copy isPrime function used in Problem 3
	-add these numbers to a slice with a loop, break out of the loop when len(slice) = 10001
*/

package main

import "fmt"

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 || n == 1 {
			return false
		}
	}
	return true
}

func main() {
	var primes []int
	for i := 2; i < 1000000; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
		if len(primes) == 10001 {
			break
		}
	}
	fmt.Println(primes, len(primes))
}
