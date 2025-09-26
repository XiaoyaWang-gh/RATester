package template

import (
	"fmt"
	"reflect"
	"testing"
)

func Testge_1(t *testing.T) {
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

			got, err := ge(tt.args.arg1, tt.args.arg2)
			if (err != nil) != tt.wantErr {
				t.Errorf("ge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ge() = %v, want %v", got, tt.want)
			}
		})
	}
}
