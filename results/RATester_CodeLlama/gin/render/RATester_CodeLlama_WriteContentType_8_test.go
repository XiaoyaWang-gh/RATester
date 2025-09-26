package render

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWriteContentType_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var r Redirect
	var w http.ResponseWriter
	r.WriteContentType(w)
}
