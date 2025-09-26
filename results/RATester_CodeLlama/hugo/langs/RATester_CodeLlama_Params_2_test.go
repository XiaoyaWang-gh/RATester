package langs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestParams_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Language{
		params: maps.Params{
			"foo": "bar",
		},
	}

	if got := l.Params(); !reflect.DeepEqual(got, l.params) {
		t.Errorf("Params() = %v, want %v", got, l.params)
	}
}
