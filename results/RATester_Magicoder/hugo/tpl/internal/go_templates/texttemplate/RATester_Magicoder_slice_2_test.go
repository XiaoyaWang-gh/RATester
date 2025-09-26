package template

import (
	"fmt"
	"reflect"
	"testing"
)

func Testslice_2(t *testing.T) {
	type args struct {
		item    reflect.Value
		indexes []reflect.Value
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

			got, err := slice(tt.args.item, tt.args.indexes...)
			if (err != nil) != tt.wantErr {
				t.Errorf("slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
