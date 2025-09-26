package filecache

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &httpCache{}
	id := "id"
	resp := []byte("resp")
	ok := true

	h.Set(id, resp)
	b, ok := h.Get(id)

	if !ok {
		t.Errorf("Get() = (%v, %v), want (%v, %v)", b, ok, resp, true)
	}

	if !reflect.DeepEqual(b, resp) {
		t.Errorf("Get() = (%v, %v), want (%v, %v)", b, ok, resp, true)
	}
}
