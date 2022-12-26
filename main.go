package main

import "fmt"

func main() {
	primes := [4]int{2, 3, 5, 7}

	slice := primes[0:2]
	// slice2 := primes[0:4]
	// slice2[3] = 11
	slice = append(slice, 11, 13)
	slice[3] = 17
	fmt.Println(primes)

}
