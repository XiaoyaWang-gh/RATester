package hugio

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewOpenReadSeekCloser_1(t *testing.T) {
	type args struct {
		r ReadSeekCloser
	}
	tests := []struct {
		name string
		args args
		want OpenReadSeekCloser
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

			if got := NewOpenReadSeekCloser(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOpenReadSeekCloser() = %v, want %v", got, tt.want)
			}
		})
	}
}
