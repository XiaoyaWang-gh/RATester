package constraints

import (
	"fmt"
	"testing"
)

func TestLabelRegexFn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "name"
	expr := "expr"
	labelRegexFn(name, expr)
}
