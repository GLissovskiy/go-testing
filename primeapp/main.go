package main

import "fmt"

func main() {

	n := 3

	fmt.Println("started....")

	_, msg := isPrime(n)
	fmt.Println(msg)

	fmt.Println("terminated!")
}

// func returns true if number is not divisible by 2
func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition

	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definitions", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "negative numbers are not prime, by definitions."
	}

	//use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)

}
