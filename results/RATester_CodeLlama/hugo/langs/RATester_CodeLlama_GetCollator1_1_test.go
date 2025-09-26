package langs

import (
	"fmt"
	"testing"
)

func TestGetCollator1_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Language{}
	l.collator1 = &Collator{}
	if got := GetCollator1(l); got != l.collator1 {
		t.Errorf("GetCollator1() = %v, want %v", got, l.collator1)
	}
}
