package hugolib

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestprepareContext_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		c    *cachedContentScope
		args args
		want context.Context
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

			if got := tt.c.prepareContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cachedContentScope.prepareContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
