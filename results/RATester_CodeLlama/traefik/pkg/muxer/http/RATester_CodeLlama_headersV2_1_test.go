package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHeadersV2_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	err := headersV2(tree, "header1", "header2")
	if err != nil {
		t.Errorf("headersV2() error = %v", err)
		return
	}

	if tree.matcher == nil {
		t.Errorf("headersV2() matcher = nil")
		return
	}

	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Errorf("http.NewRequest() error = %v", err)
		return
	}

	req.Header.Add("header1", "value1")
	req.Header.Add("header2", "value2")

	if !tree.match(req) {
		t.Errorf("headersV2() matcher = false")
		return
	}
}
