package httplib

import (
	"fmt"
	"net/http"
	"testing"
)

func TestsendRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	b := &BeegoHTTPRequest{
		url: "http://example.com",
		req: req,
		setting: BeegoHTTPSettings{
			Retries: -1,
		},
	}
	resp, err := b.sendRequest(client)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
