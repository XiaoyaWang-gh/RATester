package web

import (
	"fmt"
	"testing"
)

func TestList_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var m M
	list("", &struct {
		A int
		B string
		C struct {
			D int
			E string
		}
	}{
		A: 1,
		B: "2",
		C: struct {
			D int
			E string
		}{
			D: 3,
			E: "4",
		},
	}, m)
	if len(m) != 5 {
		t.Errorf("m length is %d, want 5", len(m))
	}
	if m["A"] != 1 {
		t.Errorf("m[\"A\"] is %v, want 1", m["A"])
	}
	if m["B"] != "2" {
		t.Errorf("m[\"B\"] is %v, want \"2\"", m["B"])
	}
	if m["C.D"] != 3 {
		t.Errorf("m[\"C.D\"] is %v, want 3", m["C.D"])
	}
	if m["C.E"] != "4" {
		t.Errorf("m[\"C.E\"] is %v, want \"4\"", m["C.E"])
	}
}
