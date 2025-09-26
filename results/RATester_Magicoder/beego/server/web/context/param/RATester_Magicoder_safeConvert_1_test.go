package param

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsafeConvert_1(t *testing.T) {
	type args struct {
		value reflect.Value
		t     reflect.Type
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

			got, err := safeConvert(tt.args.value, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("safeConvert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("safeConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}
