package metric

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewStandardDateCounter_1(t *testing.T) {
	type args struct {
		reserveDays int64
	}
	tests := []struct {
		name string
		args args
		want *StandardDateCounter
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

			got := newStandardDateCounter(tt.args.reserveDays)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStandardDateCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}
