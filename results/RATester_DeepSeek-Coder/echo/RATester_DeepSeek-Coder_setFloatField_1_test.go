package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetFloatField_1(t *testing.T) {
	type args struct {
		value   string
		bitSize int
		field   reflect.Value
	}
	tests := []struct {
		name    string
		args    args
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

			err := setFloatField(tt.args.value, tt.args.bitSize, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("setFloatField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
