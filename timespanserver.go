package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

type TimeSpan struct {
  Destination string
  Time  string
}

func (p *TimeSpan) save() error {
  csvContent = string(ioutil.ReadFile("timeSpans.csv"))
  csvContent += timeSpanItem.destination + "," + string(timeSpanItem.time) + "\r\n"
  ioutil.WriteFile(byte("timeSpans.csv", csvLine, 0600))
}

func saveTimeSpanHandler(w http.ResponseWriter, r *http.Request) {
    destination := r.FormValue("destination")
    time := r.FormValue("time")
    timeSpanItem := &TimeSpan{Destination: destination, Time: time}
    timeSpanItem.save()
}

func main() {
    http.HandleFunc("/savetimespan", saveTimeSpanHandler)
    http.ListenAndServe(":8080", nil)
}