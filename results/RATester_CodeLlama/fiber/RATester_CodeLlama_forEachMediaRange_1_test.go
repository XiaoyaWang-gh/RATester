package fiber

import (
	"fmt"
	"testing"
)

func TestForEachMediaRange_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	header := []byte("text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	functor := func([]byte) {}
	forEachMediaRange(header, functor)
}
