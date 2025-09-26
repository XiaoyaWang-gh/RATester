package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMatch_2(t *testing.T) {
	type fields struct {
		matcher  func(*http.Request) bool
		operator string
		left     *matchersTree
		right    *matchersTree
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
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

			m := &matchersTree{
				matcher:  tt.fields.matcher,
				operator: tt.fields.operator,
				left:     tt.fields.left,
				right:    tt.fields.right,
			}
			if got := m.match(tt.args.req); got != tt.want {
				t.Errorf("matchersTree.match() = %v, want %v", got, tt.want)
			}
		})
	}
}
