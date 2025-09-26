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
		pageRange:   []int{1, 2, 3},
		pageNums:    10,
		page:        5,
	}

	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Test case 1",
			want: true,
		},
		{
			name: "Test case 2",
			want: false,
		},
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
