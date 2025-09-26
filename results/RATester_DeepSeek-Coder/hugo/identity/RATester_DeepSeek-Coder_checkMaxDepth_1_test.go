package identity

import (
	"fmt"
	"testing"
)

func TestCheckMaxDepth_1(t *testing.T) {
	type args struct {
		sid   *searchID
		level int
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

			if got := tt.f.checkMaxDepth(tt.args.sid, tt.args.level); got != tt.want {
				t.Errorf("Finder.checkMaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
