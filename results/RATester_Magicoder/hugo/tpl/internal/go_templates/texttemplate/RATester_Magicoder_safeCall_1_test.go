package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsafeCall_1(t *testing.T) {
	type args struct {
		fun  reflect.Value
		args []reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantVal reflect.Value
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

			gotVal, err := safeCall(tt.args.fun, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("safeCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("safeCall() = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}
