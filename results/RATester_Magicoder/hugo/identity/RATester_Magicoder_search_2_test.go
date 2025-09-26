package identity

import (
	"fmt"
	"testing"
)

func Testsearch_2(t *testing.T) {
	type args struct {
		sid   *searchID
		m     Manager
		depth int
	}
	tests := []struct {
		name string
		f    *Finder
		args args
		want FinderResult
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

			if got := tt.f.search(tt.args.sid, tt.args.m, tt.args.depth); got != tt.want {
				t.Errorf("Finder.search() = %v, want %v", got, tt.want)
			}
		})
	}
}
