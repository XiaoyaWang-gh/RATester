package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDrain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	type K string
	type T int

	cache := &Cache[K, T]{
		m: map[K]T{
			"one":   1,
			"two":   2,
			"three": 3,
		},
	}

	expected := map[K]T{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	drained := cache.Drain()

	if !reflect.DeepEqual(drained, expected) {
		t.Errorf("expected %v, got %v", expected, drained)
	}

	if len(cache.m) != 0 {
		t.Errorf("expected empty map, got %v", cache.m)
	}
}
