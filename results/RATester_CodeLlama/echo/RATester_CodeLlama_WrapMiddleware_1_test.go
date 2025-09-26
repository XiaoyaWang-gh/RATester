package echo

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWrapMiddleware_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := func(h http.Handler) http.Handler {
		return h
	}
	wrapped := WrapMiddleware(m)
	if reflect.TypeOf(wrapped) != reflect.TypeOf(m) {
		t.Errorf("WrapMiddleware() = %v, want %v", reflect.TypeOf(wrapped), reflect.TypeOf(m))
	}
}
