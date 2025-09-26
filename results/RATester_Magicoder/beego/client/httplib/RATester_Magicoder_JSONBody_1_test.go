package httplib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestJSONBody_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{
		url: "http://example.com",
		req: &http.Request{},
	}

	obj := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "John Doe",
		Age:  30,
	}

	_, err := b.JSONBody(obj)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if b.req.Body == nil {
		t.Error("Expected request body to be set, but it is nil")
	}

	body, err := ioutil.ReadAll(b.req.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}

	var result struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Errorf("Error unmarshaling request body: %v", err)
	}

	if result.Name != obj.Name || result.Age != obj.Age {
		t.Errorf("Expected %v, but got %v", obj, result)
	}
}
