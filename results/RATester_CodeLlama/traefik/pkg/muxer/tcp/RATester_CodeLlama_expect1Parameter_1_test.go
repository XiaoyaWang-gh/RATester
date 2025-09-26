package tcp

import (
	"fmt"
	"testing"
)

func TestExpect1Parameter_1(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fn := func(route *matchersTree, s ...string) error {
			return nil
		}
		expect1Parameter(fn)(nil, "test")
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fn := func(route *matchersTree, s ...string) error {
			return fmt.Errorf("error")
		}
		err := expect1Parameter(fn)(nil, "test")
		if err == nil {
			t.Fatal("expected error")
		}
		if err.Error() != "error" {
			t.Fatalf("expected error %q, got %q", "error", err.Error())
		}
	})
}
