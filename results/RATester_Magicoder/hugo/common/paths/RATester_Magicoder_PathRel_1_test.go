package paths

import (
	"fmt"
	"testing"
)

func TestPathRel_1(t *testing.T) {
	type args struct {
		owner *Path
	}
	tests := []struct {
		name string
		p    *Path
		args args
		want string
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

			if got := tt.p.PathRel(tt.args.owner); got != tt.want {
				t.Errorf("Path.PathRel() = %v, want %v", got, tt.want)
			}
		})
	}
}
