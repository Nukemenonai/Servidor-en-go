package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServer(server string, channel chan string) {
	_, err := http.Get(server)
	if err != nil {
		channel <- server + "not available"
	}
	channel <- server + "works :D"
}

func main() {
	start := time.Now()
	channel := make(chan string) //string means we are transmitting strings

	servers := []string{
		"http://google.com",
		"http://platzi.com",
		"http://instagram.com",
	}

	for _, server := range servers {
		go checkServer(server, channel)

	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}
	passed := time.Since(start)

	fmt.Printf("execution time: %s", passed)

}
