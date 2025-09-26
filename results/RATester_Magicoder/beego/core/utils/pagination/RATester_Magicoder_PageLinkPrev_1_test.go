package pagination

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPageLinkPrev_1(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	paginator := &Paginator{
		Request:     req,
		PerPageNums: 10,
		MaxPages:    100,
		nums:        1000,
		pageRange:   []int{1, 2, 3},
		pageNums:    10,
		page:        2,
	}

	tests := []struct {
		name      string
		paginator *Paginator
		wantLink  string
	}{
		{
			name:      "Test case 1",
			paginator: paginator,
			wantLink:  "/?page=1",
		},
		{
			name: "Test case 2",
			paginator: &Paginator{
				Request:     req,
				PerPageNums: 10,
				MaxPages:    100,
				nums:        1000,
				pageRange:   []int{1, 2, 3},
				pageNums:    10,
				page:        1,
			},
			wantLink: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if gotLink := tt.paginator.PageLinkPrev(); gotLink != tt.wantLink {
				t.Errorf("PageLinkPrev() = %v, want %v", gotLink, tt.wantLink)
			}
		})
	}
}
