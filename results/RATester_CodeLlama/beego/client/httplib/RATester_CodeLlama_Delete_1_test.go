package httplib

import (
	"fmt"
	"testing"
)

func TestDelete_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Client{}
	value := &struct{}{}
	path := "/path"
	opts := []BeegoHTTPRequestOption{}
	err := c.Delete(value, path, opts...)
	if err != nil {
		t.Errorf("Delete() error = %v", err)
		return
	}
}
