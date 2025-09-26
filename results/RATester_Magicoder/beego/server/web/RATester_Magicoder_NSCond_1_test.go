package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNSCond_1(t *testing.T) {
	type args struct {
		cond namespaceCond
	}
	tests := []struct {
		name string
		args args
		want LinkNamespace
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

			if got := NSCond(tt.args.cond); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NSCond() = %v, want %v", got, tt.want)
			}
		})
	}
}
