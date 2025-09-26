package middlewares

import (
	"fmt"
	"testing"
)

func TestFlush_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &ResponseModifier{}
	r.Flush()
}
