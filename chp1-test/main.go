package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func reverseName(name string) string {
	reversed := make([]byte, 0)
	for i := len(name) - 1; i >= 0; i-- {
		reversed = append(reversed, name[i])
	}

	return string(reversed)
}

func main() {
	fmt.Println("Enter your Name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	name = strings.TrimSpace(name)
	fmt.Println(reverseName(name), ",olleH")
}
