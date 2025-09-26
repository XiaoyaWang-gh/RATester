package context

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

func TestBody_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{
		Context: &Context{
			Request: &http.Request{},
		},
		EnableGzip: true,
	}

	content := []byte("test content")
	err := output.Body(content)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if output.Context.ResponseWriter.Started != true {
		t.Errorf("Expected ResponseWriter.Started to be true, but got false")
	}

	if output.Context.ResponseWriter.Header().Get("Content-Encoding") != "gzip" {
		t.Errorf("Expected Content-Encoding to be gzip, but got %v", output.Context.ResponseWriter.Header().Get("Content-Encoding"))
	}

	if output.Context.ResponseWriter.Header().Get("Content-Length") != strconv.Itoa(len(content)) {
		t.Errorf("Expected Content-Length to be %v, but got %v", len(content), output.Context.ResponseWriter.Header().Get("Content-Length"))
	}
}
