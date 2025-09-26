package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetValidFuncs_1(t *testing.T) {
	type args struct {
		f reflect.StructField
	}
	tests := []struct {
		name    string
		args    args
		wantVfs []ValidFunc
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

			gotVfs, err := getValidFuncs(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("getValidFuncs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVfs, tt.wantVfs) {
				t.Errorf("getValidFuncs() = %v, want %v", gotVfs, tt.wantVfs)
			}
		})
	}
}
