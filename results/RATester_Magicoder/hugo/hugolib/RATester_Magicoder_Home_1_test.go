package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestHome_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	site := &Site{
		home: &pageState{
			pageCommon: &pageCommon{
				store: &maps.Scratch{},
			},
		},
	}

	expected := &pageState{
		pageCommon: &pageCommon{
			store: &maps.Scratch{},
		},
	}

	result := site.Home()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
