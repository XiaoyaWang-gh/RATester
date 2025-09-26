package tplimpl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestfromVariantsSlice_1(t *testing.T) {
	type args struct {
		variants []string
	}
	tests := []struct {
		name  string
		s     *shortcodeTemplates
		args  args
		want  shortcodeVariant
		want1 bool
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

			got, got1 := tt.s.fromVariantsSlice(tt.args.variants)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fromVariantsSlice() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("fromVariantsSlice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
