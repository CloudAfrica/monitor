package main

import (
  "fmt"
  "net/http"
  "os"
  "io"
  "io/ioutil"
  "regexp"
  "strings"
)

// Variables

type TimeSpan struct {
  Destination string
  Time string
}

// Helper functions

func (p *TimeSpan) save() error {
  f, err := os.OpenFile("timespans.csv", os.O_RDWR|os.O_APPEND, 0660)
  if err != nil {
    fmt.Println("Error opening file.")
    return err
  }

  fmt.Println("Destination: " + p.Destination)
  fmt.Println("Time: " + p.Time)

  if _, err2 := io.WriteString(f, p.Destination + "," + p.Time + "\r\n"); err2 != nil {
    fmt.Println("Error saving file: ", err2)
    return err2
  }
  f.Close()
  return nil
}

func removeAllSpaces(s string) string {
 matcher := regexp.MustCompile(`\n+|\r+`)

 s = matcher.ReplaceAllString(s, "")
 s = strings.Replace(s, " ", "", -1)
 return s
}

// HTTP Handler functions

func saveTimeSpanHandler(w http.ResponseWriter, r *http.Request) {
  destination := r.FormValue("destination")
  time := r.FormValue("time")
  if destination != "" && time != "" {
    timeSpanItem := &TimeSpan{Destination: destination, Time: time}
    err := timeSpanItem.save()
    if err == nil {
      fmt.Fprintf(w, "Time span saved.")
    } else {
      http.Error(w, "Time span could not be saved.", 500)
    }
  } else {
    http.Error(w, "Time span could not be saved.", 500)
  }
}

func returnProbesHandler(w http.ResponseWriter, r *http.Request) {
  if sitesJson, err := ioutil.ReadFile("sites.txt"); err != nil {
    http.Error(w, "Could not return a list of sites to probe.", 500)
  } else {
    fmt.Fprintf(w, removeAllSpaces(string(sitesJson)))
  }
}

func main() {
    http.HandleFunc("/savetimespan", saveTimeSpanHandler)
    http.HandleFunc("/probes", returnProbesHandler)
    http.ListenAndServe(":8080", nil)
}