package migration

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewCol_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		m    *Migration
		args args
		want *Column
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

			if got := tt.m.NewCol(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Migration.NewCol() = %v, want %v", got, tt.want)
			}
		})
	}
}
