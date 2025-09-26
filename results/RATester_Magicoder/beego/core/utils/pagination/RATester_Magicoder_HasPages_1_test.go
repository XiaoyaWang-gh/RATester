package pagination

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHasPages_1(t *testing.T) {
	paginator := &Paginator{
		Request:     &http.Request{},
		PerPageNums: 10,
		MaxPages:    100,
		nums:        1000,
		pageRange:   []int{},
		pageNums:    1,
		page:        1,
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

			if got := paginator.HasPages(); got != tt.want {
				t.Errorf("HasPages() = %v, want %v", got, tt.want)
			}
		})
	}
}
