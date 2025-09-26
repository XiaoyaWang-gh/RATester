package pagemeta

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUpdateDateAndLastmodAndPublishDateIfAfter_1(t *testing.T) {
	type args struct {
		in Dates
	}
	tests := []struct {
		name string
		d    *Dates
		args args
		want *Dates
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

			tt.d.UpdateDateAndLastmodAndPublishDateIfAfter(tt.args.in)
			if !reflect.DeepEqual(tt.d, tt.want) {
				t.Errorf("UpdateDateAndLastmodAndPublishDateIfAfter() = %v, want %v", tt.d, tt.want)
			}
		})
	}
}
