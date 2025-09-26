package api

import (
	"fmt"
	"testing"
)

func TestRule_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := udpRouterRepresentation{}
	want := ""
	got := r.rule()
	if got != want {
		t.Errorf("rule() = %v, want %v", got, want)
	}
}
