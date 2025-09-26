package client

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestacquireErrChan_1(t *testing.T) {
	type args struct {
		errChanPool sync.Pool
	}
	tests := []struct {
		name string
		args args
		want chan error
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

			if got := acquireErrChan(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("acquireErrChan() = %v, want %v", got, tt.want)
			}
		})
	}
}
