package main

import (
	"fmt"
	"net/http"
	"time"
)
//comand0: go run script.go
func main() {
	strat := time.Now()
	myChannel := make(chan string)

	urls := []string{}

	for i := 0; i < 500000; i++ {
		url := "URL"
		urls = append(urls, url)
	}
	// fmt.Println(urls)

	for _, url := range urls {
		go ping(url, myChannel)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-myChannel)
	}
	fmt.Println("TIEMPO", time.Since(strat))
}

func ping(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		c <- "ERROR"
	} else {
		c <- "OK " + url

	}

}
