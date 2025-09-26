package nathole

import (
	"fmt"
	"testing"
)

func TestCloseClient_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Controller{}
	c.CloseClient("test")
}
