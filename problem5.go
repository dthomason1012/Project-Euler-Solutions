//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//What is the smallest positive number that is ~evenly divisible~ by all of the numbers from 1 to 20?

/*
	-iterate from numbers 1 to 20
	-if a number is divisible by 1, check to see if it's divisible by 2
		-if it is not divisible by any number, start over and increment the inputted number
*/

package main

import "fmt"

func multiplier() int {
	n := 2520
	for i := 1; i <= 20; i++ {
		if n%i != 0 {
			n++
			i = 1
		}
	}
	return n
}

func main() {
	fmt.Println(multiplier())
}
