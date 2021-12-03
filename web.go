package main

import (
	"fmt"
	"io"
	"net/http"
)

//structs must be implemented that use the method write to use the reader that is returned by response

type webWriter struct{}

func (webWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	//string converts the bytes flow into a string
	return len(p), nil
}

func main() {
	response, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println(err)
	}
	w := webWriter{}
	io.Copy(w, response.Body)
	//fmt.Println(response.Body)
}
