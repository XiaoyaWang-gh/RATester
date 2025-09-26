package net

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSetAuthFailDelay_1(t *testing.T) {
	type args struct {
		delay time.Duration
	}
	tests := []struct {
		name    string
		authMid *HTTPAuthMiddleware
		args    args
		want    *HTTPAuthMiddleware
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

			if got := tt.authMid.SetAuthFailDelay(tt.args.delay); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPAuthMiddleware.SetAuthFailDelay() = %v, want %v", got, tt.want)
			}
		})
	}
}
