package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("name", "kim")

	fmt.Println(os.Getenv("name"))
	os.Clearenv()
	fmt.Println(os.Getenv("name"))

}
