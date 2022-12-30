package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("successfully wrote %d\n", i)
	}
	close(ch)

}

func main() {
	ch := make(chan int, 2)

	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Printf("read value %v\n", v)
		time.Sleep(1 * time.Second)
	}
}
