package binding

import (
	"fmt"
	"testing"
)

func TestEngine_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &defaultValidator{}
	v.Engine()
}
