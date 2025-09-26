package context

import (
	"fmt"
	"testing"
)

func TestParamsLen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	input.pnames = []string{"name1", "name2"}
	if input.ParamsLen() != 2 {
		t.Errorf("ParamsLen() = %v, want %v", input.ParamsLen(), 2)
	}
}
