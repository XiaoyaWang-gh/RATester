package template

import (
	"fmt"
	"testing"
)

func TestNextJSCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var s []byte
	var preceding jsCtx
	var expected jsCtx
	var actual jsCtx
	actual = nextJSCtx(s, preceding)
	if actual != expected {
		t.Errorf("nextJSCtx(%v, %v) = %v; expected %v", s, preceding, actual, expected)
	}
}
