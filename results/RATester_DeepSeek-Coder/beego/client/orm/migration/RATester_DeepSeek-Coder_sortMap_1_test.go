package migration

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortMap_1(t *testing.T) {
	type args struct {
		m map[string]Migrationer
	}
	tests := []struct {
		name string
		args args
		want dataSlice
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

			if got := sortMap(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
