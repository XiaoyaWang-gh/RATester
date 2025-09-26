package mirror

import (
	"fmt"
	"testing"
)

func TestFlush_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := blackHoleResponseWriter{}
	b.Flush()
}
