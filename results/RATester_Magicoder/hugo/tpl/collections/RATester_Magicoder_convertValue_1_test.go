package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconvertValue_1(t *testing.T) {
	type args struct {
		v  reflect.Value
		to reflect.Type
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

			got, err := convertValue(tt.args.v, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
