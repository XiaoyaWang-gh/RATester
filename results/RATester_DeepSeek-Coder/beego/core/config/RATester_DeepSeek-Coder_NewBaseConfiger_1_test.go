package config

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestNewBaseConfiger_1(t *testing.T) {
	type args struct {
		reader func(ctx context.Context, key string) (string, error)
	}
	tests := []struct {
		name string
		args args
		want BaseConfiger
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

			if got := NewBaseConfiger(tt.args.reader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBaseConfiger() = %v, want %v", got, tt.want)
			}
		})
	}
}
