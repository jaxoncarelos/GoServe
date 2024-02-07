package main

import (
	"github.com/jaxoncarelos/GoWebServer/template"
	"log"
	"net/http"
	"strings"
)

func Last[T any](array []T) T {
  return array[len(array)-1]
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  path := Last(strings.Split(r.URL.Path, "/")) + ".html"
  if path == ".html" {
    path = "index.html"
  }
  w.Header().Set("Content-Type", "text/html")
  contents, err := template.ReadTemplate("./templates/" + path)
  if err != nil {
    log.Printf("Error : %s\n", err)
    return
  }
  _, writeErr := w.Write(contents)  
  if writeErr != nil {
    log.Printf("Error: %s\n", err)
  }
}
func main(){
  log.Println("Hello world")
  handler := http.NewServeMux()
  handler.Handle("/", &homeHandler{}) 
  err := http.ListenAndServe("localhost:8080", handler)
  if err != nil {
    log.Printf("Error: %s\n", err)
  }
}
