package mirror

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClone_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := reusableRequest{
		req:  req,
		body: []byte("test body"),
	}

	cloneReq := rr.clone(ctx)

	if cloneReq.Method != "GET" {
		t.Errorf("Expected method to be 'GET', got %s", cloneReq.Method)
	}

	if cloneReq.URL.String() != "http://example.com" {
		t.Errorf("Expected URL to be 'http://example.com', got %s", cloneReq.URL.String())
	}

	body, err := ioutil.ReadAll(cloneReq.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "test body" {
		t.Errorf("Expected body to be 'test body', got %s", string(body))
	}
}
