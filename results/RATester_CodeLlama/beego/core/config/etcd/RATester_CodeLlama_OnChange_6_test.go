package etcd

import (
	"fmt"
	"testing"
)

func TestOnChange_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
