package task

import (
	"fmt"
	"testing"
)

func TestParseIntOrName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expr := "1"
	names := map[string]uint{
		"a": 1,
		"b": 2,
	}
	if got := parseIntOrName(expr, names); got != 1 {
		t.Errorf("parseIntOrName() = %v, want %v", got, 1)
	}
}
