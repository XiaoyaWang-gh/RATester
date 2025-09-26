package param

import (
	"fmt"
	"reflect"
	"testing"
)

func TestparseValue_1(t *testing.T) {
	type args struct {
		param      *MethodParam
		paramValue string
		paramType  reflect.Type
	}
	tests := []struct {
		name    string
		args    args
		want    reflect.Value
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

			got, err := parseValue(tt.args.param, tt.args.paramValue, tt.args.paramType)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
