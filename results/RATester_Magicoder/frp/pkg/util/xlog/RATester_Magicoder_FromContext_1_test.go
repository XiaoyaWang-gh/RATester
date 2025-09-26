package xlog

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestFromContext_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		args   args
		wantXl *Logger
		wantOk bool
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

			gotXl, gotOk := FromContext(tt.args.ctx)
			if !reflect.DeepEqual(gotXl, tt.wantXl) {
				t.Errorf("FromContext() gotXl = %v, want %v", gotXl, tt.wantXl)
			}
			if gotOk != tt.wantOk {
				t.Errorf("FromContext() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
