package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetBoolField_1(t *testing.T) {
	type args struct {
		value string
		field reflect.Value
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

			if err := setBoolField(tt.args.value, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("setBoolField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
