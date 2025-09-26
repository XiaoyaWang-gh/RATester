package collections

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestapplyFnToThis_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		fn   reflect.Value
		this reflect.Value
		args []any
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

			got, err := applyFnToThis(tt.args.ctx, tt.args.fn, tt.args.this, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("applyFnToThis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyFnToThis() = %v, want %v", got, tt.want)
			}
		})
	}
}
