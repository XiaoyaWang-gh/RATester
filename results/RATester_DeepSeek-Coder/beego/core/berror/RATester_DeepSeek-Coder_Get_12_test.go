package berror

import (
	"fmt"
	"sync"
	"testing"
)

func TestGet_12(t *testing.T) {
	cr := &codeRegistry{
		lock:  sync.RWMutex{},
		codes: map[uint32]*codeDefinition{},
	}

	testCode := &codeDefinition{
		code:   1,
		module: "test",
		desc:   "test description",
		name:   "test name",
	}

	cr.codes[testCode.code] = testCode

	t.Run("ExistingCode", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		c, ok := cr.Get(testCode.code)
		if !ok {
			t.Errorf("Expected code to exist but it did not")
		}

		if c.Code() != testCode.code || c.Module() != testCode.module || c.Desc() != testCode.desc || c.Name() != testCode.name {
			t.Errorf("Expected code %v, got %v", testCode, c)
		}
	})

	t.Run("NonExistingCode", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, ok := cr.Get(2)
		if ok {
			t.Errorf("Expected code not to exist but it did")
		}
	})
}
