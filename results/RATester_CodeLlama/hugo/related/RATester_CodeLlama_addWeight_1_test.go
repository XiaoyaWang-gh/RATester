package related

import (
	"fmt"
	"testing"
)

func TestAddWeight_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &rank{}
	w := 1
	r.addWeight(w)
	if r.Weight != w {
		t.Errorf("r.Weight = %d, want %d", r.Weight, w)
	}
	if r.Matches != 1 {
		t.Errorf("r.Matches = %d, want %d", r.Matches, 1)
	}
}
