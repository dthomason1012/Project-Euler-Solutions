//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6, and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.

package main

import "fmt"

func main() {
	arr := make([]int, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	var sum int
	for j := 0; j < len(arr); j++ {
		if arr[j]%3 == 0 || arr[j]%5 == 0 {
			sum += arr[j]
		}
	}

	fmt.Println(sum)
}
