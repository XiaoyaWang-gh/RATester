package log

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithprintableContentTypes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testTypes := []string{"text/plain", "text/html"}
	builder := &FilterChainBuilder{}
	WithprintableContentTypes(testTypes)(builder)

	if !reflect.DeepEqual(builder.printableContentTypes, testTypes) {
		t.Errorf("Expected printableContentTypes to be %v, but got %v", testTypes, builder.printableContentTypes)
	}
}
