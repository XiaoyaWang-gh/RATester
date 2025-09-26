package predicate

import (
	"fmt"
	"testing"
)

func TestOr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := P[int](func(v int) bool {
		return v > 10
	})
	ps := []P[int]{
		P[int](func(v int) bool {
			return v > 100
		}),
		P[int](func(v int) bool {
			return v > 1000
		}),
	}
	got := p.Or(ps...)(1000)
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
