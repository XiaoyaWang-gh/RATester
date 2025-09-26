package bean

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMockBool_1(t *testing.T) {
	type args struct {
		tagValue string
		pvv      reflect.Value
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

			err := mockBool(tt.args.tagValue, tt.args.pvv)
			if (err != nil) != tt.wantErr {
				t.Errorf("mockBool() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
