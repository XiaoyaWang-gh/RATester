package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetTableUnique_1(t *testing.T) {
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name string
		args args
		want [][]string
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

			if got := GetTableUnique(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTableUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
