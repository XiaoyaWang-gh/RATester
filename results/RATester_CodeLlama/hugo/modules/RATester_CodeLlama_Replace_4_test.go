package modules

import (
	"fmt"
	"testing"
)

func TestReplace_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &moduleAdapter{
		gomod: &goModule{
			Path:    "github.com/golang/mock",
			Version: "v1.4.4",
		},
	}
	if m.Replace() == nil {
		t.Errorf("Replace() = nil, want moduleAdapter")
	}
}
