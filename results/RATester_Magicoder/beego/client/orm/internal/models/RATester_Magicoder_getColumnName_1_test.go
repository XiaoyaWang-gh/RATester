package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetColumnName_1(t *testing.T) {
	type args struct {
		ft        int
		addrField reflect.Value
		sf        reflect.StructField
		col       string
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := getColumnName(tt.args.ft, tt.args.addrField, tt.args.sf, tt.args.col); got != tt.want {
				t.Errorf("getColumnName() = %v, want %v", got, tt.want)
			}
		})
	}
}
