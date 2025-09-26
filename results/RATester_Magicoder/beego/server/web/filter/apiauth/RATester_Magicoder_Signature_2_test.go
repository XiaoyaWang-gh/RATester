package apiauth

import (
	"fmt"
	"net/url"
	"testing"
)

func TestSignature_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	appsecret := "your_app_secret"
	method := "GET"
	params := url.Values{}
	params.Add("key1", "value1")
	params.Add("key2", "value2")
	RequestURL := "http://your_request_url"

	expected := "expected_signature"
	result := Signature(appsecret, method, params, RequestURL)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
