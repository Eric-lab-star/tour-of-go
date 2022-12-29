package main

import "fmt"

func main() {
	var i interface{}

	i = 10
	i = "hello"

	fmt.Println(i)

}
