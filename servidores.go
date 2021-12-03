package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServer(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "server not available :(")
	}
	fmt.Println(server, "works :D")
}

func main() {
	start := time.Now()

	servers := []string{
		"http://google.com",
		"http://platzi.com",
		"http://instagram.com",
	}

	for _, server := range servers {
		checkServer(server)
	}

	passed := time.Since(start)

	fmt.Printf("execution time: %s", passed)

}
