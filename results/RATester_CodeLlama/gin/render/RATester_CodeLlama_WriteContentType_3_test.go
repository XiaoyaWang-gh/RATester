package render

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWriteContentType_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	var jsonpJsonpJSON JsonpJSON
	var w http.ResponseWriter
	jsonpJsonpJSON.WriteContentType(w)
}
