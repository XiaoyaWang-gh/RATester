package commands

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugolib"
)

func TestIsStatic_1(t *testing.T) {
	type args struct {
		h        *hugolib.HugoSites
		filename string
	}
	tests := []struct {
		name string
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

			s := &staticSyncer{}
			if got := s.isStatic(tt.args.h, tt.args.filename); got != tt.want {
				t.Errorf("staticSyncer.isStatic() = %v, want %v", got, tt.want)
			}
		})
	}
}
