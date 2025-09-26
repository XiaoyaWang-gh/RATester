package langs

import (
	"fmt"
	"testing"
)

func TestGetCollator2_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Language{}
	l.collator2 = &Collator{}
	if got := GetCollator2(l); got != l.collator2 {
		t.Errorf("GetCollator2() = %v, want %v", got, l.collator2)
	}
}
