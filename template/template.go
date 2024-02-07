package template

import (
	"os"
	"path/filepath"
	"strings"
)

/*
	ReadTemplate Method

Takes in a filename, and attempts to read it to a byte array
Returns byte array and the potential error.
Byte array is empty if error is not nil
*/
func ReadTemplate(filename string) ([]byte, error) {
  pageNotFoundError := []byte("<div>404: Page not found</div>") 
	if !strings.HasSuffix(filename, ".html") {
		return pageNotFoundError, nil
	}
	relPath := filename

	absolutePath, err := filepath.Abs(relPath)
  if err != nil {
		return pageNotFoundError, nil
  }
	file, err := os.Open(absolutePath)
	if err != nil {
		// Return the error page if it can not get the html of the page
		return pageNotFoundError, nil
	}
	defer file.Close()
  fi, err := os.Stat(absolutePath)
  if err != nil {
		return pageNotFoundError, err
  }
	contents := make([]byte, fi.Size())
	_, err = file.Read(contents)
	if err != nil {
		return pageNotFoundError, nil
	}
	return contents, err
}
