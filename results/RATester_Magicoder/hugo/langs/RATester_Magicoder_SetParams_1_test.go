package langs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestSetParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Language{
		Lang: "en",
		params: maps.Params{
			"key1": "value1",
			"key2": "value2",
		},
	}

	newParams := maps.Params{
		"key3": "value3",
		"key4": "value4",
	}

	SetParams(l, newParams)

	if !reflect.DeepEqual(l.params, newParams) {
		t.Errorf("Expected params to be %v, but got %v", newParams, l.params)
	}
}
