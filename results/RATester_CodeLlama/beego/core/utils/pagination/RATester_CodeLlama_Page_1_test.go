package pagination

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestPage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Paginator{
		Request:     &http.Request{},
		PerPageNums: 10,
		MaxPages:    10,
		nums:        100,
	}
	if p.Page() != 1 {
		t.Errorf("p.Page() = %v, want %v", p.Page(), 1)
	}
	p.Request.Form = url.Values{"p": []string{"2"}}
	if p.Page() != 2 {
		t.Errorf("p.Page() = %v, want %v", p.Page(), 2)
	}
	p.Request.Form = url.Values{"p": []string{"1000"}}
	if p.Page() != 10 {
		t.Errorf("p.Page() = %v, want %v", p.Page(), 10)
	}
}
