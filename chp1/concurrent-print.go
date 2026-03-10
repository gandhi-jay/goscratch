package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 5)
	}
}

func main() {
	go count()
	time.Sleep(time.Millisecond * 13)
	fmt.Println("HelloWorld!")
	time.Sleep(time.Millisecond * 10)
}
