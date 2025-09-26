package middlewares

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/rs/zerolog"
)

func TestGetLogger_1(t *testing.T) {
	type args struct {
		ctx            context.Context
		middleware     string
		middlewareType string
	}
	tests := []struct {
		name string
		args args
		want *zerolog.Logger
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

			if got := GetLogger(tt.args.ctx, tt.args.middleware, tt.args.middlewareType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
