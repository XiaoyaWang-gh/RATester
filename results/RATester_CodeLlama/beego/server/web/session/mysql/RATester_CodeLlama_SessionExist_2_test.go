package mysql

import (
	"fmt"
	"testing"
)

func TestSessionExist_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
