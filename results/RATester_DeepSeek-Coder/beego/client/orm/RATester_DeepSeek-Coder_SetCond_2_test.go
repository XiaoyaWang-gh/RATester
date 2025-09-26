package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetCond_2(t *testing.T) {
	type args struct {
		cond *Condition
	}
	tests := []struct {
		name string
		o    querySet
		args args
		want QuerySeter
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

			if got := tt.o.SetCond(tt.args.cond); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("querySet.SetCond() = %v, want %v", got, tt.want)
			}
		})
	}
}
