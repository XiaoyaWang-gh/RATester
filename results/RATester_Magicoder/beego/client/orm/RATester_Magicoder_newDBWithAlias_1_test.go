package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewDBWithAlias_1(t *testing.T) {
	type args struct {
		al *alias
	}
	tests := []struct {
		name string
		args args
		want Ormer
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

			if got := newDBWithAlias(tt.args.al); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDBWithAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}
