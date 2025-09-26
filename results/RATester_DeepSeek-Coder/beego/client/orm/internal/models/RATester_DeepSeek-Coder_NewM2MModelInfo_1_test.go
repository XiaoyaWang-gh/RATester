package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewM2MModelInfo_1(t *testing.T) {
	type args struct {
		m1 *ModelInfo
		m2 *ModelInfo
	}
	tests := []struct {
		name   string
		args   args
		wantMi *ModelInfo
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

			gotMi := NewM2MModelInfo(tt.args.m1, tt.args.m2)
			if !reflect.DeepEqual(gotMi, tt.wantMi) {
				t.Errorf("NewM2MModelInfo() = %v, want %v", gotMi, tt.wantMi)
			}
		})
	}
}
