package gin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetRawData_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Request: &http.Request{
			Body: ioutil.NopCloser(bytes.NewBufferString("hello")),
		},
	}
	data, err := c.GetRawData()
	if err != nil {
		t.Errorf("GetRawData() error = %v", err)
		return
	}
	if string(data) != "hello" {
		t.Errorf("GetRawData() = %v, want %v", string(data), "hello")
	}
}
