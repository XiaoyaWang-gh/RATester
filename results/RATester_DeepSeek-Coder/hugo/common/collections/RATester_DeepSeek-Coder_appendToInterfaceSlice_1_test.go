package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppendToInterfaceSlice_1(t *testing.T) {
	type args struct {
		tov  reflect.Value
		from []any
	}
	tests := []struct {
		name    string
		args    args
		want    []any
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

			got, err := appendToInterfaceSlice(tt.args.tov, tt.args.from...)
			if (err != nil) != tt.wantErr {
				t.Errorf("appendToInterfaceSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendToInterfaceSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
