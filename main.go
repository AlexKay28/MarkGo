package main

import (
  "fmt"
  "net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "This is home page!")
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "This is contacts page!")
}

func handleRequest() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/contacts/", contacts_page)
  http.ListenAndServe(":5566", nil)
}

func main() {
  handleRequest()
}
