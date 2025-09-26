package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetNested_1(t *testing.T) {
	type args struct {
		m       map[string]any
		indices []string
	}
	tests := []struct {
		name  string
		args  args
		want  any
		want1 string
		want2 map[string]any
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

			got, got1, got2 := getNested(tt.args.m, tt.args.indices)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNested() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getNested() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("getNested() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
