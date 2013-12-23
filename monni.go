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
  q := make(chan bool)
	for _, v := range urls {
    go checkUrl(v, c, q)
	}

  i := 0

  for res := range q {
    if res {
      i++
    }

    if i == len(urls) {
      close(q)
      close(c)
    }
  }
}

func checkUrl(url string, c chan string, q chan bool) {
  start := time.Now()

  _, err := http.Get(url)

  if err != nil {
    msg := url + " Error: " + err.Error()
    c <- msg
  } else {
    lag := time.Since(start)
    msg := url + " OK : " + lag.String()
    c <- msg
  }

  q <- true
}
