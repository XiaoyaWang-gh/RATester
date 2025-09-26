package gin

import (
	"fmt"
	"testing"
)

func TestCreateStaticHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	group := &RouterGroup{}
	group.createStaticHandler("", nil)
}
