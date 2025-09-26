package helpers

import (
	"fmt"
	"testing"
)

func TestFindCWD_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path, err := FindCWD()
	if err != nil {
		t.Errorf("FindCWD() error = %v", err)
		return
	}
	if path == "" {
		t.Errorf("FindCWD() path = %v", path)
	}
}
