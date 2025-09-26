package orm

import (
	"fmt"
	"testing"
)

func TestSetOffset_1(t *testing.T) {
	type args struct {
		num interface{}
	}
	tests := []struct {
		name string
		o    *querySet
		args args
		want int64
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

			tt.o.setOffset(tt.args.num)
			if got := tt.o.offset; got != tt.want {
				t.Errorf("querySet.setOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}
