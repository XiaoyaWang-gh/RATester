package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/middlewares/accesslog"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestSetupAccessLog_1(t *testing.T) {
	type args struct {
		conf *types.AccessLog
	}
	tests := []struct {
		name string
		args args
		want *accesslog.Handler
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

			if got := setupAccessLog(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupAccessLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
