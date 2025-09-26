package paths

import (
	"fmt"
	"testing"
)

func TestidentifierIndex_1(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		p    *Path
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

			if got := tt.p.identifierIndex(tt.args.i); got != tt.want {
				t.Errorf("Path.identifierIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
