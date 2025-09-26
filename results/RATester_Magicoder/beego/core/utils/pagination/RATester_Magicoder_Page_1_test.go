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

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	paginator := &Paginator{
		Request: req,
	}

	// Test case 1: p is not in the form, should return 1
	paginator.page = 0
	paginator.Request.Form = nil
	page := paginator.Page()
	if page != 1 {
		t.Errorf("Expected 1, got %d", page)
	}

	// Test case 2: p is in the form, should return the parsed value
	paginator.Request.Form = url.Values{"p": {"2"}}
	page = paginator.Page()
	if page != 2 {
		t.Errorf("Expected 2, got %d", page)
	}

	// Test case 3: p is greater than PageNums, should return PageNums
	paginator.Request.Form = url.Values{"p": {"10"}}
	page = paginator.Page()
	if page != paginator.PageNums() {
		t.Errorf("Expected %d, got %d", paginator.PageNums(), page)
	}

	// Test case 4: p is less than or equal to 0, should return 1
	paginator.Request.Form = url.Values{"p": {"-1"}}
	page = paginator.Page()
	if page != 1 {
		t.Errorf("Expected 1, got %d", page)
	}
}
