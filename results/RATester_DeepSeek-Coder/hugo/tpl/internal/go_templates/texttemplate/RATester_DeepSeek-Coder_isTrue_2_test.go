package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsTrue_2(t *testing.T) {
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name      string
		args      args
		wantTruth bool
		wantOk    bool
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

			gotTruth, gotOk := isTrue(tt.args.val)
			if gotTruth != tt.wantTruth {
				t.Errorf("isTrue() gotTruth = %v, want %v", gotTruth, tt.wantTruth)
			}
			if gotOk != tt.wantOk {
				t.Errorf("isTrue() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
