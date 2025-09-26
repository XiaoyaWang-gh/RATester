package web

import (
	"fmt"
	"testing"
)

func TestRun_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	admin := &adminApp{}
	admin.Run()
}
