package tplimpl

import (
	"fmt"
	"testing"
)

func TestIndexOf_1(t *testing.T) {
	type args struct {
		variants []string
	}
	tests := []struct {
		name string
		s    *shortcodeTemplates
		args args
		want int
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

			if got := tt.s.indexOf(tt.args.variants); got != tt.want {
				t.Errorf("shortcodeTemplates.indexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
