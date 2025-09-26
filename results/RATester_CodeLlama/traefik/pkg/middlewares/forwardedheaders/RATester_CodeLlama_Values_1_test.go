package forwardedheaders

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValues_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := unsafeHeader{}
	h.Set("key", "value")
	h.Set("key", "value2")
	h.Set("key", "value3")
	if got := h.Values("key"); !reflect.DeepEqual(got, []string{"value", "value2", "value3"}) {
		t.Errorf("Values() = %v, want %v", got, []string{"value", "value2", "value3"})
	}
}
