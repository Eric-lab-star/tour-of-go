package main

import "fmt"

type student struct {
	first_name, last_name string
}

var m map[string]student

func main() {
	m = make(map[string]student)
	m["korean"] = student{
		first_name: "Kyungsub",
		last_name:  "Kim",
	}
	fmt.Println(m["korean"])

}
