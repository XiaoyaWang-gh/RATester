package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestOpen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := OnlyFilesFS{http.Dir(".")}

	_, err := fs.Open("test.txt")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
