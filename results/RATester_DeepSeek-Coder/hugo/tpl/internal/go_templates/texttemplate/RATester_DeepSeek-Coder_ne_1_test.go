package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNe_1(t *testing.T) {
	type args struct {
		arg1 reflect.Value
		arg2 reflect.Value
	}
	tests := []struct {
		name    string
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

			got, err := ne(tt.args.arg1, tt.args.arg2)
			if (err != nil) != tt.wantErr {
				t.Errorf("ne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ne() = %v, want %v", got, tt.want)
			}
		})
	}
}
