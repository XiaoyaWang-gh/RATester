package es

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/logs"
)

func Testinit_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logs.Register(logs.AdapterEs, NewES)

	logger := NewES()
	if logger == nil {
		t.Error("Expected logger to be not nil")
	}
}
