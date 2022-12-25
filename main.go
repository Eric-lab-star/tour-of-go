package main

import (
	"fmt"
	"os"
)

func main() {
	fs := os.DirFS("/main.go")
	fmt.Printf("type %T, value %v\n", fs, fs)
}
