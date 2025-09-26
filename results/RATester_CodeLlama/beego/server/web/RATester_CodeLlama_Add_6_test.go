package web

import (
	"fmt"
	"testing"
)

func TestAdd_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &ControllerRegister{}
	p.Add("", nil, nil)
}
