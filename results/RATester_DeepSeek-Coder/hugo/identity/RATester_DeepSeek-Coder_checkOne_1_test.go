package identity

import (
	"fmt"
	"testing"
)

func TestCheckOne_1(t *testing.T) {
	type args struct {
		sid   *searchID
		v     Identity
		depth int
	}
	tests := []struct {
		name string
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

			f := &Finder{}
			if got := f.checkOne(tt.args.sid, tt.args.v, tt.args.depth); got != tt.want {
				t.Errorf("Finder.checkOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
