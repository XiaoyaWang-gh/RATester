package common

import (
	"fmt"
	"testing"
)

func TestDestroy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lg := &StoreMsg{}
	lg.Destroy()
}
