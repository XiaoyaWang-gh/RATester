package schema

import (
	"errors"
	"fmt"
	"testing"
)

func TestMerge_1(t *testing.T) {
	t.Run("Test merge function", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		e1 := make(MultiError)
		e2 := make(MultiError)

		e1["key1"] = errors.New("error1")
		e2["key2"] = errors.New("error2")

		e1.merge(e2)

		if len(e1) != 2 {
			t.Errorf("Expected length of e1 to be 2, got %d", len(e1))
		}

		if e1["key1"] == nil || e1["key2"] == nil {
			t.Errorf("Expected e1 to contain both errors, got %v", e1)
		}
	})
}
