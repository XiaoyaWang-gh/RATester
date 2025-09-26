package common

import (
	"fmt"
	"testing"
)

func TestInit_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lg := &StoreMsg{}
	err := lg.Init("")
	if err != nil {
		t.Errorf("Init() error = %v", err)
		return
	}
	lg.Destroy()
}
