package cache

import (
	"fmt"
	"testing"
)

func TestGetString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var v interface{}
	if GetString(v) != "" {
		t.Errorf("GetString(%v) = %v, want %v", v, GetString(v), "")
	}
	if GetString("") != "" {
		t.Errorf("GetString(%v) = %v, want %v", v, GetString(v), "")
	}
	if GetString("hello") != "hello" {
		t.Errorf("GetString(%v) = %v, want %v", v, GetString(v), "hello")
	}
	if GetString([]byte("hello")) != "hello" {
		t.Errorf("GetString(%v) = %v, want %v", v, GetString(v), "hello")
	}
	if GetString(1) != "1" {
		t.Errorf("GetString(%v) = %v, want %v", v, GetString(v), "1")
	}
	if GetString(1.0) != "1" {
		t.Errorf("GetString(%v) = %v, want %v", v, GetString(v), "1")
	}
	if GetString(true) != "true" {
		t.Errorf("GetString(%v) = %v, want %v", v, GetString(v), "true")
	}
	if GetString(false) != "false" {
		t.Errorf("GetString(%v) = %v, want %v", v, GetString(v), "false")
	}
	if GetString(nil) != "" {
		t.Errorf("GetString(%v) = %v, want %v", v, GetString(v), "")
	}
}
