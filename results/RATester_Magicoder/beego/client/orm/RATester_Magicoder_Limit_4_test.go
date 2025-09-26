package orm

import (
	"fmt"
	"testing"
)

func TestLimit_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := &querySet{
		limit: 10,
	}

	qs.Limit(20, 10)

	if qs.limit != 20 {
		t.Errorf("Expected limit to be 20, but got %d", qs.limit)
	}

	if qs.offset != 10 {
		t.Errorf("Expected offset to be 10, but got %d", qs.offset)
	}

	qs.Limit(0, 20)

	if qs.limit != 1000 {
		t.Errorf("Expected limit to be 1000, but got %d", qs.limit)
	}

	qs.Limit(10)

	if qs.limit != 10 {
		t.Errorf("Expected limit to be 10, but got %d", qs.limit)
	}

	if qs.offset != 0 {
		t.Errorf("Expected offset to be 0, but got %d", qs.offset)
	}
}
