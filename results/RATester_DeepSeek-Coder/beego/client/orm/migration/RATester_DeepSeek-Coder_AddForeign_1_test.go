package migration

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddForeign_1(t *testing.T) {
	type args struct {
		foreign *Foreign
	}
	tests := []struct {
		name string
		m    *Migration
		args args
		want *Migration
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

			if got := tt.m.AddForeign(tt.args.foreign); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Migration.AddForeign() = %v, want %v", got, tt.want)
			}
		})
	}
}
