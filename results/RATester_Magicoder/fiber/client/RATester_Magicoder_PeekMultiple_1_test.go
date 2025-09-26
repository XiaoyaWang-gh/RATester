package client

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestPeekMultiple_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	header := &Header{
		RequestHeader: &fasthttp.RequestHeader{},
	}

	header.Set("Key", "Value1")
	header.Set("Key", "Value2")
	header.Set("Key", "Value3")

	expected := []string{"Value1", "Value2", "Value3"}
	result := header.PeekMultiple("Key")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
