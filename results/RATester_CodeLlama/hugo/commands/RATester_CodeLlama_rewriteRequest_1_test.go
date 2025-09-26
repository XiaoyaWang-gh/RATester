package commands

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRewriteRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &fileServer{}
	r := &http.Request{}
	toPath := "toPath"
	r2 := f.rewriteRequest(r, toPath)
	if r2.URL.Path != toPath {
		t.Errorf("r2.URL.Path = %v, want %v", r2.URL.Path, toPath)
	}
}
