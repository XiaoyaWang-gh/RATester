package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetNestedParamFn_1(t *testing.T) {
	type args struct {
		keyStr    string
		separator string
		lookupFn  func(key string) any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		want1   string
		want2   map[string]any
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

			got, got1, got2, err := GetNestedParamFn(tt.args.keyStr, tt.args.separator, tt.args.lookupFn)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNestedParamFn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNestedParamFn() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetNestedParamFn() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetNestedParamFn() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
