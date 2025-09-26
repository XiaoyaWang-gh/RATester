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
	files, err := nrf.Readdir(0)
	if files != nil {
		t.Errorf("Expected nil, got %v", files)
	}
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
