package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestParams_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &genericResource{
		params: map[string]any{
			"key1": "value1",
			"key2": "value2",
		},
	}

	expectedParams := maps.Params{
		"key1": "value1",
		"key2": "value2",
	}

	result := l.Params()

	if !reflect.DeepEqual(result, expectedParams) {
		t.Errorf("Expected %v, but got %v", expectedParams, result)
	}
}
