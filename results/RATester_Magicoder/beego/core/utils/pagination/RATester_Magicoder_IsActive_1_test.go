package pagination

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsActive_1(t *testing.T) {
	paginator := &Paginator{
		Request:     &http.Request{},
		PerPageNums: 10,
		MaxPages:    100,
		nums:        1000,
		pageRange:   []int{1, 2, 3},
		pageNums:    10,
		page:        5,
	}

	tests := []struct {
		name string
		page int
		want bool
	}{
		{"Test case 1", 5, true},
		{"Test case 2", 10, false},
		{"Test case 3", 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := paginator.IsActive(tt.page); got != tt.want {
				t.Errorf("IsActive() = %v, want %v", got, tt.want)
			}
		})
	}
}
