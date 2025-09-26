package middleware

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/containous/alice"
)

func TestBuildConstructor_1(t *testing.T) {
	type args struct {
		ctx            context.Context
		middlewareName string
	}
	tests := []struct {
		name    string
		args    args
		want    alice.Constructor
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

			b := &Builder{}
			got, err := b.buildConstructor(tt.args.ctx, tt.args.middlewareName)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildConstructor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildConstructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
