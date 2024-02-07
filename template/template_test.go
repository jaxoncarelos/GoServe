package template

import (
	"strings"
	"testing"
)


func TestReadTemplate(t *testing.T) {
  idealResponse := "<div>Hi from index</div>\n"
  response, err := ReadTemplate("../templates/index.html")
  if err != nil {
    t.Fatal("Error: ", err)
  }
  stringResponse := strings.TrimRight(string(response), "\x00")
  if stringResponse != idealResponse {
    t.Fatalf("Expected:\n \"%s\", got:\n \"%s\"", idealResponse, response)
  }
}
