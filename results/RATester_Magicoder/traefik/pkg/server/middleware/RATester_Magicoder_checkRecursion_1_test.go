package middleware

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestCheckRecursion_1(t *testing.T) {
	type args struct {
		ctx            context.Context
		middlewareName string
	}
	tests := []struct {
		name    string
		args    args
		want    context.Context
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

			got, err := checkRecursion(tt.args.ctx, tt.args.middlewareName)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkRecursion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
