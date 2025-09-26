package hydratation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFill_1(t *testing.T) {
	type args struct {
		field reflect.Value
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

			if err := fill(tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("fill() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
