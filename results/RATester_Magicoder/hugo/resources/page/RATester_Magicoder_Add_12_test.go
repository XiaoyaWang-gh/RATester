package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdd_12(t *testing.T) {
	type args struct {
		el []string
	}
	tests := []struct {
		name string
		p    *pagePathBuilder
		args args
		want []string
	}{
		{
			name: "Test case 1",
			p:    &pagePathBuilder{els: []string{"a", "b", "c"}},
			args: args{el: []string{"d", "e", "f"}},
			want: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			name: "Test case 2",
			p:    &pagePathBuilder{els: []string{"g", "h", "i"}},
			args: args{el: []string{"j", "k", "l"}},
			want: []string{"g", "h", "i", "j", "k", "l"},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.p.Add(tt.args.el...)
			if !reflect.DeepEqual(tt.p.els, tt.want) {
				t.Errorf("Add() = %v, want %v", tt.p.els, tt.want)
			}
		})
	}
}
