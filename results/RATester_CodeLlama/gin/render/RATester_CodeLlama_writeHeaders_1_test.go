package render

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWriteHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var r Reader
	var w http.ResponseWriter
	var headers map[string]string
	r.writeHeaders(w, headers)
}
