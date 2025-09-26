package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := &context{
		request: req,
	}

	result := ctx.Request()

	if result != req {
		t.Errorf("Expected %v, got %v", req, result)
	}
}
