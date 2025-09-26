package pagination

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHasNext_1(t *testing.T) {
	paginator := &Paginator{
		Request:     &http.Request{},
		PerPageNums: 10,
		MaxPages:    100,
		nums:        1000,
		pageRange:   []int{1, 2, 3, 4, 5},
		pageNums:    5,
		page:        1,
	}

	tests := []struct {
		name string
		want bool
	}{
		{"TestHasNext_1", true},
		{"TestHasNext_2", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := paginator.HasNext(); got != tt.want {
				t.Errorf("Paginator.HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
