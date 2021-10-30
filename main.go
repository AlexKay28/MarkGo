package main

import (
  "fmt"
  "net/http"
  "html/template"
)

type Graph struct {
  X_axis_name string
  Y_axis_name string
  Legend string
  X_points []int
  Y_points []int
}

func (g *Graph) getPoints() ([]int, []int) {
  return g.X_points, g.Y_points
}

func home_page(w http.ResponseWriter, r *http.Request) {

  needed_graph := Graph{
    X_axis_name: "X", Y_axis_name: "Y", Legend: "This is the legend of the graph",
    X_points: []int{1, 2, 3 ,4 ,5},
    Y_points: []int{1, 2, 3 ,4 ,5},
  }

  tmpl, _ := template.ParseFiles("templates/home_page.html")
  tmpl.Execute(w, needed_graph)
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
