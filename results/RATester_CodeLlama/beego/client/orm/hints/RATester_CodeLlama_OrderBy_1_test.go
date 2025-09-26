package hints

import (
	"fmt"
	"testing"
)

func TestOrderBy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := "order by id desc"
	h := OrderBy(s)
	if h.GetKey() != KeyOrderBy {
		t.Errorf("Expected %v, but got %v", KeyOrderBy, h.GetKey())
	}
	if h.GetValue() != s {
		t.Errorf("Expected %v, but got %v", s, h.GetValue())
	}
}
