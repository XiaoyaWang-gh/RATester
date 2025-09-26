package fiber

import (
	"fmt"
	"testing"
)

func Testregister_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := new(App)

	methods := []string{"GET", "POST"}
	pathRaw := "/test"
	group := new(Group)
	handler := func(c Ctx) error { return nil }
	middleware := []func(Ctx) error{func(c Ctx) error { return nil }}

	app.register(methods, pathRaw, group, handler, middleware...)

	// Test cases
	// ...
}
