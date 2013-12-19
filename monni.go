package main

//https://bitbucket.org/cmccabe314/netcanary/src/master/main.go

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"https://www.cloudafrica.net/",
	"https://www.sensepost.com/",
	"http://www.afrihost.co.za"}

func main() {
	ch := make(chan string)

	go checkUrls(urls, ch)

	for res := range ch {
		fmt.Println(res)
	}

}

func checkUrls(urls []string, c chan string) {
	for _, v := range urls {
		start := time.Now()

		_, err := http.Get(v)

		if err != nil {
			msg := v + " Error: " + err.Error()
			c <- msg
		} else {
			lag := time.Since(start)
			msg := v + " OK : " + lag.String()
			c <- msg
		}
	}
	close(c)
}
