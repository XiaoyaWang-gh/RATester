package plugins

import (
	"fmt"
	"testing"

	"github.com/traefik/yaegi/interp"
)

func TestNewYaegiMiddlewareBuilder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := &interp.Interpreter{}
	basePkg := "test"
	imp := "test/test"

	_, err := newYaegiMiddlewareBuilder(i, basePkg, imp)
	if err != nil {
		t.Errorf("newYaegiMiddlewareBuilder() error = %v", err)
		return
	}

	// Add more test cases here
}
