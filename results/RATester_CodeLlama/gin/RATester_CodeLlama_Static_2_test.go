package gin

import (
	"fmt"
	"testing"
)

func TestStatic_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	group := &RouterGroup{}
	group.Static("relativePath", "root")
}
