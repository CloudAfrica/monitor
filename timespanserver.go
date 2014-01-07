package main

import (
  "fmt"
  "net/http"
  "os"
  "io"
)

type TimeSpan struct {
  Destination string
  Time  string
}

func (p *TimeSpan) save() {
  f, err := os.OpenFile("timespans.csv", os.O_APPEND, 0666)
  if err != nil {
    fmt.Println("Error opening file.")
    return
  }

  fmt.Println("Destination:" + p.Destination)
  fmt.Println("Time:" + p.Time)

  _, err2 := io.WriteString(f, p.Destination + "," + p.Time + "\r\n")
  if err2 != nil {
    fmt.Println("Error saving file.")
    return
  }
  f.Close()
}

func saveTimeSpanHandler(w http.ResponseWriter, r *http.Request) {
    destination := r.FormValue("destination")
    time := r.FormValue("time")
    timeSpanItem := &TimeSpan{Destination: destination, Time: time}
    timeSpanItem.save()
    fmt.Fprintf(w, "Time span saved.")
}

func main() {
    http.HandleFunc("/savetimespan", saveTimeSpanHandler)
    http.ListenAndServe(":8080", nil)
}