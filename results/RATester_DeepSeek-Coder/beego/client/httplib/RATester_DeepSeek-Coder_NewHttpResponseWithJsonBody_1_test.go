package httplib

import (
	"fmt"
	"io"
	"testing"
)

func TestNewHttpResponseWithJsonBody_1(t *testing.T) {
	t.Run("Test with string data", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		data := `{"name": "test", "age": 25}`
		resp := NewHttpResponseWithJsonBody(data)
		if resp.ContentLength != int64(len(data)) {
			t.Errorf("Expected ContentLength to be %d, got %d", len(data), resp.ContentLength)
		}
	})

	t.Run("Test with []byte data", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		data := []byte(`{"name": "test", "age": 25}`)
		resp := NewHttpResponseWithJsonBody(data)
		if resp.ContentLength != int64(len(data)) {
			t.Errorf("Expected ContentLength to be %d, got %d", len(data), resp.ContentLength)
		}
	})

	t.Run("Test with struct data", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		type Person struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		data := Person{Name: "test", Age: 25}
		resp := NewHttpResponseWithJsonBody(data)
		expectedBody := `{"name":"test","age":25}`
		body, _ := io.ReadAll(resp.Body)
		if resp.ContentLength != int64(len(expectedBody)) {
			t.Errorf("Expected ContentLength to be %d, got %d", len(expectedBody), resp.ContentLength)
		}
		if string(body) != expectedBody {
			t.Errorf("Expected body to be %s, got %s", expectedBody, string(body))
		}
	})
}
