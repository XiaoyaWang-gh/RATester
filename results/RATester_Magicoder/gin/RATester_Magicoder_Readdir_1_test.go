package gin

import (
	"fmt"
	"testing"
)

func TestReaddir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nrf := neutralizedReaddirFile{}
	_, err := nrf.Readdir(0)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
