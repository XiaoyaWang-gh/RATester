package param

import (
	"fmt"
	"reflect"
	"testing"
)

func Testparse_15(t *testing.T) {
	type args struct {
		value  string
		toType reflect.Type
	}
	tests := []struct {
		name    string
		f       parserFunc
		args    args
		want    interface{}
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

			got, err := tt.f.parse(tt.args.value, tt.args.toType)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
