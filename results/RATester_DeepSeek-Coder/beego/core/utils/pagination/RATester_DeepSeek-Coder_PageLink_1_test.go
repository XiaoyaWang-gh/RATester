package pagination

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPageLink_1(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/?p=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	p := &Paginator{
		Request: req,
	}
	tests := []struct {
		name string
		p    *Paginator
		page int
		want string
	}{
		{
			name: "first page",
			p:    p,
			page: 1,
			want: "http://example.com/?p=1",
		},
		{
			name: "second page",
			p:    p,
			page: 2,
			want: "http://example.com/?p=2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.PageLink(tt.page); got != tt.want {
				t.Errorf("Paginator.PageLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
