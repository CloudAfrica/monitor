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

func (p *TimeSpan) save() error {
  f, err := os.OpenFile("timespans.csv", os.O_RDWR|os.O_APPEND, 0660)
  if err != nil {
    fmt.Println("Error opening file.")
    return err
  }

  fmt.Println("Destination: " + p.Destination)
  fmt.Println("Time: " + p.Time)

  _, err2 := io.WriteString(f, p.Destination + "," + p.Time + "\r\n")
  if err2 != nil {
    fmt.Println("Error saving file: ", err2)
    return err2
  }
  f.Close()
  return nil
}

func saveTimeSpanHandler(w http.ResponseWriter, r *http.Request) {
  destination := r.FormValue("destination")
  time := r.FormValue("time")
  if destination != "" && time != "" {
    timeSpanItem := &TimeSpan{Destination: destination, Time: time}
    err := timeSpanItem.save()
    if err == nil { fmt.Fprintf(w, "Time span saved.") } else { fmt.Fprintf(w, "Time span could not be saved:", err)}
  } else {
    fmt.Fprintf(w, "Time span could not be saved.")
  }
}

func main() {
    http.HandleFunc("/savetimespan", saveTimeSpanHandler)
    http.ListenAndServe(":8080", nil)
}