package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewWeightedPage_1(t *testing.T) {
	type args struct {
		weight int
		p      Page
		owner  Page
	}
	tests := []struct {
		name string
		args args
		want WeightedPage
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

			if got := NewWeightedPage(tt.args.weight, tt.args.p, tt.args.owner); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWeightedPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
