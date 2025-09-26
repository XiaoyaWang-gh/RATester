package alils

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/logs"
)

func Testinit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logs.Register(logs.AdapterAliLS, NewAliLS)

	logger := NewAliLS()
	if logger == nil {
		t.Error("Failed to create a new logger")
	}
}
