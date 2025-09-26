package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Request: &http.Request{
			Header: http.Header{
				"Cookie": []string{"name=value"},
			},
		},
	}
	val, err := c.Cookie("name")
	if err != nil {
		t.Errorf("Cookie() error = %v", err)
		return
	}
	if val != "value" {
		t.Errorf("Cookie() val = %v, want %v", val, "value")
	}
}
