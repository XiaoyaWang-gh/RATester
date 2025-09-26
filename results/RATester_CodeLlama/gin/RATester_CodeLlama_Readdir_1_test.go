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

	var n neutralizedReaddirFile
	_, err := n.Readdir(1)
	if err != nil {
		t.Errorf("Readdir() error = %v", err)
		return
	}
}
