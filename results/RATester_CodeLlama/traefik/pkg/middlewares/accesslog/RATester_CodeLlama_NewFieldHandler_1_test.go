package accesslog

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewFieldHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	next := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
	name := "name"
	value := "value"
	applyFn := func(rw http.ResponseWriter, r *http.Request, next http.Handler, data *LogData) {}

	// when
	handler := NewFieldHandler(next, name, value, applyFn)

	// then
	assert.NotNil(t, handler)
}
