package httplib

import (
	"fmt"
	"testing"
)

func TestGet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Client{}
	value := &struct{}{}
	path := "/path"
	opts := []BeegoHTTPRequestOption{}
	err := c.Get(value, path, opts...)
	if err != nil {
		t.Errorf("Get() error = %v", err)
		return
	}
}
