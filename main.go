package main

import (
	"fmt"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fs := os.DirFS(wd)
	fmt.Printf("type %T, value %v\n", fs, fs)
	file, err := fs.Open("main.go")
	if err != nil {
		panic(err)
	}
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
