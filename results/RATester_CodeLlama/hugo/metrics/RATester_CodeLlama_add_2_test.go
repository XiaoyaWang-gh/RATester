package metrics

import (
	"fmt"
	"testing"
)

func TestAdd_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &diff{}
	v := 1
	d.add(v)
	if d.baseline != v {
		t.Errorf("baseline is not equal to v")
	}
	if d.count != 1 {
		t.Errorf("count is not equal to 1")
	}
	if d.simSum != 100 {
		t.Errorf("simSum is not equal to 100")
	}
}
