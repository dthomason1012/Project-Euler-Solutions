//The prime factors of 13195 are 5, 7, 13, and 29.
//What is the largest prime factor of the number 600851475143

/*
	+Determine the factors of a number
		+use a for loop, add values to an array that meet the condition 600851475143 % n == 0
	-Determine which factors are prime
		-add those to their own array
	-find the max value of that array
*/
package main

import (
	"fmt"
)

func factors(n int) []int {
	var arr []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			arr = append(arr, i)
		}
	}
	return arr
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeFactors(arr []int) []int {
	var primes []int

	for i := 0; i < len(arr); i++ {
		if isPrime(arr[i]) {
			primes = append(primes, arr[i])
		}
	}
	return primes
}

func main() {

	fmt.Println(primeFactors(factors(197584355555)))

}
