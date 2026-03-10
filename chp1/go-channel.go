package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c // Waits for value to come in
		fmt.Println(num, " ")
	}
}

func main() {
	c := make(chan int)
	a := []int{8, 7, 6, 5, 4, 3, 2, 1, 0, 9}
	go printCount(c)
	go printCount(c)
	go printCount(c)

	for _, v := range a {
		c <- v
	}

	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")
}
