package client

import (
	"fmt"
	"testing"
)

func TestSetURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	url := "http://example.com"
	r.SetURL(url)
	if r.url != url {
		t.Errorf("Expected url to be %s, but got %s", url, r.url)
	}
}
