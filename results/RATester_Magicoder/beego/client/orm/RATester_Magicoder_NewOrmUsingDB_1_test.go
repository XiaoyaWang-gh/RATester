package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewOrmUsingDB_1(t *testing.T) {
	type args struct {
		aliasName string
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

			if got := NewOrmUsingDB(tt.args.aliasName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrmUsingDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
