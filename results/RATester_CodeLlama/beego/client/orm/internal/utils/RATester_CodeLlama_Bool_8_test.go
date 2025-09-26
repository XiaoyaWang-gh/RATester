package utils

import (
	"fmt"
	"testing"
)

func TestBool_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var f StrTo
	f.Set("true")
	b, err := f.Bool()
	if err != nil {
		t.Errorf("Bool() error = %v", err)
		return
	}
	if !b {
		t.Errorf("Bool() = %v, want %v", b, true)
	}

	f.Set("false")
	b, err = f.Bool()
	if err != nil {
		t.Errorf("Bool() error = %v", err)
		return
	}
	if b {
		t.Errorf("Bool() = %v, want %v", b, false)
	}

	f.Set("1")
	b, err = f.Bool()
	if err != nil {
		t.Errorf("Bool() error = %v", err)
		return
	}
	if !b {
		t.Errorf("Bool() = %v, want %v", b, true)
	}

	f.Set("0")
	b, err = f.Bool()
	if err != nil {
		t.Errorf("Bool() error = %v", err)
		return
	}
	if b {
		t.Errorf("Bool() = %v, want %v", b, false)
	}

	f.Set("")
	_, err = f.Bool()
	if err == nil {
		t.Errorf("Bool() error = %v, want %v", err, "error")
		return
	}
}
