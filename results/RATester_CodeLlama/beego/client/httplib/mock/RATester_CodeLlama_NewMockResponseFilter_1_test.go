package mock

import (
	"fmt"
	"testing"
)

func TestNewMockResponseFilter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	NewMockResponseFilter()
}
