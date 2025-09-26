package nathole

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewMakeHoleRecords_1(t *testing.T) {
	type args struct {
		c *NatFeature
		v *NatFeature
	}
	tests := []struct {
		name string
		args args
		want *MakeHoleRecords
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

			if got := NewMakeHoleRecords(tt.args.c, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMakeHoleRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
