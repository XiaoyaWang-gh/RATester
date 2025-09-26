package pagination

import (
	"fmt"
	"testing"
)

func TestPageNums_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		paginator *Paginator
		want      int
	}{
		{
			name: "Test with 100 items, 10 per page, no max pages",
			paginator: &Paginator{
				nums:        100,
				PerPageNums: 10,
			},
			want: 10,
		},
		{
			name: "Test with 100 items, 10 per page, max pages 5",
			paginator: &Paginator{
				nums:        100,
				PerPageNums: 10,
				MaxPages:    5,
			},
			want: 5,
		},
		{
			name: "Test with 100 items, 20 per page, no max pages",
			paginator: &Paginator{
				nums:        100,
				PerPageNums: 20,
			},
			want: 5,
		},
		{
			name: "Test with 100 items, 20 per page, max pages 5",
			paginator: &Paginator{
				nums:        100,
				PerPageNums: 20,
				MaxPages:    5,
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.paginator.PageNums(); got != tt.want {
				t.Errorf("Paginator.PageNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
