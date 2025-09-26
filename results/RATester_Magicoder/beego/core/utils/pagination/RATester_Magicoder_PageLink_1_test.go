package pagination

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPageLink_1(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/foo?p=1", nil)
	paginator := &Paginator{Request: req, PerPageNums: 10}

	tests := []struct {
		name string
		page int
		want string
	}{
		{"First page", 1, "http://example.com/foo?p=1"},
		{"Second page", 2, "http://example.com/foo?p=2"},
		{"Third page", 3, "http://example.com/foo?p=3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := paginator.PageLink(tt.page); got != tt.want {
				t.Errorf("PageLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
