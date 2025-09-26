package media

import (
	"fmt"
	"testing"
)

func TestInit_54(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Type{}
	m.init()
	if m.FirstSuffix.FullSuffix != "" {
		t.Errorf("m.FirstSuffix.FullSuffix = %q, want %q", m.FirstSuffix.FullSuffix, "")
	}
	if m.FirstSuffix.Suffix != "" {
		t.Errorf("m.FirstSuffix.Suffix = %q, want %q", m.FirstSuffix.Suffix, "")
	}
}
