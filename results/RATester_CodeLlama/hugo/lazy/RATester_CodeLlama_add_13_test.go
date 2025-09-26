package lazy

import (
	"context"
	"fmt"
	"testing"
)

func TestAdd_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ini := &Init{}
	branch := true
	initFn := func(context.Context) (any, error) {
		return nil, nil
	}
	ini.add(branch, initFn)
}
