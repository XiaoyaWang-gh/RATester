package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTryToSetValue_1(t *testing.T) {
	type args struct {
		value  reflect.Value
		field  reflect.StructField
		setter setter
		tag    string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
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

			got, err := tryToSetValue(tt.args.value, tt.args.field, tt.args.setter, tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("tryToSetValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("tryToSetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
