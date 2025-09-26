package gin

import (
	"fmt"
	"testing"
)

func TestAddRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &node{}
	path := "/test"
	handlers := HandlersChain{}
	n.addRoute(path, handlers)
}
