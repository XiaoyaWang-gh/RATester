package orm

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestConnMaxIdletime_1(t *testing.T) {
	type args struct {
		v time.Duration
	}
	tests := []struct {
		name string
		args args
		want DBOption
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

			if got := ConnMaxIdletime(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnMaxIdletime() = %v, want %v", got, tt.want)
			}
		})
	}
}
