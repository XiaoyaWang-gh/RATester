package releaser

import (
	"fmt"
	"testing"
)

func TestLogln_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	args := []interface{}{"hello", "world"}
	logln(args...)
}
