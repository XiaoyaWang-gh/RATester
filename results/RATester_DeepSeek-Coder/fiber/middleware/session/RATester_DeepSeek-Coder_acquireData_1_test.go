package session

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestAcquireData_1(t *testing.T) {
	type args struct {
		dataPool sync.Pool
	}
	tests := []struct {
		name string
		args args
		want *data
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

			if got := acquireData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("acquireData() = %v, want %v", got, tt.want)
			}
		})
	}
}
