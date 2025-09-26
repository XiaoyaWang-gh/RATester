package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsApplicableTableForDB_1(t *testing.T) {
	type args struct {
		val reflect.Value
		db  string
	}
	tests := []struct {
		name string
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

			if got := IsApplicableTableForDB(tt.args.val, tt.args.db); got != tt.want {
				t.Errorf("IsApplicableTableForDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
