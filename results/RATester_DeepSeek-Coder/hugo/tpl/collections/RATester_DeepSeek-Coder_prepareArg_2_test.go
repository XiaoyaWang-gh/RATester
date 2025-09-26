package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrepareArg_2(t *testing.T) {
	type args struct {
		value   reflect.Value
		argType reflect.Type
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

			got, err := prepareArg(tt.args.value, tt.args.argType)
			if (err != nil) != tt.wantErr {
				t.Errorf("prepareArg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepareArg() = %v, want %v", got, tt.want)
			}
		})
	}
}
