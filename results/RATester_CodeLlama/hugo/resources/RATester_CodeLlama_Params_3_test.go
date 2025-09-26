package resources

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParams_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &genericResource{
		params: map[string]any{
			"foo": "bar",
		},
	}

	if got := l.Params(); !reflect.DeepEqual(got, l.params) {
		t.Errorf("Params() = %v, want %v", got, l.params)
	}
}
