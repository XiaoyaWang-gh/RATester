package http

import (
	"fmt"
	"net/http"
	"testing"
)

func Testmatch_2(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		m    *matchersTree
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.match(tt.args.req); got != tt.want {
				t.Errorf("matchersTree.match() = %v, want %v", got, tt.want)
			}
		})
	}
}
