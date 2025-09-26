package try

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHasHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	header := "Content-Type"
	res := &http.Response{
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}
	if err := HasHeader(header)(res); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
