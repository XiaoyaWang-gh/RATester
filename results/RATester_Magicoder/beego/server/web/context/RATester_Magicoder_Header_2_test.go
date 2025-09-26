package context

import (
	"fmt"
	"testing"
)

func TestHeader_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{
		Context: &Context{
			ResponseWriter: &Response{},
		},
	}

	key := "Content-Type"
	val := "application/json"

	output.Header(key, val)

	if output.Context.ResponseWriter.Header().Get(key) != val {
		t.Errorf("Expected %s, got %s", val, output.Context.ResponseWriter.Header().Get(key))
	}
}
