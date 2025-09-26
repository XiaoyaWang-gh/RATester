package requestdecorator

import (
	"fmt"
	"testing"
	"time"
)

func TestLen_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := byTTL{
		&cnameResolv{TTL: 1 * time.Second, Record: "test1"},
		&cnameResolv{TTL: 2 * time.Second, Record: "test2"},
		&cnameResolv{TTL: 3 * time.Second, Record: "test3"},
	}

	expected := 3
	actual := a.Len()

	if actual != expected {
		t.Errorf("Expected Len to be %d, but got %d", expected, actual)
	}
}
