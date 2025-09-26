package fiber

import (
	"fmt"
	"testing"
)

func TestOnListen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	h := &Hooks{}
	handler := []func(ListenData) error{}
	h.OnListen(handler...)
	// Act
	h.executeOnListenHooks(ListenData{})
	// Assert
	// ...
}
