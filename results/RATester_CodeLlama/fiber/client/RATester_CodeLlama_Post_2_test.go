package client

import (
	"context"
	"fmt"
	"testing"
)

func TestPost_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	url := "http://localhost:3000/post"
	r := &Request{
		ctx: context.Background(),
	}
	// when
	res, err := r.Post(url)
	// then
	if err != nil {
		t.Errorf("Post() error = %v", err)
		return
	}
	if res.StatusCode() != 200 {
		t.Errorf("Post() status code = %v, want %v", res.StatusCode(), 200)
	}
}
