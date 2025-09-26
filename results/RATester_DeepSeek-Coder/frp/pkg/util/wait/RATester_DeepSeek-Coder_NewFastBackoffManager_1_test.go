package wait

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFastBackoffManager_1(t *testing.T) {
	type args struct {
		options FastBackoffOptions
	}
	tests := []struct {
		name string
		args args
		want BackoffManager
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

			if got := NewFastBackoffManager(tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFastBackoffManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
