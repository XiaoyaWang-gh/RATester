package metrics

import (
	"fmt"
	"net/http"
	"testing"
)

func TestContainsHeader_1(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{"text/html", "text/html", true},
		{"application/xhtml+xml", "application/xhtml+xml", true},
		{"application/xml", "application/xml", true},
		{"*/*", "*/*", true},
		{"text/plain", "text/plain", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := containsHeader(req, "Accept", tt.value); got != tt.want {
				t.Errorf("containsHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
