package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://go.dev")
	if err != nil {
		log.Fatal("couldnot retrieve", err)
	}

	// Any defer line is executed at the end of that code block—in this case, the main() function.
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("couldn't read the body", err)
	}

	log.Println(string(body))
}
