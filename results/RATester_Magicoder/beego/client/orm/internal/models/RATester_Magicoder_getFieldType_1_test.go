package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetFieldType_1(t *testing.T) {
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantFt  int
		wantErr bool
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

			gotFt, err := getFieldType(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFieldType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFt != tt.wantFt {
				t.Errorf("getFieldType() = %v, want %v", gotFt, tt.wantFt)
			}
		})
	}
}
