package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrySet_3(t *testing.T) {
	type args struct {
		value reflect.Value
		field reflect.StructField
		key   string
		opt   setOptions
	}
	tests := []struct {
		name    string
		r       *multipartRequest
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

			got, err := tt.r.TrySet(tt.args.value, tt.args.field, tt.args.key, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("TrySet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TrySet() = %v, want %v", got, tt.want)
			}
		})
	}
}
