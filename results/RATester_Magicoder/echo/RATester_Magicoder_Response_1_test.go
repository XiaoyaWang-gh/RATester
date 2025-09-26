package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &context{
		response: &Response{},
	}

	expected := &Response{}
	actual := ctx.Response()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
