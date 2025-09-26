package orm

import (
	"fmt"
	"testing"
)

func TestExist_4(t *testing.T) {
	type args struct {
		md interface{}
	}
	tests := []struct {
		name string
		o    *queryM2M
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

			if got := tt.o.Exist(tt.args.md); got != tt.want {
				t.Errorf("queryM2M.Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
