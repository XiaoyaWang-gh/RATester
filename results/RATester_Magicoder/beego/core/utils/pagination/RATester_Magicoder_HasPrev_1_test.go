package pagination

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHasPrev_1(t *testing.T) {
	paginator := &Paginator{
		Request:     &http.Request{},
		PerPageNums: 10,
		MaxPages:    100,
		nums:        1000,
		pageRange:   []int{},
		pageNums:    10,
		page:        1,
	}

	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Test case 1",
			want: false,
		},
		{
			name: "Test case 2",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := paginator.HasPrev(); got != tt.want {
				t.Errorf("Paginator.HasPrev() = %v, want %v", got, tt.want)
			}
		})
	}
}
