package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestAddData_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	formData := &FormData{
		Args: fasthttp.AcquireArgs(),
	}
	defer fasthttp.ReleaseArgs(formData.Args)

	key := "testKey"
	val := "testVal"

	formData.AddData(key, val)

	if formData.Args.Len() != 1 {
		t.Errorf("Expected formData.Args.Len() to be 1, but got %d", formData.Args.Len())
	}

	if string(formData.Args.Peek(key)) != val {
		t.Errorf("Expected formData.Args.Peek(%s) to be %s, but got %s", key, val, formData.Args.Peek(key))
	}
}
