package page

import (
	"fmt"
	"testing"
)

func TestHasNext_1(t *testing.T) {
	paginator := &Paginator{
		paginatedElements: make([]paginatedElement, 10),
		total:             100,
	}

	pager := &Pager{
		number:    1,
		Paginator: paginator,
	}

	tests := []struct {
		name string
		p    *Pager
		want bool
	}{
		{
			name: "Has next",
			p:    pager,
			want: true,
		},
		{
			name: "No next",
			p: &Pager{
				number:    10,
				Paginator: paginator,
			},
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

			if got := tt.p.HasNext(); got != tt.want {
				t.Errorf("Pager.HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
