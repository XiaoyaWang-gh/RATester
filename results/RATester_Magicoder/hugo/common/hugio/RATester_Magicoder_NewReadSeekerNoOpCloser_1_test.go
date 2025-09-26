package hugio

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewReadSeekerNoOpCloser_1(t *testing.T) {
	type args struct {
		r ReadSeeker
	}
	tests := []struct {
		name string
		args args
		want ReadSeekCloser
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

			if got := NewReadSeekerNoOpCloser(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReadSeekerNoOpCloser() = %v, want %v", got, tt.want)
			}
		})
	}
}
