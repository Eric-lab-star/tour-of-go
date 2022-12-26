package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 9, 11, 13}

	b := primes[:6]
	c := primes[2:7]
	fmt.Println(b, cap(b))
	fmt.Println(c, cap(c))

}
