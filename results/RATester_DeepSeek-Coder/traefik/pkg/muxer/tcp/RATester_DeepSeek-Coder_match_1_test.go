package tcp

import (
	"fmt"
	"testing"
)

func TestMatch_1(t *testing.T) {
	type fields struct {
		matcher  func(ConnData) bool
		operator string
		left     *matchersTree
		right    *matchersTree
	}
	type args struct {
		meta ConnData
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
			if got := m.match(tt.args.meta); got != tt.want {
				t.Errorf("matchersTree.match() = %v, want %v", got, tt.want)
			}
		})
	}
}
