//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 * 99
//Find the largest palindrome made from the product of two 3-digit numbers

/*
	-two for loops, one for each 3 digit number, counting backwards
		-multiply i * j in each iteration
	+look up how to treat ints like strings
		+see if there's a way to reverse a string
	+check if the reversal is equal to the given number
*/

package main

import "fmt"

func threeDigitProducts() []int {
	var products []int
	for i := 1; i <= 999; i++ {
		for j := 1; j <= 999; j++ {
			products = append(products, i*j)
		}
	}
	return products
}

func palindrome(n int) bool {
	var remainder, temp int
	var reverse int = 0

	temp = n

	for {
		remainder = n % 10
		reverse = reverse*10 + remainder
		n /= 10

		if n == 0 {
			break
		}
	}
	if temp == reverse {
		return true
	}
	return false
}

func main() {
	products := threeDigitProducts()
	var palindromes []int
	for i := 0; i < len(products); i++ {
		if palindrome(products[i]) {
			palindromes = append(palindromes, products[i])
		}
	}
	max := palindromes[0]
	for j := 0; j < len(palindromes); j++ {
		if palindromes[j] > max {
			max = palindromes[j]
		}
	}
	fmt.Println(max)
}
